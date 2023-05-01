package material

import "game/model/level"

type MaterialType int64

const (
	EquipmentUpgrade MaterialType = iota + 1
	LevelUpgrade
	CityUpgrade
)

func (m MaterialType) PickCheckerConditions(currentLevel *level.Level) []IsSufficienter {
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
