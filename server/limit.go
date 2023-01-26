package server

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"net/http"
	"ratelimit/config"
)

var ctx = context.Background()
var rdc *redis.Client

func InitRedis() {
	rdc = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Password,
		DB:       config.Redis.Db,
	})
}

func IsReached(key string) (limReached bool, err error) {
	var cnt *redis.IntCmd
	_, err = rdc.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		cnt = pipe.Incr(ctx, key)
		pipe.Expire(ctx, key, config.Limit.Time)

		return nil
	})
	if err != nil {
		log.Error().Err(err).Msg("error connecting to redis")
		return
	}

	if cnt.Val() > config.Limit.MaxConnect {
		limReached = true
	}
	return
}

func reset(w http.ResponseWriter, r *http.Request) {
	vals := mux.Vars(r)
	key := vals["key"]
	if key == "" {
		http.Error(w, "key error", http.StatusBadRequest)
		log.Error().Msg("error get key")
		return
	}

	err := deleteKey(key)
	if err != nil {
		http.Error(w, "failed key deletion", http.StatusBadRequest)
		log.Error().Err(err).Msg("error key delete")
		return
	}

	fmt.Fprintf(w, "удаление %s успешно", key)
}

func deleteKey(key string) (err error) {
	key, err = prepareIp(key)
	if err != nil {
		return
	}
	key += "*"
	keys := rdc.Keys(ctx, key).Val()
	for i := 0; i < len(keys); i++ {
		err = rdc.Del(ctx, keys[i]).Err()
		return
	}
	return
}
