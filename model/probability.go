package model

type Probabilityer interface {
	IsSucceed(level int64) bool
}
