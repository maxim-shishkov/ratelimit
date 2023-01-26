package server

import (
	"errors"
	"net"
	"ratelimit/config"
	"strings"
	"time"
)

func buildKey(ip string) (key string, err error) {
	key, err = prepareIp(ip)
	if err != nil {
		return
	}

	// example 12:54 127.0.0.1 -> 127.0.0.0:12:54
	key += "0:" + time.Now().Format("15:04")
	return
}

func prepareIp(str string) (ip string, err error) {
	rawIp := net.ParseIP(str)

	switch true {
	case rawIp.To4() != nil:
		ip = parseIP(rawIp.To4().String(), ".", config.Limit.LenPrefix4)
	case rawIp.To16() != nil:
		ip = parseIP(rawIp.To16().String(), ":", config.Limit.LenPrefix6)
	default:
		err = errors.New("неверный формат IP адреса")
	}

	return
}

func parseIP(raw, sep string, length int) (ip string) {
	ips := strings.Split(raw, sep)
	for i := 0; i < length; i++ {
		ip += ips[i] + sep
	}
	return
}
