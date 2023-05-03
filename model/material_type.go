package model

type MaterialType int64

type IsSufficienter interface {
	IsSufficient(user *User) error
	Value() int64
}

const (
	EquipmentUpgrade MaterialType = iota + 1
	LevelUpgrade
	CityUpgrade
)

func (m MaterialType) PickCheckerConditions(currentLevel *Level) []IsSufficienter {
	var checkers []IsSufficienter

	switch m {
	case EquipmentUpgrade:
		checkers = append(checkers, NewGoldCheck(currentLevel), NewDiamondCheck(currentLevel))
	case LevelUpgrade:
		checkers = append(checkers, NewGoldCheck(currentLevel), NewDiamondCheck(currentLevel))
	case CityUpgrade:
		checkers = append(checkers, NewGoldCheck(currentLevel), NewDiamondCheck(currentLevel))
	}

	return checkers
}
