package model

type EquipmentInUse struct {
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

func (ue *EquipmentInUse) Atk() float64 {
	if ue == nil {
		return 0
	}

	closeAtk := ue.Clothes.Atk * ue.Clothes.Quality.Multiple() * ue.Clothes.EnhancementLevel.Multiple()
	trousersAtk := ue.Trousers.Atk * ue.Trousers.Quality.Multiple() * ue.Trousers.EnhancementLevel.Multiple()
	shoesAtk := ue.Shoes.Atk * ue.Shoes.Quality.Multiple() * ue.Shoes.EnhancementLevel.Multiple()
	helmetAtk := ue.Helmet.Atk * ue.Helmet.Quality.Multiple() * ue.Helmet.EnhancementLevel.Multiple()
	bracersAtk := ue.Bracers.Atk * ue.Bracers.Quality.Multiple() * ue.Bracers.EnhancementLevel.Multiple()
	weaponAtk := ue.Weapon.Atk * ue.Weapon.Quality.Multiple() * ue.Weapon.EnhancementLevel.Multiple()

	return closeAtk + trousersAtk + shoesAtk + helmetAtk + bracersAtk + weaponAtk
}

func (ue *EquipmentInUse) Def() float64 {
	if ue == nil {
		return 0
	}

	closeDef := ue.Clothes.Def * ue.Clothes.Quality.Multiple() * float64(1+ue.Clothes.EnhancementLevel/10)
	trousersDef := ue.Trousers.Def * ue.Trousers.Quality.Multiple() * float64(1+ue.Trousers.EnhancementLevel/10)
	shoesDef := ue.Shoes.Def * ue.Shoes.Quality.Multiple() * float64(1+ue.Shoes.EnhancementLevel/10)
	helmetDef := ue.Helmet.Def * ue.Helmet.Quality.Multiple() * float64(1+ue.Helmet.EnhancementLevel/10)
	bracersDef := ue.Bracers.Def * ue.Bracers.Quality.Multiple() * float64(1+ue.Bracers.EnhancementLevel/10)
	weaponDef := ue.Weapon.Def * ue.Weapon.Quality.Multiple() * float64(1+ue.Weapon.EnhancementLevel/10)

	return closeDef + trousersDef + shoesDef + helmetDef + bracersDef + weaponDef
}

func (ue *EquipmentInUse) Hp() float64 {
	if ue == nil {
		return 0
	}

	closeHp := ue.Clothes.Hp * ue.Clothes.Quality.Multiple() * float64(1+ue.Clothes.EnhancementLevel/10)
	trousersHp := ue.Trousers.Hp * ue.Trousers.Quality.Multiple() * float64(1+ue.Trousers.EnhancementLevel/10)
	shoesHp := ue.Shoes.Hp * ue.Shoes.Quality.Multiple() * float64(1+ue.Shoes.EnhancementLevel/10)
	helmetHp := ue.Helmet.Hp * ue.Helmet.Quality.Multiple() * float64(1+ue.Helmet.EnhancementLevel/10)
	bracersHp := ue.Bracers.Hp * ue.Bracers.Quality.Multiple() * float64(1+ue.Bracers.EnhancementLevel/10)
	weaponHp := ue.Weapon.Hp * ue.Weapon.Quality.Multiple() * float64(1+ue.Weapon.EnhancementLevel/10)

	return closeHp + trousersHp + shoesHp + helmetHp + bracersHp + weaponHp
}

func (ue *EquipmentInUse) Crit() float64 {
	if ue == nil {
		return 0
	}

	return float64(ue.Clothes.Crit + ue.Trousers.Crit + ue.Shoes.Crit + ue.Helmet.Crit + ue.Bracers.Crit + ue.Weapon.Crit)
}

func (ue *EquipmentInUse) Pierce() float64 {
	if ue == nil {
		return 0
	}

	return float64(ue.Clothes.Pierce + ue.Trousers.Pierce + ue.Shoes.Pierce + ue.Helmet.Pierce + ue.Bracers.Pierce + ue.Weapon.Pierce)
}

func (ue *EquipmentInUse) Agile() float64 {
	if ue == nil {
		return 0
	}

	return float64(ue.Clothes.Agile + ue.Trousers.Agile + ue.Shoes.Agile + ue.Helmet.Agile + ue.Bracers.Agile + ue.Weapon.Agile)
}
