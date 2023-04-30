package level

import (
	"fmt"
	"testing"
	"time"
)

func TestCityUpgradeLevel_Decr(t *testing.T) {
	c := NewCityUpgradeLevel()
	n := 0
	for {
		n++
		if c.Incr() {
			fmt.Println("强化成功，当前等级：", c.currentLevel.Value())
		} else {
			c.Decr()
			fmt.Println("强化失败，当前等级：", c.currentLevel.Value())
		}

		if c.currentLevel.IsHighest() {
			fmt.Println("已强化至最高级")
			break
		}

		time.Sleep(time.Second / 4)
	}
	fmt.Println("强化次数：", n)
}
