package posts

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

//OpenConnection gets connection string and do connect
func OpenConnection(connectionString string) {
	var err error
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}
