package main

import (
	"docsevrv/internal/config"
	"docsevrv/internal/controller"
	"docsevrv/internal/database"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/labstack/echo/v4"
	"log"
)

// если честно, я не очень понял, насколько серьезная нужна авторизация, можно было сделать через jwt токен и прочее...

func main() {
	config.InitConfig()

	_, err := database.InitUsersDatabase()
	if err != nil {
		log.Fatal(err)
	}

	_, err = database.InitDocsDatabase()
	if err != nil {
		log.Fatal(err)
	}

	ch, err := migrate.New("file://./migrations/userbase", config.C.Database.Userdb.URL)
	if err != nil {
		log.Fatalln(err)
	}

	err = ch.Up()
	if err != nil && err.Error() != "no changes in userbase" {
		log.Fatalln(err)
	}

	ch, err = migrate.New("file://./migrations/docbase", config.C.Database.Docsdb.URL)
	if err != nil {
		log.Fatalln(err)
	}

	err = ch.Up()
	if err != nil && err.Error() != "no changes in docsbase" {
		log.Fatalln(err)
	}

	e := echo.New()

	api := e.Group("/api")

	controller.Add(api)

	e.Logger.Fatal(e.Start(config.C.Server.URL))
}
