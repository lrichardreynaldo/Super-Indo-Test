package main

import (
	controllers "superIndo/controller"
	middlewares "superIndo/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := initRouter()
	router.Run(":8080")
}
func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	v1 := api.Group("/v1")
	v2 := api.Group("/v2").Use(middlewares.Auth())
	user := v1.Group("/user")
	user.POST("/login", controllers.Login)
	user.POST("/register", controllers.Register)
	v2.GET("/cart", controllers.GetCurrentUserCartDetail)
	v2.GET("/categories", controllers.ListCategory)
	v2.GET("/categories/:id", controllers.GetProductByCategoryId)
	v2.GET("/product/:id", controllers.GetProductDetailByProductId)
	v2.POST("/cart/add", controllers.AddProductToCart)
	return router
}
