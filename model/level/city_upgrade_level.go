package level

import "game/model/probability"

type CityUpgradeLevel struct {
	probability  probability.Probabilityer
	currentLevel Level
}

func NewCityUpgradeLevel() *CityUpgradeLevel {
	return &CityUpgradeLevel{
		probability:  probability.NewCityUpgradeProbability(0),
		currentLevel: 0,
	}
}

func (c CityUpgradeLevel) Incr() bool {
	if c.probability.IsSucceed() {
		return c.currentLevel.Incr()
	}
	return false
}

func (c CityUpgradeLevel) Decr() bool {
	return c.currentLevel.Decr()
}
