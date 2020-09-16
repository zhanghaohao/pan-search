package route

import (
	"net/http"
	"search"
	"plugin/weixin"
	"plugin/config"
	"util/logger"
	"plugin/db"
)

func initData() (searchFactory search.SearchFactory, weixinFactory weixin.WeiXinFactory, err error) {
	configFactory, err := config.LoadConfig()
	if err != nil {
		return
	}
	mysqlEngine, err := db.NewMysqlEngineForConfig(configFactory.GetMysqlConfig())
	if err != nil {
		logger.Error.Println(err)
		return
	}
	// sync table structure
	err = mysqlEngine.Sync()
	if err != nil {
		return
	}
	redisPool, err := db.NewRedisPoolForConfig(configFactory.GetRedisConfig())
	if err != nil {
		logger.Error.Println(err)
		return
	}
	searchFactory, err = search.NewForConfig(configFactory)
	if err != nil {
		return
	}
	weixinFactory = weixin.New(mysqlEngine, redisPool, configFactory.GetWeiXinConfig())
	return
}

func RegisterRoutes(mux *http.ServeMux) (*http.ServeMux, error) {
	// get all config
	searchFactory, weixinFactory, err := initData()
	if err != nil {
		return nil, err
	}
	mux.HandleFunc("/", searchFactory.Index)
	//mux.Handle("/static/", http.FileServer(http.Dir("./src")))
	mux.HandleFunc("/pan/search", searchFactory.Search)
	mux.HandleFunc("/pan/search/resource/", searchFactory.SearchMetaData)
	mux.HandleFunc("/weixin/getticket", weixinFactory.GetTmpTicket)
	mux.HandleFunc("/weixin/buildwebsocket", weixinFactory.WebSocketHandler)
	mux.HandleFunc("/weixin/receiveevent", weixinFactory.EventReceiver)
	mux.HandleFunc("/pan/verifypanurl", searchFactory.VerifyPanURL)
	return mux, nil
}
