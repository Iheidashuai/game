package db

import (
	"context"
	"game/model"
	"testing"
)

func TestUser(t *testing.T) {
	db, err := NewDB()
	if err != nil {
		t.Errorf("db conn error %v", err)
	}
	uid := 100

	if err = db.AddUser(context.Background(), &model.User{
		Name:     "heidashuai",
		Uid:      int64(uid),
		Password: "123456",
	}); err != nil {
		t.Errorf("db AddUser error %v", err)
	}

	user, err := db.UserByUid(context.Background(), int64(uid))
	if err != nil {
		t.Errorf("db UserByUid error %v", err)
	}
	if !user.CheckPassword("123456") {
		t.Errorf("user password error")
	}
}
