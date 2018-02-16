package service

import (
	"golang/dao"
	"testing"
)

func TestMainFn(t *testing.T) {
	client := dao.InitRedis()
	info, _ := client.Ping().Result()
	t.Log(info)
}
