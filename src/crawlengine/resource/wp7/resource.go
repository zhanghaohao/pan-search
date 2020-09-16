package wp7

import (
	"util/logger"
	"strconv"
	"sync"
	"util/httpclient"
	"os"
	"github.com/PuerkitoBio/goquery"
	"bufio"
	"strings"
	"fmt"
	"regexp"
	"crawlengine/resource/common"
	"util"
	"plugin/db"
	"time"
	"crawlengine/process"
	"plugin/config"
)

type ResourceKind string

const (
	Baidu ResourceKind = "百度网盘"
	Xinlang ResourceKind = "新浪微盘"
)

type mysql struct {
	mysqlEngine 				*db.MysqlEngine
}

func loadResources() (err error) {
	/*
	get original urls from file, process them and get the valid pan url or xinlangpan url.
	 */
	// get mysql engine
	mysql, err := initData()
	if err != nil {
		return
	}
	filePath := "./urls.txt"
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0766)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	defer f.Close()
	var batchSize = 10
	var batchURLs []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		url := scanner.Text()
		batchURLs = append(batchURLs, url)
		if len(batchURLs) == batchSize {
			logger.Info.Printf("handling batch urls with last url %s", url)
			err := mysql.handleBatchURLs(batchURLs)
			if err != nil {
				logger.Error.Println(err)
				logger.Error.Printf("last handled url is %s", batchURLs[len(batchURLs)-1])
				return err
			}
			batchURLs = nil
			time.Sleep(5*time.Second)
		}
	}
	if err := scanner.Err(); err != nil {
		logger.Error.Printf("cannot read file: %s, err: [%v]", filePath, err)
		return err
	}
	return
}

func initData() (c *mysql, err error) {
	configFactory, err := config.ReadConfigFile("../../../../plugin/config/config.yaml")
	if err != nil {
		return
	}
	mysqlEngine, err := db.NewMysqlEngineForConfig(configFactory.GetMysqlConfig())
	if err != nil {
		return
	}
	c = &mysql{
		mysqlEngine: mysqlEngine,
	}
	return
}

func (c *mysql) handleBatchURLs(batchURLs []string) (err error) {
	wg := sync.WaitGroup{}
	for _, url := range batchURLs {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			c.handleSingleURL(url)
		}(url)
	}
	wg.Wait()
	return
}

func (c *mysql) handleSingleURL(url string) (err error) {
	redirectURL := assembleRedirectURL(url)
	panURL, err := getPanURL(redirectURL)
	if err != nil {
		return
	}
	isValid, resourceKind, err := verifyPanURL(panURL)
	if err != nil {
		return
	}
	if isValid == false {
		return
	}
	//logger.Info.Printf("valid url %s", panURL)
	metaData, err := getMetaData(url, resourceKind)
	if err != nil {
		//logger.Error.Println(err)
		return
	}
	metaData.Url = panURL
	logger.Info.Printf("%+v", metaData)
	err = c.writeMetaData(metaData)
	if err != nil {
		return
	}
	return
}

func assembleRedirectURL(url string) (redirectURL string) {
	// convert original url to redirect url
	// get id
	segments := strings.Split(url, "/")
	id := segments[len(segments)-1]
	// assemble redirect url
	redirectURL = "https://wp7.net/redirect/file?id=" + id
	return
}

func getPanURL(redirectURL string) (panURL string, err error) {
	// get pan url or xinlangpan url by parsing redirect url
	request, err := httpclient.NewRequest(redirectURL, nil, "Get", nil, "")
	if err != nil {
		return
	}
	resp, err := httpclient.NewClient().Do(request.WithDefaultUserAgent())
	if err != nil {
		return
	}
	dom, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	//logger.Info.Println(dom.Html())
	panURL = dom.Find("div#tip_msg").Find("p:contains(http)").Text()
	if len(panURL) == 0 {
		err = fmt.Errorf("cannot find pan url by redirect url %s", redirectURL)
		logger.Error.Println(err)
		return
	}
	return
}

func verifyPanURL(panURL string) (isValid bool, resourceKind ResourceKind, err error) {
	// pan url may be baidu pan url or xinglang pan url
	matched, err := regexp.MatchString("baidu", panURL)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	if matched == true {
		resourceKind = Baidu
		isValid, err = process.VerifyBaiduPanURL(panURL)
		if err != nil {
			logger.Error.Println(err)
			return
		}
		return
	}
	//matched, err = regexp.MatchString("vdisk", panURL)
	//if err != nil {
	//	logger.Error.Println(err)
	//	return
	//}
	//if matched == true {
	//	resourceKind = Xinlang
	//	isValid, err = crawlengine.VerifyXinlangPanURL(panURL)
	//	if err != nil {
	//		logger.Error.Println(err)
	//		return
	//	}
	//	return
	//}
	// abandan other kind of resource
	err = fmt.Errorf("unknown kind of resource %s", panURL)
	return
}

