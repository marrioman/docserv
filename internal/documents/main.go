package documents

import (
	"docsevrv/internal/common"
	"fmt"
	"github.com/labstack/gommon/log"
)

func CreateDocument(ctx common.Context, id, password, doctype, docname, document string) (err error) {
	err = checkAuth(ctx, id, password)
	if err != nil {
		log.Error(err)
		return
	}

	sqlStatement := `INSERT INTO docs.documents_ltree (user_id, doc_type, document_name, document) VALUES ($1, $2, $3, $4)`
	_, err = ctx.DOCSDB.Exec(sqlStatement, id, doctype, docname, document)
	if err != nil {
		log.Error(err)
		return
	}

	return
}

func UpdateDocument(ctx common.Context, id, password, doctype, docname, document string) (err error) {
	err = checkAuth(ctx, id, password)
	if err != nil {
		log.Error(err)
		return
	}

	sqlStatement := `UPDATE docs.documents_ltree SET doc_type = $1, document_name = $2, document = $3 WHERE user_id = $4`
	_, err = ctx.DOCSDB.Exec(sqlStatement, doctype, docname, document, id)
	if err != nil {
		log.Error(err)
		return
	}

	return
}

func GetDocument(ctx common.Context, id, password string) (directions []DocumentsResponse, err error) {
	err = checkAuth(ctx, id, password)
	if err != nil {
		log.Error(err)
		return
	}

	seldirection := make([]DocumentsResponse, 0)

	sqlStatement := `SELECT user_id, doc_type, document_name, document FROM docs.documents_ltree WHERE user_id = $1`
	err = ctx.DOCSDB.Select(&seldirection, sqlStatement, id)
	if err != nil {
		log.Error(err)
		return
	}

	directions = seldirection

	return
}

func checkAuth(ctx common.Context, id, password string) (err error) {
	sqlStatement := `SELECT user_name FROM users.user_profile WHERE id = $1 and password = $2 LIMIT 1`
	result, err := ctx.USERDB.Exec(sqlStatement, id, password)
	if err != nil {
		log.Error(err)
		return
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected != 1 {
		err = fmt.Errorf("wrong pass/no such user")
		return
	}

	return
}
