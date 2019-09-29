package sopan

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"strconv"
	"regexp"
	logger "util/log"
	"util/httpclient"
	"util"
	"net/url"
	"strings"
	"bytes"
	"baidupan/pansearch/collect"
	"time"
	"util/db"
	"util/coreconfig"
)

type Crawl52sopanInterface interface {
	// get count of bdp
	CrawlCount() (*int, error)
	Crawl() ([]collect.BDP, error)
	GetPassword(id int) (*string, error)
}

type CrawlInfo collect.CrawlInfo

type Crawl52sopanCrawlCountResponse struct {
	Success 	bool	`json:"success"`
	Count 		string	`json:"count"`
}

type Crawl52sopanGetPasswordResponse struct {
	Success 	bool	`json:"success"`
	Password 	string	`json:"pwd"`
}

type Crawl52sopanCrawlElement struct {
	Id 		string			`json:"id"`
	Url 	string			`json:"url"`
	Pwd 	string			`json:"pwd"`
	CTime 	string			`json:"ctime"`
	Size 	string			`json:"size"`
	Context string			`json:"context"`
	User 	string			`json:"user"`
	Type 	string			`json:"type"`
	Ext 	string			`json:"ext"`
	Valid 	string			`json:"valid"`
	Report 	string			`json:"report"`
	Engine 	string			`json:"engine"`
	Tags 	string			`json:"tags"`
	HasPassword bool		`json:"has_pwd, omitempty"`
}

type Crawl52sopanCrawlResponse []Crawl52sopanCrawlElement

func (c *CrawlInfo) CrawlCount() (*int, error) {
	url := c.BaseUrl + "?mode=count&q=" + url.QueryEscape(c.Keyword)
	resp, err := httpclient.WithTimeout().Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//logger.Info.Println(url)
	//logger.Info.Println(string(body))
	var data Crawl52sopanCrawlCountResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		logger.Error.Println("Format CrawlCountResponse failed\n", err.Error())
		return nil, err
	}
	//logger.Info.Println("data.Success: ", data.Success)
	if data.Success == false {
		return nil, fmt.Errorf("CrawlCount return false")
	}
	totalCount, err := strconv.Atoi(data.Count)
	if err != nil {
		logger.Error.Println("Cannot convert string to int")
		return nil, err
	}
	logger.Info.Printf("Total count is: %d", totalCount)
	return &totalCount, nil
}

//func (c *Crawl52sopan) Crawl1() (collect.BDPS, error) {
//	totalCount, err := c.CrawlCount()
//	if err != nil {
//		return nil, err
//	}
//	if *totalCount == 0 {
//		return nil, nil
//	}
//	if *totalCount > collect.MaxTotalCount {
//		*totalCount = collect.MaxTotalCount
//	}
//	// calculate total page number
//	pageCount := (*totalCount + c.CountPerPage - 1)/c.CountPerPage
//	var bdps collect.BDPS
//	channelBdps := make(chan collect.BDPS)
//	//channelCount := make(chan int, pageCount)
//	// loop each page
//	for i := 0; i < pageCount; i++ {
//		go c.CrawlPage(i, channelBdps)
//	}
//	for tmpBdps := range channelBdps {
//		bdps = append(bdps, tmpBdps...)
//
//		pageCount --
//		if pageCount == 0 {
//			close(channelBdps)
//			break
//		}
//	}
//	bdps = c.WithPassword(bdps)
//	//logger.Info.Printf("Crawl bdps is: %v", bdps)
//	return bdps, nil
//}

func (c *CrawlInfo) Crawl() (collect.BDPS) {
	var bdps collect.BDPS
	channelBdps := make(chan collect.BDPS)
	pageCount := 3
	for i := 0; i < pageCount; i++ {
		go c.CrawlPage(i, channelBdps)
	}
	for tmpBdps := range channelBdps {
		bdps = append(bdps, tmpBdps...)
		pageCount --
		if pageCount == 0 {
			close(channelBdps)
			break
		}
	}
	//bdps = c.WithPassword(bdps)
	//logger.Info.Printf("Crawl bdps is: %v", bdps)
	return bdps
}

