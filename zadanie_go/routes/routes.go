package routes

import (
	"zadanie_go/controllers"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/products", controllers.GetProducts)
	e.POST("/products", controllers.CreateProduct)
	e.PUT("/products/:id", controllers.UpdateProduct)
	e.DELETE("products/:id", controllers.DeleteProduct)
	e.POST("/cart", controllers.CreateCart)
	e.GET("/cart/:id", controllers.GetCart)
	e.POST("/cart/:id/items", controllers.AddItemToCart)
	e.DELETE("/cart/:cartId/items/:itemId", controllers.RemoveItemFromCart)
	e.DELETE("/cart/:id", controllers.DeleteCart)
}

