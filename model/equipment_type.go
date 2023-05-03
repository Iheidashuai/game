package model

type EquipmentType int64

const (
	// 武器类型
	WEAPON EquipmentType = 1
	// 上衣类型
	CLOTHES = 2
	// 裤子类型
	TROUSERS = 3
	// 鞋子类型
	SHOES = 4
	// 头盔类型
	HELMET = 5
	// 护腕类型
	BRACERS = 6
	// 项链类型
	NECKLACE = 7
	// 戒指类型
	RING = 8
	// 腰带类型
	BELT = 9
	// 护符类型
	AMULET = 10
)

func (e EquipmentType) Name() string {
	switch e {
	case WEAPON:
		return "武器"
	case CLOTHES:
		return "上衣"
	case TROUSERS:
		return "裤子"
	case SHOES:
		return "鞋子"
	case HELMET:
		return "头盔"
	case BRACERS:
		return "护腕"
	case NECKLACE:
		return "项链"
	case RING:
		return "戒指"
	case BELT:
		return "腰带"
	case AMULET:
		return "护符"
	default:
		return "未知类型"
	}
}
