package model

type EquipmentEnhancementLevel int64

const (
	// 无强化
	ENHANCEMENT_LEVEL_0 EquipmentEnhancementLevel = 0
	// 一级强化
	ENHANCEMENT_LEVEL_1 = 1
	// 二级强化
	ENHANCEMENT_LEVEL_2 = 2
	// 三级强化
	ENHANCEMENT_LEVEL_3 = 3
	// 四级强化
	ENHANCEMENT_LEVEL_4 = 4
	// 五级强化
	ENHANCEMENT_LEVEL_5 = 5
	// 六级强化
	ENHANCEMENT_LEVEL_6 = 6
	// 七级强化
	ENHANCEMENT_LEVEL_7 = 7
	// 八级强化
	ENHANCEMENT_LEVEL_8 = 8
	// 九级强化
	ENHANCEMENT_LEVEL_9 = 9
	// 十级强化
	ENHANCEMENT_LEVEL_10 = 10
	// 十一级强化
	ENHANCEMENT_LEVEL_11 = 11
	// 十二级强化
	ENHANCEMENT_LEVEL_12 = 12
	// 十三级强化
	ENHANCEMENT_LEVEL_13 = 13
	// 十四级强化
	ENHANCEMENT_LEVEL_14 = 14
)

// 获取对应强化等级可增加的倍数
func (el EquipmentEnhancementLevel) Multiple() float64 {
	switch el {
	case ENHANCEMENT_LEVEL_0:
		return 1
	case ENHANCEMENT_LEVEL_1:
		return 1.1
	case ENHANCEMENT_LEVEL_2:
		return 1.2
	case ENHANCEMENT_LEVEL_3:
		return 1.3
	case ENHANCEMENT_LEVEL_4:
		return 1.4
	case ENHANCEMENT_LEVEL_5:
		return 1.5
	case ENHANCEMENT_LEVEL_6:
		return 1.6
	case ENHANCEMENT_LEVEL_7:
		return 1.7
	case ENHANCEMENT_LEVEL_8:
		return 1.8
	case ENHANCEMENT_LEVEL_9:
		return 1.9
	case ENHANCEMENT_LEVEL_10:
		return 2
	case ENHANCEMENT_LEVEL_11:
		return 2.1
	case ENHANCEMENT_LEVEL_12:
		return 2.2
	case ENHANCEMENT_LEVEL_13:
		return 2.3
	case ENHANCEMENT_LEVEL_14:
		return 2.4
	default:
		return 1
	}
}
