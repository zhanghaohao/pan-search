package pansearch

import (
	"baidupan/pansearch/collect"
	"github.com/go-redis/redis"
	"encoding/json"
	logger "util/log"
)

func SearchCache(keyword string) (collect.BDPS, error) {
	client := redis.NewClient(&redis.Options{
		Addr: RedisHost + ":" + RedisPort,
		Password: RedisPassword,
		DB: RedisDB,
	})
	defer client.Close()
	result, err := client.Get(keyword).Result()
	if err == redis.Nil {
		//logger.Info.Printf("Key doesnot exist in Redis")
		return nil, nil
	} else if err != nil {
		logger.Error.Println("Get key from Redis failed.", err.Error())
		return nil, err
	}

	var bdps collect.BDPS
	err = json.Unmarshal([]byte(result), &bdps)
	if err != nil {
		logger.Error.Println("Format value from Redis failed\n", err.Error())
		return nil, err
	}
	return bdps, nil
}
