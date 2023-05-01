package exp

import "game/model/level"

func isUpgrade(level *level.Level, currentExp int64) bool {
	if level.Value() <= 10 {
		return currentExp >= 100*level.Value()
	}

	if level.Value() > 10 && level.Value() <= 20 {
		return currentExp >= 200*level.Value()
	}

	if level.Value() > 20 && level.Value() <= 30 {
		return currentExp >= 500*level.Value()
	}

	if level.Value() > 30 && level.Value() <= 40 {
		return currentExp >= 1000*level.Value()
	}

	if level.Value() > 40 && level.Value() <= 50 {
		return currentExp >= 2000*level.Value()
	}

	if level.Value() > 50 && level.Value() <= 60 {
		return currentExp >= 5000*level.Value()
	}

	return false
}
