package config

import "time"

type limit struct {
	Time       time.Duration
	MaxConnect int64
	LenPrefix4 int
	LenPrefix6 int
}

var Limit = &limit{}

func (l *limit) Init() {
}
