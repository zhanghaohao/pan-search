package cacheengine

import (
	"util"
	"github.com/yunge/sphinx"
	"util/logger"
	"fmt"
	"strconv"
	"plugin/db"
	"xorm.io/builder"
	"strings"
)

const (
	maxItems = 10000000
)

type SphinxConfig struct {
	Host 					string 					`yaml:"host"`
	Port 					int 					`yaml:"port"`
}

type Filter struct {
	Category 				string
}

type engine struct {
	sphinxConfig 			*SphinxConfig
	mysqlEngine 			*db.MysqlEngine
}

type EngineFactory interface {
	GetMysqlEngine() *db.MysqlEngine
	Search(keyword string, pageSize int, currentPageNum int, filter *Filter) (totalItems int, metaDatas []util.MetaData, err error)
}

func (e *engine) GetMysqlEngine() *db.MysqlEngine {
	return e.mysqlEngine
}

func NewForConfig(sphinxConfig *SphinxConfig, mysqlConfig *db.MysqlConfig) (EngineFactory, error) {
	mysqlEngine, err := db.NewMysqlEngineForConfig(mysqlConfig)
	if err != nil {
		return nil, err
	}
	engine := &engine{
		sphinxConfig: sphinxConfig,
		mysqlEngine: mysqlEngine,
	}
	return engine, nil
}

func newClient(sphinxConfig *SphinxConfig) (sc *sphinx.Client) {
	opts := &sphinx.Options{
		Host:    sphinxConfig.Host,
		Port:    sphinxConfig.Port,
		Timeout: 5000,
		MaxMatches: maxItems,
		Offset: 0,
		Limit: 10,
	}
	sc = sphinx.NewClient(opts)
	return
}

func (e *engine) Search(keyword string, pageSize int, currentPageNum int, filter *Filter) (totalItems int, metaDatas []util.MetaData, err error) {
	// get client
	sc := newClient(e.sphinxConfig)
	/*
	search cached docs
	  */
	if pageSize < 0 {
		err = fmt.Errorf("invalid pageSize %d", pageSize)
		logger.Error.Println(err)
		return
	}
	if currentPageNum < 1 {
		err = fmt.Errorf("invalid currentpageNum %d", currentPageNum)
		logger.Error.Println(err)
		return
	}

	// set offset and limit
	offset := pageSize * (currentPageNum - 1)
	//sc := e.sphinxClient
	sc.Offset = offset
	sc.Limit = pageSize
	index := "pan"
	sc.MatchMode = sphinx.SPH_MATCH_EXTENDED2
	// query
	var query = ""
	// construct matcher segment for title
	var titleMatcher = ""
	var keywords []string
	rawKeywords := strings.Split(keyword, " ")
	for _, e := range rawKeywords {
		if len(e) != 0 {
			keywords = append(keywords, e)
		}
	}
	for i, e := range keywords {
		if i == 0 {
			titleMatcher += " @title (" + e + ")"
		} else {
			titleMatcher += " |" + " @title (" + e + ")"
		}
	}
	titleMatcher = "(" + titleMatcher + ")"
	//logger.Info.Println(titleMatcher)

	if len(filter.Category) > 0 && filter.Category != "all" {
		query = fmt.Sprintf("@category %s & %s", filter.Category, titleMatcher)
	} else {
		query = titleMatcher
	}
	logger.Info.Println(query)
	// ensure connection had been closed
	if sc.GetConn() != nil {
		sc.Close()
	}
	res, err := sc.Query(query, index, "")
	if err != nil {
		logger.Error.Println(err)
		return
	}
	defer func() {
		if sc.GetConn() != nil {
			sc.Close()
		}
	}()
	totalItems = res.Total
	if totalItems == 0 {
		err = fmt.Errorf("搜索不到任何结果！")
		logger.Info.Println(err)
		return
	}
	// verify currentPageNum is not out of range
	pageCount := totalItems/pageSize
	if totalItems%pageSize != 0 {
		pageCount ++
	}
	if currentPageNum > pageCount {
		err = fmt.Errorf("invalid currentpageNum %d", currentPageNum)
		logger.Error.Println(err)
		return
	}
	//logger.Info.Printf("%+v", res)
	var ids []string
	for _, item := range res.Matches {
		//logger.Info.Printf("%+v", item)
		ids = append(ids, strconv.FormatUint(item.DocId, 10))
	}
	if len(ids) == 0 {
		return
	}
	// get metadatas from mysql by ids
	var pans []db.Pan
	err = e.mysqlEngine.Where(builder.In("id", ids)).Find(&pans)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	metaDatas, err = handle(pans, keyword)
	return
}