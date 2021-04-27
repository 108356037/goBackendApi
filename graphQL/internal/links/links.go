package links

import (
	database "graphql/example.com/internal/pkg/db/mysql"
	"graphql/example.com/internal/users"
	"log"
)

type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

func (link Link) Save() (int64, error) {
	stmt, err := database.Db.Prepare("INSERT INTO Links(Title,Address) VALUES(?,?)")
	if err != nil {
		return -1, err
	}

	res, err := stmt.Exec(link.Title, link.Address)
	if err != nil {
		return -1, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	log.Print("Row inserted!")
	return id, err
}

func GetAll() ([]Link, error) {
	stmt, err := database.Db.Prepare("SELECT Id, Title, Address FROM Links;")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	resultLinks := []Link{}

	for rows.Next() {
		link := Link{}
		err := rows.Scan(&link.ID, &link.Title, &link.Address)
		if err != nil {
			return nil, err
		}
		resultLinks = append(resultLinks, link)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return resultLinks, nil
}
