package httpclient

import (
	"net/http"
	"time"
	"strings"
	"util/logger"
	"net/url"
	"crypto/tls"
	"io/ioutil"
	"math/rand"
)

var userAgents = []string{
	"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_8; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Windows; U; Windows NT 6.1; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0;",
	"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0)",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0)",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.6; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",
	"Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; en) Presto/2.8.131 Version/11.11",
	"Opera/9.80 (Windows NT 6.1; U; en) Presto/2.8.131 Version/11.11",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; 360SE)",
	"Mozilla/5.0 (iPhone; U; CPU iPhone OS 4_3_3 like Mac OS X; en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
}

type httpClient struct {
	*http.Client
}

type httpRequest struct {
	*http.Request
}

type HttpClientFactory interface {
	WithTimeout(timeout time.Duration) *httpClient
	WithProxy(proxyURL string) *httpClient
	Do(request *httpRequest) (resp *http.Response, err error)
	DoForBody(request *httpRequest) (body []byte, err error)
}

type HttpRequestFactory interface {
	WithUserAgent(userAgent string) *httpRequest
	WithDefaultUserAgent() *httpRequest
	WithRandomUserAgent() *httpRequest
	WithReferer(referer string) *httpRequest
	WithCookie(cookie map[string]string) *httpRequest
}

func NewClient() (client *httpClient) {
	return &httpClient{
		&http.Client{},
	}
}

func NewRequest(url string, header map[string]string, httpMethod string, httpGetParams map[string]string, httpPostBody string) (request *httpRequest, err error) {
	request = new(httpRequest)
	req, err := http.NewRequest(strings.ToUpper( httpMethod ), url, strings.NewReader(httpPostBody))
	if err != nil {
		logger.Info.Println(err)
		return
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	if httpGetParams != nil {
		q := req.URL.Query()
		for k, v := range httpGetParams {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}
	request.Request = req
	return
}

func (c *httpClient) WithTimeout(timeout time.Duration) *httpClient {
	c.Client.Timeout = timeout
	return c
}

func (c *httpClient) WithProxy(proxyURL string) *httpClient {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(proxyURL)
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		Proxy: proxy,
	}
	c.Client.Transport = tr
	return c
}

func (r *httpRequest) WithUserAgent(userAgent string) *httpRequest {
	r.Request.Header.Add("User-Agent", userAgent)
	return r
}

func (r *httpRequest) WithDefaultUserAgent() *httpRequest {
	userAgent := "Mozilla/5.0 (Linux; Android 5.0) AppleWebKit/537.36 (KHTML, like Gecko) Mobile Safari/537.36 (compatible; Bytespider; https://zhanzhang.toutiao.com/)"
	r.Request.Header.Add("User-Agent", userAgent)
	return r
}

func (r *httpRequest) WithRandomUserAgent() *httpRequest {
	random := rand.Intn(len(userAgents))
	r.Request.Header.Add("User-Agent", userAgents[random])
	return r
}

func (r *httpRequest) WithReferer(referer string) *httpRequest {
	r.Request.Header.Add("Referer", referer)
	return r
}

func (r *httpRequest) WithCookie(cookie map[string]string) *httpRequest {
	for key, value := range cookie {
		cookie := &http.Cookie{
			Name: key,
			Value: value,
			HttpOnly: false,
		}
		r.Request.AddCookie(cookie)
	}
	return r
}

func (c *httpClient) Do(request *httpRequest) (resp *http.Response, err error) {
	resp, err = c.Client.Do(request.Request)
	if err != nil {
		logger.Info.Println(err)
		return
	}
	return
}

func (c *httpClient) DoForBody(request *httpRequest) (body []byte, err error) {
	resp, err := c.Client.Do(request.Request)
	if err != nil {
		logger.Info.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Info.Println(err)
		return
	}
	return
}