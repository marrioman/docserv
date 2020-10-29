package common

import (
	"docsevrv/internal/database"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Context struct {
	USERDB *sqlx.DB
	DOCSDB *sqlx.DB
}

func NewContext(c echo.Context) (ctx Context) {
	ctx.USERDB, _ = database.InitUsersDatabase()
	ctx.DOCSDB, _ = database.InitDocsDatabase()
	return
}
