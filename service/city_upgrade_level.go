package service

import (
	"game/db"
	"game/model"
)

type CityUpgradeLevel struct {
	probability  model.Probabilityer
	currentLevel *model.Level
	db           *db.DB
}

func NewCityUpgradeLevel() *CityUpgradeLevel {
	return &CityUpgradeLevel{
		probability:  model.NewCommonProbability(),
		currentLevel: model.ZeroLevel(),
	}
}

func (c *CityUpgradeLevel) Incr() error {
	materialDecrClient := NewMaterialDecrClient(model.CityUpgrade, c.db, c.currentLevel)
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
