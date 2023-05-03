package service

import (
	"context"
	"game/db"
	"game/model"
)

var (
	dbClient *db.DB
	err      error
)

func init() {
	dbClient, err = db.NewDB()
	if err != nil {
		panic(err)
	}
}

func GetUser(ctx context.Context, userId int64) (*model.UserVo, error) {
	user, err := dbClient.UserByUid(ctx, userId)
	if err != nil {
		return nil, err
	}

	equipments, err := dbClient.EquipmentByIds(ctx, user.EquipmentIds())
	if err != nil {
		return nil, err
	}

	return user.Vo(&model.UserEquipment{
		Clothes:  equipments[user.ClothesEquipmentId],
		Trousers: equipments[user.TrousersEquipmentId],
		Shoes:    equipments[user.ShoesEquipmentId],
		Helmet:   equipments[user.HelmetEquipmentId],
		Bracers:  equipments[user.BracersEquipmentId],
		Weapon:   equipments[user.WeaponEquipmentId],
	}), nil
}
