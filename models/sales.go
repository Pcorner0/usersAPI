package models

import (
	"gorm.io/gorm"
)

type Sales struct {
	gorm.Model
	IdUser    uint  `json:"idUser" gorm:"not null"`
	IdProduct uint  `json:"idProduct" gorm:"not null"`
	Amount    int64 `json:"amount" gorm:"not null"`
}

// CreateSale function creates a new sales
func CreateSale(db *gorm.DB, sales *Sales) error {
	err := db.Create(sales).Error
	if err != nil {
		return err
	}
	return nil
}

// GetSale function returns a sales by id
func GetSale(db *gorm.DB, sales *Sales, id string) (err error) {
	db.Raw("SELECT * FROM sales WHERE id = ? AND deleted_at IS NULL LIMIT 1", id).Scan(sales)
	if err != nil {
		return err
	}
	return nil
}

// GetAllSales function returns all sales
func GetAllSales(db *gorm.DB, sales *[]Sales) (err error) {
	db.Raw("SELECT * FROM sales WHERE deleted_at IS NULL ORDER BY created_at DESC").Scan(sales)
	if err != nil {
		return err
	}
	return nil
}

// DeleteSale function deletes a sales by id
func DeleteSale(db *gorm.DB, sales *Sales, id string) (err error) {
	db.Where("id = ?", id).Delete(sales)
	if err != nil {
		return err
	}
	return nil
}

// UpdateSale function updates a sales by id
func UpdateSale(db *gorm.DB, sales *Sales, id string) (err error) {
	result := db.Model(sales).Where("id = ?", id).Updates(sales)
	err = result.Error
	if err != nil {
		return result.Error
	}
	return nil
}
