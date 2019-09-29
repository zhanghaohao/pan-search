package pansearch

import (
	"baidupan/pansearch/collect/sopan"
	logger "util/log"
	"net/http"
	"time"
	"encoding/json"
	"fmt"
	"strings"
	"html/template"
	"util/db"
	"util/coreconfig"
	"strconv"
	"baidupan/pansearch/collect"
	"util"
	"baidupan/pansearch/collect/soyunpan"
)

func Search(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("src/template/pansearch.html", "src/template/public/header.html", "src/template/public/footer.html", "src/template/public/navi.html", "src/template/public/pansearcher.html", "src/template/public/magnetsearcher.html")
	if err != nil {
		logger.Error.Println("Canot find pansearch.html, ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Page not found"))
		return
	}
	keyword := strings.TrimPrefix(r.URL.Path, "/pan/search/")
	t.ExecuteTemplate(w, "pansearch", keyword)
	return
}

func SearchBdp(w http.ResponseWriter, r *http.Request)  {
	var bdp collect.BDP
	t, err := template.ParseFiles("src/template/pansearchbdp.html", "src/template/public/header.html", "src/template/public/footer.html", "src/template/public/navi.html", "src/template/public/pansearcher.html", "src/template/public/magnetsearcher.html")
	t404, err := template.ParseFiles("src/template/404.html", "src/template/public/header.html", "src/template/public/footer.html", "src/template/public/navi.html", "src/template/public/pansearcher.html", "src/template/public/magnetsearcher.html")
	if err != nil {
		logger.Error.Println("Canot find pansearchbdp.html, ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Page not found"))
		return
	}
	id := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/pan/search/bdp/"), ".html")
	db, err := db.DBConnection()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer db.Close()
	// convert string id to int64
	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	sqlStr := fmt.Sprintf("select * from %s where id = %d", coreconfig.CC.Mysql.TBaidupan, idInt64)
	row := db.QueryRow(sqlStr)
	err = row.Scan(&bdp.Id, &bdp.Url, &bdp.Title, &bdp.Ext, &bdp.CTime, &bdp.Size, &bdp.HasPwd, &bdp.Password, &bdp.Category, &bdp.Resource)
	if err != nil {
		logger.Error.Println(err.Error())
		t404.ExecuteTemplate(w, "404", nil)
		return
	}
	t.ExecuteTemplate(w, "pansearchbdp", bdp)
	return
}

