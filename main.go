package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pcorner0/usersAPI/database"
	"github.com/pcorner0/usersAPI/routes"
)

func main() {
	e := echo.New()
	database.Connect()

	// Root level middleware
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	//Users routes
	routes.MyRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
