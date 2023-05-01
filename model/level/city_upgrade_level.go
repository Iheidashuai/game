package level

import (
	"game/db"
	"game/model/material"
	"game/model/probability"
)

type CityUpgradeLevel struct {
	probability  probability.Probabilityer
	currentLevel *Level
	db           *db.DB
}

func NewCityUpgradeLevel() *CityUpgradeLevel {
	return &CityUpgradeLevel{
		probability:  probability.NewCommonProbability(),
		currentLevel: ZeroLevel(),
	}
}

func (c *CityUpgradeLevel) Incr() error {
	materialDecrClient := material.NewMaterialDecrClient(material.CityUpgrade, c.db, c.currentLevel)
	if err := materialDecrClient.Check(); err != nil {
		return err
	}
	defer func() {
		materialDecrClient.Consume()
	}()
	if c.probability.IsSucceed(c.currentLevel.Value()) {
		return c.currentLevel.Incr()
	}
	return nil
}

func (c *CityUpgradeLevel) Decr() error {
	return c.currentLevel.Decr()
}
