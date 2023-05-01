package model

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
	Crit     int64  `json:"crit"`
	Pierce   int64  `json:"pierce"`
	Agile    int64  `json:"agile"`
}

func (u *User) CheckPassword(password string) bool {
	if u == nil {
		return false
	}
	return u.Password == password
}
