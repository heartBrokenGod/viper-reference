package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/heartBrokenGod/viper-reference/api/handler"
)

type ApiServerImpl struct {
	config     *Config
	handler    handler.UserHandler
	router     *mux.Router
	httpserver *http.Server
}

func New(config *Config, handler handler.UserHandler) (*ApiServerImpl, error) {
	if config == nil {
		return nil, errors.New("config dependency is nil")
	}
	if handler == nil {
		return nil, errors.New("handler.UserHandler dependency is nil")
	}
	router := mux.NewRouter()
	// register routes
	router.HandleFunc("/users", handler.CreateNewProfile).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", handler.GetProfileDetails).Methods(http.MethodGet)

	return &ApiServerImpl{
		handler:    handler,
		router:     router,
		httpserver: &http.Server{IdleTimeout: config.TimeOut, Handler: router, Addr: fmt.Sprint(":", config.Port)},
		config:     config,
	}, nil

}

func (apiserver *ApiServerImpl) Start() error {
	errChan := make(chan error, 1)
	go func() {
		errChan <- apiserver.httpserver.ListenAndServe()
	}()
	log.Println("app apiserver started successfully on port:", apiserver.config.Port)
	return <-errChan
}
