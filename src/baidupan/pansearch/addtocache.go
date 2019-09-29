package pansearch

import (
	"github.com/garyburd/redigo/redis"
	"encoding/json"
	logger "util/log"
	"baidupan/pansearch/collect"
)

const (
	RedisHost = "127.0.0.1"
	RedisPort = "6379"
	RedisPassword = ""
	RedisDB = 0
)

func AddToCache(bdps collect.BDPS, keyword string) error {
	client, err := redis.Dial("tcp", RedisHost + ":" + RedisPort)
	if err != nil {
		logger.Error.Println(err.Error())
		return err
	}
	defer client.Close()
	byteBdps, err := json.Marshal(bdps)
	if err != nil {
		logger.Error.Println(err.Error())
		return err
	}
	//jsonBdps := string(byteBdps)
	_, err = client.Do("SET", keyword, byteBdps)
	if err != nil {
		logger.Error.Println("Set key error, ", err.Error())
		return err
	}
	// cache result for 30 mins
	_, _ = client.Do("EXPIRE", keyword, 1800)
	return nil
}
