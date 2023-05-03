package model

type UserEquipment struct {
	// 上衣装备
	Clothes *Equipment `json:"clothes"`
	// 裤子装备
	Trousers *Equipment `json:"trousers"`
	// 鞋子装备
	Shoes *Equipment `json:"shoes"`
	// 头盔装备
	Helmet *Equipment `json:"helmet"`
	// 护腕装备
	Bracers *Equipment `json:"bracers"`
	// 武器装备
	Weapon *Equipment `json:"weapon"`
}

func (ue *UserEquipment) Atk() float64 {
	if ue == nil {
		return 0
	}
	origin := ue.Clothes.Atk + ue.Trousers.Atk + ue.Shoes.Atk + ue.Helmet.Atk + ue.Bracers.Atk + ue.Weapon.Atk
	return float64(origin) * ue.Clothes.Quality.Multiple()
}

func (ue *UserEquipment) Def() float64 {
	if ue == nil {
		return 0
	}
	origin := ue.Clothes.Def + ue.Trousers.Def + ue.Shoes.Def + ue.Helmet.Def + ue.Bracers.Def + ue.Weapon.Def
	return float64(origin) * ue.Clothes.Quality.Multiple()
}

func (ue *UserEquipment) Hp() float64 {
	if ue == nil {
		return 0
	}
	origin := ue.Clothes.Hp + ue.Trousers.Hp + ue.Shoes.Hp + ue.Helmet.Hp + ue.Bracers.Hp + ue.Weapon.Hp
	return float64(origin) * ue.Clothes.Quality.Multiple()
}

func (ue *UserEquipment) Crit() float64 {
	if ue == nil {
		return 0
	}
	origin := ue.Clothes.Crit + ue.Trousers.Crit + ue.Shoes.Crit + ue.Helmet.Crit + ue.Bracers.Crit + ue.Weapon.Crit
	return float64(origin) * ue.Clothes.Quality.Multiple()
}

func (ue *UserEquipment) Pierce() float64 {
	if ue == nil {
		return 0
	}
	origin := ue.Clothes.Pierce + ue.Trousers.Pierce + ue.Shoes.Pierce + ue.Helmet.Pierce + ue.Bracers.Pierce + ue.Weapon.Pierce
	return float64(origin) * ue.Clothes.Quality.Multiple()
}

func (ue *UserEquipment) Agile() float64 {
	if ue == nil {
		return 0
	}
	origin := ue.Clothes.Agile + ue.Trousers.Agile + ue.Shoes.Agile + ue.Helmet.Agile + ue.Bracers.Agile + ue.Weapon.Agile
	return float64(origin) * ue.Clothes.Quality.Multiple()
}
