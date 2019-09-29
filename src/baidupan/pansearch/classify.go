package pansearch

import "baidupan/pansearch/collect"

func Classify(category string, bdps collect.BDPS) collect.BDPS {
	var videoBdps, documentBdps, audioBdps, seedBdps, pictureBdps, folderBdps, archiveBdps, unknownBdps collect.BDPS
	for _, bdp := range bdps {
		if bdp.Category == "video" {
			videoBdps = append(videoBdps, bdp)
		} else if bdp.Category == "document" {
			documentBdps = append(documentBdps, bdp)
		} else if bdp.Category == "audio" {
			audioBdps = append(audioBdps, bdp)
		} else if bdp.Category == "seed" {
			seedBdps = append(seedBdps, bdp)
		} else if bdp.Category == "picture" {
			pictureBdps = append(pictureBdps, bdp)
		} else if bdp.Category == "folder" {
			folderBdps = append(folderBdps, bdp)
		} else if bdp.Category == "archive" {
			archiveBdps = append(archiveBdps, bdp)
		} else {
			unknownBdps = append(unknownBdps, bdp)
		}
	}
	if category == "video" {
		return videoBdps
	} else if category == "document" {
		return documentBdps
	} else if category == "audio" {
		return audioBdps
	} else if category == "seed" {
		return seedBdps
	} else if category == "picture" {
		return pictureBdps
	} else if category == "folder" {
		return folderBdps
	} else if category == "archive" {
		return archiveBdps
	} else if category == "all" {
		return bdps
	} else {
		return unknownBdps
	}
}

