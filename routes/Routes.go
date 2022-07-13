package routes

import (
	"github.com/labstack/echo/v4"
	controller "github.com/pcorner0/usersAPI/controllers"
)

func MyRoutes(e *echo.Echo) {
	U := e.Group("/users")
	userRepo := controller.NewRepoUsers()
	U.POST("/create", userRepo.CreateUser)
	U.GET("/get", userRepo.GetUser)
	U.GET("/getall", userRepo.GetAllUsers)
	U.DELETE("/delete", userRepo.DeleteUser)
	U.PUT("/update", userRepo.UpdateUser)

	P := e.Group("/products")
	productRepo := controller.NewRepoProducts()
	P.POST("/create", productRepo.CreateProduct)
	P.GET("/get", productRepo.GetProduct)
	P.GET("/getall", productRepo.GetAllProducts)
	P.DELETE("/delete", productRepo.DeleteProduct)
	P.PUT("/update", productRepo.UpdateProduct)

	S := e.Group("/sales")
	saleRepo := controller.NewRepoSales()
	S.POST("/create", saleRepo.CreateSale)
	S.GET("/get", saleRepo.GetSale)
	S.GET("/getall", saleRepo.GetAllSales)
	S.DELETE("/delete", saleRepo.DeleteSale)
	S.PUT("/update", saleRepo.UpdateSale)

}
