package service

import (
	"encoding/json"
	// "io/ioutil"
	"log"
	"net/http"
)

type requestUrl struct{}

func NewRequestUrl() *requestUrl {
	return &requestUrl{}
}

func (this *requestUrl) GetData() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://robot.pgc.panda.tv/api/show/list", nil)
	if err != nil {
		log.Printf("err: %v", err)
		return
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	type ResultSub struct {
		Roomid        string
		Status        string
		Stream_status string
	}
	var result struct {
		Errmsg string      `json:"errmsg"`
		Errno  int         `json:"errno"`
		Data   []ResultSub `json:"data"`
	}
	// t, _ := ioutil.ReadAll(res.Body)
	// log.Printf("res.body: %T, %v", t, string(t))
	err1 := json.NewDecoder(res.Body).Decode(&result)
	log.Printf("result: %v, %v, %v", result, result.Errmsg, err1)
}
