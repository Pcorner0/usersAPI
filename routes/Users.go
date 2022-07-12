package routes

import (
	"github.com/labstack/echo/v4"
	controller "github.com/pcorner0/usersAPI/controllers"
	"net/http"
)

func Users(e *echo.Echo) {
	U := e.Group("/users")
	userRepo := controller.NewRepoUsers()
	U.POST("/create", userRepo.CreateUser)
	U.GET("/get", userRepo.GetUser)
	U.DELETE("/delete", func(c echo.Context) error {
		name := c.QueryParam("id")
		return c.String(http.StatusOK, name)
	})
}
