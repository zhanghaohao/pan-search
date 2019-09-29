package pansearch

import (
	"baidupan/pansearch/collect"
)

func Combine(bdpss ...collect.BDPS) collect.BDPS {
	var bigBdps collect.BDPS
	for _, bdps := range bdpss {
		if len(bdps) != 0 {
			bigBdps = append(bigBdps, bdps...)
		}
	}
	if len(bigBdps) == 0 {
		return nil
	}
	return bigBdps
}
