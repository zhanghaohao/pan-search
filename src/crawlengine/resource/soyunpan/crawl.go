package soyunpan

import (
	"time"
	"net/url"
	"github.com/PuerkitoBio/goquery"
	logger "util/logger"
	"strconv"
	"strings"
	"util"
	"crawlengine/resource/common"
)

const (
	ChannelName = "soyunpan"
)

type CrawlImpl struct {
	Keyword 		string
	BaseUrl 		string
}

func New(keyword string) common.Crawler {
	return &CrawlImpl{
		Keyword: keyword,
		BaseUrl: "http://www.soyunpan.com",
	}
}

func (c *CrawlImpl) Crawl() (common.BDPS) {
	var bdps common.BDPS
	var bdpCount = 0
	channelBdp := make(chan common.BDP)
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

func (c *CrawlImpl) CrawlPage(i int, channelBdp chan <- common.BDP, chanCount chan <- int) {
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

func (c *CrawlImpl) CrawlBdp(bdpUrl string, channelBdp chan <- common.BDP) {
	var bdp common.BDP
	dom, err := goquery.NewDocument(bdpUrl)
	if err != nil {
		logger.Error.Println(err.Error())
		channelBdp <- common.BDP{}
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
		logger.Error.Println("Cannot find pan url from: ", bdpUrl)
		channelBdp <- common.BDP{}
		return
	}
	s, err := url.QueryUnescape(strings.TrimPrefix(u, "http://www.soyunpan.com/down.php?url="))
	if err != nil {
		logger.Error.Printf("cannot unescape url %s", u)
		channelBdp <- common.BDP{}
		return
	}
	bdp.Url = common.HandleURL(s)
	bdp.Category = util.ExtToCategory(bdp.Ext)
	bdp.Resource = ChannelName
	channelBdp <- bdp
}