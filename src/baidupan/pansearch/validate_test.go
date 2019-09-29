package pansearch

import (
	"testing"
	"time"
	"net/http"
	"fmt"
	"io/ioutil"
	"github.com/PuerkitoBio/goquery"
)

func TestValidate(t *testing.T) {
	//urls := []string{
	//	// 80ms 获取百度网盘链接页面有验证码
	//	"http://www.panduoduo.net/s/name/%E8%89%B2%E6%88%92/1",
	//	"http://www.panduoduo.net/s/name/%E8%89%B2%E6%88%92/2",
	//	"http://www.panduoduo.net/s/name/%E8%89%B2%E6%88%92/3",
	//	"http://www.panduoduo.net/s/name/%E8%89%B2%E6%88%92/4",
	//	"http://www.panduoduo.net/s/name/%E8%89%B2%E6%88%92/5",
	//	"http://www.panduoduo.net/s/name/%E8%89%B2%E6%88%92/6",
	//	"http://www.panduoduo.net/s/name/%E8%89%B2%E6%88%92/7",
	//	"http://www.panduoduo.net/s/name/%E8%89%B2%E6%88%92/8",
	//	"http://www.panduoduo.net/s/name/%E8%89%B2%E6%88%92/9",
	//	"http://www.panduoduo.net/s/name/%E8%89%B2%E6%88%92/10",
	//	"http://www.panduoduo.net/s/name/%E8%89%B2%E6%88%92/11",
	//	"http://www.panduoduo.net/s/name/%E8%89%B2%E6%88%92/12",
	//}
	//urls := []string{
	//	// 200ms
	//	"http://www.soyunpan.com/search/%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9B-0-%E5%85%A8%E9%83%A8-1.html",
	//	"http://www.soyunpan.com/search/%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9B-0-%E5%85%A8%E9%83%A8-2.html",
	//	"http://www.soyunpan.com/search/%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9B-0-%E5%85%A8%E9%83%A8-3.html",
	//	"http://www.soyunpan.com/search/%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9B-0-%E5%85%A8%E9%83%A8-4.html",
	//	"http://www.soyunpan.com/search/%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9B-0-%E5%85%A8%E9%83%A8-5.html",
	//	"http://www.soyunpan.com/search/%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9B-0-%E5%85%A8%E9%83%A8-6.html",
	//	"http://www.soyunpan.com/search/%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9B-0-%E5%85%A8%E9%83%A8-7.html",
	//	"http://www.soyunpan.com/search/%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9B-0-%E5%85%A8%E9%83%A8-8.html",
	//	"http://www.soyunpan.com/search/%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9B-0-%E5%85%A8%E9%83%A8-9.html",
	//}
	urls := []string{
		// 503 Service Temporarily Unavailable
		"http://uzi8.cn/search/kw%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9Bpg1",
		"http://uzi8.cn/search/kw%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9Bpg2",
		"http://uzi8.cn/search/kw%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9Bpg3",
		"http://uzi8.cn/search/kw%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9Bpg4",
		"http://uzi8.cn/search/kw%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9Bpg5",
		"http://uzi8.cn/search/kw%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9Bpg6",
		"http://uzi8.cn/search/kw%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9Bpg7",
	}
	//urls := []string{
	//	// 需要验证码
	//	"http://www.panmeme.com/query?key=%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9B&type=0&p=1",
	//	"http://www.panmeme.com/query?key=%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9B&type=0&p=2",
	//	"http://www.panmeme.com/query?key=%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9B&type=0&p=3",
	//	"http://www.panmeme.com/query?key=%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9B&type=0&p=4",
	//	"http://www.panmeme.com/query?key=%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9B&type=0&p=5",
	//	"http://www.panmeme.com/query?key=%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9B&type=0&p=6",
	//	"http://www.panmeme.com/query?key=%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9B&type=0&p=7",
	//	"http://www.panmeme.com/query?key=%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9B&type=0&p=8",
	//	"http://www.panmeme.com/query?key=%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9B&type=0&p=9",
	//}
	//urls := []string{
	//	// 需要验证码
	//	"http://www.xgsoso.com/s?wd=%E8%89%B2%E8%AF%AB%E6%B1%A4%E5%94%AF%E7%89%88&st=4&p=1",
	//	"http://www.xgsoso.com/s?wd=%E8%89%B2%E8%AF%AB%E6%B1%A4%E5%94%AF%E7%89%88&st=4&p=2",
	//	"http://www.xgsoso.com/s?wd=%E8%89%B2%E8%AF%AB%E6%B1%A4%E5%94%AF%E7%89%88&st=4&p=3",
	//	"http://www.xgsoso.com/s?wd=%E8%89%B2%E8%AF%AB%E6%B1%A4%E5%94%AF%E7%89%88&st=4&p=4",
	//	"http://www.xgsoso.com/s?wd=%E8%89%B2%E8%AF%AB%E6%B1%A4%E5%94%AF%E7%89%88&st=4&p=5",
	//	"http://www.xgsoso.com/s?wd=%E8%89%B2%E8%AF%AB%E6%B1%A4%E5%94%AF%E7%89%88&st=4&p=6",
	//	"http://www.xgsoso.com/s?wd=%E8%89%B2%E8%AF%AB%E6%B1%A4%E5%94%AF%E7%89%88&st=4&p=7",
	//	"http://www.xgsoso.com/s?wd=%E8%89%B2%E8%AF%AB%E6%B1%A4%E5%94%AF%E7%89%88&st=4&p=8",
	//}
	//urls := []string{
	//	// 403 forbidden
	//	"http://www.pan58.com/s?wd=%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9B&s=0&t=1&p=1",
	//	"http://www.pan58.com/s?wd=%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9B&s=0&t=1&p=2",
	//	"http://www.pan58.com/s?wd=%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9B&s=0&t=1&p=3",
	//	"http://www.pan58.com/s?wd=%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9B&s=0&t=1&p=4",
	//	"http://www.pan58.com/s?wd=%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9B&s=0&t=1&p=5",
	//	"http://www.pan58.com/s?wd=%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9B&s=0&t=1&p=6",
	//	"http://www.pan58.com/s?wd=%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9B&s=0&t=1&p=7",
	//	"http://www.pan58.com/s?wd=%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9B&s=0&t=1&p=8",
	//	"http://www.pan58.com/s?wd=%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9B&s=0&t=1&p=9",
	//}
	//urls := []string{
	//	// 需要验证码
	//	"http://www.repanso.com/q?wd=%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9B&type=0&p=1",
	//	"http://www.repanso.com/q?wd=%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9B&type=0&p=2",
	//	"http://www.repanso.com/q?wd=%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9B&type=0&p=3",
	//	"http://www.repanso.com/q?wd=%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9B&type=0&p=4",
	//	"http://www.repanso.com/q?wd=%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9B&type=0&p=5",
	//	"http://www.repanso.com/q?wd=%E6%96%B0%E5%9C%B0%E7%BE%A4%E5%B2%9B&type=0&p=6",
	//}
	//urls := []string{
	//	// 800ms
	//	"http://www.bdyso.com/t/KGTnA83d65c6Q/?p=1",
	//	"http://www.bdyso.com/t/KGTnA83d65c6Q/?p=2",
	//	"http://www.bdyso.com/t/KGTnA83d65c6Q/?p=3",
	//	"http://www.bdyso.com/t/KGTnA83d65c6Q/?p=4",
	//	"http://www.bdyso.com/t/KGTnA83d65c6Q/?p=5",
	//	"http://www.bdyso.com/t/KGTnA83d65c6Q/?p=6",
	//	"http://www.bdyso.com/t/KGTnA83d65c6Q/?p=7",
	//	"http://www.bdyso.com/t/KGTnA83d65c6Q/?p=8",
	//}
	//urls := []string{
	//	// 1000ms
	//	"http://www.yunpangou.com/21563787884136976?p=1",
	//	"http://www.yunpangou.com/21563787884136976?p=2",
	//	"http://www.yunpangou.com/21563787884136976?p=3",
	//	"http://www.yunpangou.com/21563787884136976?p=4",
	//	"http://www.yunpangou.com/21563787884136976?p=5",
	//	"http://www.yunpangou.com/21563787884136976?p=6",
	//	"http://www.yunpangou.com/21563787884136976?p=7",
	//	"http://www.yunpangou.com/21563787884136976?p=8",
	//	"http://www.yunpangou.com/21563787884136976?p=9",
	//	"http://www.yunpangou.com/21563787884136976?p=10",
	//}
	//url := "http://pdd.19mi.net/go/45486193"
	//for i := 1; i <= 100; i++ {
	for _, url := range urls {
		go func(url string) {
			start := time.Now()
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err.Error())
			}
			body, err := ioutil.ReadAll(resp.Body)
			strBody := string(body)
			resp.Body.Close()
			fmt.Println(strBody)
			stop := time.Now()
			cost := stop.Sub(start)
			fmt.Println(cost)
		}(url)
	}
	time.Sleep(2* time.Second)
}

func aTestValidate(t *testing.T) {
	url := "http://pan.baidu.com/wap/link?uk=3702969232&shareid=4097939565"
	dom, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println(dom.Html())
	selection := dom.Find(".share-error")
	if len(selection.Nodes) == 0 {
		fmt.Println("no error")
	} else {
		fmt.Println("error")
	}
}
