package db

import (
	"context"
	"game/model"
)

func (db *DB) UserByUid(ctx context.Context, uid int64) (*model.User, error) {
	row := db.QueryRow("select id,name,uid,password,atk,def,hp,gold,diamond,exp,level,crit,pierce,agile from user where uid = ?", uid)
	user := &model.User{}
	if err := row.Scan(&user.Id, &user.Name, &user.Uid, &user.Password); err != nil {
		return nil, err
	}
	user.SetPower()
	return user, nil
}

func (db *DB) AddUser(ctx context.Context, user *model.User) error {
	_, err := db.Exec("insert into user(`name`,`uid`,`password`) values (?,?,?)", user.Name, user.Uid, user.Password)
	return err
}

func (db *DB) UpdateUserByUid(ctx context.Context, user *model.User) error {
	_, err := db.Exec("update user set atk = ?,def = ?,hp = ?,gold = ?,diamond = ?,exp = ?,level = ?,crit = ?,pierce = ?,agile = ? where uid = ?",
		user.Atk, user.Def, user.Hp, user.Gold, user.Diamond, user.Exp, user.Level, user.Crit, user.Pierce, user.Agile, user.Uid)
	if err != nil {
		return err
	}
	return nil
}
