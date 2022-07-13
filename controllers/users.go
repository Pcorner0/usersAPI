package controllers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pcorner0/usersAPI/database"
	"github.com/pcorner0/usersAPI/models"
	"github.com/pcorner0/usersAPI/utils"
	"gorm.io/gorm"
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
		c.JSON(http.StatusBadRequest, err)
		return err
	}

	// encrypt the password
	encriptedPassword, err := utils.HashPassword(User.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	User.Password = encriptedPassword
	err = models.CreateUser(u.DB, User)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, User)
}

// GetUser shows a user by id
func (u *RepoUsers) GetUser(c echo.Context) error {
	id := c.FormValue("id")

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

// GetAllUsers method using to get all users
func (u *RepoUsers) GetAllUsers(c echo.Context) error {
	Users := &[]models.Users{}
	err := models.GetUsers(u.DB, Users)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, Users)
}

// DeleteUser method using to delete a user by id
func (u *RepoUsers) DeleteUser(c echo.Context) error {
	id := c.FormValue("id")
	User := &models.Users{}

	err := models.DeleteUser(u.DB, User, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, User)
}

// UpdateUser method using to update a user by id
func (u *RepoUsers) UpdateUser(c echo.Context) error {

	id := c.FormValue("id")
	User := &models.Users{}

	if err := c.Bind(User); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return err
	}

	err := models.UpdateUser(u.DB, User, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, User)
}
