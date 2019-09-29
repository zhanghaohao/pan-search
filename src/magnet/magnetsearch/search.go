package magnetsearch

import (
	"net/http"
	"strings"
	"html/template"
	logger "util/log"
	"encoding/json"
	"fmt"
	"time"
)

func Search(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("src/template/magnetsearch.html", "src/template/public/header.html", "src/template/public/footer.html", "src/template/public/navi.html", "src/template/public/pansearcher.html", "src/template/public/magnetsearcher.html")
	if err != nil {
		logger.Error.Println("Canot find magnetsearch.html, ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Page not found"))
		return
	}
	keyword := strings.TrimPrefix(r.URL.Path, "/magnet/search/")
	t.ExecuteTemplate(w, "magnetsearch", keyword)
	return
}

func GetMagnets(w http.ResponseWriter, r *http.Request)  {
	keyword := r.URL.Query().Get("keyword")
	w.Header().Set("Content-Type", "application/json")
	// convert nil to json
	null, _ := json.Marshal(nil)
	magnets, err := SearchEngine(keyword)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if magnets == nil {
		w.WriteHeader(http.StatusOK)
		w.Write(null)
		return
	}
	result, err := json.Marshal(magnets)
	if err != nil {
		logger.Error.Println("Serialize magnets struct to json []byte failed\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Serialize magnets struct to json []byte failed"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)
	return
}

func SearchEngine(keyword string) ([]MAG, error) {
	logger.Info.Println("Magnet search keyword is: ", keyword)
	logger.Info.Println("Progressing to Crawl Stage ...")
	startCrawlTime := time.Now()
	magnets, err := Crawl(keyword)
	stopCrawlTime := time.Now()
	costCrawlTime := stopCrawlTime.Sub(startCrawlTime)
	logger.Info.Println("Crawl Stage cost: ", costCrawlTime)
	if err != nil {
		logger.Error.Println("Crawl magnets failed\n", err.Error())
		return nil, fmt.Errorf("crawl magnets failed")
	}
	logger.Info.Println("Progressing to AddToDB Stage ...")
	err = AddToDB(magnets)
	if err != nil {
		logger.Error.Println(err.Error())
		return nil, err
	}
	return magnets, nil
}

