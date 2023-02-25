package level

import "sync/atomic"

type Level int64

func (l Level) Int64() int64 {
	return int64(l)
}

func (l Level) Incr() bool {
	origin := l
	return atomic.AddInt64((*int64)(&l), 1) != origin.Int64()
}

func (l Level) Decr() bool {
	origin := l
	return atomic.AddInt64((*int64)(&l), -1) != origin.Int64()
}

type Leveler interface {
	Incr() bool
	Decr() bool
}
