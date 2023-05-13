package db

import (
	"context"

	"game/model"
)

func (db *DB) UserByUid(ctx context.Context, uid int64) (*model.User, error) {
	row := db.QueryRow("select id,name,uid,password,atk,def,hp,gold,diamond,exp,level,crit,pierce,agile,clothes_equipment_id,weapon_equipment_id,trousers_equipment_id,shoes_equipment_id,helmet_equipment_id,bracers_equipment_id from user where uid = ?", uid)
	user := &model.User{}
	if err := row.Scan(&user.Id, &user.Name, &user.Uid, &user.Password, &user.Atk,
		&user.Def, &user.Hp, &user.Gold, &user.Diamond, &user.Exp, &user.Level,
		&user.Crit, &user.Pierce, &user.Agile, &user.ClothesEquipmentId,
		&user.WeaponEquipmentId, &user.TrousersEquipmentId, &user.ShoesEquipmentId, &user.HelmetEquipmentId,
		&user.BracersEquipmentId); err != nil {
		return nil, err
	}
	return user, nil
}

func (db *DB) AddUser(ctx context.Context, user *model.User) error {
	_, err := db.Exec("insert into user(`name`,`uid`,`password`) values (?,?,?)", user.Name, user.Uid, user.Password)
	return err
}

func (db *DB) UpdateUserByUid(ctx context.Context, user *model.User) error {
	_, err := db.Exec("update user set atk = ?,def = ?,hp = ?,gold = ?,diamond = ?,exp = ?,level = ?,crit = ?,pierce = ?,agile = ?,clothes_equipment_id = ?,weapon_equipment_id = ?,trousers_equipment_id = ?,shoes_equipment_id = ?,helmet_equipment_id = ?,bracers_equipment_id = ? where uid = ?",
		user.Atk, user.Def, user.Hp, user.Gold, user.Diamond, user.Exp, user.Level, user.Crit, user.Pierce, user.Agile, user.ClothesEquipmentId, user.WeaponEquipmentId, user.TrousersEquipmentId, user.ShoesEquipmentId, user.HelmetEquipmentId, user.BracersEquipmentId, user.Uid)
	if err != nil {
		return err
	}
	return nil
}
