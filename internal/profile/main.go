package profile

import (
	"docsevrv/internal/common"
	"github.com/labstack/gommon/log"
)

func CreateProfile(ctx common.Context, username, lastname, password string) (err error) {
	sqlStatement := `INSERT INTO users.user_profile (user_name, last_name, password) VALUES ($1, $2, $3)`
	result, err := ctx.USERDB.Exec(sqlStatement, username, lastname, password)
	if err != nil {
		log.Error(err)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected != 1 {
		log.Infof("0 rows affected(")
		return
	}

	return
}
