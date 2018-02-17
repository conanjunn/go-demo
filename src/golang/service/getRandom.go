package service

import (
	"math/rand"
	"time"
)

type random struct {
	max    int
	source rand.Source
	rand   *rand.Rand
}

func (this *random) getRandom() int {
	return this.rand.Intn(this.max)
}

func newRandom(max int) *random {
	rdm := &random{
		max: max,
	}
	rdm.source = rand.NewSource(time.Now().Unix())
	rdm.rand = rand.New(rdm.source)
	return rdm
}
