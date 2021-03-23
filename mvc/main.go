package main

import (
	"fmt"
	"os"

	"github.com/108356037/goBackendMvc/app"
	"github.com/108356037/goBackendMvc/config"
)

type user struct {
	Id        uint64
	firstname string
	lastname  string
	email     string
	team      string
}

func main() {
	db := config.Init()
	rows, err := db.Query("SELECT * FROM tbl_userinfo")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for rows.Next() {
		nullUser := user{}
		if err := rows.Scan(&nullUser.Id, &nullUser.firstname, &nullUser.lastname, &nullUser.email, &nullUser.team); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println(nullUser)
	}
	app.StartApp()
}