func getMetaData(url string, resourceKind ResourceKind) (metaData common.BDP, err error) {
	// get resource meta data by parsing original url if pan url is valid
	// get title
	request, err := httpclient.NewRequest(url, nil, "get", nil, "")
	if err != nil {
		return
	}
	res, err := httpclient.NewClient().Do(request.WithDefaultUserAgent())
	if err != nil {
		logger.Error.Println(err)
		return
	}
	dom, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	metaData.Title = dom.Find("div.title-box").Find("h1").Text()
	if len(metaData.Title) == 0 {
		err = fmt.Errorf("can not parse url %s", url)
		logger.Error.Println(err)
		return
	}
	// get ext, rawExt is like (.rmvb)
	rawExt := dom.Find("ul.info-basic").Find("li:contains(文件类型)").Find("span").Text()
	regexpExt := regexp.MustCompile(`\(\.(\w+)\)`)
	metaData.Ext = regexpExt.FindStringSubmatch(rawExt)[1]
	// get ctime, raw ctime is like 分享时间：2014-12-11
	rawCTime := dom.Find("ul.info-basic").Find("li:contains(分享时间)").Text()
	regexpCTime := regexp.MustCompile(`分享时间：(\S+)`)
	metaData.CTime = regexpCTime.FindStringSubmatch(rawCTime)[1]
	// get size, raw size is like 文件大小：104MB
	rawSize := dom.Find("ul.info-basic").Find("li:contains(文件大小)").Text()
	regexpSize := regexp.MustCompile(`文件大小：(\S+)`)
	rawSize = regexpSize.FindStringSubmatch(rawSize)[1]
	metaData.Size = rawSize[0:len(rawSize)-2] + "000000"
	// get category
	metaData.Category = util.ExtToCategory(metaData.Ext)
	// get password if exist
	metaData.HasPwd = false
	metaData.Password = ""
	metaData.Resource = string(resourceKind)
	return
}

func (c *mysql) writeMetaData(metaData common.BDP) (err error) {
	// write resource meta data into mysql, ignore if already exist
	var pan db.Pan
	pan.Url = metaData.Url
	pan.Title = metaData.Title
	pan.Ext = metaData.Ext
	pan.CTime = metaData.CTime
	pan.Size = metaData.Size
	pan.HasPwd = metaData.HasPwd
	pan.Password = metaData.Password
	pan.Category = metaData.Category
	pan.Resource = metaData.Resource
	_, err = c.mysqlEngine.InsertOne(&pan)
	if err != nil {
		logger.Error.Println(err)
		return
	}

	//_, err = coreconfig.YamlParser()
	//if err != nil {
	//	logger.Error.Println(err)
	//	return
	//}
	//db, err := db.DBConnection()
	//if err != nil {
	//	return
	//}
	//defer db.Close()
	//sqlStr := fmt.Sprintf("insert ignore into %s (url, title, ext, ctime, size, haspwd, password, category, resource)  values ('%s', '%s', '%s', '%s', '%s', '%t', '%s', '%s', '%s')", coreconfig.CC.Mysql.TBaidupan, metaData.Url, metaData.Title, metaData.Ext, metaData.CTime, metaData.Size, metaData.HasPwd, metaData.Password, metaData.Category, metaData.Resource)
	//_, err = db.Exec(sqlStr)
	//if err != nil {
	//	logger.Error.Println(err)
	//	logger.Error.Println(sqlStr)
	//	return
	//}
	return
}

func getIDs() (err error) {
	// get all urls with id
	startID := int64(20394899)
	endID := int64(50000000)
	var batch []int64
	var batchSize = 50
	for id:=startID;id<endID;id++ {
		//logger.Info.Println(id)
		batch = append(batch, id)
		if len(batch) == batchSize {
			logger.Info.Printf("handling batch with last id %d ...", id)
			wg := sync.WaitGroup{}
			urlChan := make(chan string, batchSize)
			for _, id := range batch {
				wg.Add(1)
				go func(id int64) {
					defer wg.Done()
					checkURL(id, urlChan)
				}(id)
			}
			wg.Wait()
			close(urlChan)
			for url := range urlChan {
				err := writeURLToFile(url)
				if err != nil {
					return err
				}
			}
			batch = nil
		}
	}
	return 
}

func writeURLToFile(url string) (err error) {
	// write url into file
	//logger.Info.Printf("writing url %s into file", url)
	filePath := "./urls.txt"
	content := []byte(url + "\n")
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0766)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	defer f.Close()
	_, err = f.Write(content)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	return
}

func checkURL(id int64, urlChan chan string) (err error) {
	url := "https://wp7.net/share/file/" + strconv.FormatInt(id, 10)
	resp, err := httpclient.NewClient().WithTimeout(5 * time.Second).Get(url)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	dom, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	selection := dom.Find("title")
	title, err := selection.Html()
	if err != nil {
		logger.Error.Println(err)
		return
	}
	if title != "404 not found" {
		urlChan <- url
	}
	return
}

func getHttpStatusCode(url string) (statusCode int, err error) {
	resp, err := httpclient.NewClient().WithTimeout(5 * time.Second).Get(url)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	statusCode = resp.StatusCode
	return
}
