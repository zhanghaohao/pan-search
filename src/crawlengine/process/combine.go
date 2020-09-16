package process

import (
	"crawlengine/resource/common"
)

func Combine(bdpss ...common.BDPS) common.BDPS {
	var bigBdps common.BDPS
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
