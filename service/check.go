package service

import (
	"game/db"
	"game/model"
)

type Checker struct {
	user *model.User
	db   *db.DB
}

func NewChecker(db *db.DB, user *model.User) *Checker {
	return &Checker{
		db:   db,
		user: user,
	}
}

func (c *Checker) Check(conditions ...model.IsSufficienter) error {
	for _, checker := range conditions {
		if err := checker.IsSufficient(c.user); err != nil {
			return err
		}
	}

	return nil
}
