package config

type redis struct {
	Addr     string
	Password string
	Db       int
}

var Redis = &redis{}

func (r *redis) Init() {
}
