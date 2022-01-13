package main

import (
	"fmt"
	"log"
	"os"
	"product-api/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	err := godotenv.Load(".env")

    if err != nil {
        log.Fatal("Error loading .env file")
    }

	db := database.InitDB()
	defer db.Close()

	port := fmt.Sprintf(":%s",os.Getenv("PORT"))

	productController := InitProductController(db)

	r := gin.Default()

	r.GET("/products", productController.FindAll)
	r.POST("/products", productController.Create)
	r.GET("/index", productController.Index)

	err = r.Run(port)

	if err != nil {
		panic(err)
	}
}
