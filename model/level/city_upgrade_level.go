package level

import "game/model/probability"

type CityUpgradeLevel struct {
	probability  probability.Probabilityer
	currentLevel *Level
}

func NewCityUpgradeLevel() *CityUpgradeLevel {
	return &CityUpgradeLevel{
		probability:  probability.NewCityUpgradeProbability(),
		currentLevel: ZeroLevel(),
	}
}

func (c CityUpgradeLevel) Incr() bool {
	if c.probability.IsSucceed(c.currentLevel.Value()) {
		return c.currentLevel.Incr()
	}
	return false
}

func (c CityUpgradeLevel) Decr() bool {
	return c.currentLevel.Decr()
}
