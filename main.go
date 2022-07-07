package main

import (
	"rank/dao"
)

// 主入口函数
func main() {

}

func init() {
	redisClient, err := dao.NewRedisClient()
	if err != nil {
		panic(err)
	}
	dao.RedisClient = redisClient
}
