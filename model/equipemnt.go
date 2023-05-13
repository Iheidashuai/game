package model

import (
	"math"
)

type Equipment struct {
	Id               int64                     `json:"id"`
	Name             string                    `json:"name"`
	EquipmentId      int64                     `json:"equipment_id"`
	Atk              float64                   `json:"atk"`
	Def              float64                   `json:"def"`
	Hp               float64                   `json:"hp"`
	Level            int64                     `json:"level"`
	Crit             int64                     `json:"crit"`
	Pierce           int64                     `json:"pierce"`
	Agile            int64                     `json:"agile"`
	EnhancementLevel EquipmentEnhancementLevel `json:"enhancement_level"`
	Quality          EquipmentQuality          `json:"quality"`
	Type             EquipmentType             `json:"type"`
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
		Atk:    int64(math.Ceil(e.Atk)),
		Def:    int64(math.Ceil(e.Def)),
		Hp:     int64(math.Ceil(e.Hp)),
		Level:  e.Level,
		Crit:   e.Crit,
		Pierce: e.Pierce,
		Agile:  e.Agile,
		Type:   e.Type.Name(),
	}
}
