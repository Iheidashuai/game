package service

import (
	"context"
	"errors"

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

	userEquipments, err := svc.dbClient.QueryUserEquipment(ctx, userId)
	if err != nil {
		return nil, err
	}

	for _, userEquipment := range userEquipments {
		equipments[userEquipment.EquipmentId].EnhancementLevel = userEquipment.EnhancementLevel
	}

	return user.Vo(&model.EquipmentInUse{
		Clothes:  equipments[user.ClothesEquipmentId],
		Trousers: equipments[user.TrousersEquipmentId],
		Shoes:    equipments[user.ShoesEquipmentId],
		Helmet:   equipments[user.HelmetEquipmentId],
		Bracers:  equipments[user.BracersEquipmentId],
		Weapon:   equipments[user.WeaponEquipmentId],
	}), nil
}

func StrongerEquipment(ctx context.Context, userId int64, equipmentId int64) error {
	user, err := svc.dbClient.UserByUid(ctx, userId)
	if err != nil {
		return err
	}

	equipment, err := svc.dbClient.QueryUserEquipment(ctx, userId)
	if err != nil {
		return err
	}

	upgradeLevel := NewEquipmentUpgradeLevel(model.NewLevel(equipment[equipmentId].EnhancementLevel.Value()), user, svc.dbClient)
	if err = upgradeLevel.Incr(); err != nil {
		return err
	}

	if err = svc.dbClient.UpdateUserEquipment(ctx, upgradeLevel.currentLevel.Value(), userId, equipmentId); err != nil {
		return err
	}

	if upgradeLevel.currentLevel.Value() != equipment[equipmentId].EnhancementLevel.Value() {
		return errors.New("强化成功")
	}

	return errors.New("强化失败")
}
