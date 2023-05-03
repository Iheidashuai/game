package model

import (
	"errors"
)

var (
	GoldUnSufficientError = errors.New("gold is not sufficient")
)

type GoldCheck struct {
	currentLevel *Level
}

func NewGoldCheck(currentLevel *Level) *GoldCheck {
	return &GoldCheck{currentLevel: currentLevel}
}

func (g *GoldCheck) IsSufficient(user *User) error {
	if user.Gold < g.Value() {
		return GoldUnSufficientError
	}
	return nil
}

func (g *GoldCheck) Value() int64 {
	return 100
}
