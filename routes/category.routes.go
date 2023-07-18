package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/projectlac/golang-gorm-postgres/controllers"
	"github.com/projectlac/golang-gorm-postgres/middleware"
)

type CategoryRouteController struct {
	categoryController controllers.CategoryController
}

func NewRouteCategoryController(categoryController controllers.CategoryController) CategoryRouteController {
	return CategoryRouteController{categoryController}
}

func (pc *CategoryRouteController) CategoryRoute(rg *gin.RouterGroup) {

	router := rg.Group("categories")
	router.Use(middleware.DeserializeUser())
	router.POST("/", pc.categoryController.CreateCategory)
	router.GET("/", pc.categoryController.FindCategories)
	router.PUT("/:categoryId", pc.categoryController.UpdateCategory)
	router.GET("/:categoryId", pc.categoryController.FindCategoryById)
	router.DELETE("/:categoryId", pc.categoryController.DeleteCategory)
}
