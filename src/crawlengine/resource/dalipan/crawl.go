package dalipan

import (
	"strconv"
	"util/httpclient"
	"util/logger"
	"fmt"
	"crawlengine/resource/common"
	"strings"
	"encoding/json"
	"util"
	"time"
	"os"
	"bufio"
	"plugin/db"
	"io/ioutil"
	"sync"
	"regexp"
)

const (
	maxTry = 2
)

var (
	mysqlConfig = &db.MysqlConfig{
		Host: "",
		Port: "3306",
		Database: "",
		User: "",
		Password: "",
	}
)

type MysqlEngine struct {
	*db.MysqlEngine
}

type silverBullet struct {
	mysqlEngine *MysqlEngine
	ipGenerator *ipGenerator
}

type rawMetaData struct {
	ID 					string					`json:"id"`
	URL 				string					`json:"url"`
	FileName 			string					`json:"filename"`
	Size 				int64					`json:"size"`
	IsDir 				int						`json:"isdir"`
	Ext 				string					`json:"ext"`
	HasPwd 				bool					`json:"haspwd"`
	Pwd 				string					`json:"pwd"`
	Valid 				int						`json:"valid"`
	CTime 				string					`json:"ctime"`
}

type searchRespBody struct {
	Resources 			[]searchRespBodyResource				`json:"resources"`
	Total 				int										`json:"total"`
}

type searchRespBodyResource struct {
	Res 				searchRespBodyRes 						`json:"res"`
}

type searchRespBodyRes struct {
	ID 					string 									`json:"id"`
}

type ipGenerator struct {
	a 					int
	b 					int
	c					int
	d 					int
}

func (o *ipGenerator) getIDs(keyword string, pageNumber int) (ids []string, err error) {
	url := "https://dalipan.com/api/search"
	method := "GET"
	header := map[string]string{
		"Content-Type": "application/json",
	}
	pageNum := strconv.Itoa(pageNumber)
	getParams := map[string]string{
		"kw": keyword,
		"page": pageNum,
		"ip": o.next(),
	}
	resp, err := httpDo(url, header, method, getParams, "")
	if err != nil {
		return
	}
	var body searchRespBody
	err = json.Unmarshal(resp, &body)
	if err != nil {
		logger.Error.Println(err)
		logger.Error.Println(string(resp))
		return
	}
	for _, resource := range body.Resources {
		ids = append(ids, resource.Res.ID)
	}

	return
}

func (o *ipGenerator) getTotalNumber(keyword string) (total int, err error)  {
	url := "https://dalipan.com/api/search"
	method := "GET"
	header := map[string]string{
		"Content-Type": "application/json",
	}
	getParams := map[string]string{
		"kw": keyword,
		"page": "1",
		"ip": o.next(),
	}
	resp, err := httpDo(url, header, method, getParams, "")
	if err != nil {
		// skip if error
		return 0, nil
	}
	if len(resp) == 0 {
		total = 0
		return
	}
	var body searchRespBody
	err = json.Unmarshal(resp, &body)
	if err != nil {
		logger.Error.Println(err)
		logger.Error.Println(string(resp))
		return
	}
	total = body.Total
	return
}

