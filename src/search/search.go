package search

import (
	"util/logger"
	"net/http"
	"time"
	"strings"
	"html/template"
	"plugin/db"
	"crawlengine/resource/common"
	"crawlengine"
	"util/encrypt"
	"util"
	"cacheengine"
	"strconv"
	"plugin/config"
	"crawlengine/process"
)

const (
	pageSize int = 20
)

type search struct {
	cacheEngine 			cacheengine.EngineFactory
}

type SearchResult struct {
	Keyword 					string
	TotalCount 					int
	MetaDatas 					[]util.MetaData
	CurrentPageNumber			int
	Pages 						[]Page
}

type Page struct {
	Text 						string
	Class 						string
	Href 						string
}

type SearchFactory interface {
	Index(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
	SearchMetaData(w http.ResponseWriter, r *http.Request)
	VerifyPanURL(w http.ResponseWriter, r *http.Request)
}

func NewForConfig(config config.ConfigFactory) (SearchFactory, error) {
	searchFactory, err := cacheengine.NewForConfig(config.GetSphinxConfig(), config.GetMysqlConfig())
	if err != nil {
		return nil, err
	}
	return &search{
		cacheEngine: searchFactory,
	}, nil
}

func (s *search) Index(w http.ResponseWriter, r *http.Request)  {
	// check if url is valid
	t404, err := template.ParseFiles(TemplateNotFound1, TemplateHeader, TemplateFooter, TemplateNavi, TemplatePanSearcher, TemplateHotKeys, TemplateMagnetSearcher)
	if err != nil {
		logger.Error.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Page not found"))
		return
	}
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		t404.ExecuteTemplate(w, "404-1", nil)
		return
	}
	// render index page
	t, err := template.ParseFiles(TemplatePanIndex, TemplateHeader, TemplateFooter, TemplateNavi, TemplatePanSearcher, TemplateHotKeys, TemplateMagnetSearcher)
	if err != nil {
		logger.Error.Println("Canot find panindex.html, ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Page not found"))
		return
	}
	t.ExecuteTemplate(w, "panindex", nil)
}

func (s *search) Search(w http.ResponseWriter, r *http.Request)  {
	t404, err := template.ParseFiles(TemplateNotFound, TemplateHeader, TemplateFooter, TemplateNavi, TemplatePanSearcher, TemplateHotKeys, TemplateMagnetSearcher)
	if err != nil {
		logger.Error.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Page not found"))
		return
	}
	// get parameters
	keyword := r.URL.Query().Get("keyword")
	currentPageNumStr := r.URL.Query().Get("page")
	category := r.URL.Query().Get("category")
	if len(keyword) != 0 {
		searchType := "pan"
		sourceIP := r.Header.Get("X-Real-IP")
		searchTime := time.Now().Format("2006-01-02 15:04:05")
		err := s.RecordKeyword(keyword, searchType, sourceIP, searchTime)
		if err != nil {
			logger.Error.Println(err)
		}
	}
	var searchResult = new(SearchResult)
	searchResult.Keyword = keyword
	if len(currentPageNumStr) == 0 {
		searchResult.CurrentPageNumber = 1
	} else {
		searchResult.CurrentPageNumber, err = strconv.Atoi(currentPageNumStr)
		if err != nil {
			logger.Error.Println(err)
			t404.ExecuteTemplate(w, "404", nil)
			return
		}
	}
	// get meta data
	filter := &cacheengine.Filter{
		Category: category,
	}
	totalItems, metaDatas, err := s.cacheSearch(keyword, pageSize, searchResult.CurrentPageNumber, filter)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		t404.ExecuteTemplate(w, "404", nil)
		return
	}
	searchResult.MetaDatas = metaDatas
	// get total count
	if totalItems == 0 {
		t404.ExecuteTemplate(w, "404", nil)
		return
	}
	searchResult.TotalCount = totalItems
	/*
	render pagination
	 */
	pages := renderPagination(r, totalItems, searchResult.CurrentPageNumber)
	searchResult.Pages = pages
	/*
	render template
	 */
	t, err := template.ParseFiles(TemplatePanSearch, TemplateHeader, TemplateFooter, TemplateNavi, TemplatePanSearcher, TemplateHotKeys, TemplateMagnetSearcher)
	if err != nil {
		logger.Error.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Page not found"))
		return
	}
	t.ExecuteTemplate(w, "search", searchResult)
	return
}

func updateURLGetParam(r *http.Request, key string, value string) (newURL string) {
	path := r.URL.Path
	query := r.URL.Query()
	v := query.Get(key)
	if len(v) == 0 {
		query.Add(key, value)
	} else {
		query.Set(key, value)
	}
	newURL = path + "?" + query.Encode()
	return
}

