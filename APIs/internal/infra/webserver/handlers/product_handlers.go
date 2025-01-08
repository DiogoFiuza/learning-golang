package handlers

import (
	"encoding/json"
	"github.com/DiogoFiuza/learning-golang/APIs/internal/dto"
	"github.com/DiogoFiuza/learning-golang/APIs/internal/entity"
	"github.com/DiogoFiuza/learning-golang/APIs/internal/infra/database"
	"github.com/DiogoFiuza/learning-golang/APIs/pkg/clock"
	"net/http"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (p *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductDTO
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newProduct, err := entity.NewProduct(product.Name, product.Price, clock.NewClock())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = p.ProductDB.Create(newProduct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
