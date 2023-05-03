package model

type EquipmentVo struct {
	Name   string `json:"name"`
	Color  string `json:"color"`
	Atk    int64  `json:"atk"`
	Def    int64  `json:"def"`
	Hp     int64  `json:"hp"`
	Level  int64  `json:"level"`
	Crit   int64  `json:"crit"`
	Pierce int64  `json:"pierce"`
	Agile  int64  `json:"agile"`
	Type   string `json:"type"`
}
