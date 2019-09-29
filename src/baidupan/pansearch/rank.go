package pansearch

import (
	"baidupan/pansearch/collect"
	"sort"
)

func Rank(bdps collect.BDPS) collect.BDPS {
	// reorder according to bdp.Resource
	// the first weight: www.52sopan.com, the second weight: www.dupanbang.com

	//var sopanBdps collect.BDPS
	//var dupanbangBdps collect.BDPS
	//var newBdps collect.BDPS
	//for _, bdp := range bdps {
		//if bdp.Resource == "www.52sopan.com" {
		//	sopanBdps = append(sopanBdps, bdp)
		//} else {
		//	dupanbangBdps = append(dupanbangBdps, bdp)
		//}
	//}
	//newBdps = append(newBdps, sopanBdps...)
	//newBdps = append(newBdps, dupanbangBdps...)
	//return newBdps
	var newBdps collect.BDPS
	var weight int
	weightBdps := make(map[int]collect.BDPS)
	for _, bdp := range bdps {
		weight = bdp.Weight()
		weightBdps[weight] = append(weightBdps[weight], bdp)
	}
	weights := make([]int, len(weightBdps))
	for k := range weightBdps {
		weights = append(weights, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(weights)))
	for _, w := range weights {
		newBdps = append(newBdps, weightBdps[w]...)
	}
	return newBdps
}
