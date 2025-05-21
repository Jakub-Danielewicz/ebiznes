package routes

import (
	"backend/controllers"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/products", controllers.GetProducts)
	e.POST("/products", controllers.CreateProduct)
	e.PUT("/products/:id", controllers.UpdateProduct)
	e.DELETE("products/:id", controllers.DeleteProduct)
	e.POST("/cart", controllers.GetOrCreateCart)
	e.GET("/cart", controllers.GetOrCreateCart)
	e.POST("/cart/:id/items", controllers.AddItemToCart)
	e.DELETE("/cart/:cartId/items/:itemId", controllers.RemoveItemFromCart)
	e.DELETE("/cart/:id", controllers.DeleteCart)
}

