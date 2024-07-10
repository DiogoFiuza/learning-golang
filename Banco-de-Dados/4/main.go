package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm: "primaryKey; autoIncrement:true"`
	Name string
}

type Product struct {
	ID           int `gorm:"primaryKey; autoIncrement:true`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey; autoIncrement:true`
	Number    string
	ProductID int
}

// HasOne
func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{}, SerialNumber{})

	//Create category
	category := Category{Name: "Eletronics"}
	db.Create(&category)

	//Create product
	product := Product{Name: "Smartphone", Price: 1000, CategoryID: category.ID}
	db.Create(&product)

	db.Create(&SerialNumber{Number: "123456", ProductID: product.ID})

	//Read product
	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, product := range products {
		fmt.Println("Product:", product.Name, "Category:", product.Category.Name, "Serial Number", product.SerialNumber.ID)
	}
}
