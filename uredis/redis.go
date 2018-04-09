package uredis

import (
	"github.com/garyburd/redigo/redis"
	"github.com/zhengjianwen/gohrrap/config"
	"github.com/zhengjianwen/gohrrap/log"
)


func ReadCode(key string) error {
	c, err := redis.Dial("tcp", config.G.Redis.Addr)
	defer c.Close()
	if err != nil {
		log.Println("连接Redis 错误", err)
		return err
	}

	syscode, err := redis.String(c.Do("GET", key))
	if err != nil{return  err}
	if syscode == "HaiRui"{
		return nil
	}
	return err
}

func WiteCode(key string) bool {
	c, err := redis.Dial("tcp", config.G.Redis.Addr)
	if err != nil {
		log.Println("Connect to redis error", err)
		return false
	}
	defer c.Close()
	_, err = c.Do("SET", key, "unicloud")
	if err != nil{return  false}
	_, err = c.Do("EXPIRE", key, 60 * config.G.Redis.Timeout)
	if err != nil{return  false}
	return true


}
