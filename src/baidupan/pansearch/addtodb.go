package pansearch

import (
	"baidupan/pansearch/collect"
	_ "github.com/go-sql-driver/mysql"
	"util/db"
	"util/coreconfig"
	logger "util/log"
	"fmt"
)

func AddToDB(bdps collect.BDPS) error {
	db, err := db.DBConnection()
	if err != nil {
		return err
	}
	defer db.Close()
	// Begin will get a connection with mysql before insert, so this method has high effect
	tx, _ := db.Begin()
	for _, bdp := range bdps {
		sqlStr := fmt.Sprintf("INSERT IGNORE INTO %s (url, title, ext, ctime, size, haspwd, password, category, resource) values ('%s', '%s', '%s', '%s', '%s', %t, '%s', '%s', '%s')", coreconfig.CC.Mysql.TBaidupan, bdp.Url, bdp.Title, bdp.Ext, bdp.CTime, bdp.Size, bdp.HasPwd, bdp.Password, bdp.Category, bdp.Resource)
		//logger.Info.Println(sqlStr)
		_, err := tx.Exec(sqlStr)
		if err != nil {
			logger.Error.Println(err.Error())
		}
		//id, _ := result.LastInsertId()
		//logger.Error.Println(id)
		//ids = append(ids, id)
	}
	tx.Commit()
	return nil
}
