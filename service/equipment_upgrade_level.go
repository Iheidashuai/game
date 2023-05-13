package service

import (
	"game/db"
	"game/model"
)

type EquipmentUpgradeLevel struct {
	probability  model.Probabilityer
	currentLevel *model.Level
	user         *model.User
	db           *db.DB
}

func NewEquipmentUpgradeLevel(currentLevel *model.Level, user *model.User, db *db.DB) *EquipmentUpgradeLevel {
	return &EquipmentUpgradeLevel{
		probability:  model.NewCommonProbability(),
		currentLevel: currentLevel,
		db:           db,
		user:         user,
	}
}

func (c *EquipmentUpgradeLevel) Incr() error {
	materialDecrClient := NewMaterialDecrClient(model.EquipmentUpgrade, c.db, c.currentLevel, c.user)
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

func (c *EquipmentUpgradeLevel) Decr() error {
	return c.currentLevel.Decr()
}
