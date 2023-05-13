package db

import (
	"context"

	"game/model"
)

func (db *DB) QueryUserEquipment(ctx context.Context, userId int64) (map[int64]*model.UserEquipment, error) {
	res := make(map[int64]*model.UserEquipment, 0)

	rows, err := db.QueryContext(ctx, "SELECT user_id,equipment_id,enhancement_level,is_permanent,end_time FROM user_equipment WHERE user_id = ?", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var userEquipment model.UserEquipment
		err := rows.Scan(&userEquipment.UserId, &userEquipment.EquipmentId, &userEquipment.EnhancementLevel, &userEquipment.IsPermanent, &userEquipment.EndTime)
		if err != nil {
			return nil, err
		}
		res[userEquipment.EquipmentId] = &userEquipment
	}

	return res, nil
}
