package database

import (
	"fmt"
	"github.com/DiogoFiuza/learning-golang/APIs/internal/entity"
	"github.com/DiogoFiuza/learning-golang/APIs/pkg/clock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
	"time"
)

// This test has some lines that repeat, we can extract this lines to a function, but I repeat this lines to exercise the code
// Because I believe the repetition is good for memorization

func TestProduct_Create(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Smartphone", 1000.00, clock.FakeClock{})
	if err != nil {
		t.Error(err)
	}
	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)

	var p entity.Product
	db.First(&p, "name = ?", "Smartphone")

	assert.Equal(t, product.ID, p.ID)
	assert.Equal(t, product.Name, p.Name)
	assert.Equal(t, product.Price, p.Price)
	assert.Equal(t, product.CreatedAt.Format(time.RFC3339), p.CreatedAt.Format(time.RFC3339))
}

func TestProduct_FindAll(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	productDB := NewProduct(db)
	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprint("Product ", i), 1000.00, clock.FakeClock{})
		if err != nil {
			t.Error(err)
		}
		productDB.Create(product)
	}
	products, err := productDB.FindAll(1, 10, "asc")
	assert.Len(t, products, 10)
	assert.NoError(t, err)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	productsPage2, err := productDB.FindAll(2, 10, "asc")
	assert.Len(t, productsPage2, 10)
	assert.NoError(t, err)
	assert.Equal(t, "Product 11", productsPage2[0].Name)
	assert.Equal(t, "Product 20", productsPage2[9].Name)

	productsPage3, err := productDB.FindAll(3, 10, "asc")
	assert.Len(t, productsPage3, 3)
	assert.NoError(t, err)
	assert.Equal(t, "Product 21", productsPage3[0].Name)
	assert.Equal(t, "Product 23", productsPage3[2].Name)
}

func TestProduct_FindByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	productDB := NewProduct(db)

	product, err := entity.NewProduct("Notebook", 1000.00, clock.FakeClock{})
	if err != nil {
		t.Error(err)
	}
	productDB.Create(product)

	productFound, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestProduct_Update(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	productDB := NewProduct(db)

	product, err := entity.NewProduct("Notebook", 1000.00, clock.FakeClock{})
	if err != nil {
		t.Error(err)
	}
	productDB.Create(product)

	product2, err := entity.NewProduct("Notebook", 500.00, clock.FakeClock{})
	assert.NoError(t, err)
	err = productDB.Update(*product2)
	assert.NoError(t, err)
	var p entity.Product
	db.First(&p, "id = ?", product.ID)

	assert.NoError(t, err)
	assert.Equal(t, p.ID, product.ID)
	assert.Equal(t, p.Name, product2.Name)
}

func TestProduct_Delete(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Smartphone", 1000.00, clock.FakeClock{})
	if err != nil {
		t.Error(err)
	}
	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)

	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)

	_, err = productDB.FindByID(product.ID.String())
	assert.Error(t, err)
}
