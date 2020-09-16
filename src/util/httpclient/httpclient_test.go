package httpclient

import (
	"testing"
	"github.com/PuerkitoBio/goquery"
	"util/logger"
)

func TestHttpDoWithProxy(t *testing.T) {
	url := "https://pan.baidu.com/"
	resp, err := HttpDoWithProxy(url)
	if err != nil {
		t.Error(err)
	}
	dom, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		t.Error(err)
	}
	logger.Info.Println(dom.Html())
}
