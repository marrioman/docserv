package controller

import (
	"docsevrv/internal/common"
	"docsevrv/internal/profile"
	"github.com/labstack/echo/v4"
	"net/http"
)

func createProfile(c echo.Context) (err error) {
	var (
		req struct {
			UserName string `json:"user_name"`
			LastName string `json:"last_name"`
			Password string `json:"password"`
		}
		ctx = common.NewContext(c)
	)

	err = c.Bind(&req)
	if err != nil {
		return err
	}

	err = profile.CreateProfile(ctx, req.UserName, req.LastName, req.Password)
	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": "success",
	})
}
