package controller

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pcorner0/usersAPI/database"
	"github.com/pcorner0/usersAPI/models"
	"gorm.io/gorm"
	"net/http"
)

type RepoUsers struct {
	DB *gorm.DB
}

func NewRepoUsers() *RepoUsers {
	db := database.InitDB()
	err := db.AutoMigrate(&models.Users{})
	if err != nil {
		return nil
	}
	return &RepoUsers{DB: db}
}

// CreateUser method CreateUser using the echo.Context
func (u *RepoUsers) CreateUser(c echo.Context) error {
	User := &models.Users{}
	if err := c.Bind(User); err != nil {
		return err
	}
	err := models.CreateUser(u.DB, User)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, User)
}

/*func (u *RepoUsers) CreateUser(c echo.Context) error {
	User := &models.Users{}

	if err := c.Bind(User); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return err
	}

	err := models.CreateUser(u.DB, User)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, User)
}*/

// GetUser shows a user by id
func (u *RepoUsers) GetUser(c echo.Context) error {
	id := c.Param("id")
	User := &models.Users{}
	err := models.GetUser(u.DB, User, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, err)
		}
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, User)
}

func (u *RepoUsers) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	fmt.Println(id, "leeeento")
	User := &models.Users{}

	err := models.DeleteUser(u.DB, User, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, User)
}
