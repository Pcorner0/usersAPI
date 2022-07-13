package controllers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pcorner0/usersAPI/database"
	"github.com/pcorner0/usersAPI/models"
	"gorm.io/gorm"
)

type RepoSales struct {
	DB *gorm.DB
}

func NewRepoSales() *RepoSales {
	db := database.InitDB()
	err := db.AutoMigrate(&models.Sales{})
	if err != nil {
		return nil
	}
	return &RepoSales{DB: db}
}

// CreateSale method using the echo.Context
func (s *RepoSales) CreateSale(c echo.Context) error {
	Sale := &models.Sales{}

	if err := c.Bind(Sale); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return err
	}
	err := models.CreateSale(s.DB, Sale)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, Sale)
}

// GetSale shows a sale by id
func (s *RepoSales) GetSale(c echo.Context) error {
	id := c.FormValue("id")

	Sale := &models.Sales{}
	err := models.GetSale(s.DB, Sale, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, err)
		}
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, Sale)
}

// GetAllSales method using to get all sales
func (s *RepoSales) GetAllSales(c echo.Context) error {
	Sales := &[]models.Sales{}
	err := models.GetAllSales(s.DB, Sales)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, Sales)
}

// DeleteSale method using the echo.Context
func (s *RepoSales) DeleteSale(c echo.Context) error {
	id := c.FormValue("id")
	Sale := &models.Sales{}
	err := models.DeleteSale(s.DB, Sale, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, err)
		}
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, Sale)
}

// UpdateSale method using the echo.Context
func (s *RepoSales) UpdateSale(c echo.Context) error {
	id := c.FormValue("id")
	Sale := &models.Sales{}

	if err := c.Bind(Sale); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return err
	}
	err := models.UpdateSale(s.DB, Sale, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, Sale)
}
