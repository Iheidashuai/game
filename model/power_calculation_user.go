package model

type UserPowerCalculator struct {
	user *User
}

func NewUserPowerCalculator(user *User) *UserPowerCalculator {
	return &UserPowerCalculator{
		user: user,
	}
}

func (u *UserPowerCalculator) Power() int64 {
	user := u.user
	atk := user.Atk
	def := user.Def
	hp := user.Hp
	crit := user.Crit
	pierce := user.Pierce
	agile := user.Agile

	atkPower := atk * 4
	defPower := def * 4
	hpPower := hp * 1
	critPower := crit * 1
	piercePower := pierce * 16
	agilePower := agile * 10

	return atkPower + defPower + hpPower + critPower + piercePower + agilePower
}
