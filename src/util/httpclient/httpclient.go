package httpclient

import (
	"net/http"
	"time"
	logger "util/log"
	"crypto/tls"
)

func WithTimeout() *http.Client {
	timeout := time.Duration(time.Second * 2)
	client := http.Client{
		Timeout: timeout,
	}
	return &client
}

func HttpDoWithUserAgent(url string) (*http.Response, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 2,
	}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
	//req.Header.Add("Referer", baseUrl)
	res, err := client.Do(req)
	if err != nil {
		logger.Error.Println(err.Error())
		return nil, err
	}
	return res, nil
}

func HttpDoWithReferer(url string, referer string) (*http.Response, error) {
	client := &http.Client{
		Timeout:   time.Second * 2,
	}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Referer", referer)
	res, err := client.Do(req)
	if err != nil {
		logger.Error.Println(err.Error())
		return nil, err
	}
	return res, nil
}