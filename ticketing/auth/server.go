package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/108356037/goticketapp/auth/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/users/signup", routes.SignUpHandler).Methods("POST")
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:3000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("server up at at port: %v\n", srv.Addr)
	log.Fatal(srv.ListenAndServe())

}
