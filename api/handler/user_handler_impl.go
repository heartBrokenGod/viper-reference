package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/heartBrokenGod/viper-reference/entity"
	"github.com/heartBrokenGod/viper-reference/service"
)

type UserHandlerImpl struct {
	config      *Config
	userService service.UserService
}

func New(config *Config, userService service.UserService) (*UserHandlerImpl, error) {
	if config == nil {
		return nil, errors.New("config depencency is nil")
	}
	if userService == nil {
		return nil, errors.New("service.UserService dependency is nil")
	}
	return &UserHandlerImpl{
		config:      config,
		userService: userService,
	}, nil
}

func (userHandler *UserHandlerImpl) CreateNewProfile(w http.ResponseWriter, r *http.Request) {
	user := &entity.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// if validate request config specifies true then validate the request
	if userHandler.config.ValidateRequest {
		if strings.ContainsAny(user.Name, `!@#$%^&*`) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("validation failed on name field of user"))
			return
		}
	}

	user, err = userHandler.userService.Add(*user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	userJson, _ := json.Marshal(user)
	w.WriteHeader(http.StatusCreated)
	w.Write(userJson)

}

func (userHandler *UserHandlerImpl) GetProfileDetails(w http.ResponseWriter, r *http.Request) {

	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("validation failed on name field of user"))
		return
	}
	user, err := userHandler.userService.Get(uint32(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	userJson, _ := json.Marshal(user)
	w.WriteHeader(http.StatusOK)
	w.Write(userJson)

}
