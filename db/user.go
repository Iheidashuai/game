package db

import (
	"context"
	"game/model"
)

func (db *DB) UserByUid(ctx context.Context, uid int64) (*model.User, error) {
	row := db.QueryRow("select id,name,uid,password from user where uid = ?", uid)
	user := &model.User{}
	if err := row.Scan(&user.Id, &user.Name, &user.Uid, &user.Password); err != nil {
		return nil, err
	}
	return user, nil
}

func (db *DB) AddUser(ctx context.Context, user *model.User) error {
	_, err := db.Exec("insert into user(`name`,`uid`,`password`) values (?,?,?)", user.Name, user.Uid, user.Password)
	return err
}
