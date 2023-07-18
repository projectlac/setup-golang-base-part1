package main

import (
	"fmt"
	"log"

	"github.com/projectlac/golang-gorm-postgres/initializers"
	"github.com/projectlac/golang-gorm-postgres/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{}, &models.TableOrder{}, &models.Product{}, &models.Category{}, &models.Order{})
	fmt.Println("? Migration complete")
}
