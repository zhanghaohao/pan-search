package db

import (
	 "util/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type MysqlEngine struct {
	*xorm.Engine
}

type Pan struct {
	Id 				int64 			`xorm:"not null pk autoincr int 'id'"`
	Url 			string  		`xorm:"not null varchar(128) unique 'url'"`
	Title 			string 			`xorm:"not null varchar(512) 'title'"`
	Ext 			string 			`xorm:"varchar(8) 'ext'"`
	CTime 			string 			`xorm:"date 'ctime'"`
	Size 			string 			`xorm:"default 0 bigint 'size'"`
	HasPwd 			bool 			`xorm:"bool 'haspwd'"`
	Password 		string 			`xorm:"varchar(8) 'password'"`
	Category 		string 			`xorm:"varchar(128) 'category'"`
	Resource 		string			`xorm:"varchar(128) resource"`
	IsInvalid		bool 			`xorm:"default false bool 'isinvalid'"`
}

type Keyword struct {
	Id 				int64 			`xorm:"not null pk autoincr int 'id'"`
	Keyword 		string 			`xorm:"not null varchar(128) 'keyword'"`
	SearchType 		string 			`xorm:"varchar(32) 'searchtype'"`
	SourceIP 		string 			`xorm:"varchar(128) 'sourceip'"`
	SearchTime 		string 			`xorm:"varchar(64) 'searchtime'"`
}

type Qrcode struct {
	Id 				int64 			`xorm:"not null pk autoincr int 'id'"`
	TmpTicket		string 			`xorm:"varchar(512) 'tmpTicket'"`
	ValidMinutes	int 			`xorm:"int(32) 'validMinutes'"`
	IsScanned 		bool 			`xorm:"bool 'isScanned'"`
	IsExpired 		bool 			`xorm:"bool 'isExpired'"`
	IsLocked 		bool 			`xorm:"bool 'isLocked'"`
	CreateAt 		string 			`xorm:"varchar(32) 'createAt'"`
	FromUser 		string 			`xorm:"varchar(128) 'fromUser'"`
	ScannedAt 		string 			`xorm:"varchar(32) 'scannedAt'"`
}

type MysqlConfig struct {
	Host 		string				`yaml:"host"`
	Port 		string				`yaml:"port"`
	Database 	string				`yaml:"database"`
	User 		string				`yaml:"user"`
	Password 	string				`yaml:"password"`
}

type MysqlEngineInterface interface {
	Sync() (err error)
}

func NewMysqlEngineForConfig(config *MysqlConfig) (mysqlEngine *MysqlEngine, err error) {
	mysqlEngine = new(MysqlEngine)
	conStr := config.User + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.Database + "?charset=utf8"
	engine, err := xorm.NewEngine("mysql", conStr)
	if err != nil {
		logger.Error.Println(err)
		return 
	}
	mysqlEngine.Engine = engine
	return 
}

func (e *MysqlEngine) Sync() (err error) {
	err = e.Sync2(new(Pan))
	if err != nil {
		logger.Error.Println(err)
		return
	}
	err = e.Sync2(new(Keyword))
	if err != nil {
		logger.Error.Println(err)
		return
	}
	err = e.Sync2(new(Qrcode))
	if err != nil {
		logger.Error.Println(err)
		return
	}
	return
}