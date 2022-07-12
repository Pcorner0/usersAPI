package models

import "gorm.io/gorm"

type Sales struct {
	gorm.Model
	idUser    int64
	idProduct int64
	amount    int64
}
