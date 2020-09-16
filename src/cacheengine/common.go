package cacheengine

import (
	"util/logger"
	"util/encrypt"
	"plugin/db"
	"strconv"
	"regexp"
	"util"
	"html/template"
)

func handle(rawMetaDatas []db.Pan, keyword string) (cookedMetaDatas []util.MetaData, err error) {
	for _, rawMetaData := range rawMetaDatas {
		var cookedMetaData util.MetaData
		// convert id to encrypted format
		idStr, err := encrypt.EncryptID(rawMetaData.Id)
		if err != nil {
			return nil, err
		}
		cookedMetaData.ID = idStr
		cookedMetaData.CTime = rawMetaData.CTime
		// (?i) ignore upper and lower case
		var title string
		reg, err := regexp.Compile("(?i)" + keyword)
		if err != nil {
			//logger.Error.Println(err)
			title = rawMetaData.Title
		} else {
			title = reg.ReplaceAllString(rawMetaData.Title, "<span class='highlight'>" + keyword + "</span>")
		}
		cookedMetaData.Title = template.HTML(title)
		cookedMetaData.Size, err = calSize(rawMetaData.Size)
		if err != nil {
			return nil, err
		}
		cookedMetaData.Icon = getIconByCategory(rawMetaData.Category)
		cookedMetaDatas = append(cookedMetaDatas, cookedMetaData)
	}
	return
}

func getIconByCategory(category string) (icon string) {
	if category == "video" {
		icon = "video.png"
	} else if category == "document" {
		icon = "document.png"
	} else if category == "audio" || category == "picture" {
		icon = "file.png"
	} else if category == "seed" {
		icon = "seed.png"
	} else if category == "folder" {
		icon = "folder.png"
	} else if category == "archive" {
		icon = "archive.png"
	} else {
		icon = "file.png"
	}
	return
}

func calSize(size string) (cookedSize string, err error) {
	sizeInt64, err := strconv.ParseInt(size, 10, 64)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	if sizeInt64 == 0 {
		cookedSize = "未知"
		return
	}
	sizeInt64 = sizeInt64/1024
	if sizeInt64 < 1024 {
		cookedSize = strconv.FormatInt(sizeInt64, 10) + " KB"
		return
	}
	sizeInt64 = sizeInt64/1024
	if sizeInt64 < 1024 {
		cookedSize = strconv.FormatInt(sizeInt64, 10) + " MB"
		return
	}
	sizeInt64 = sizeInt64/1024
	if sizeInt64 < 1024 {
		cookedSize = strconv.FormatInt(sizeInt64, 10) + " GB"
		return
	}
	sizeInt64 = sizeInt64/1024
	cookedSize = strconv.FormatInt(sizeInt64, 10) + " TB"
	return
}