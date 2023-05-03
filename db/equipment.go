package db

import (
	"context"
	"game/model"
)

// EquipmentById 使用 equipmentId 查询装备
func (db *DB) EquipmentById(ctx context.Context, equipmentId int64) (*model.Equipment, error) {
	row := db.QueryRow("select id,name,equipment_id,atk,def,hp,level,crit,pierce,agile,quality,type from equipment where equipment_id = ?", equipmentId)
	equipment := &model.Equipment{}
	if err := row.Scan(&equipment.Id, &equipment.Name, &equipment.EquipmentId, &equipment.Atk, &equipment.Def, &equipment.Hp,
		&equipment.Level, &equipment.Crit, &equipment.Pierce, &equipment.Agile, &equipment.Quality, &equipment.Type); err != nil {
		return nil, err
	}
	return equipment, nil
}

func (db *DB) EquipmentByIds(ctx context.Context, equipmentIds []int64) (map[int64]*model.Equipment, error) {
	equipments := make(map[int64]*model.Equipment)
	for _, equipmentId := range equipmentIds {
		row := db.QueryRow("select id,name,equipment_id,atk,def,hp,level,crit,pierce,agile,quality,type from equipment where equipment_id = ?", equipmentId)
		equipment := &model.Equipment{}
		if err := row.Scan(&equipment.Id, &equipment.Name, &equipment.EquipmentId, &equipment.Atk, &equipment.Def, &equipment.Hp,
			&equipment.Level, &equipment.Crit, &equipment.Pierce, &equipment.Agile, &equipment.Quality, &equipment.Type); err != nil {
			return nil, err
		}
		equipments[equipmentId] = equipment
	}
	return equipments, nil
}
