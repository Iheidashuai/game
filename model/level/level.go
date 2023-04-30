package level

import "sync/atomic"

type Level struct {
	value int64
}

func NewLevel(level int64) *Level {
	return &Level{value: level}
}

func ZeroLevel() *Level {
	return &Level{value: 0}
}

func HighestLevel() *Level {
	return &Level{value: 13}
}

func (l *Level) IsHighest() bool {
	return l.Equal(HighestLevel())
}

func (l *Level) Value() int64 {
	return l.value
}

// Incr 增加 level 的值，返回是否增加成功
func (l *Level) Incr() bool {
	origin := l.Value()
	return atomic.AddInt64(&l.value, 1) != origin
}

func (l *Level) Decr() bool {
	origin := l.Value()
	return atomic.AddInt64(&l.value, -1) != origin
}

func (l *Level) Equal(other *Level) bool {
	return l.Value() == other.Value()
}

type Leveler interface {
	Incr() bool
	Decr() bool
}
