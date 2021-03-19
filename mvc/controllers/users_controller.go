package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/108356037/goBackendMvc/services"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	fmt.Printf("Getting request %v \n", req)
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		resp.WriteHeader(http.StatusBadGateway)
		resp.Write([]byte("user_id must be int"))
		return
	}

	user, err := services.GetUser(userId)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))
		return
	}

	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
