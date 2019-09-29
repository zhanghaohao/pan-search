package pansearch

import (
	"baidupan/pansearch/collect"
)

func DeDup(bdps collect.BDPS) collect.BDPS {
	if len(bdps) == 0 || bdps == nil {
		return nil
	}
	var newBdps collect.BDPS
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