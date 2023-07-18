package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/projectlac/golang-gorm-postgres/controllers"
	"github.com/projectlac/golang-gorm-postgres/middleware"
)

type TableRouteController struct {
	tableController controllers.TableController
}

func NewRouteTableController(tableController controllers.TableController) TableRouteController {
	return TableRouteController{tableController}
}

func (pc *TableRouteController) TableRoute(rg *gin.RouterGroup) {

	router := rg.Group("tables")
	router.Use(middleware.DeserializeUser())
	router.POST("/", pc.tableController.CreateTable)
	router.GET("/", pc.tableController.FindTables)
	router.PUT("/:categoryId", pc.tableController.UpdateTable)
	router.GET("/:categoryId", pc.tableController.FindTableById)
	router.DELETE("/:categoryId", pc.tableController.DeleteTable)
}
