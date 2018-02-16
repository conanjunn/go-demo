package service

import (
	"fmt"
	"golang/dao"
)

func MainFn() string {
	redis := dao.InitRedis()
	result, err := redis.Ping().Result()
	fmt.Println(result, err)
	return result
}
