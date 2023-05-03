package model

type Equipment struct {
	Id          int64            `json:"id"`
	Name        string           `json:"name"`
	EquipmentId int64            `json:"equipment_id"`
	Atk         int64            `json:"atk"`
	Def         int64            `json:"def"`
	Hp          int64            `json:"hp"`
	Level       int64            `json:"level"`
	Crit        int64            `json:"crit"`
	Pierce      int64            `json:"pierce"`
	Agile       int64            `json:"agile"`
	Quality     EquipmentQuality `json:"quality"`
	Type        EquipmentType    `json:"type"`
}

func (e *Equipment) Color() string {
	if e == nil {
		return ""
	}
	return e.Quality.Color()
}

func (e *Equipment) Vo() *EquipmentVo {
	if e == nil {
		return nil
	}
	return &EquipmentVo{
		Name:   e.Name,
		Color:  e.Color(),
		Atk:    e.Atk,
		Def:    e.Def,
		Hp:     e.Hp,
		Level:  e.Level,
		Crit:   e.Crit,
		Pierce: e.Pierce,
		Agile:  e.Agile,
		Type:   e.Type.Name(),
	}
}
