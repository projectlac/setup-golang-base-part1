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

type OrderController struct {
	DB *gorm.DB
}

func NewOrderController(DB *gorm.DB) OrderController {
	return OrderController{DB}
}

func (pc *OrderController) CreateOrder(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.CreateOrderRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	newOrder := models.Order{
		TableId:   payload.TableId,
		OrderItem: payload.OrderItem,
		User:      currentUser.ID,
		CreatedAt: now,
		UpdatedAt: now,
	}
	result := pc.DB.Create(&newOrder)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Post with that title already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newOrder})
}

// [...] Update Post Handler
func (pc *OrderController) UpdateOrder(ctx *gin.Context) {
	orderId := ctx.Param("orderId")
	currentUser := ctx.MustGet("currentUser").(models.User)

	var payload *models.UpdateOrder
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updatedOrder models.Order
	result := pc.DB.First(&updatedOrder, "id = ?", orderId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}
	now := time.Now()
	orderToUpdate := models.Order{
		OrderItem: payload.OrderItem,
		User:      currentUser.ID,
		CreatedAt: updatedOrder.CreatedAt,
		UpdatedAt: now,
	}

	pc.DB.Model(&updatedOrder).Updates(orderToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedOrder})
}

// [...] Get Single Post Handler
func (pc *OrderController) FindOrderById(ctx *gin.Context) {
	orderId := ctx.Param("orderId")

	var order models.Order
	result := pc.DB.First(&order, "id = ?", orderId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": order})
}

// [...] Get All Posts Handler
func (pc *OrderController) FindOrders(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")
	var order = ctx.DefaultQuery("order", "asc")
	tableId := ctx.Param("tableId")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var products []models.Product

	query := pc.DB

	if tableId != "" {
		query = query.Where("table_id = ?", tableId)
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
func (pc *OrderController) DeleteOrder(ctx *gin.Context) {
	orderId := ctx.Param("orderId")

	result := pc.DB.Delete(&models.Order{}, "id = ?", orderId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
