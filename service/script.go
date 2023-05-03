package service

import (
	"fmt"
	"game/model"
	"time"
)

type BaseAttribute struct {
	Name   string `json:"name"`
	Atk    int64  `json:"atk"`
	Def    int64  `json:"def"`
	Hp     int64  `json:"hp"`
	Crit   int64  `json:"crit"`   // 暴击率
	Pierce int64  `json:"pierce"` // 穿透
	Agile  int64  `json:"agile"`  // 敏捷
}

type TestScript struct {
	Boss     *BaseAttribute   `json:"boss"`
	Soldiers []*BaseAttribute `json:"soldiers"`
}

type ScriptResult struct {
	Win    bool    `json:"win"`
	Reward *Reward `json:"reward"`
}

type Reward struct {
	EquipmentIds []int64 `json:"equipment_ids"`
	Gold         int64   `json:"gold"`
	Exp          int64   `json:"exp"`
	Diamond      int64   `json:"diamond"`
}

func FightScript(user *model.User, script *TestScript) *ScriptResult {
	isWin := true
	reward := &Reward{
		EquipmentIds: []int64{1, 2, 3, 4, 5, 6},
		Gold:         100,
		Exp:          100,
		Diamond:      100,
	}

	for _, soldier := range script.Soldiers {
		for soldier.Hp > 0 && user.Hp > 0 {
			if user.Atk-soldier.Def > 0 {
				soldier.Hp -= user.Atk - soldier.Def
			} else {
				soldier.Hp -= 1
			}

			fmt.Printf("你攻击了敌人 %s,产生了 %d 点伤害,敌人剩余血量 %d\n", soldier.Name, user.Atk-soldier.Def, soldier.Hp)

			if soldier.Hp <= 0 {
				fmt.Printf("你击败了敌人 %s\n", soldier.Name)
				break
			}

			if soldier.Atk-user.Def > 0 {
				user.Hp -= soldier.Atk - user.Def
			} else {
				user.Hp -= 1
			}

			fmt.Printf("敌人 %s 攻击了你,产生了 %d 点伤害,你剩余血量 %d\n", soldier.Name, soldier.Atk-user.Def, user.Hp)
			if user.Hp <= 0 {
				fmt.Printf("你被敌人 %s 击败了\n", soldier.Name)
				isWin = false
				break
			}

			time.Sleep(time.Second)
		}

		if !isWin {
			break
		}
	}

	if isWin {
		fmt.Printf("你通过了副本,获得了 %d 金币, %d 经验, %d 钻石\n", reward.Gold, reward.Exp, reward.Diamond)
	}
	return &ScriptResult{Win: isWin, Reward: reward}
}
