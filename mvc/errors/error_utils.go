package errors

type MiddlewareError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Code       string `json:"code"`
}

type ServiceError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Code       string `json:"code"`
}

func (miderr *MiddlewareError) Error() string {
	return "from middleware err"
}
