package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"time"
	"encoding/json"
	"redigoTest/package/redisdb"

)


func init() {

}


func main() {

	//简单的 KEY-VALUE
	//testKey()

	//读写 json 
	//testJson()


	// 从池里获取连接
	rc := redisdb.Get()
	// 用完后将连接放回连接池
	defer rc.Close()

}

func testKey()  {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	_, err = c.Do("SET", "mykey", "superzhang", "EX", "5")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}

	time.Sleep(8*time.Second)

	username, err = redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}
}

func testJson()  {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()


	key := "profile"
	imap := map[string]string{"username":"zhang", "mobile":"888"}
	value, _ := json.Marshal(imap)

	n, err := c.Do("SETNX", key, value)
	if err != nil {
		fmt.Println(err)
	}
	if n == int64(1) {
		fmt.Println("sucess")
	}

	var imapGet map[string]string
	valueGet, err := redis.Bytes(c.Do("GET", key))
	if err != nil {
		fmt.Println(err)
	}

	errShal := json.Unmarshal(valueGet, &imapGet)
	if errShal != nil {
		fmt.Println(errShal)
	}
	fmt.Println(imapGet["username"])
	fmt.Println(imapGet["mobile"])

}