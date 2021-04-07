package app

import (
	"net/http"
	"time"

	"log"

	"github.com/gorilla/mux"

	"github.com/108356037/goBackendMvc/controllers"
)

func StartApp() {
	router := mux.NewRouter()
	router.StrictSlash(true)
	router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id:[0-9a-z]+}", controllers.GetUserById).Methods("GET")

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:3000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
