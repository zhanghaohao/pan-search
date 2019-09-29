package coreconfig

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	logger "util/log"
)

var CC *CoreConfig

type CoreConfig struct {
	Mysql 		MysqlConfig			`yaml:"mysql"`
	Redis 		RedisConfig			`yaml:"redis"`
}

type MysqlConfig struct {
	Host 		string				`yaml:"host"`
	Port 		string				`yaml:"port"`
	Database 	string				`yaml:"database"`
	User 		string				`yaml:"user"`
	Password 	string				`yaml:"password"`
	TBaidupan 	string				`yaml:"tablebaidupan"`
	TMagnet 	string				`yaml:"tablemagnet"`
}

type RedisConfig struct {
	Host 		string				`yaml:"host"`
	Port 		string				`yaml:"port"`
	Password 	string				`yaml:"password"`
	DB 			string				`yaml:"db"`
}

func YamlParser() (*CoreConfig, error) {
	config, err := ioutil.ReadFile("src/util/coreconfig/config.yaml")
	if err != nil {
		logger.Error.Println("Cannot find core config file ", err.Error())
		return nil, err
	}
	var cc CoreConfig
	err = yaml.Unmarshal(config, &cc)
	if err != nil {
		logger.Error.Println("Cannot deserialize core config file to struct")
		return nil, err
	}
	CC = &cc
	return &cc, nil
}

