package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm: "primaryKey; autoIncrement:true"`
	Name     string
	Products []Product
}

type Product struct {
	ID         int `gorm:"primaryKey; autoIncrement:true`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
}

// HasMany
func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	//Create category
	category := Category{Name: "Eletronics"}
	db.Create(&category)

	//Create product
	product := Product{Name: "Smartphone", Price: 1000, CategoryID: category.ID}
	db.Create(&product)

	var categories []Category

	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	//Read product
	for _, category := range categories {
		fmt.Println("Category:", category.Name)
		for _, product := range category.Products {
			fmt.Println("Product:", product.Name)
		}
	}
}
