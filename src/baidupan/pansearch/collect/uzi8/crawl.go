package uzi8

import (
	"baidupan/pansearch/collect"
	"time"
	"net/url"
	"github.com/PuerkitoBio/goquery"
	logger "util/log"
	"strings"
	"util"
	"regexp"
	"util/httpclient"
	"strconv"
)

type CrawlInfo collect.CrawlInfo

func (c *CrawlInfo) Crawl() collect.BDPS {
	var bdps collect.BDPS
	var bdpCount = 0
	channelBdp := make(chan collect.BDP)
	chanCount := make(chan int)
	pageCount := 1
	for i := 0; i < pageCount; i++ {
		go c.CrawlPage(i, channelBdp, chanCount)
	}
	timeout := time.After(3 * time.Second)
	for c := range chanCount {
		bdpCount += c
		pageCount --
		if pageCount == 0 {
			break
		}
	}
	logger.Info.Println(bdpCount)
	close(chanCount)
	LOOP:
		for {
			select {
			case <- timeout:
				break LOOP
			case bdp := <- channelBdp:
				//logger.Info.Println(bdp)
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
	logger.Info.Println(bdps)
	return bdps
}

func (c *CrawlInfo) CrawlPage(i int, channelBdp chan <- collect.BDP, chanCount chan <- int) {
	pageUrl := c.BaseUrl + "/search/kw" + url.QueryEscape(c.Keyword) + "pg" + strconv.Itoa(i)
	dom, err := goquery.NewDocument(pageUrl)
	if err != nil {
		logger.Error.Println(err.Error())
		chanCount <- 0
		return
	}
	//logger.Info.Println(dom.Html())
	var bdpUrls []string
	dom.Find("li.clear").Find(".title").Each(func(i int, selection *goquery.Selection) {
		bdpUrl, ok := selection.Find("a[href]").Attr("href")
		if ok == false {
			logger.Error.Println("Cannot find bdpUrl")
			logger.Error.Println(pageUrl)
		} else {
			bdpUrls = append(bdpUrls, c.BaseUrl + bdpUrl)
		}
	})
	for _, bdpUrl := range bdpUrls {
		go c.CrawlBdp(bdpUrl, channelBdp)
	}
	chanCount <- len(bdpUrls)
}

func (c *CrawlInfo) CrawlBdp(bdpUrl string, channelBdp chan <- collect.BDP)  {
	var bdp collect.BDP
	dom, err := goquery.NewDocument(bdpUrl)
	if err != nil {
		logger.Error.Println(err.Error())
		channelBdp <- collect.BDP{}
		return
	}
	bdp.Title = strings.TrimSpace(dom.Find("h1.title").Text())
	size := strings.TrimSpace(dom.Find("li:contains(大小)").Find("span").Text())
	bdp.Size = util.SizeConv(size)
	bdp.CTime = strings.TrimSpace(dom.Find("li:contains(时间)").Find("span").Text())
	ext := strings.TrimSpace(dom.Find("li:contains(类型)").Find("span").Text())
	if ext == "目录" || ext == "" {
		bdp.Ext = ""
	} else {
		reg := regexp.MustCompile(`.*\.(.*)\).*`)
		match := reg.FindStringSubmatch(ext)
		//logger.Info.Println(ext)
		bdp.Ext = match[1]
	}
	bdp.Category = util.ExtToCategory(bdp.Ext)
	bdp.Resource = "3"
	selectionFile := dom.Find("a#btn_r_file")
	selectionDir := dom.Find("a#btn_r_dir")
	var btnType string
	if len(selectionFile.Nodes) == 0 && len(selectionDir.Nodes) == 0 {
		logger.Info.Println("btn_r_file and btn_r_dir are not found")
		logger.Info.Println(dom.Html())
		channelBdp <- collect.BDP{}
		return
	} else if len(selectionFile.Nodes) != 0 {
		btnType = "file"
	} else {
		btnType = "dir"
	}
	url, hasPwd, password := c.getBdpUrl(bdpUrl, btnType)
	if url == "" {
		logger.Info.Println("url is empty")
		channelBdp <- collect.BDP{}
		return
	}
	bdp.Url = url
	bdp.HasPwd = hasPwd
	bdp.Password = password
	channelBdp <- bdp
}

func (c *CrawlInfo) getBdpUrl(url string, btnType string) (string, bool, string) {
	// url is of format http://uzi8.cn/file/33814477
	var bdpUrl string
	var hasPwd bool
	var password string
	referer := url
	reg := regexp.MustCompile(`.*file/(.*)$`)
	match := reg.FindStringSubmatch(url)
	fileId := match[1]
	var jumpUrl string
	if btnType == "file" {
		jumpUrl = c.BaseUrl + "/redirect/file?id=" + fileId
	} else {
		jumpUrl = c.BaseUrl + "/redirect/file?id=" + fileId + "&t=dir"
	}
	res, err := httpclient.HttpDoWithReferer(jumpUrl, referer)
	if err != nil {
		return "", false, ""
	}
	dom, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		logger.Error.Println(err.Error())
		return "", false, ""
	}
	//logger.Info.Println(dom.Html())
	selectionPwd := dom.Find(".tip_msg").Find("p:contains(文件提取码)")
	if len(selectionPwd.Nodes) != 0 {
		password = strings.TrimSpace(dom.Find("#tip_msg").Find("p:contains(文件提取码)").Find("span").Text())
		bdpUrl = strings.TrimSpace(dom.Find("#tip_msg").Find("p[style]").Find("a").Text())
		hasPwd = true
	} else {
		//logger.Info.Println(dom.Html())
		bdpUrl = strings.TrimSpace(dom.Find("#tip_msg").Find("p:contains(baidu)").Text())
		hasPwd = false
		password = ""
	}
	//logger.Info.Println(bdpUrl)
	return bdpUrl, hasPwd, password
}