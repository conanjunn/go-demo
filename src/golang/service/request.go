package service

import (
	"encoding/json"
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
	var result struct {
		errno  string `json:"errmsg"`
		errmsg int    `json:"errno"`
	}
	err1 := json.NewDecoder(res.Body).Decode(&result)
	log.Printf("result: %v, %v, %v", result.errno, result.errmsg, err1)
}
