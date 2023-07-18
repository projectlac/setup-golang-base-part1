package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/projectlac/golang-gorm-postgres/controllers"
	"github.com/projectlac/golang-gorm-postgres/middleware"
)

type OrderRouteController struct {
	orderController controllers.OrderController
}

func NewRouteOrderController(orderController controllers.OrderController) OrderRouteController {
	return OrderRouteController{orderController}
}

func (pc *OrderRouteController) OrderRoute(rg *gin.RouterGroup) {

	router := rg.Group("orders")
	router.Use(middleware.DeserializeUser())
	router.POST("/", pc.orderController.CreateOrder)
	router.GET("/", pc.orderController.FindOrders)
	router.PUT("/:tableId", pc.orderController.UpdateOrder)
	router.GET("/:tableId", pc.orderController.FindOrderById)
	router.DELETE("/:tableId", pc.orderController.DeleteOrder)
}
