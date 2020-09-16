package db

import (
	"github.com/garyburd/redigo/redis"
	"time"
	"util/logger"
)

type RedisConfig struct {
	Host 		string				`yaml:"host"`
	Port 		string				`yaml:"port"`
	Password 	string				`yaml:"password"`
}

type RedisPool struct {
	*redis.Pool
}

func NewRedisPoolForConfig(config *RedisConfig) (redisPool *RedisPool, err error) {
	redisPool = new(RedisPool)
	pool := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 80 * time.Second,
		Dial: func() (redis.Conn, error) {
			cli, err := redis.Dial("tcp", config.Host + ":" + config.Port)
			if err != nil {
				logger.Error.Println(err)
				return nil, err
			}
			if config.Password != "" {
				if _, err := cli.Do("AUTH", config.Password); err != nil {
					logger.Error.Println(err)
					cli.Close()
					return nil, err
				}
			}
			return cli, err
		},
		TestOnBorrow: func(cli redis.Conn, t time.Time) error {
			if _, err := cli.Do("PING"); err != nil {
				logger.Error.Println(err)
				return err
			}
			return nil
		},
	}
	redisPool.Pool = pool
	return
}
