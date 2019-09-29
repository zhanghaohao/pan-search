package pansearch

import (
	"baidupan/pansearch/collect"
	"github.com/PuerkitoBio/goquery"
	logger "util/log"
	"time"
	"util"
)

func Validate(bdps collect.BDPS) (collect.BDPS, error) {
	defer util.PrintCostTime(time.Now())
	if len(bdps) == 0 || bdps == nil {
		return nil, nil
	}
	var newBdps collect.BDPS
	invalidBdpChannel := make(chan collect.BDP)
	validBdpChannel := make(chan collect.BDP)
	count := len(bdps)
	for index, bdp := range bdps {

		go func(index int, bdp collect.BDP) {
			url := bdp.Url
			dom, err := goquery.NewDocument(url)
			if err != nil {
				logger.Error.Printf("validate url error, %+v\n %s", bdp, err.Error())
				invalidBdpChannel <- bdp
				return
			}
			selection1 := dom.Find(".share-error-left")
			selection2 := dom.Find(".share-error")
			selection3 := dom.Find(".error-404")  // 页面不存在
			if len(selection1.Nodes) != 0 || len(selection2.Nodes) != 0 || len(selection3.Nodes) != 0 {
				// write bdps into channel
				invalidBdpChannel <- bdp
			} else {
				validBdpChannel <- bdp
				//logger.Info.Println(url)
			}
			//fmt.Printf("%d, %s\n", index, url)
		}(index, bdp)

	}
	timeout := time.After(2 * time.Second)
	LOOP:
		for {
			select {
				// limit validating stage time to 2 seconds
				case <- timeout:
					//logger.Info.Println("timeout .........")
					break LOOP
				case validTmpBdp := <- validBdpChannel :
					//logger.Info.Println("valid bdp", validTmpBdp)
					newBdps = append(newBdps, validTmpBdp)
					count --
					if count == 0 {
						close(invalidBdpChannel)
						close(validBdpChannel)
						break LOOP
					}
				case  <- invalidBdpChannel :
					//logger.Info.Println("invalid bdp", invalidTmpBdp)
					count --
					if count == 0 {
						close(invalidBdpChannel)
						close(validBdpChannel)
						break LOOP
					}
			}

		}
	return newBdps, nil
}
