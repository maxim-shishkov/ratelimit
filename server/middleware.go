package server

import (
	"github.com/rs/zerolog/log"
	"net"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := log.Error().Str("RemoteAddr", r.RemoteAddr)

		ip, _, errAddr := net.SplitHostPort(r.RemoteAddr)
		if errAddr != nil {
			logger.Err(errAddr).Msg("invalid format for ip")
			http.Error(w, "error building key", http.StatusBadRequest)
			return
		}

		key, buildErr := buildKey(ip)
		if buildErr != nil {
			logger.Err(buildErr).Msg("building key")
			http.Error(w, "error building key", http.StatusBadRequest)
			return
		}

		isReached, errReach := IsReached(key)
		if errReach != nil {
			http.Error(w, "limit check error", http.StatusBadRequest)
			logger.Err(errReach).Msg("limit check error")
			return
		}

		if isReached {
			http.Error(w, "too many requests", http.StatusTooManyRequests)
			logger.Str("RemoteAddr", r.RemoteAddr).Msg("limit check error")
			return
		}

		next.ServeHTTP(w, r)
	})
}
