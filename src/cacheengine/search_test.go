package cacheengine

import (
	"testing"
	"util/logger"
	"plugin/db"
)

func Test_query(t *testing.T) {
	mysqlConfig := &db.MysqlConfig{
		Host: "127.0.0.1",
		Port: "3306",
		Database: "",
		User: "",
		Password: "",
	}
	sphinxConfig := &SphinxConfig{
		Host: "localhost",
		Port: 9312,
	}
	cacheEngine, err := NewForConfig(sphinxConfig, mysqlConfig)
	if err != nil {
		t.Error(err)
		return
	}
	totalItems, metaDatas, err := cacheEngine.Search("色戒", 20, 1, &Filter{Category: "all"})
	if err != nil {
		t.Error(err)
		return
	}
	logger.Info.Println(totalItems)
	logger.Info.Printf("%+v", metaDatas)
}
