package entity

import (
	"errors"
	"testing"
	"time"

	"github.com/DiogoFiuza/learning-golang/APIs/pkg/clock"
	"github.com/DiogoFiuza/learning-golang/APIs/pkg/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {

	fc := clock.FakeClock{}

	type NewProductPrams struct {
		name  string
		price float64
		c     clock.Clock
	}

	tests := []struct {
		name            string
		product         *NewProductPrams
		expectedProduct *Product
		err             error
	}{
		{
			name: "OK",
			product: &NewProductPrams{
				name:  "Bicicleta",
				price: 999.00,
				c:     fc,
			},
			expectedProduct: &Product{
				ID:        entity.NewID(),
				Name:      "Bicicleta",
				Price:     999.00,
				CreatedAt: fc.Now(),
			},
			err: nil,
		},
		{
			name: "Fail/Error name is required",
			product: &NewProductPrams{
				name:  "",
				price: 999.00,
				c:     fc,
			},
			err: errors.New("name is required"),
		},
	}

	for _, tt := range tests {
		product, err := NewProduct(tt.product.name, tt.product.price, tt.product.c)
		if err != nil {
			assert.Equal(t, tt.err, err)
			continue
		}
		tt.expectedProduct.ID = product.ID
		assert.Equal(t, product, tt.expectedProduct)

	}
}

func TestProduct_Validate(t *testing.T) {
	type fields struct {
		ID       entity.ID
		Name     string
		Price    float64
		CreateAt time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		{
			name: "OK",
			fields: fields{
				ID:       entity.NewID(),
				Name:     "Bicicleta",
				Price:    999.00,
				CreateAt: time.Now(),
			},
			wantErr: nil,
		},
		{
			name: "Fail/Error name is required",
			fields: fields{
				ID:       entity.NewID(),
				Name:     "",
				Price:    999.00,
				CreateAt: time.Now(),
			},
			wantErr: errors.New("name is required"),
		},
		{
			name: "Fail/Error price is required",
			fields: fields{
				ID:       entity.NewID(),
				Name:     "Bicicleta",
				Price:    0,
				CreateAt: time.Now(),
			},
			wantErr: errors.New("price is required"),
		},
		{
			name: "Fail/Error price is invalid",
			fields: fields{
				ID:       entity.NewID(),
				Name:     "Bicicleta",
				Price:    -1,
				CreateAt: time.Now(),
			},
			wantErr: errors.New("price is invalid"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Product{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				Price:     tt.fields.Price,
				CreatedAt: tt.fields.CreateAt,
			}
			err := p.Validate()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
