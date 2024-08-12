package dbmanager

import (
	"database/sql"
	"log"
	"math/rand"

	_ "github.com/mattn/go-sqlite3"
)

const file = "app.db"

func GetString() (string, error) {
	sqlite, err := sql.Open("sqlite3", file)
	if err != nil {
		return "", err
	}
	defer sqlite.Close()
	var number int

	err = sqlite.QueryRow("select count(*) from letters").Scan(&number)
	if err != nil {
		log.Printf("error getting letters count: %v", err)
		return "", err
	}
	var result string
	err = sqlite.QueryRow("SELECT text FROM letters WHERE ROWID == $1", 1+rand.Intn(number)).Scan(&result)
	if err != nil {
		log.Printf("error getting letter text: %v", err)
		return "", err
	}
	return result, nil
}