func (c *CrawlInfo) CrawlPage(i int, channelBdps chan <- collect.BDPS)  {
	url := c.BaseUrl + "?mode=so&q=" + url.QueryEscape(c.Keyword) + "&page_size=30&page_number=" + strconv.Itoa(i)
	resp, err := httpclient.WithTimeout().Get(url)
	if err != nil {
		logger.Error.Println(err.Error())
		channelBdps <- nil
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error.Println(err.Error())
		channelBdps <- nil
		return
	}
	resp.Body.Close()
	var crawlResponse Crawl52sopanCrawlResponse
	var bdp collect.BDP
	var tmpBdps collect.BDPS
	// if pansearch keyword is emtpy, resource returns empty , not []
	if len(body) == 0 {
		channelBdps <- nil
		return
	}
	err = json.Unmarshal(body, &crawlResponse)
	if err != nil {
		logger.Error.Println("Format BDP data failed\n", err.Error())
		channelBdps <- nil
		return
	}
	// skip nil slice
	if len(crawlResponse) == 0 {
		// push nil to channel or looping channel will hang forever
		channelBdps <- nil
		return
	}
	// loop each item of page
	for _, crawlElement := range crawlResponse {
		bdp.Id = crawlElement.Id
		bdp.Resource = "1"
		var baidupanUrl string
		if strings.HasPrefix(crawlElement.Url, "https") == false {
			baidupanUrl = "https://" + crawlElement.Url
		} else {
			baidupanUrl = crawlElement.Url
		}
		bdp.Url = collect.HandleURL(baidupanUrl)
		// convert unicode to string
		//bdp.Title = CorrectStr(crawlElement.Context)
		bdp.Title = crawlElement.Context
		// remove . if ext start with . , use `` instead of "" due to interpreted string literal problem
		re := regexp.MustCompile(`\.(.*)`)
		match := re.MatchString(crawlElement.Ext)
		if match == true {
			matcher := re.FindStringSubmatch(crawlElement.Ext)
			bdp.Ext = matcher[1]
		} else {
			bdp.Ext = crawlElement.Ext
		}
		bdp.Category = util.ExtToCategory(bdp.Ext)
		bdp.CTime = crawlElement.CTime
		if crawlElement.CTime == "" || crawlElement.CTime == "0000-00-00 00:00:00" {
			bdp.CTime = "2000-01-01 00:00:00"
		}
		bdp.Size = crawlElement.Size
		if crawlElement.HasPassword == true {
			//go c.GetPassword(crawlElement.Id, bdp.Url, channelUrlPassword)

			bdp.Password = ""
			bdp.HasPwd = true
		} else {
			bdp.HasPwd = false
			bdp.Password = ""
		}
		pattern := "/share/link"
		matched, err := regexp.MatchString(pattern, crawlElement.Url)
		if err != nil {
			logger.Error.Println(err.Error())
			channelBdps <- nil
			return
		}
		if matched == true {
			continue
		}
		tmpBdps = append(tmpBdps, bdp)
	}

	channelBdps <- tmpBdps
}

func WithPassword(bdps collect.BDPS) error {
	defer util.PrintCostTime(time.Now())
	// get password for all required bdps
	channelUrlPassword := make(chan map[string]string)
	passwordCount := 0
	for _, bdp := range bdps {
		if bdp.HasPwd == true {
			passwordCount ++
			go GetPassword(bdp.Id, bdp.Url, channelUrlPassword)
		}
	}
	if passwordCount == 0 {
		close(channelUrlPassword)
		return nil
	}
	ups := make(map[string]string)
	for up := range channelUrlPassword {
		//for i, bdp := range bdps {
		//	password, found := up[bdp.Id]
		//	if found == true {
		//		bdps[i].Password = password
		//	}
		//}
		for k, v := range up {
			ups[k] = v
		}
		passwordCount --
		if passwordCount == 0 {
			close(channelUrlPassword)
			break
		}
	}
	return filterPassword(ups)
}

func filterPassword(ups map[string]string) error {
	db, err := db.DBConnection()
	if err != nil {
		return err
	}
	defer db.Close()
	// Begin will get a connection with mysql before insert, so this method has high effect
	tx, _ := db.Begin()
	for url, password := range ups {
		sqlStr := fmt.Sprintf("update %s set password = '%s' where url = '%s'", coreconfig.CC.Mysql.TBaidupan, password, url)
		//logger.Info.Println(sqlStr)
		_, err := tx.Exec(sqlStr)
		if err != nil {
			logger.Error.Println(err.Error())
		}
	}
	tx.Commit()
	return nil
}

func GetPassword(id string, bdpUrl string, channel chan <- map[string]string) {
	url := "http://www.52sopan.com/search.php" + "?mode=get-password&id=" + id
	resp, err := httpclient.WithTimeout().Get(url)
	if err != nil {
		logger.Error.Println(err.Error())
		channel <- map[string]string{bdpUrl: ""}
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error.Println(err.Error())
		channel <- map[string]string{bdpUrl: ""}
		return
	}
	resp.Body.Close()
	var data Crawl52sopanGetPasswordResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		logger.Error.Printf("Format GetPasswordResponse data failed: %s", err.Error())
		channel <- map[string]string{bdpUrl: ""}
		return
	}
	if data.Success == false {
		logger.Error.Println("GetPassword return false")
		channel <- map[string]string{bdpUrl: ""}
		return
	}
	//logger.Info.Printf("Password is: %v", data.Password)
	channel <- map[string]string{bdpUrl: data.Password}
}

func CorrectStr(old string) string {
	str := strings.Replace(old, "u", "\\u", -1)
	buf := bytes.NewBuffer(nil)

	i, j := 0, len(str)
	for i < j {
		x := i + 6
		if x > j {
			buf.WriteString(str[i:])
			break
		}
		if str[i] == '\\' && str[i+1] == 'u' {
			hex := str[i+2 : x]
			r, err := strconv.ParseUint(hex, 16, 64)
			if err == nil {
				buf.WriteRune(rune(r))
			} else {
				buf.WriteString(str[i:x])
			}
			i = x
		} else {
			buf.WriteByte(str[i])
			i++
		}
	}
	return buf.String()
}

func UnicodeToString(u string) string {
	sUnicodev := strings.Split(u, "u")
	var context string
	for _, v := range sUnicodev {
		if len(v) < 1 {
			continue
		}
		temp, err := strconv.ParseInt(v, 16, 32)
		if err != nil {
			logger.Error.Println(err.Error())
		}
		context += fmt.Sprintf("%c", temp)
	}
	return context
}






