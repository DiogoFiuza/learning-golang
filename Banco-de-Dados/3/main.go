package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm: "primaryKey; autoIncrement:true"`
	Name string
}

type Product struct {
	ID         int `gorm:"primaryKey; autoIncrement:true`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
}

// BelongsTo
func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	// Create category
	//category := Category{Name: "Eletronics"}
	//db.Create(&category)

	// Create product
	//product := Product{Name: "Smartphone", Price: 1000, CategoryID: category.ID}
	//db.Create(&product)

	// Read product
	//var products []Product
	//db.Preload("Category").Find(&products)
	//for _, product := range products {
	//	fmt.Println("Product:", product.Name, "Category:", product.Category.Name)
	//}
}
