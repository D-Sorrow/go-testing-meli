package handler_test

import (
	"app/internal"
	"app/internal/handler"
	"app/internal/repository"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProducts_Success(t *testing.T) {
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

	mockService := repository.ProductMock{
		SearchProductsFunc: func(query internal.ProductQuery) (p map[int]internal.Product, err error) {
			return mockProduct, nil
		},
	}

	handler := handler.NewProductsDefault(&mockService)

	router := chi.NewRouter()
	router.Get("/products", handler.Get())

	req := httptest.NewRequest("GET", "/products", nil)
	rec := httptest.NewRecorder()
	stringResponse := `{"data":{"1":{"id":1,"description":"Hola","price":12.4,"seller_id":1}},"message":"success"}`

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, stringResponse, rec.Body.String())
	assert.EqualValues(t, 1, mockService.Call)

}

func TestProducts_SuccessById(t *testing.T) {
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

	mockService := repository.ProductMock{
		SearchProductsFunc: func(query internal.ProductQuery) (p map[int]internal.Product, err error) {
			return mockProduct, nil
		},
	}

	handler := handler.NewProductsDefault(&mockService)

	router := chi.NewRouter()
	router.Get("/products", handler.Get())

	req := httptest.NewRequest("GET", "/products?id=1", nil)
	rec := httptest.NewRecorder()
	stringResponse := `{"data":{"1":{"id":1,"description":"Hola","price":12.4,"seller_id":1}},"message":"success"}`

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, stringResponse, rec.Body.String())
	assert.EqualValues(t, 1, mockService.Call)

}

func TestProducts_InvalidId(t *testing.T) {

	mockService := repository.ProductMock{
		SearchProductsFunc: func(query internal.ProductQuery) (p map[int]internal.Product, err error) {
			return nil, nil
		},
	}

	handler := handler.NewProductsDefault(&mockService)

	router := chi.NewRouter()
	router.Get("/products", handler.Get())

	req := httptest.NewRequest("GET", "/products?id=A1", nil)
	rec := httptest.NewRecorder()
	stringResponse := `{"status":"Bad Request","message":"invalid id"}`

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, stringResponse, rec.Body.String())
	assert.EqualValues(t, 0, mockService.Call)

}

func TestProducts_GetError(t *testing.T) {

	mockService := repository.ProductMock{
		SearchProductsFunc: func(query internal.ProductQuery) (p map[int]internal.Product, err error) {
			return nil, errors.New("error")
		},
	}

	handler := handler.NewProductsDefault(&mockService)

	router := chi.NewRouter()
	router.Get("/products", handler.Get())

	req := httptest.NewRequest("GET", "/products", nil)
	rec := httptest.NewRecorder()
	stringResponse := `{"status":"Internal Server Error","message":"internal error"}`

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Equal(t, stringResponse, rec.Body.String())
	assert.EqualValues(t, 1, mockService.Call)

}
