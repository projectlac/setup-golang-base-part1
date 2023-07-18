package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/projectlac/golang-gorm-postgres/controllers"
	"github.com/projectlac/golang-gorm-postgres/initializers"
	"github.com/projectlac/golang-gorm-postgres/routes"
)

var (
	server              *gin.Engine
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	TableOrderController      controllers.TableOrderController
	TableOrderRouteController routes.TableOrderRouteController

	ProductController      controllers.ProductController
	ProductRouteController routes.ProductRouteController

	OrderController      controllers.OrderController
	OrderRouteController routes.OrderRouteController

	CategoryController      controllers.CategoryController
	CategoryRouteController routes.CategoryRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(initializers.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	ProductController = controllers.NewProductController(initializers.DB)
	ProductRouteController = routes.NewRouteProductController(ProductController)

	CategoryController = controllers.NewCategoryController(initializers.DB)
	CategoryRouteController = routes.NewRouteCategoryController(CategoryController)

	TableOrderController = controllers.NewTableOrderController(initializers.DB)
	TableOrderRouteController = routes.NewRouteTableOrderController(TableOrderController)

	OrderController = controllers.NewOrderController(initializers.DB)
	OrderRouteController = routes.NewRouteOrderController(OrderController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	TableOrderRouteController.TableRoute(router)
	CategoryRouteController.CategoryRoute(router)
	ProductRouteController.ProductRoute(router)
	OrderRouteController.OrderRoute(router)

	log.Fatal(server.Run(":" + config.ServerPort))
}
