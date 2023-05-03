package service

import (
	"fmt"
	"game/model"
	"testing"
)

func TestFightScript(t *testing.T) {
	user := &model.User{
		Name: "test",
		Atk:  20,
		Def:  5,
		Hp:   100,
	}

	script := &TestScript{
		Boss: &BaseAttribute{
			Name: "boss",
			Atk:  100,
			Def:  100,
			Hp:   100,
		},
		Soldiers: []*BaseAttribute{
			{
				Name: "soldier1",
				Atk:  10,
				Def:  5,
				Hp:   100,
			},
			{
				Name: "soldier2",
				Atk:  10,
				Def:  5,
				Hp:   100,
			},
		},
	}

	result := FightScript(user, script)
	fmt.Println(result)
}
