package product

import (
	"net/http"
	"strconv"

	"product-api/common/dto"
	"product-api/common/util"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService ProductService
}

func NewProductController(p ProductService) ProductController {
	return ProductController{ProductService: p}
}

func (p *ProductController) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product := p.ProductService.FindById(uint(id))

	c.JSON(http.StatusOK, gin.H{"product": ToProductDto(product)})
}

func (p *ProductController) Index(c *gin.Context) {
	var args dto.PaginationDTO

	args.Sort = c.DefaultQuery("sort", "ID")
	args.Order = c.DefaultQuery("order", "DESC")
	args.Offset = c.DefaultQuery("offset", "0")
	args.Limit = c.DefaultQuery("limit", "25")
	args.Search = c.DefaultQuery("search", "")

	products, filteredData, totalData, err := p.ProductService.FindAllPaginate(args)

	if err != nil {
		c.AbortWithStatus(404)
	}

	data := PaginationProductData{
		TotalData:    totalData,
		FilteredData: filteredData,
		Data:         ToProductDtos(products),
	}

	response := dto.BaseResponseDTO("Get Product List Success.", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}

func (p *ProductController) FindAll(c *gin.Context) {
	products := p.ProductService.FindAll()

	response := dto.BaseResponseDTO("Get Product List Success.", http.StatusOK, "success", products)

	c.JSON(http.StatusOK, response)
}

func (p *ProductController) Create(c *gin.Context) {
	var productDTO ProductDTO

	err := c.ShouldBindJSON(&productDTO)

	if err != nil {
		errors := util.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := dto.BaseResponseDTO("Failed to create product", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	createdProduct, err := p.ProductService.Create(ToProduct(productDTO))

	if err != nil {
		c.AbortWithStatus(400)
	}

	response := dto.BaseResponseDTO("Create Product Success.", http.StatusOK, "success", ToProductDto(createdProduct))

	c.JSON(http.StatusOK, response)
}

func (p *ProductController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product := p.ProductService.FindById(uint(id))

	if product == (Product{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	p.ProductService.Delete(product)

	c.Status(http.StatusOK)
}
