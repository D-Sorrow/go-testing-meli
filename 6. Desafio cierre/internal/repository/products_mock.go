package repository

import "app/internal"

type ProductMock struct {
	Call               int
	SearchProductsFunc func(query internal.ProductQuery) (p map[int]internal.Product, err error)
}

func (r *ProductMock) SearchProducts(query internal.ProductQuery) (p map[int]internal.Product, err error) {
	r.Call++
	return r.SearchProductsFunc(query)
}
