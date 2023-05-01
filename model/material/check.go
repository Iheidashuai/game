package material

import (
	"game/db"
	"game/model"
)

type IsSufficienter interface {
	IsSufficient(user *model.User) error
	Value() int64
}

type Checker struct {
	user *model.User
	db   *db.DB
}

func NewChecker(db *db.DB) *Checker {
	return &Checker{
		db: db,
	}
}

func (c *Checker) Check(conditions ...IsSufficienter) error {
	for _, checker := range conditions {
		if err := checker.IsSufficient(c.user); err != nil {
			return err
		}
	}

	return nil
}
