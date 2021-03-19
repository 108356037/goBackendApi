package app

import (
	"net/http"

	"github.com/108356037/goBackendMvc/controllers"
)

const (
	port string = ":8888"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
