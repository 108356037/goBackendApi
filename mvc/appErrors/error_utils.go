package appErrors

import (
	"encoding/json"
	"net/http"
)

type CommonError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Code       string `json:"code"`
}

// type MiddlewareError struct {
// 	Message    string `json:"message"`
// 	StatusCode int    `json:"status_code"`
// 	Code       string `json:"code"`
// }

// type ServiceError struct {
// 	Message    string `json:"message"`
// 	StatusCode int    `json:"status_code"`
// 	Code       string `json:"code"`
// }

func MiddlewareError(err error, resp http.ResponseWriter) {
	jsonValue, _ := json.Marshal(
		CommonError{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
			Code:       "Bad Request",
		},
	)
	resp.WriteHeader(http.StatusBadRequest)
	resp.Write([]byte(jsonValue))
	return
}

func SerivceError(err error, resp http.ResponseWriter) {
	jsonValue, _ := json.Marshal(
		CommonError{
			Message:    err.Error(),
			StatusCode: http.StatusNotFound,
			Code:       "Not found",
		},
	)
	resp.WriteHeader(http.StatusNotFound)
	resp.Write([]byte(jsonValue))
	return
}
