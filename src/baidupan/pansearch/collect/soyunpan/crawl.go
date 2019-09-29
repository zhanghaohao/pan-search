package soyunpan

import (
	"baidupan/pansearch/collect"
	"time"
	"net/url"
	"github.com/PuerkitoBio/goquery"
	logger "util/log"
	"strconv"
	"strings"
	"util"
)

type CrawlInfo collect.CrawlInfo

func (c *CrawlInfo) Crawl() (collect.BDPS) {
	var bdps collect.BDPS
	var bdpCount = 0
	channelBdp := make(chan collect.BDP)
	chanCount := make(chan int)
	pageCount := 2
	for i := 0; i < pageCount; i++ {
		go c.CrawlPage(i, channelBdp, chanCount)
	}
	timeout := time.After(2 * time.Second)
	for bdpPageCount := range chanCount {
		bdpCount += bdpPageCount
		pageCount --
		if pageCount == 0 {
			break
		}
	}
	//logger.Info.Println(bdpCount)
	close(chanCount)
	LOOP:
		for {
			select {
			case <- timeout:
				break LOOP
			case bdp := <- channelBdp:
				if bdp.Url != "" {
					bdps = append(bdps, bdp)
				}
				bdpCount --
				if bdpCount == 0 {
					close(channelBdp)
					break LOOP
				}
			}
		}
	return bdps
}

func (c *CrawlInfo) CrawlPage(i int, channelBdp chan <- collect.BDP, chanCount chan <- int) {
	var bdpUrls []string
	pageUrl := c.BaseUrl + "/search/" + url.QueryEscape(c.Keyword + "-0-" + "全部" + "-") + strconv.Itoa(i) + ".html"
	dom, err := goquery.NewDocument(pageUrl)
	if err != nil {
		logger.Error.Println(err.Error())
		chanCount <- 0
		return
	}
	dom.Find(".main-x").Each(func(i int, selection *goquery.Selection) {
		bdpUrl, ok := selection.Find("a[href]").Attr("href")
		if ok == false {
			logger.Error.Println("Cannot find bdpUrl")
		} else {
			bdpUrls = append(bdpUrls, bdpUrl)
		}
	})
	for _, bdpUrl := range bdpUrls {
		go c.CrawlBdp(bdpUrl, channelBdp)
	}
	chanCount <- len(bdpUrls)
}

func (c *CrawlInfo) CrawlBdp(bdpUrl string, channelBdp chan <- collect.BDP) {
	var bdp collect.BDP
	dom, err := goquery.NewDocument(bdpUrl)
	if err != nil {
		logger.Error.Println(err.Error())
		channelBdp <- collect.BDP{}
		return
	}
	bdp.Title = strings.TrimSpace(dom.Find(".resource-h2").Text())
	bdp.CTime = strings.TrimSpace(dom.Find("li.x-right-li:contains(时间)").Find("span").Text())
	size := strings.TrimSpace(dom.Find("li.x-right-li:contains(大小)").Find("span").Text())
	bdp.Size = util.SizeConv(size)
	ext := strings.TrimSpace(dom.Find("li.x-right-li:contains(格式)").Find("span").Text())
	if ext == "" {
		bdp.Ext = ext
	} else if strings.HasPrefix(ext, ".") {
		bdp.Ext = strings.TrimPrefix(ext, ".")
	} else {
		bdp.Ext = ext
	}
	bdp.HasPwd = false
	u, ok := dom.Find("a.main-xzfx-a:contains(进入百度网盘下载)").Attr("href")
	if ok == false {
		logger.Error.Println("Cannot find baidupan url from: ", bdpUrl)
		channelBdp <- collect.BDP{}
		return
	}
	s, err := url.QueryUnescape(strings.TrimPrefix(u, "http://www.soyunpan.com/down.php?url="))
	if err != nil {
		logger.Error.Printf("cannot unescape url %s", u)
		channelBdp <- collect.BDP{}
		return
	}
	bdp.Url = collect.HandleURL(s)
	bdp.Category = util.ExtToCategory(bdp.Ext)
	bdp.Resource = "2"
	channelBdp <- bdp
}