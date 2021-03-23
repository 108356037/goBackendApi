package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/108356037/goBackendMvc/errors"
	"github.com/108356037/goBackendMvc/services"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	// fmt.Printf("Getting request from %v \n", req)
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := &errors.MiddlewareError{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
			Code:       "Bad Request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(jsonValue))
		return
	}

	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(jsonValue))
		return
	}

	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
