package database

import (
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
)

func InitMySQL() (*sqlx.DB, error) {

	db, err := sqlx.Open(os.Getenv("MYSQL_DIALEG"), os.Getenv("MYSQL_DSN"))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	db.SetConnMaxLifetime(time.Second * 10)
	db.SetMaxIdleConns(0)
	db.SetMaxOpenConns(100)

	return db, nil
}
