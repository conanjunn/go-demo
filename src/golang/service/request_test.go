package service

import (
	"testing"
)

func TestGetData(t *testing.T) {
	obj := NewRequestUrl()
	obj.GetData()
	t.Logf("xxx")
}
