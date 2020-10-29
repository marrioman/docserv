package controller

import (
	"github.com/labstack/echo/v4"
)

func Add(api *echo.Group) {

	api.POST("/documents/create", createDoc)
	api.POST("/documents/update", updateDoc)
	//api.POST("/documents/delete", deleteDoc)
	api.POST("/documents/getdocs", getDoc)

	api.POST("/user/create", createProfile)
	//api.POST("/user/update", updateProfile)
	//api.PUT("/user/delete", deleteProfile)
	//api.GET("/user/profile", getProfile)

	return
}
