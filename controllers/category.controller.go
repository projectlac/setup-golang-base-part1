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

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController(DB *gorm.DB) CategoryController {
	return CategoryController{DB}
}

func (pc *CategoryController) CreateCategory(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.CreateCategoryRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	newPost := models.Category{
		CategoryName: payload.CategoryName,
		User:         currentUser.ID,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	result := pc.DB.Create(&newPost)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Post with that title already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newPost})
}

// [...] Update Post Handler
func (pc *CategoryController) UpdateCategory(ctx *gin.Context) {
	categoryId := ctx.Param("categoryId")
	currentUser := ctx.MustGet("currentUser").(models.User)

	var payload *models.UpdateCategory
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updatedCategory models.Category
	result := pc.DB.First(&updatedCategory, "id = ?", categoryId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}
	now := time.Now()
	categoryToUpdate := models.Category{
		CategoryName: payload.CategoryName,
		User:         currentUser.ID,
		CreatedAt:    updatedCategory.CreatedAt,
		UpdatedAt:    now,
	}

	pc.DB.Model(&updatedCategory).Updates(categoryToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedCategory})
}

// [...] Get Single Post Handler
func (pc *CategoryController) FindCategoryById(ctx *gin.Context) {
	categoryId := ctx.Param("categoryId")

	var category models.Category
	result := pc.DB.First(&category, "id = ?", categoryId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": category})
}

// [...] Get All Posts Handler
func (pc *CategoryController) FindCategories(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var categories []models.Category
	results := pc.DB.Limit(intLimit).Offset(offset).Find(&categories)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(categories), "data": categories})
}

// [...] Delete Post Handler
func (pc *CategoryController) DeleteCategory(ctx *gin.Context) {
	categoryId := ctx.Param("categoryId")

	result := pc.DB.Delete(&models.Category{}, "id = ?", categoryId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
