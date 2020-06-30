package utils

import (
	"math/rand"
	"time"
)

type MyRand struct {
	rawRand *rand.Rand
}

//var Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func NewMyRand() *MyRand {
	return &MyRand{
		rawRand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (mr *MyRand) NextInt64() int64 {
	return mr.rawRand.Int63()
}

func (mr *MyRand) NextFloat64() float64 {
	return mr.rawRand.Float64()
}

func (mr *MyRand) Shuffle(n int, swap func(i, j int)) {
	mr.rawRand.Shuffle(n, swap)
}
