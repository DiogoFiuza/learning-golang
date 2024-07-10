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

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	//db.Create(&Product{
	//	Name:  "Laptop",
	//	Price: 1000.0,
	//})

	// Create batch
	//db.Create([]*Product{
	//	{Name: "Laptop", Price: 1000.0},
	//	{Name: "Mouse", Price: 50.0},
	//	{Name: "Keyboard", Price: 100.0},
	//})

	// Select one
	//var product Product
	//db.First(&product, "name = ?", "Laptop")
	//fmt.Print(product)

	// Select all
	//var products []Product
	//db.Find(&products)
	//
	//for _, product := range products {
	//	fmt.Println(product)
	//}

	// Working with limit and offset
	//var products []Product
	//db.Limit(2).Offset(0).Find(&products)
	//for _, products := range products {
	//	fmt.Println(products)
	//}

	// Where
	//var products []Product
	//db.Where("price > ?", 100).Find(&products)
	//for _, products := range products {
	//	fmt.Println(products)
	//}

	// Like
	//var products []Product
	//db.Where("name LIKE ?", "%top%").Find(&products)
	//for _, products := range products {
	//	fmt.Println(products)
	//}

	//Update values
	//var p Product
	//db.First(&p)
	//p.Name = "Laptop 2"
	//db.Save(&p)

	// Delete value
	//var p2 Product
	//db.First(&p2, "name = ?", "Laptop 2")
	//fmt.Println(p2)
	//
	//db.Delete(&p2)

	// Delete to clean the table
	//db.Exec("Delete from products")
}
