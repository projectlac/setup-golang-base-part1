package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/projectlac/golang-gorm-postgres/models"
	"gorm.io/gorm"
)

type ProductController struct {
	DB *gorm.DB
}

func NewProductController(DB *gorm.DB) ProductController {
	return ProductController{DB}
}

func (pc *ProductController) CreateProduct(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.CreateProductRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	newProduct := models.Product{
		Name:      payload.Name,
		User:      currentUser.ID,
		Image:     payload.Image,
		Price:     payload.Price,
		Category:  payload.Category,
		CreatedAt: now,
		UpdatedAt: now,
	}
	result := pc.DB.Create(&newProduct)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Post with that title already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newProduct})
}

// [...] Update Post Handler
func (pc *ProductController) UpdateProduct(ctx *gin.Context) {
	productId := ctx.Param("productId")
	currentUser := ctx.MustGet("currentUser").(models.User)

	var payload *models.UpdateProduct
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updatedProduct models.Product
	result := pc.DB.First(&updatedProduct, "id = ?", productId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}
	now := time.Now()
	productToUpdate := models.Product{
		Name:      payload.Name,
		User:      currentUser.ID,
		CreatedAt: updatedProduct.CreatedAt,
		UpdatedAt: now,
	}

	pc.DB.Model(&updatedProduct).Updates(productToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedProduct})
}

// [...] Get Single Post Handler
func (pc *ProductController) FindProductById(ctx *gin.Context) {
	productId := ctx.Param("productId")

	var product models.Product
	result := pc.DB.First(&product, "id = ?", productId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": product})
}

// [...] Get All Posts Handler
func (pc *ProductController) FindProducts(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")
	var order = ctx.DefaultQuery("order", "asc")
	categoryId := ctx.Param("categoryId")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var products []models.Product

	query := pc.DB

	if categoryId != "" {
		query = query.Where("category_id = ?", categoryId)
	}
	if order == "desc" {
		query = query.Order("price desc")
	} else {
		query = query.Order("price asc")
	}

	results := query.Limit(intLimit).Offset(offset).Find(&products)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(products), "data": products})
}

// [...] Delete Post Handler
func (pc *ProductController) DeleteProduct(ctx *gin.Context) {
	productId := ctx.Param("productId")

	result := pc.DB.Delete(&models.Product{}, "id = ?", productId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
