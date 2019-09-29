package collect

import "regexp"

// set max total count or too many request could break remote server
const (
	MaxTotalCount = 200
	MaxTotalPage1 = 10
	MaxTotalPage2 = 5
)

type BDPS []BDP

type BDP struct {
	Id			string		`json:"id"`
	Url 		string		`json:"url"`
	Title 		string		`json:"title"`
	Ext 		string		`json:"ext"`
	CTime 		string		`json:"ctime"`
	Size 		string		`json:"size"`
	HasPwd 		bool		`json:"haspwd"`
	Password 	string		`json:"password"`
	Category	string		`json:"category"`
	Resource 	string		`json:"resource"`
}

type CrawlInfo struct {
	Keyword 		string
	BaseUrl 		string
}

type Crawl interface {
	Crawl() BDPS
}

func (bdp *BDP) Weight() int {
	// 100 is the highest weight
	if bdp.Resource == "www.52sopan.com" {
		return 100
	} else if bdp.Resource == "www.soyunpan.com" {
		return 80
	} else {
		return 60
	}
}

func HandleURL(url string) string {
	// handle baidupan url
	// trim /wap in the middle of url
	reg := regexp.MustCompile(`(.*)(/wap)(.*)`)
	return reg.ReplaceAllString(url, "$1$3")
}