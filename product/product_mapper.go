package product

func ToProduct(productDTO ProductDTO) Product {
	return Product{
		Name:        productDTO.Name,
		Price:       productDTO.Price,
		Description: productDTO.Description,
		Quantity:    productDTO.Quantity,
	}
}

func ToProductDtos(products []Product) []ProductDTO {
	productsDtos := make([]ProductDTO, len(products))

	for i, product := range products {
		productsDtos[i] = ToProductDto(product)
	}

	return productsDtos
}

func ToProductDto(product Product) ProductDTO {
	productDto := ProductDTO{}
	productDto.ID = product.ID
	productDto.Name = product.Name
	productDto.Price = product.Price
	productDto.Description = product.Description
	productDto.Quantity = product.Quantity
	productDto.CreatedAt = product.CreatedAt
	productDto.UpdatedAt = product.UpdatedAt
	productDto.DeletedAt = product.DeletedAt

	return productDto
}
