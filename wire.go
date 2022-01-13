package main

import (
	"product-api/product"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func initProductController(db *gorm.DB) product.ProductController {
	wire.Build(product.NewProductRepostiory, product.NewProductService, product.NewProductController)

	return product.ProductController{}
}