func httpDo(url string, header map[string]string, httpMethod string, httpGetParams map[string]string, httpPostBody string) (body []byte, err error) {
	request, err := httpclient.NewRequest(url, header, httpMethod, httpGetParams, httpPostBody)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	resp, err := httpclient.NewClient().Do(request.WithDefaultUserAgent())
	if err != nil {
		logger.Error.Println(err)
		return
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	//logger.Info.Println(string(body))
	if string(body) == "porn" || string(body) == "privacy" || string(body) == "politics" || string(body) == "contraband" || string(body) == "limited" || string(body) == "涉嫌隐私" {
		// skip these errors
		logger.Info.Println(string(body))
		return nil, nil
	} else if string(body) == "" {
		// error
		err = fmt.Errorf("返回内容为空")
		logger.Error.Println(err)
		return
	} else if resp.StatusCode == 502 {
		err = fmt.Errorf("返回状态码为502")
		logger.Error.Println(err)
		return
	} else if resp.StatusCode == 504 {
		err = fmt.Errorf("返回状态码为504")
		logger.Error.Println(err)
		return
	} else {
		return
	}
}

func httpDoWithRetry(url string, header map[string]string, httpMethod string, httpGetParams map[string]string, httpPostBody string) (body []byte, err error) {
	var retry = 1
	for {
		body, err = httpDo(url, header, httpMethod, httpGetParams, httpPostBody)
		if err != nil {
			waitTime := time.Duration(3 * retry) * time.Minute
			time.Sleep(waitTime)
		} else {
			return
		}
		retry ++
		if retry >= maxTry {
			err = fmt.Errorf("重试多次后还是无法获取数据, %s", url)
			logger.Error.Println(err)
			logger.Error.Println(string(body))
			return nil, err
		}
	}
	return
}

func (o *ipGenerator) getMetaData(id string) (metaData *common.BDP, err error) {
	metaData = new(common.BDP)
	url := "https://dalipan.com/api/detail"
	header := map[string]string{
		"Content-Type": "application/json",
	}
	method := "GET"
	getParams := map[string]string{
		"id": id,
		"size": "15",
		"ip": o.next(),
	}
	resp, err := httpDo(url, header, method, getParams, "")
	if err != nil {
		return
	}
	var raw rawMetaData
	err = json.Unmarshal(resp, &raw)
	if err != nil {
		logger.Error.Println(err)
		logger.Error.Println(string(resp))
		return
	}
	//logger.Info.Printf("%+v", raw)
	metaData.Url = raw.URL
	metaData.Title = raw.FileName
	metaData.Ext = raw.Ext
	metaData.CTime = strings.Split(raw.CTime, " ")[0]
	metaData.Size = strconv.FormatInt(raw.Size, 10)
	metaData.HasPwd = raw.HasPwd
	metaData.Password = raw.Pwd
	metaData.Category = util.ExtToCategory(metaData.Ext)
	metaData.Resource = "百度网盘"
	return
}

func (o *ipGenerator) getTotalPageNumber(total int) (totalPageNumber int) {
	pageSize := 30
	totalPageNumber = 0
	maxPageNumber := 300
	// 页面最多显示300页
	if total > (pageSize * maxPageNumber) {
		totalPageNumber = maxPageNumber
	} else {
		if total%pageSize == 0 {
			totalPageNumber = total/pageSize
		} else {
			totalPageNumber = total/pageSize + 1
		}
	}
	return
}

func (o *silverBullet) getMetaDataPerPagePerKeyword(keyword string, pageNumber int, metaDataCh chan *common.BDP, doneCh chan bool, errCh chan error) {
	ids, err := o.ipGenerator.getIDs(keyword, pageNumber)
	if err != nil {
		errCh <- err
		return
	}
	wg := sync.WaitGroup{}
	for _, id := range ids {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			metaData, err := o.ipGenerator.getMetaData(id)
			if err != nil {
				return
			}
			// check if metaData already existed
			exist, err := o.mysqlEngine.hasExisted(metaData)
			if err != nil {
				return
			}
			if exist == false {
				metaDataCh <- metaData
				logger.Info.Printf("%s %+v", keyword, metaData)
			} else {
				logger.Info.Printf("已经存在，%s %+v", keyword, metaData)
			}
		}(id)
	}
	wg.Wait()
	doneCh <- true
	return
}

