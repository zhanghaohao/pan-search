package process

import (
	"crawlengine/resource/common"
)

func DeDup(bdps common.BDPS) common.BDPS {
	if len(bdps) == 0 || bdps == nil {
		return nil
	}
	var newBdps common.BDPS
	var urls []string
	for _, bdp := range bdps {
		if FindUrl(bdp.Url, urls) == true {
			continue
		} else {
			urls = append(urls, bdp.Url)
			newBdps = append(newBdps, bdp)
		}
	}
	return newBdps
}

func FindUrl(url string, urls []string) bool {
	for _, e := range urls {
		if url == e {
			return true
		}
	}
	return false
}