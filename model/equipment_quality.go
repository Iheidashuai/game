package model

type EquipmentQuality int64

const (
	// 白色
	WHITE EquipmentQuality = 1
	// 绿色
	GREEN = 2
	// 蓝色
	BLUE = 3
	// 紫色
	PURPLE = 4
	// 橙色
	ORANGE = 5
	// 红色
	RED = 6
)

// 获取对应颜色可增加的倍数
func (eq EquipmentQuality) Multiple() float64 {
	switch eq {
	case WHITE:
		return 1
	case GREEN:
		return 2
	case BLUE:
		return 3
	case PURPLE:
		return 4
	case ORANGE:
		return 5
	case RED:
		return 6
	default:
		return 1
	}
}

// 获取对应颜色 argb
func (eq EquipmentQuality) Color() string {
	switch eq {
	case WHITE:
		return "#FFFFFF"
	case GREEN:
		return "#00FF00"
	case BLUE:
		return "#0000FF"
	case PURPLE:
		return "#FF00FF"
	case ORANGE:
		return "#FFA500"
	case RED:
		return "#FF0000"
	default:
		return "#FFFFFF"
	}
}
