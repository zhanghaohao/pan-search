package magnetsearch

import (
	"util/db"
	"fmt"
	"util/coreconfig"
	logger "util/log"
)

func AddToDB(mags []MAG) error {
	db, err := db.DBConnection()
	if err != nil {
		return err
	}
	defer db.Close()
	// Begin will get a connection with mysql before insert, so this method has high effect
	tx, _ := db.Begin()
	for _, mag := range mags {
		sqlStr := fmt.Sprintf("INSERT IGNORE INTO %s (magnet, title, size, date, hot) values ('%s', '%s', '%s', '%s', '%s')", coreconfig.CC.Mysql.TMagnet, mag.Magnet, mag.Title, mag.Size, mag.Date, mag.Hot)
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
