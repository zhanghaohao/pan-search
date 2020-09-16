package search

import (
	"io/ioutil"
	"sync"
	"testing"
	"util/httpclient"
	"util/logger"
)

func Test_constructCondition(t *testing.T) {
	// 并行
	wg := sync.WaitGroup{}
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			url := "https://dalipan.com/api/detail?id=4ae7c3fcb6184ed0d94a06c49279f258"
			request, err := httpclient.NewRequest(url, nil, "GET", nil, "")
			if err != nil {
				return
			}
			resp, err := httpclient.NewClient().Do(request.WithDefaultUserAgent())
			if err != nil {
				logger.Error.Println(err)
				return
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				logger.Error.Println(err)
				return
			}
			logger.Info.Println(resp.StatusCode)
			logger.Info.Println(string(body))
		}()
	}
	wg.Wait()
	logger.Info.Println("finish")
}

func Test_updateURLGetParam(t *testing.T) {

}
