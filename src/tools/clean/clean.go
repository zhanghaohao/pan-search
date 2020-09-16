package clean

import (
	"plugin/db"
	"util/logger"
	"sync"
	"crawlengine/process"
	"plugin/config"
)

type idurl struct {
	ID 				int64
	URL 			string
}

type mysql struct {
	mysqlEngine 				*db.MysqlEngine
}

/*
clean invalid pan data from mysql
*/
func Clean() (err error) {
	configFactory, err := config.ReadConfigFile("../../plugin/config/config.yaml")
	if err != nil {
		logger.Error.Println(err)
		return
	}
	mysqlEngine, err := db.NewMysqlEngineForConfig(configFactory.GetMysqlConfig())
	if err != nil {
		logger.Error.Println(err)
		return
	}
	var records []db.Pan
	err = mysqlEngine.Cols("id", "url").Find(&records)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	var idurls []idurl
	for _, record := range records {
		var idurl idurl
		idurl.ID = record.Id
		idurl.URL = record.Url
		idurls = append(idurls, idurl)
	}
	mysql := &mysql{
		mysqlEngine: mysqlEngine,
	}
	sortedidurls := sortByID(idurls)
	logger.Info.Println(sortedidurls)
	var batchidurl []idurl
	for _, idurl := range sortedidurls {
		logger.Info.Printf("validating %d %s", idurl.ID, idurl.URL)
		batchidurl = append(batchidurl, idurl)
		if len(batchidurl) == 50 {
			wg := sync.WaitGroup{}
			for _, idurl := range batchidurl {
				wg.Add(1)
				go func(url string) {
					defer wg.Done()
					mysql.handleURL(url)
				}(idurl.URL)
			}
			wg.Wait()
			batchidurl = nil
			continue
		}
	}
	return
}

func (c *mysql) handleURL(url string) (err error) {
	isValid, err := process.VerifyBaiduPanURL(url)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	if isValid == false {
		var pan db.Pan
		_, err := c.mysqlEngine.Where("url = ?", url).Delete(&pan)
		if err != nil {
			logger.Error.Println(err)
			return err
		}
		logger.Info.Printf("invalid url %s has been deleted", url)
	}
	return
}

func sortByID(raw []idurl) (cooked []idurl) {
	// sort by id from small to big
	cooked = raw
	for j:=0;j<len(cooked)-2;j++ {
		for i:=0;i<len(cooked)-1-j;i++ {
			if cooked[i].ID > cooked[i+1].ID {
				tmp := cooked[i]
				cooked[i] = cooked[i+1]
				cooked[i+1] = tmp
			}
		}
	}
	return cooked
}