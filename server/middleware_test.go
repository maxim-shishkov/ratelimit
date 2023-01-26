package server

import (
	"errors"
	"github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"ratelimit/config"
	"testing"
	"time"
)

func Test_deleteKey(t *testing.T) {
	db, mock := redismock.NewClientMock()
	rdc = db

	mock.ExpectDel("127.0.0.*").RedisNil()

	err := deleteKey("127.0.0.")
	assert.False(t, errors.Is(err, redis.Nil))
}

func Test_IsReached(t *testing.T) {
	db, mock := redismock.NewClientMock()
	rdc = db

	tests := []struct {
		count      int
		maxConnect int64
		expected   bool
	}{
		{
			maxConnect: 10,
			count:      10,
			expected:   false,
		},
		{
			maxConnect: 10,
			count:      12,
			expected:   true,
		},
	}

	key := "key"
	ttl, _ := time.ParseDuration("10s")
	config.Limit.Time = ttl

	for _, tt := range tests {
		config.Limit.MaxConnect = tt.maxConnect

		var limit bool
		var err error

		for i := 0; i < tt.count; i++ {
			mock.ExpectTxPipeline()
			mock.ExpectIncr(key).SetVal(int64(i))
			mock.ExpectExpire(key, ttl).SetVal(true)
			mock.ExpectTxPipelineExec()

			limit, err = IsReached(key)
			assert.False(t, errors.Is(err, redis.Nil))
		}
		assert.Equal(t, tt.expected, limit)
	}
}
