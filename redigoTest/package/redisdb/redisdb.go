package redisdb

import (
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego"
	"time"
)

var (
	redisClient 	*redis.Pool
	redisHost		string
	redisDB			int
)

func init() {
	redisHost = beego.AppConfig.String("redisHost")
	redisDB, _ = beego.AppConfig.Int("redisDB")
	redisClient = &redis.Pool{
		MaxIdle: beego.AppConfig.DefaultInt("redisMaxidle", 1),
		MaxActive: beego.AppConfig.DefaultInt("redisMaxactive", 10),
		IdleTimeout: 180*time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisHost)
			if err != nil {
				return nil, err
			}
			c.Do("SELECT", redisDB)
			return c,nil
		},
	}
}

func Get() redis.Conn {
	rc := redisClient.Get()
	return rc
}







