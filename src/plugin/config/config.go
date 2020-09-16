package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"util/logger"
	"plugin/db"
	"plugin/weixin"
	"flag"
	"fmt"
	"cacheengine"
)

type config struct {
	MysqlConfig 		db.MysqlConfig				`yaml:"mysql"`
	RedisConfig 		db.RedisConfig 				`yaml:"redis"`
	WeiXinConfig		weixin.WeiXinConfig			`yaml:"weixin"`
	SphinxConfig 		cacheengine.SphinxConfig	`yaml:"sphinx"`
}

type ConfigFactory interface {
	GetMysqlConfig() *db.MysqlConfig
	GetRedisConfig() *db.RedisConfig
	GetWeiXinConfig() *weixin.WeiXinConfig
	GetSphinxConfig() *cacheengine.SphinxConfig
}

func (c *config) GetMysqlConfig() *db.MysqlConfig {
	return &c.MysqlConfig
}

func (c *config) GetRedisConfig() *db.RedisConfig {
	return &c.RedisConfig
}

func (c *config) GetWeiXinConfig() *weixin.WeiXinConfig {
	return &c.WeiXinConfig
}

func (c *config) GetSphinxConfig() *cacheengine.SphinxConfig {
	return &c.SphinxConfig
}

func GetConfigPath() (path string, err error) {
	flag.StringVar(&path, "config", "", "path of config file with yaml format")
	flag.Parse()
	if len(path) == 0 {
		flag.Usage()
		err = fmt.Errorf("param for config file is needed")
		return
	}
	return
}

func ReadConfigFile(path string) (configFactory ConfigFactory, err error) {
	configByte, err := ioutil.ReadFile(path)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	var c config
	err = yaml.Unmarshal(configByte, &c)
	if err != nil {
		logger.Error.Println("cannot deserialize config file")
		return
	}
	configFactory = &c
	return
}

func LoadConfig() (configFactory ConfigFactory, err error) {
	path, err := GetConfigPath()
	if err != nil {
		return
	}
	return ReadConfigFile(path)
}