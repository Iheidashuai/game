package service

import (
	"context"

	"game/db"
	"game/model"
)

type MaterialDecrClient interface {
	Check() error
	Consume() error
}

type MaterialDecr struct {
	materialType model.MaterialType
	currentLevel *model.Level
	user         *model.User
	checker      *Checker
	db           *db.DB
}

func NewMaterialDecrClient(materialType model.MaterialType, db *db.DB, currentLevel *model.Level, user *model.User) MaterialDecrClient {
	return &MaterialDecr{materialType: materialType, currentLevel: currentLevel, db: db, checker: NewChecker(db, user), user: user}
}

func (m *MaterialDecr) Check() error {
	conditions := m.materialType.PickCheckerConditions(m.currentLevel)
	if err := m.checker.Check(conditions...); err != nil {
		return err
	}
	return nil
}

func (m *MaterialDecr) Consume() error {
	conditions := m.materialType.PickCheckerConditions(m.currentLevel)
	originUser := m.user

	for _, condition := range conditions {
		if _, ok := condition.(*model.DiamondCheck); ok {
			originUser.Diamond -= condition.Value()
		}
		if _, ok := condition.(*model.GoldCheck); ok {
			originUser.Gold -= condition.Value()
		}
	}

	return m.db.UpdateUserByUid(context.Background(), originUser)
}
