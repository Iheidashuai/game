package material

import (
	"errors"
	"game/model"
	"game/model/level"
)

var (
	DiamondUnSufficientError = errors.New("diamond is not sufficient")
)

type DiamondCheck struct {
	currentLevel *level.Level
}

func NewDiamondCheck(currentLevel *level.Level) *DiamondCheck {
	return &DiamondCheck{currentLevel: currentLevel}
}

func (d *DiamondCheck) IsSufficient(user *model.User) error {
	if user.Diamond < d.Value() {
		return DiamondUnSufficientError
	}
	return nil
}

func (d *DiamondCheck) Value() int64 {
	return 100
}