func (o *silverBullet) getMetaDataPerKeyword(keyword string) (err error) {
	total, err := o.ipGenerator.getTotalNumber(keyword)
	if err != nil {
		return
	}
	logger.Info.Printf("共搜索到 %d 条结果", total)
	if total == 0 {
		return
	}
	var metaDatas []*common.BDP
	batchSize := 150
	metaDataCh := make(chan *common.BDP)
	doneCh := make(chan bool)
	errCh := make(chan error)
	totalPageNumber := o.ipGenerator.getTotalPageNumber(total)
	for i:=1;i<=totalPageNumber;i++ {
		go func(i int) {
			o.getMetaDataPerPagePerKeyword(keyword, i, metaDataCh, doneCh, errCh)
		}(i)
	}
	// close doneCh when every page is finished
	go func() {
		doneNum := 0
		for {
			select {
			case <- doneCh:
				//logger.Info.Println("one page done")
				doneNum ++
				if doneNum == totalPageNumber {
					close(metaDataCh)
					close(doneCh)
					close(errCh)
					return
				}
			case <- errCh:
				//logger.Error.Println("one page error")
				doneNum ++
				if doneNum == totalPageNumber {
					close(metaDataCh)
					close(doneCh)
					close(errCh)
					return
				}
			}
		}
	}()
	// write metaData to mysql in batch
	num := 0
	for metaData := range metaDataCh {
		num ++
		//logger.Info.Println(num)
		metaDatas = append(metaDatas, metaData)
		if len(metaDatas) == batchSize {
			logger.Info.Println("write one batch to mysql")
			err = o.mysqlEngine.writeToMysql(metaDatas)
			if err != nil {
				return
			}
			metaDatas = []*common.BDP{}
		}
	}
	// handle left metaDatas
	if len(metaDatas) > 0 {
		logger.Info.Println("writing to mysql")
		err = o.mysqlEngine.writeToMysql(metaDatas)
		if err != nil {
			return
		}
	}

	return
}

func initMysqlEngine() (engine *MysqlEngine, err error) {
	mysqlEngine, err := db.NewMysqlEngineForConfig(mysqlConfig)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	engine = &MysqlEngine{
		mysqlEngine,
	}
	return
}

func (engine *MysqlEngine) hasExisted(metaData *common.BDP) (exist bool, err error) {
	// check if url already exist
	exist, err = engine.Exist(&db.Pan{
		Url: metaData.Url,
	})
	if err != nil {
		logger.Error.Println(err)
		return
	}
	return
}

func (engine *MysqlEngine) writeToMysql(metaDatas []*common.BDP) (err error) {
	var pans []db.Pan
	for _, metaData := range metaDatas {
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
		pan.IsInvalid = false
		pans = append(pans, pan)
	}
	_, err = engine.Insert(&pans)
	if err != nil {
		logger.Error.Println(err)
		matched, err := regexp.MatchString("Duplicate entry", err.Error())
		if err != nil {
			logger.Error.Println(err)
			return err
		}
		if matched == true {
			return nil
		} else {
			return err
		}
	}
	return
}

func run() (err error) {
	// get all keywords from txt
	filePath := "keyword.txt"
	file, err := os.Open(filePath)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	defer file.Close()
	silverBullet, err := construct()
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		keyword := scanner.Text()
		// remove whitespaces at beginning and end
		keyword = strings.TrimSpace(keyword)
		if len(keyword) == 0 {
			continue
		}
		logger.Info.Printf("爬取关键词 %s", keyword)
		err := silverBullet.getMetaDataPerKeyword(keyword)
		if err != nil {
			return err
		}
	}
	return
}

func constructIPGenerator() *ipGenerator {
	return &ipGenerator{
		a: 58,
		b: 247,
		c: 1,
		d: 1,
	}
}

func constructMysqlEngine() (*MysqlEngine, error) {
	// initialize mysql engine
	engine, err := initMysqlEngine()
	return engine, err
}

func construct() (*silverBullet, error) {
	engine, err := constructMysqlEngine()
	if err != nil {
		return nil, err
	}
	return &silverBullet{
		mysqlEngine: engine,
		ipGenerator: constructIPGenerator(),
	}, nil
}

func (o *ipGenerator) next() (ip string) {
	if o.c == 255 {
		o.c = 1
		o.d = 1
	}
	if o.d == 255 {
		o.d = 1
		o.c++
	} else {
		o.d++
	}
	ipSlice := []string{
		strconv.Itoa(o.a),
		strconv.Itoa(o.b),
		strconv.Itoa(o.c),
		strconv.Itoa(o.d),
	}
	ip = strings.Join(ipSlice, ".")
	return
}