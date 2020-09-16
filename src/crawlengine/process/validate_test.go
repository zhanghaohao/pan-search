package process

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func TestValidate(t *testing.T) {
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
	time.Sleep(2 * time.Second)
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

func TestVerifyXinlangPanURL(t *testing.T) {
	panURL := "https://vdisk.weibo.com/s/aJivMjX9r081F"
	isValid, err := VerifyXinlangPanURL(panURL)
	if err != nil {
		t.Error(err)
	}
	if isValid == true {
		t.Error("wanted invalid, but get valid")
	}
	panURL = "https://vdisk.weibo.com/s/vGB2PLMb-QTU"
	isValid, err = VerifyXinlangPanURL(panURL)
	if err != nil {
		t.Error(err)
	}
	if isValid == false {
		t.Error("wanted valid, but get invalid")
	}
}

func TestVerifyBaiduPanURL(t *testing.T) {
	url := "https://pan.baidu.com/s/1miv1kj2https://pan.baidu.com/share/link?shareid=4156794016&uk=2570243430&fid=916088277265528"
	isValid, err := VerifyBaiduPanURL(url)
	if err != nil {
		t.Error(err)
	}
	if isValid == true {
		t.Errorf("want invalid, get valid")
	}
}
