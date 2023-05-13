package model

type UserEquipment struct {
	UserId           int64                     `json:"user_id"`
	EquipmentId      int64                     `json:"equipment_id"`
	EnhancementLevel EquipmentEnhancementLevel `json:"enhancement_level"`
	IsPermanent      bool                      `json:"is_permanent"`
	EndTime          int64                     `json:"end_time"`
}
