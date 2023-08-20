package config

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func NewDatabase() *sql.DB {
	connstr := viper.GetString("Database")

	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatalln(err.Error())
	}

	if err := db.Ping(); err != nil {
		log.Fatalln(err.Error())
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db
}
