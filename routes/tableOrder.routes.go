package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/projectlac/golang-gorm-postgres/controllers"
	"github.com/projectlac/golang-gorm-postgres/middleware"
)

type TableOrderRouteController struct {
	tableOrderController controllers.TableOrderController
}

func NewRouteTableOrderController(tableOrderController controllers.TableOrderController) TableOrderRouteController {
	return TableOrderRouteController{tableOrderController}
}

func (pc *TableOrderRouteController) TableRoute(rg *gin.RouterGroup) {

	router := rg.Group("tables")
	router.Use(middleware.DeserializeUser())
	router.POST("/", pc.tableOrderController.CreateTableOrder)
	router.GET("/", pc.tableOrderController.FindTableOrders)
	router.PUT("/:tableOrderId", pc.tableOrderController.UpdateTableOrder)
	router.GET("/:tableOrderId", pc.tableOrderController.FindTableOrderById)
	router.DELETE("/:tableOrderId", pc.tableOrderController.DeleteTableOrder)
}
