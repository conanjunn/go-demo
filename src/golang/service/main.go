package service

import (
	"golang/dao"
	"log"
	"strconv"
	"sync"
)

var r = newRandom(1000)

func MainFn() *[2]string {
	redis := dao.InitRedis()
	wg := &sync.WaitGroup{}
	arr := new([2]string)
	num := r.getRandom()
	str := strconv.Itoa(num)

	wg.Add(1)
	go func() {
		tmp, _ := redis.Get("aaa" + str).Result()
		arr[0] = tmp
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		tmp, _ := redis.Set("aaa"+str, 123, 0).Result()
		arr[1] = tmp
		wg.Done()
	}()

	wg.Wait()
	redis.Del("aaa" + str).Result()
	return arr
}
