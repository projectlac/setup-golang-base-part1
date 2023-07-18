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

type TableOrderController struct {
	DB *gorm.DB
}

func NewTableOrderController(DB *gorm.DB) TableOrderController {
	return TableOrderController{DB}
}

func (pc *TableOrderController) CreateTableOrder(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.CreateTableOrderRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	newPost := models.TableOrder{
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
func (pc *TableOrderController) UpdateTableOrder(ctx *gin.Context) {
	tableOrderId := ctx.Param("tableOrderId")
	currentUser := ctx.MustGet("currentUser").(models.User)

	var payload *models.UpdateTableOrder
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updatedTableOrder models.TableOrder
	result := pc.DB.First(&updatedTableOrder, "id = ?", tableOrderId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}
	now := time.Now()
	categoryToUpdate := models.TableOrder{
		TableName: payload.TableName,
		User:      currentUser.ID,
		CreatedAt: updatedTableOrder.CreatedAt,
		UpdatedAt: now,
	}

	pc.DB.Model(&updatedTableOrder).Updates(categoryToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedTableOrder})
}

// [...] Get Single Post Handler
func (pc *TableOrderController) FindTableOrderById(ctx *gin.Context) {
	tableOrderId := ctx.Param("tableOrderId")

	var category models.TableOrder
	result := pc.DB.First(&category, "id = ?", tableOrderId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": category})
}

// [...] Get All Posts Handler
func (pc *TableOrderController) FindTableOrders(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var categories []models.TableOrder
	results := pc.DB.Limit(intLimit).Offset(offset).Find(&categories)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(categories), "data": categories})
}

// [...] Delete Post Handler
func (pc *TableOrderController) DeleteTableOrder(ctx *gin.Context) {
	tableOrderId := ctx.Param("tableOrderId")

	result := pc.DB.Delete(&models.TableOrder{}, "id = ?", tableOrderId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
