package product

import (
	"product-api/common/dto"
	"product-api/common/util"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepostiory(DB *gorm.DB) ProductRepository {
	return ProductRepository{DB: DB}
}

func (p *ProductRepository) FindById(id uint) Product {
	var product Product
	p.DB.Find(&product, id)

	return product
}

func (p *ProductRepository) FindAll() []Product {
	var products []Product
	p.DB.Find(&products)

	return products
}

func (p *ProductRepository) FindAllPaginate(args dto.PaginationDTO) ([]Product, int64, error) {
	products := []Product{}
	var totalData int64

	table := "products"
	query := p.DB.Select(table + ".*")
	query = query.Offset(util.Offset(args.Offset))
	query = query.Limit(util.Limit(args.Limit))
	query = query.Order(util.SortOrder(table, args.Sort, args.Order))
	query = query.Scopes(util.Search(args.Search))

	if err := query.Find(&products).Error; err != nil {
		return products, totalData, err
	}

	p.DB.Table(table).Count(&totalData)

	return products, totalData, nil
}

func (p *ProductRepository) Create(product Product) (Product, error) {
	err := p.DB.Create(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (p *ProductRepository) Delete(product Product) {
	p.DB.Delete(&product)
}
