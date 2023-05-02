package power

import "game/model"

type Calculator interface {
	Power() int64
}

type CalculationCollector struct {
	user *model.User
}

func NewCalculationCollector(user *model.User) *CalculationCollector {
	return &CalculationCollector{
		user: user,
	}
}

func (c *CalculationCollector) Collect() int64 {
	var result int64

	if c.user != nil {
		result += NewUserPowerCalculator(c.user).Power()
	}

	return result
}
