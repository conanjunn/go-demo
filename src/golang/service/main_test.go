package service

import (
	"testing"
)

func TestMainFn(t *testing.T) {
	for i := 1; i < 10; i++ {
		info := MainFn()
		t.Log(info)
	}
}

func TestRandom(t *testing.T) {
	r := newRandom(1000)
	for i := 0; i < 10; i++ {
		num := r.getRandom()
		t.Log(num)
	}
}
