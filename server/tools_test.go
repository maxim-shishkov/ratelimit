package server

import (
	"github.com/stretchr/testify/assert"
	"ratelimit/config"
	"testing"
	"time"
)

func Test_prepareIp(t *testing.T) {
	tests := []struct {
		ip       string
		isError  bool
		expected string
	}{
		{
			ip:       "10.10.10.2",
			isError:  false,
			expected: "10.10.10.",
		},
		{
			ip:       "2222:fb8::68",
			isError:  false,
			expected: "2222:fb8::68:",
		},
		{
			ip:       "2a02:06b8:0000:0001:0000:0000:feed:0a11",
			isError:  false,
			expected: "2a02:6b8:0:1:",
		},
		{
			ip:       "2a02:6b8:0:1:0:0:feed:a11",
			isError:  false,
			expected: "2a02:6b8:0:1:",
		},
		{
			ip:       "2a02:6b8:0:1::feed:a11",
			isError:  false,
			expected: "2a02:6b8:0:1:",
		},
		{
			ip:       "localhost:52000",
			isError:  true,
			expected: "",
		},
	}

	config.Limit.LenPrefix4 = 3
	config.Limit.LenPrefix6 = 4

	for _, tt := range tests {
		ip, err := prepareIp(tt.ip)
		assert.Equal(t, tt.isError, err != nil)
		assert.Equal(t, tt.expected, ip)
	}
}

func Test_buildKey(t *testing.T) {
	n := "0:" + time.Now().Format("15:04")
	key, err := buildKey("192.168.0.1")
	assert.NoError(t, err)
	assert.Equal(t, "192.168.0."+n, key)
}
