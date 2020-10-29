package controller

import (
	"docsevrv/internal/common"
	"docsevrv/internal/documents"
	"github.com/labstack/echo/v4"
	"net/http"
)

func createDoc(c echo.Context) (err error) {
	var (
		req struct {
			ID       string `json:"id"`
			Password string `json:"password"`
			DocType  string `json:"doc_type"`
			DocName  string `json:"doc_name"`
			Document string `json:"document"`
		}
		ctx = common.NewContext(c)
	)

	err = c.Bind(&req)
	if err != nil {
		return err
	}

	err = documents.CreateDocument(ctx, req.ID, req.Password, req.DocType, req.DocName, req.Document)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": "success",
	})
}

func updateDoc(c echo.Context) (err error) {
	var (
		req struct {
			ID       string `json:"id"`
			Password string `json:"password"`
			DocType  string `json:"doc_type,omitempty"`
			DocName  string `json:"doc_name,omitempty"`
			Document string `json:"document,omitempty"`
		}
		ctx = common.NewContext(c)
	)

	err = c.Bind(&req)
	if err != nil {
		return err
	}

	err = documents.UpdateDocument(ctx, req.ID, req.Password, req.DocType, req.DocName, req.Document)
	if err != nil {
		return c.JSON(401, echo.Map{
			"data": "wrong pass/no such user",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": "success",
	})
}

func getDoc(c echo.Context) (err error) {
	var (
		req struct {
			ID       string `json:"id"`
			Password string `json:"password"`
		}
		ctx = common.NewContext(c)
	)

	err = c.Bind(&req)
	if err != nil {
		return err
	}

	response, err := documents.GetDocument(ctx, req.ID, req.Password)
	if err != nil {
		return c.JSON(401, echo.Map{
			"data": "wrong pass/no such user",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": response,
	})
}
