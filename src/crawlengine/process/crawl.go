package process

import (
	"crawlengine/resource/common"
	"crawlengine/resource/sopan"
	logger "util/logger"
	"sync"
)

type Resource struct {
	resourceCh 					chan common.BDPS
}

func (r *Resource) Register(crawlers []common.Crawler) {
	wg := sync.WaitGroup{}
	for _, crawler := range crawlers {
		wg.Add(1)
		go func(crawler common.Crawler) {
			defer wg.Done()
			resources:= crawler.Crawl()
			r.resourceCh <- resources
		}(crawler)
	}
	wg.Wait()
	close(r.resourceCh)
}

func Crawl(keyword string) (bdps common.BDPS, err error) {
	//defer util.PrintCostTime(time.Now())
	var crawlers = []common.Crawler{
		sopan.New(keyword),
		//soyunpan.New(keyword),
	}
	resource := &Resource{
		resourceCh: make(chan common.BDPS, len(crawlers)),
	}
	resource.Register(crawlers)

	for r := range resource.resourceCh {
		bdps = append(bdps, r...)
	}
	if len(bdps) == 0 {
		logger.Warn.Println("Crawl nothing from all channels")
		return
	}
	return
}
