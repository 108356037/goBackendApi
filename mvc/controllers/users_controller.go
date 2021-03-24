package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/108356037/goBackendMvc/appErrors"
	"github.com/108356037/goBackendMvc/services"
	"github.com/gorilla/mux"
)

func GetUserById(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	userId := vars["id"]
	user, err := services.GetUserById(userId)
	if err != nil {
		appErrors.SerivceError(err, resp)
		return
	}
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}

func GetAllUsers(resp http.ResponseWriter, req *http.Request) {
	users, err := services.GetAllUsers()
	if err != nil {
		appErrors.SerivceError(err, resp)
		return
	}
	jsonValue, _ := json.Marshal(users)
	resp.Write(jsonValue)
}
