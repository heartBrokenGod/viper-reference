package handler

import "net/http"

type UserHandler interface {
	CreateNewProfile(w http.ResponseWriter, r *http.Request)
	GetProfileDetails(w http.ResponseWriter, r *http.Request)
}
