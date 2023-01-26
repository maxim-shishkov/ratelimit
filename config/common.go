package config

import "time"

type common struct {
	ServerPort   int
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
}

var Common = &common{}

func (c *common) Init() {
}
