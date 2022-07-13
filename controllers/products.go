package controllers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pcorner0/usersAPI/database"
	"github.com/pcorner0/usersAPI/models"
	"gorm.io/gorm"
)

type RepoProducts struct {
	DB *gorm.DB
}

func NewRepoProducts() *RepoProducts {
	db := database.InitDB()
	err := db.AutoMigrate(&models.Products{})
	if err != nil {
		return nil
	}
	return &RepoProducts{DB: db}
}

// CreateProduct method using the echo.Context
func (p *RepoProducts) CreateProduct(c echo.Context) error {
	Product := &models.Products{}

	if err := c.Bind(Product); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return err
	}
	err := models.CreateProduct(p.DB, Product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, Product)
}

// GetProduct shows a product by id
func (p *RepoProducts) GetProduct(c echo.Context) error {
	id := c.FormValue("id")

	Product := &models.Products{}
	err := models.GetProduct(p.DB, Product, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, err)
		}
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, Product)
}

// GetAllProducts method using to get all products
func (p *RepoProducts) GetAllProducts(c echo.Context) error {
	Products := &[]models.Products{}
	err := models.GetAllProducts(p.DB, Products)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, Products)
}

// DeleteProduct method using the echo.Context
func (p *RepoProducts) DeleteProduct(c echo.Context) error {
	id := c.FormValue("id")
	Product := &models.Products{}
	err := models.DeleteProduct(p.DB, Product, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, Product)
}

// UpdateProduct method using the echo.Context
func (p *RepoProducts) UpdateProduct(c echo.Context) error {
	id := c.FormValue("id")
	Product := &models.Products{}

	if err := c.Bind(Product); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return err
	}
	err := models.UpdateProduct(p.DB, Product, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, Product)
}
