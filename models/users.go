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

func CreateUser(db *gorm.DB, User *Users) error {
	err := db.Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUser(db *gorm.DB, User *Users, id string) (err error) {
	db.Raw("SELECT * FROM users WHERE id = ? AND deleted_at IS NULL LIMIT 1", id).Scan(User)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(db *gorm.DB, User *Users, id string) (err error) {
	fmt.Println(id)
	db.Where("id = ?", id).Delete(User)
	if err != nil {
		return err
	}
	return nil
}
