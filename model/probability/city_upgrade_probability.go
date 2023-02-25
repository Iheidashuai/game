package probability

import (
	"math/rand"
	"time"
)

type CityUpgradeProbability struct {
	level int
	rand  *rand.Rand
}

func NewCityUpgradeProbability(level int) *CityUpgradeProbability {
	return &CityUpgradeProbability{
		level: level,
		rand:  rand.New(rand.NewSource(time.Now().Unix())),
	}
}

func (c *CityUpgradeProbability) IsSucceed() bool {
	p := 1 - float64(c.level)*0.05
	return c.rand.Intn(1) > int(p)
}
