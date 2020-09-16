package common

import (
	"regexp"
)

// set max total count or too many request could break remote server
const (
	MaxTotalCount = 200
	MaxTotalPage1 = 10
	MaxTotalPage2 = 5
)

type BDPS []BDP

type BDP struct {
	ID 						int64					`json:"id"`
	RawID    				string					`json:"rawID"`
	Url      				string					`json:"url"`
	Title    				string					`json:"title"`
	Ext      				string					`json:"ext"`
	CTime    				string					`json:"cTime"`
	Size     				string					`json:"size"`
	HasPwd   				bool					`json:"hasPwd"`
	Password 				string					`json:"password"`
	Category 				string					`json:"category"`
	Resource 				string					`json:"resource"`
}

type Crawler interface {
	Crawl() (bdps BDPS)
}

func HandleURL(url string) string {
	// handle pan url
	// trim /wap in the middle of url
	reg := regexp.MustCompile(`(.*)(/wap)(.*)`)
	return reg.ReplaceAllString(url, "$1$3")
}