func renderPagination(r *http.Request, totalItems int, currentPageNum int) (pages []Page) {
	// the page number around current page
	interval := 4
	halfInterval := interval/2
	prevPage := Page{
		Text: "上一页",
		Class: "prev",
	}
	nextPage := Page{
		Text: "下一页",
		Class: "next",
	}
	skipPage := Page{
		Text: "...",
	}
	firstPage := Page{
		Text: "1",
		Href: updateURLGetParam(r, "page", "1"),
	}
	pageCount := totalItems/pageSize
	if totalItems % pageSize != 0 {
		pageCount ++
	}
	lastPage := Page{
		Text: strconv.Itoa(pageCount),
		Href: updateURLGetParam(r, "page", strconv.Itoa(pageCount)),
	}
	var start, stop int
	var startFlag, stopFlag int
	prevPoint := currentPageNum - halfInterval
	if prevPoint == -1 {
		// current page is first page
		start = 1
		startFlag = 1
	} else if prevPoint > 2 {
		start = prevPoint
		startFlag = 2
	} else {
		start = 1
		startFlag = 3
	}
	nextPoint := currentPageNum + halfInterval
	if nextPoint == pageCount + 2 {
		// current page is last page
		stop = pageCount
		stopFlag = 1
	} else if nextPoint < pageCount - 2 {
		stop = nextPoint
		stopFlag = 2
	} else {
		stop = pageCount
		stopFlag = 3
	}
	// generate pages
	switch startFlag {
	case 2:
		prevPage.Href = updateURLGetParam(r, "page", strconv.Itoa(currentPageNum-1))
		pages = append(pages, prevPage)
		pages = append(pages, firstPage)
		pages = append(pages, skipPage)
	case 3:
		prevPage.Href = updateURLGetParam(r, "page", strconv.Itoa(currentPageNum-1))
		pages = append(pages, prevPage)
	}
	// append interval
	for i:=start;i<=stop;i++ {
		var class string
		if i == currentPageNum {
			class = "current"
		} else {
			class = ""
		}
		page := Page{
			Text: strconv.Itoa(i),
			Href: updateURLGetParam(r, "page", strconv.Itoa(i)),
			Class: class,
		}
		pages = append(pages, page)
	}
	switch stopFlag {
	case 2:
		pages = append(pages, skipPage)
		pages = append(pages, lastPage)
		nextPage.Href = updateURLGetParam(r, "page", strconv.Itoa(currentPageNum+1))
		pages = append(pages, nextPage)
	case 3:
		nextPage.Href = updateURLGetParam(r, "page", strconv.Itoa(currentPageNum+1))
		pages = append(pages, nextPage)
	}
	return
}

func (s *search) cacheSearch(keyword string, pageSize int, currentPageNum int, filter *cacheengine.Filter) (totalItems int, metaDatas []util.MetaData, err error) {
	// search from riot
	totalItems, metaDatas, err = s.cacheEngine.Search(keyword, pageSize, currentPageNum, filter)
	if err != nil {
		return
	}
	return
}

func (s *search) deepSearch(keyword string) (resources common.BDPS, err error) {
	return s.crawlSearchEngine(keyword)
}

func (s *search) crawlSearchEngine(keyword string) (cookedBdps common.BDPS, err error) {
	cookedBdps, err = crawlengine.CrawlSearch(keyword)
	if err != nil {
		return
	}
	err = s.RecordPanMetaData(cookedBdps)
	if err != nil {
		return
	}
	return
}

func (s *search) SearchMetaData(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(TemplatePanMetaData, TemplateHeader, TemplateFooter, TemplateNavi, TemplatePanSearcher, TemplateMagnetSearcher)
	t404, err := template.ParseFiles(TemplateNotFound, TemplateHeader, TemplateFooter, TemplateNavi, TemplatePanSearcher, TemplateHotKeys, TemplateMagnetSearcher)
	if err != nil {
		logger.Error.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Page not found"))
		return
	}
	id := strings.TrimPrefix(r.URL.Path, "/pan/search/resource/")
	// decrypt id
	id, err = encrypt.DecryptID(id)
	if err != nil {
		t404.ExecuteTemplate(w, "404", nil)
		return
	}
	var pan db.Pan
	has, err := s.cacheEngine.GetMysqlEngine().Where("id = ?", id).Get(&pan)
	if err != nil {
		logger.Error.Println(err)
	}
	if has == false || err != nil {
		t404.ExecuteTemplate(w, "404", nil)
		return
	}
	t.ExecuteTemplate(w, "panmetadata", pan)
	return
}

func (s *search) VerifyPanURL(w http.ResponseWriter, r *http.Request)  {
	panURL := r.PostFormValue("panURL")
	if len(panURL) == 0 {
		w.Write([]byte("invalid"))
		return
	}
	isValid, err := process.VerifyBaiduPanURL(panURL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error"))
		return
	}
	if isValid == true {
		w.Write([]byte("valid"))
		return
	} else {
		w.Write([]byte("invalid"))
	}
	// set invalid flag
	pan := &db.Pan{
		Url: panURL,
	}
	_, err = s.cacheEngine.GetMysqlEngine().Cols("isinvalid").Update(&db.Pan{IsInvalid: true}, pan)
	if err != nil {
		logger.Error.Println(err)
		w.Write([]byte("valid"))
	}
}
