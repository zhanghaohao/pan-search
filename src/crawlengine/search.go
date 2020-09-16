package crawlengine

import (
	"crawlengine/resource/common"
	"util"
	"time"
	"util/logger"
	"crawlengine/process"
)

func CrawlSearch(keyword string) (cookedBdps common.BDPS, err error) {
	defer util.PrintCostTime(time.Now())
	logger.Info.Println("start crawling search engine...")
	// resource
	//logger.Info.Println("Progressing to Crawl Stage ...")
	bdps, err := process.Crawl(keyword)
	// combine
	//logger.Info.Println("Progressing to Combine Stage ...")
	bdps = process.Combine(bdps)
	// deduplicate
	//logger.Info.Println("Progressing to DeDup Stage ...")
	bdps = process.DeDup(bdps)
	// validate
	//logger.Info.Println("Progressing to Validate Stage ...")
	bdps, err = process.Validate(bdps)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	if len(bdps) == 0 || bdps == nil {
		return
	}
	// rank bdps by weight
	cookedBdps = process.Rank(bdps)
	return
}
