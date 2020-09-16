package process

import (
	"sort"
	"crawlengine/resource/common"
	"crawlengine/resource/sopan"
	"crawlengine/resource/soyunpan"
)

func Rank(bdps []common.BDP) (rankedBdps []common.BDP) {
	// reorder according to bdp.Resource
	// the first weight: www.52sopan.com, the second weight: www.dupanbang.com
	var newBdps []common.BDP
	var weight int
	weightBdps := make(map[int][]common.BDP)
	for _, bdp := range bdps {
		weight = getWeight(bdp)
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
	// convert to slice of pointer
	for _, newBdp := range newBdps {
		// use copy, because the pointers of original struct are the same
		bdpCopy := newBdp
		rankedBdps = append(rankedBdps, bdpCopy)
	}
	return
}

func getWeight(bdp common.BDP) int {
	// 100 is the highest weight
	if bdp.Resource == sopan.ChannelName {
		return 100
	} else if bdp.Resource == soyunpan.ChannelName {
		return 80
	} else {
		return 60
	}
}