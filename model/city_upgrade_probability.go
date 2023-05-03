package model

import (
	"fmt"
	"math/rand"
	"time"
)

type CommonProbability struct {
	rand *rand.Rand
}

func NewCommonProbability() *CommonProbability {
	return &CommonProbability{
		rand: rand.New(rand.NewSource(time.Now().Unix())),
	}
}

func (c *CommonProbability) IsSucceed(level int64) bool {
	p := 1 - float64(level)*0.065
	// 获取 0 - 1 之间的随机数，如果小于 p，则返回 true
	pp := c.rand.Float64()
	fmt.Println("pp:", pp, "p:", p)
	return pp < p
}
