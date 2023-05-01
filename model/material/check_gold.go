package material

import (
	"errors"
	"game/model"
	"game/model/level"
)

var (
	GoldUnSufficientError = errors.New("gold is not sufficient")
)

type GoldCheck struct {
	currentLevel *level.Level
}

func NewGoldCheck(currentLevel *level.Level) *GoldCheck {
	return &GoldCheck{currentLevel: currentLevel}
}

func (g *GoldCheck) IsSufficient(user *model.User) error {
	if user.Gold < g.Value() {
		return GoldUnSufficientError
	}
	return nil
}

func (g *GoldCheck) Value() int64 {
	return 100
}
