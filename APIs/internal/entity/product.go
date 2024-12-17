package entity

import (
	"errors"
	"time"

	"github.com/DiogoFiuza/learning-golang/APIs/pkg/clock"
	"github.com/DiogoFiuza/learning-golang/APIs/pkg/entity"
)

type Product struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Price    int       `json:"price"`
	CreateAt time.Time `json:"create_at"`
}

var (
	ErrIDIsRequired    = errors.New("id is required")
	ErrIDInvalid       = errors.New("id is invalid")
	ErrNameIsRequired  = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrInvalidPrice    = errors.New("price is invalid")
)

func NewProduct(name string, price int, c clock.Clock) (*Product, error) {
	product := &Product{
		ID:       entity.NewID(),
		Name:     name,
		Price:    price,
		CreateAt: c.Now(),
	}

	err := product.Validate()

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}

	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrIDInvalid
	}

	if p.Name == "" {
		return ErrNameIsRequired
	}

	if p.Price == 0 {
		return ErrPriceIsRequired
	}

	if p.Price < 0 {
		return ErrInvalidPrice
	}

	return nil
}
