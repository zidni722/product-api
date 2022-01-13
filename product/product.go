package product

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Price       uint
	Description string
	Quantity    int
}
