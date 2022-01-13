package product

import "product-api/common/dto"

type ProductService struct {
	ProductRepository ProductRepository
}

func NewProductService(p ProductRepository) ProductService {
	return ProductService{ProductRepository: p}
}

func (p *ProductService) FindById(id uint) Product {
	return p.ProductRepository.FindById(id)
}

func (p *ProductService) FindAll() []Product {
	return p.ProductRepository.FindAll()
}

func (p *ProductService) Create(product Product) (Product, error) {
	product, err := p.ProductRepository.Create(product)

	if err != nil {
		return product, err
	}

	return product, nil
}

func (p *ProductService) Delete(product Product) {
	p.ProductRepository.Delete(product)
}

func (p *ProductService) FindAllPaginate(args dto.PaginationDTO) ([]Product, int, int64, error) {
	products, totalData, err := p.ProductRepository.FindAllPaginate(args)

	filteredData := len(products)

	return products, filteredData, totalData, err
}
