package model

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Uid      int64  `json:"uid"`
	Password string `json:"password"`
}

func (u *User) CheckPassword(password string) bool {
	if u == nil {
		return false
	}
	return u.Password == password
}
