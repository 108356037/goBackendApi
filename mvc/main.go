package main

import (
	"fmt"
	"os"

	"github.com/108356037/goBackendMvc/app"
	"github.com/108356037/goBackendMvc/database"
)

func main() {
	if err := database.PostgresConnect(); err != nil {
		fmt.Printf("Error in connecting to datbase: %v\n", err)
		os.Exit(1)
	}

	app.StartApp()
}
