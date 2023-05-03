package model

type UserVo struct {
	Username string `json:"username"`

	Power int64 `json:"power"`

	Clothes  *EquipmentVo `json:"clothes"`
	Weapon   *EquipmentVo `json:"weapon"`
	Trousers *EquipmentVo `json:"trousers"`
	Shoes    *EquipmentVo `json:"shoes"`
	Helmet   *EquipmentVo `json:"helmet"`
	Bracers  *EquipmentVo `json:"bracers"`
	Atk      int64        `json:"atk"`
	Def      int64        `json:"def"`
	Hp       int64        `json:"hp"`
	Gold     int64        `json:"gold"`
	Diamond  int64        `json:"diamond"`
	Exp      int64        `json:"exp"`
	Level    int64        `json:"level"`
	// 暴击率
	Crit int64 `json:"crit"`
	// 穿透
	Pierce int64 `json:"pierce"`
	// 敏捷
	Agile int64 `json:"agile"`
}
