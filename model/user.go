package model

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Uid      int64  `json:"uid"`
	Password string `json:"password"`
	Atk      int64  `json:"atk"`
	Def      int64  `json:"def"`
	Hp       int64  `json:"hp"`
	Gold     int64  `json:"gold"`
	Diamond  int64  `json:"diamond"`
	Exp      int64  `json:"exp"`
	Level    int64  `json:"level"`
	// 暴击率
	Crit int64 `json:"crit"`
	// 穿透
	Pierce int64 `json:"pierce"`
	// 敏捷
	Agile int64 `json:"agile"`

	ClothesEquipmentId  int64 `json:"clothes_equipment_id"`
	WeaponEquipmentId   int64 `json:"weapon_equipment_id"`
	TrousersEquipmentId int64 `json:"trousers_equipment_id"`
	ShoesEquipmentId    int64 `json:"shoes_equipment_id"`
	HelmetEquipmentId   int64 `json:"helmet_equipment_id"`
	BracersEquipmentId  int64 `json:"bracer_equipment_id"`

	// 单独计算
	Power int64 `json:"power"`
}

func (u *User) Vo(userEquipment *EquipmentInUse) *UserVo {
	collector := NewCalculationCollector(u.WithEquipment(userEquipment))
	vo := &UserVo{
		Username: u.Name,
		Clothes:  userEquipment.Clothes.Vo(),
		Weapon:   userEquipment.Weapon.Vo(),
		Trousers: userEquipment.Trousers.Vo(),
		Shoes:    userEquipment.Shoes.Vo(),
		Helmet:   userEquipment.Helmet.Vo(),
		Bracers:  userEquipment.Bracers.Vo(),
		Atk:      u.Atk,
		Def:      u.Def,
		Hp:       u.Hp,
		Gold:     u.Gold,
		Diamond:  u.Diamond,
		Exp:      u.Exp,
		Level:    u.Level,
		Crit:     u.Crit,
		Pierce:   u.Pierce,
		Agile:    u.Agile,
		Power:    collector.Collect(),
	}

	return vo
}

func (u *User) EquipmentIds() []int64 {
	return []int64{u.ClothesEquipmentId, u.WeaponEquipmentId, u.TrousersEquipmentId, u.ShoesEquipmentId, u.HelmetEquipmentId, u.BracersEquipmentId}
}

func (u *User) CheckPassword(password string) bool {
	if u == nil {
		return false
	}
	return u.Password == password
}

func (u *User) WithEquipment(equipmentInUse *EquipmentInUse) *User {
	if u == nil || equipmentInUse == nil {
		return u
	}
	u.Atk += int64(equipmentInUse.Atk())
	u.Def += int64(equipmentInUse.Def())
	u.Hp += int64(equipmentInUse.Hp())
	u.Crit += int64(equipmentInUse.Crit())
	u.Pierce += int64(equipmentInUse.Pierce())
	u.Agile += int64(equipmentInUse.Agile())
	return u
}
