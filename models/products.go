package models

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	name  string
	price float64
}
