package model

import "game/model/power"

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Uid      int64  `json:"uid"`
	Password string `json:"password"`
	Atk      int64  `json:"atk"`
	Def      int64  `json:"def"`
	Hp       int64  `json:"hp"`
	Gold     int64  `json:"gold"`
	Diamond  int64  `json:"diamond"`
	Exp      int64  `json:"exp"`
	Level    int64  `json:"level"`
	// 暴击率
	Crit int64 `json:"crit"`
	// 穿透
	Pierce int64 `json:"pierce"`
	// 敏捷
	Agile int64 `json:"agile"`

	// 单独计算
	Power int64 `json:"power"`
}

func (u *User) CheckPassword(password string) bool {
	if u == nil {
		return false
	}
	return u.Password == password
}

func (u *User) SetPower() {
	u.Power = power.NewCalculationCollector(u).Collect()
}
