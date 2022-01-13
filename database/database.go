package database

import (
	"fmt"
	"os"
	"product-api/product"
	"strconv"

	"github.com/jinzhu/gorm"
)

func InitDB() *gorm.DB {
	port, _ := strconv.Atoi(os.Getenv("APP_DB_PORT"))

	db, err := gorm.Open(os.Getenv("APP_DB_DRIVER"), fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("APP_DB_USER"),
		os.Getenv("APP_DB_PASS"),
		os.Getenv("APP_DB_HOST"),
		port,
		os.Getenv("APP_DB_NAME"),
	))

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&product.Product{})

	return db
}