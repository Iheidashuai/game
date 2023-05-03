package model

import "testing"

func TestLevel_Incr(t *testing.T) {
	l := ZeroLevel()
	l.Incr()
	if l.Value() != 1 {
		t.Error("Level should be 2")
	}
	l.Incr()
	if l.Value() != 2 {
		t.Error("Level should be 2")
	}
}