func GetBdps(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("keyword")
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Accept", "application/json")
	category := r.URL.Query().Get("category")
	// convert nil to json
	null, _ := json.Marshal(nil)
	/*
	if query with category, then get data from cache only
	*/
	if category != "" {
		CacheBdps, err := SearchCache(keyword)
		if err != nil {
			logger.Error.Printf("Get value from Redis failed")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(null)
			return
		}
		if CacheBdps != nil {
			filteredCacheBdps := Classify(category, CacheBdps)
			result, err := json.Marshal(filteredCacheBdps)
			if err != nil {
				logger.Error.Println("Serialize BDPS struct to json []byte failed\n", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(null)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(result)
			return
		}
	}

	/*
	pansearch cache first
	*/
	CacheBdps, err := SearchCache(keyword)
	if err != nil {
		logger.Error.Printf("Get value from Redis failed")
	}
	if CacheBdps != nil {
		result, err := json.Marshal(CacheBdps)
		if err != nil {
			logger.Error.Println("Serialize BDPS struct to json []byte failed\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(null)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(result)
		return
	}
	/*
	pansearch db
	*/
	rbdps, err := SearchEngine(keyword)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if rbdps == nil {
		w.WriteHeader(http.StatusOK)
		w.Write(null)
		return
	}
	result, err := json.Marshal(rbdps)
	if err != nil {
		logger.Error.Println("Serialize BDPS struct to json []byte failed\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Serialize BDPS struct to json []byte failed"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)
	return
}

func SearchEngine(keyword string) (collect.BDPS, error) {
	defer util.PrintCostTime(time.Now())
	logger.Info.Println("Baidupan search keyword is: ", keyword)
	// collect
	logger.Info.Println("Progressing to Crawl Stage ...")
	bdps, err := Crawl(keyword)
	// combine
	logger.Info.Println("Progressing to Combine Stage ...")
	bdps = Combine(bdps)
	// deduplicate
	logger.Info.Println("Progressing to DeDup Stage ...")
	bdps = DeDup(bdps)
	// validate
	logger.Info.Println("Progressing to Validate Stage ...")
	bdps, err = Validate(bdps)
	if err != nil {
		logger.Error.Println("Validate BDPS failed\n", err.Error())
		return nil, fmt.Errorf("validate BDPS failed")
	}
	if len(bdps) == 0 || bdps == nil {
		return nil, nil
	}
	// rank bdps according to resource's weight
	bdps = Rank(bdps)
	logger.Info.Println("Progressing to AddToDB Stage ...")
	err = AddToDB(bdps)
	if err != nil {
		return nil, err
	}

	logger.Info.Println("Progressing to AddToCache Stage ...")
	AddToCache(bdps, keyword)
	go sopan.WithPassword(bdps)
	return bdps, nil
}

func Crawl(keyword string) (collect.BDPS, error) {
	defer util.PrintCostTime(time.Now())
	// collect
	channelBdps := make(chan collect.BDPS)
	channelError := make(chan error)
	var errors []error
	var bdps collect.BDPS
	// crawl each resource
	sopan := sopan.CrawlInfo{
		Keyword:      keyword,
		BaseUrl:      "http://www.52sopan.com/search.php",
	}
	soyunpan := soyunpan.CrawlInfo{
		Keyword: keyword,
		BaseUrl: "http://www.soyunpan.com",
	}
	//uzi8 := uzi82.CrawlInfo{
	//	Keyword: keyword,
	//	BaseUrl: "http://uzi8.cn",
	//}
	resourceCount := len([]interface{}{sopan, soyunpan})
	go func() {
		bdps := sopan.Crawl()
		//logger.Info.Println(bdps)
		channelBdps <- bdps
	}()
	go func() {
		bdps := soyunpan.Crawl()
		//logger.Info.Println(bdps)
		channelBdps <- bdps
	}()
	//go func() {
	//	bdps := uzi8.Crawl()
	//	channelBdps <- bdps
	//}()
	//go func() {
	//	dupanpang := dupanpang.Crawldupanpang{
	//		Keyword: keyword,
	//		BaseUrl: "http://www.dupanbang.com",
	//	}
	//	bdps, err := dupanpang.Crawl()
	//	if err != nil {
	//		logger.Error.Println("Crawl dupanpang failed\n", err.Error())
	//		channelError <- err
	//	} else {
	//		//logger.Info.Println(bdps)
	//		channelBdps <- bdps
	//	}
	//}()

	LOOPRESOURCES:
		for {
			select {
			case error := <- channelError:
				errors = append(errors, error)
				resourceCount --
				if resourceCount == 0 && len(bdps) > 0 {
					close(channelError)
					close(channelBdps)
					break LOOPRESOURCES
				} else if resourceCount == 0 && len(bdps) == 0  {
					close(channelError)
					close(channelBdps)
					logger.Error.Println("No resource is available!!!")
					return nil, fmt.Errorf("no resource is available")
				} else {

				}
			case tmpBdps := <- channelBdps:
				bdps = append(bdps, tmpBdps...)
				resourceCount --
				if resourceCount == 0 {
					close(channelError)
					close(channelBdps)
					break LOOPRESOURCES
				}
			}
		}

	if len(bdps) == 0 {
		logger.Info.Println("Crawl nothing from all resources")
		return nil, nil
	}
	return bdps, nil
}