package model

import (
	"errors"
)

var (
	DiamondUnSufficientError = errors.New("diamond is not sufficient")
)

type DiamondCheck struct {
	currentLevel *Level
}

func NewDiamondCheck(currentLevel *Level) *DiamondCheck {
	return &DiamondCheck{currentLevel: currentLevel}
}

func (d *DiamondCheck) IsSufficient(user *User) error {
	if user.Diamond < d.Value() {
		return DiamondUnSufficientError
	}
	return nil
}

func (d *DiamondCheck) Value() int64 {
	return 100
}
