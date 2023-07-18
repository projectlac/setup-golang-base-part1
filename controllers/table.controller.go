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

type TableController struct {
	DB *gorm.DB
}

func NewTableController(DB *gorm.DB) TableController {
	return TableController{DB}
}

func (pc *TableController) CreateTable(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.CreateTableRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	newPost := models.Table{
		TableName: payload.TableName,
		User:      currentUser.ID,
		CreatedAt: now,
		UpdatedAt: now,
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
func (pc *TableController) UpdateTable(ctx *gin.Context) {
	tableId := ctx.Param("tableId")
	currentUser := ctx.MustGet("currentUser").(models.User)

	var payload *models.UpdateTable
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updatedTable models.Table
	result := pc.DB.First(&updatedTable, "id = ?", tableId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}
	now := time.Now()
	categoryToUpdate := models.Table{
		TableName: payload.TableName,
		User:      currentUser.ID,
		CreatedAt: updatedTable.CreatedAt,
		UpdatedAt: now,
	}

	pc.DB.Model(&updatedTable).Updates(categoryToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedTable})
}

// [...] Get Single Post Handler
func (pc *TableController) FindTableById(ctx *gin.Context) {
	tableId := ctx.Param("tableId")

	var category models.Table
	result := pc.DB.First(&category, "id = ?", tableId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": category})
}

// [...] Get All Posts Handler
func (pc *TableController) FindTables(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var categories []models.Table
	results := pc.DB.Limit(intLimit).Offset(offset).Find(&categories)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(categories), "data": categories})
}

// [...] Delete Post Handler
func (pc *TableController) DeleteTable(ctx *gin.Context) {
	tableId := ctx.Param("tableId")

	result := pc.DB.Delete(&models.Table{}, "id = ?", tableId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
