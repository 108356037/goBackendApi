package postgres

import (
	"database/sql"
	// "log"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var Db *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

func InitDb() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	Db = db
	fmt.Println("Successfully connected to postgres!")
	return nil
}

func Migrate() error {
	driver, err := postgres.WithInstance(Db, &postgres.Config{})
	if err != nil {
		return nil
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/pkg/db/migrations/postgres",
		"postgres",
		driver,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
