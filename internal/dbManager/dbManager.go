package dbmanager

import (
	"database/sql"
	"log"
	"math/rand"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const file = "app.db"

func GetString() (string, error) {
	var number int
	var result string
	randStr := rand.New(rand.NewSource(time.Now().UnixNano()))

	sqlite, err := sql.Open("sqlite3", file)
	if err != nil {
		return "", err
	}
	defer sqlite.Close()

	err = sqlite.QueryRow("select count(*) from letters").Scan(&number)
	if err != nil {
		log.Printf("error getting letters count: %v", err)
		return "", err
	}

	err = sqlite.QueryRow("SELECT text FROM letters WHERE ROWID == $1", 1+randStr.Intn(number)).Scan(&result)
	if err != nil {
		log.Printf("error getting letter text: %v", err)
		return "", err
	}

	return result, nil
}
