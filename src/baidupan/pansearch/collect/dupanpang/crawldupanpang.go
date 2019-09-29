package dupanpang

import (
	"baidupan/pansearch/collect"
	"net/url"
	"strconv"
	"github.com/PuerkitoBio/goquery"
	logger "util/log"
	"time"
	"strings"
	"util"
)

type Crawldupanpang struct {
	Keyword 		string
	BaseUrl 		string
}

func (c *Crawldupanpang) Crawl() (collect.BDPS, error) {
	var bdps collect.BDPS
	channelBdp := make(chan collect.BDP)
	for i := 0; i < collect.MaxTotalPage2; i++ {
		go c.CrawlPage(i, channelBdp)
	}
	timeout := time.After(3 * time.Second)
	LOOP:
		for {
			select {
			case <- timeout:
				break LOOP
			case bdp := <- channelBdp:
				bdps = append(bdps, bdp)
			}
		}
	return bdps, nil
}

func (c *Crawldupanpang) CrawlPage(i int, channelBdp chan <- collect.BDP) {
	var bdpUrls []string
	url := c.BaseUrl + "/q/" + url.QueryEscape(c.Keyword) + "?page=" + strconv.Itoa(i)
	dom, err := goquery.NewDocument(url)
	if err != nil {
		logger.Error.Println(err.Error())
		return
	}
	dom.Find(".pansearch-item").Each(func(i int, selection *goquery.Selection) {
		bdpUrl, ok := selection.Find("a[href]").Attr("href")
		if ok == false {
			logger.Error.Println("Cannot find bdpUrl")
		} else {
			bdpUrl := c.BaseUrl + bdpUrl
			bdpUrls = append(bdpUrls, bdpUrl)
		}
	})
	for _, bdpUrl := range bdpUrls {
		go c.CrawlBdp(bdpUrl, channelBdp)
	}
}

func (c *Crawldupanpang) CrawlBdp(bdpUrl string, channelBdp chan <- collect.BDP)  {
	var bdp collect.BDP
	dom, err := goquery.NewDocument(bdpUrl)
	if err != nil {
		logger.Error.Println(err.Error())
		return
	}
	bdp.Ext = strings.TrimSpace(strings.TrimPrefix(dom.Find("dd:contains(扩展名：)").Text(), "扩展名："))
	if bdp.Ext == "文件夹" {
		bdp.Ext = ""
	}
	title := strings.TrimSpace(dom.Find("h1[class=text-center]").Text())
	bdp.Title = title + "." + bdp.Ext
	bdp.Category = util.ExtToCategory(bdp.Ext)
	bdp.CTime = strings.TrimSpace(strings.TrimPrefix(dom.Find("dd:contains(分享日期：)").Text(), "分享日期："))
	size := strings.TrimSpace(strings.TrimPrefix(dom.Find("dd:contains(文件大小：)").Text(), "文件大小："))
	bdp.Size = util.SizeConv(size)
	bdp.Resource = "www.dupanbang.com"
	url, ok := dom.Find("div[id=linkModal]").Find("div[id=link-body]").Find("a").Attr("href")
	if ok == false {
		logger.Error.Println("Cannot find baidupan url from: ", bdpUrl)
		return
	}
	if url == "https://pan.baidu.com/s/" {
		url, ok := dom.Find(".grey-cascade").Attr("href")
		if ok == false {
			logger.Error.Println("Cannot find baidupan url from: ", bdpUrl)
			return
		}
		url = c.BaseUrl + url
		url, err := CrawlUrl(url)
		if err != nil {
			return
		}
	}
	bdp.Url = url
	channelBdp <- bdp
}

func CrawlUrl(bdpUrl string) (string, error) {
	dom, err := goquery.NewDocument(bdpUrl)
	if err != nil {
		logger.Error.Println(err.Error())
		return "", nil
	}
	url, ok := dom.Find(".page-go").Find("a").Attr("href")
	if ok == false {
		logger.Error.Println("Cannot find baidupan url from: ", bdpUrl)
		return "", nil
	}
	return url, nil
}

