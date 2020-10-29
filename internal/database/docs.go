package database

import (
	"docsevrv/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var DB *sqlx.DB

func InitDocsDatabase() (*sqlx.DB, error) {
	var err error
	DB, err = sqlx.Open(config.C.Database.Docsdb.Dialect, config.C.Database.Docsdb.URL)
	if err != nil {
		log.Fatal(err)
	}
	DB.SetMaxIdleConns(config.C.Database.Docsdb.Poolsizemin)
	DB.SetMaxOpenConns(config.C.Database.Docsdb.Poolsizemax)
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return DB, err
}
