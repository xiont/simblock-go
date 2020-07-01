/*
   Copyright 2020 LittleBear(1018589158@qq.com)

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
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
