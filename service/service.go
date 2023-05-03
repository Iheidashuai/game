package service

import (
	"context"
	"game/db"
	"game/model"
)

var (
	svc *Service
)

type Service struct {
	dbClient *db.DB
}

func init() {
	var err error
	svc = &Service{}
	svc.dbClient, err = db.NewDB()
	if err != nil {
		panic(err)
	}
}

func GetUser(ctx context.Context, userId int64) (*model.UserVo, error) {
	user, err := svc.dbClient.UserByUid(ctx, userId)
	if err != nil {
		return nil, err
	}

	equipments, err := svc.dbClient.EquipmentByIds(ctx, user.EquipmentIds())
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
