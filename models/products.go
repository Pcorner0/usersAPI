package models

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	name  string
	price float64
}

func CreateProduct(db *gorm.DB, Product *Products) error {
	err := db.Create(Product).Error
	if err != nil {
		return err
	}
	return nil
}

func GetProduct(db *gorm.DB, product *Products, id string) (err error) {
	db.Raw("SELECT * FROM products WHERE id = ? AND deleted_at IS NULL LIMIT 1", id).Scan(product)
	if err != nil {
		return err
	}
	return nil
}

// GetAllProducts function returns all products
func GetAllProducts(db *gorm.DB, products *[]Products) (err error) {
	db.Raw("SELECT * FROM products WHERE deleted_at IS NULL ORDER BY created_at DESC").Scan(products)
	if err != nil {
		return err
	}
	return nil
}

// DeleteProduct function deletes a product by id
func DeleteProduct(db *gorm.DB, product *Products, id string) (err error) {
	db.Where("id = ?", id).Delete(product)
	if err != nil {
		return err
	}
	return nil
}

// UpdateProduct function updates a product by id
func UpdateProduct(db *gorm.DB, product *Products, id string) (err error) {
	result := db.Model(product).Where("id = ?", id).Updates(product)
	err = result.Error
	if err != nil {
		return result.Error
	}
	return nil
}
