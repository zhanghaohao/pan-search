package search

import (
	_ "github.com/go-sql-driver/mysql"
	"plugin/db"
	"util/logger"
	"crawlengine/resource/common"
)

func (s *search) RecordPanMetaData(bdps []common.BDP) (err error) {
	// todo: check if metaData already exist. donot insert if already exist
	var metaDatas []*db.Pan
			for _, bdp := range bdps {
			metaData := &db.Pan{
			Url:   bdp.Url,
			Title: bdp.Title,
			Ext:   bdp.Ext,
			CTime: bdp.CTime,
			Size:  bdp.Size,
			HasPwd: bdp.HasPwd,
			Password: bdp.Password,
			Category: bdp.Category,
			Resource: bdp.Resource,
		}
		metaDatas = append(metaDatas, metaData)
	}
	_, err = s.cacheEngine.GetMysqlEngine().Insert(metaDatas)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	return
}

func (s *search) RecordKeyword(keyword, searchType, sourceIP, searchTime string) (err error) {
	record := &db.Keyword{
		Keyword: keyword,
		SearchType: searchType,
		SourceIP: sourceIP,
		SearchTime: searchTime,
	}
	_, err = s.cacheEngine.GetMysqlEngine().Insert(record)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	return
}