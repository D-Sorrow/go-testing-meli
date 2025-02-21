package repository_test

import (
	"app/internal"
	"app/internal/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductsMap_CreateWithNil(t *testing.T) {
	productsMap := repository.NewProductsMap(nil)
	assert.NotNil(t, productsMap)
}
func TestProductsMap_Create(t *testing.T) {
	mockProduct := map[int]internal.Product{
		1: internal.Product{
			Id: 1,
			ProductAttributes: internal.ProductAttributes{
				Description: "Hola",
				Price:       12.4,
				SellerId:    1,
			},
		},
	}
	productsMap := repository.NewProductsMap(mockProduct)
	assert.NotNil(t, productsMap)
}

func TestProductsMap_SearchProducts(t *testing.T) {
	id := internal.ProductQuery{Id: 1}
	mockProduct := map[int]internal.Product{
		1: internal.Product{
			Id: 1,
			ProductAttributes: internal.ProductAttributes{
				Description: "Hola",
				Price:       12.4,
				SellerId:    1,
			},
		},
	}

	productsMap := repository.NewProductsMap(mockProduct)
	product, _ := productsMap.SearchProducts(id)

	assert.NotNil(t, product)
	assert.Equal(t, id.Id, product[1].Id)
}
