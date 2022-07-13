package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username string `json:"username" `
	Email    string `json:"email"`
	Password string `json:"password"`
	Rol      string `json:"rol"`
}

// CreateUser function creates a new user
func CreateUser(db *gorm.DB, User *Users) error {
	err := db.Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

// GetUser function returns a user by id
func GetUser(db *gorm.DB, User *Users, id string) (err error) {
	db.Raw("SELECT * FROM users WHERE id = ? AND deleted_at IS NULL LIMIT 1", id).Scan(User)
	if err != nil {
		return err
	}
	return nil
}

// GetUsers function returns all users
func GetUsers(db *gorm.DB, users *[]Users) (err error) {
	db.Raw("SELECT * FROM users WHERE deleted_at IS NULL ORDER BY created_at DESC").Scan(users)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser function deletes a user by id
func DeleteUser(db *gorm.DB, User *Users, id string) (err error) {
	fmt.Println(id)
	db.Where("id = ?", id).Delete(User)
	if err != nil {
		return err
	}
	return nil
}

// UpdateUser function updates a user by id
func UpdateUser(db *gorm.DB, User *Users, id string) (err error) {
	result := db.Model(User).Where("id = ?", id).Updates(User)
	err = result.Error
	if err != nil {
		return result.Error
	}
	return nil
}
