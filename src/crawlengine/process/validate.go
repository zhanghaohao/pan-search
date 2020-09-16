package process

import (
	"github.com/PuerkitoBio/goquery"
	logger "util/logger"
	"time"
	"crawlengine/resource/common"
	"util/httpclient"
)

func Validate(bdps common.BDPS) (common.BDPS, error) {
	//defer util.PrintCostTime(time.Now())
	if len(bdps) == 0 || bdps == nil {
		return nil, nil
	}
	var newBdps common.BDPS
	invalidBdpChannel := make(chan common.BDP)
	validBdpChannel := make(chan common.BDP)
	count := len(bdps)
	for index, bdp := range bdps {

		go func(index int, bdp common.BDP) {
			url := bdp.Url
			isValid, err := VerifyBaiduPanURL(url)
			if err != nil {
				invalidBdpChannel <- bdp
				return
			}
			if isValid == false {
				// write bdps into channel
				invalidBdpChannel <- bdp
			} else {
				validBdpChannel <- bdp
			}
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

func VerifyBaiduPanURL(url string) (isValid bool, err error) {
	res, err := httpclient.NewClient().WithTimeout(5 * time.Second).Get(url)
	if err != nil {
		logger.Error.Println(err)
		return true, nil
	}
	dom, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		logger.Error.Printf("validate url error, %v", err)
		return
	}
	//logger.Info.Println(dom.Find("body").Html())
	selection := dom.Find("div.share-error-left")
	if len(selection.Nodes) != 0 {
		isValid = false
		return
	}
	selection = dom.Find(".share-error")
	if len(selection.Nodes) != 0 {
		isValid = false
		return
	}
	selection = dom.Find("div#share_nofound_des")
	if len(selection.Nodes) != 0 {
		isValid = false
		return
	}
	selection = dom.Find("div#app")
	if len(selection.AppendHtml("test").Nodes) != 0 {
		isValid = false
		return
	}
	selection = dom.Find("body.error-404")  // 页面不存在
	if len(selection.Nodes) != 0 {
		//logger.Info.Println("页面不存在")
		isValid = false
		return
	}
	isValid = true
	return
}

func VerifyXinlangPanURL(url string) (isValid bool, err error)  {
	request, err := httpclient.NewRequest(url, nil, "get", nil, "")
	if err != nil {
		return
	}
	res, err := httpclient.NewClient().Do(request.WithDefaultUserAgent())
	if err != nil {
		logger.Error.Println(err)
		return true, nil
	}
	dom, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		logger.Error.Printf("validate url error, %v", err)
		return
	}
	//logger.Info.Println(dom.Html())
	selection := dom.Find("div.vd_nobrowser_wrap")
	if len(selection.Nodes) != 0 {
		isValid = false
		return
	} else {
		isValid = true
		return
	}
}