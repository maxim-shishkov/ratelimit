package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
	"ratelimit/config"
)

const baseURL = "/api/v1/"

func Run() {
	router := mux.NewRouter()
	router.HandleFunc(baseURL+"reset"+"/{key}/", reset)

	subRouter := router.PathPrefix(baseURL).Subrouter()
	subRouter.Use(Middleware)
	subRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, "static content")
	})

	addr := fmt.Sprintf(":%d", config.Common.ServerPort)

	srv := &http.Server{
		Handler:      router,
		Addr:         addr,
		WriteTimeout: config.Common.WriteTimeout,
		ReadTimeout:  config.Common.ReadTimeout,
	}

	log.Debug().Str("Addr", srv.Addr).Msg("Listener Server")
	log.Error().Err(srv.ListenAndServe()).Msg("Error listener server")
}
