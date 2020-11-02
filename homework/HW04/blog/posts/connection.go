package posts

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

//Connector make connect to db an
type Connector interface {
	OpenConnection()
}

//OpenConnection gets connection string and does connect
func (v ConnData) OpenConnection() {
	var err error
	db, err = sql.Open(v.Driver, v.ConnStr)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

}
