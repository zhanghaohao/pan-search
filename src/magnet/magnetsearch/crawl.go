package magnetsearch

import (
	"net/url"
	logger "util/log"
	"strconv"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"regexp"
	"util"
	"util/httpclient"
)

type MAG struct {
	Magnet		string		`json:"magnet"`
	Title 		string 		`json:"title"`
	Size 		string		`json:"size"`
	Date 		string		`json:"date"`
	Hot 		string		`json:"hot"`
}

func Crawl(keyword string) ([]MAG, error) {
	var pageCount = 10
	chanMagnets := make(chan []MAG)
	var magnets []MAG
	for i := 0; i < pageCount; i ++ {
		go CrawlPage(i, keyword, chanMagnets)
	}
	for tmpMagnets := range chanMagnets {
		magnets = append(magnets, tmpMagnets...)
		pageCount --
		if pageCount == 0 {
			close(chanMagnets)
			break
		}
	}
	return magnets, nil
}

func CrawlPage(i int, keyword string, chanMagnets chan <- []MAG)  {
	var mags []MAG
	var mag MAG
	url := "https://www.zhongzilou.com/list/" + url.QueryEscape(keyword) + "/" + strconv.Itoa(i)
	res, err := httpclient.HttpDoWithUserAgent(url)
	dom, err := goquery.NewDocumentFromResponse(res)
	//defer res.Body.Close()
	//dom, err := goquery.NewDocument(url)
	if err != nil {
		//logger.Error.Println(err.Error())
		chanMagnets <- nil
		return
	}
	dom.Find("table").Each(func(i int, selection *goquery.Selection) {
		htmlTitle := strings.TrimSpace(selection.Find("h4").Find("a").Text())
		re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
		title := re.ReplaceAllString(htmlTitle, "")
		re, _ = regexp.Compile("\"")
		title = re.ReplaceAllString(htmlTitle, "")
		mag.Title = title
		mag.Date = selection.Find("td:contains(创建日期：)").Find("strong").Text()
		formatSize := selection.Find("td:contains(大小：)").Find("strong").Text()
		mag.Size = util.SizeConv(formatSize)
		mag.Hot = selection.Find("td:contains(热度：)").Find("strong").Text()
		magnet, ok := selection.Find("td[class=ls-magnet]").Find("a").Attr("href")
		if ok == false {
			logger.Error.Println("Cannot find magnet link from: ", url)
			return
		}
		mag.Magnet = magnet
		mags = append(mags, mag)
	})
	chanMagnets <- mags
}