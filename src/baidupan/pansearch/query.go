package pansearch

import (
	"net/http"
	"util/db"
	"util/coreconfig"
	"fmt"
	logger "util/log"
	"strconv"
	"io/ioutil"
	"encoding/json"
)

type GetIDRespBody struct {
	Url		string			`json:"url"`
}

func GetId(w http.ResponseWriter, r *http.Request) {
	var id int64
	//url := r.URL.Query().Get("url")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	var url GetIDRespBody
	if err := json.Unmarshal(body, &url); err != nil {
		logger.Error.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	db, err := db.DBConnection()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer db.Close()
	//logger.Info.Println(url.Url)
	sqlStr := fmt.Sprintf("select id from %s where url = '%s'", coreconfig.CC.Mysql.TBaidupan, url.Url)
	row := db.QueryRow(sqlStr)
	err = row.Scan(&id)
	if err != nil {
		logger.Error.Println(err.Error(), sqlStr)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	idStr := strconv.FormatInt(id, 10)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(idStr))
	return
}
