package database

import (
	"docsevrv/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

func InitUsersDatabase() (*sqlx.DB, error) {
	var err error
	DB, err = sqlx.Open(config.C.Database.Userdb.Dialect, config.C.Database.Userdb.URL)
	if err != nil {
		log.Fatal(err)
	}
	DB.SetMaxIdleConns(config.C.Database.Userdb.Poolsizemin)
	DB.SetMaxOpenConns(config.C.Database.Userdb.Poolsizemax)
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return DB, err
}
