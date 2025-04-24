package controllers

import (
	"zadanie_go/db"
	"zadanie_go/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CartRequest struct {
	UserID uint `json:"userID"`
	Items []struct {
		ProductID uint `json:"productID"`
		Quantity int`json:"quantity"`
	} `json:items"`
}

func CreateCart(c echo.Context) error {
	var req CartRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	cart := models.Cart{
		UserID: req.UserID,
	}

	if err := db.DB.Create(&cart).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	
	for _, item := range req.Items {
		cartItem := models.CartItem{
			CartID: cart.ID,
			ProductID: item.ProductID,
			Quantity: item.Quantity,
		}
		db.DB.Create(&cartItem)
	}

	return c.JSON(http.StatusCreated, cart)
}

func GetCart(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var cart models.Cart

	err := db.DB.Preload("CartItems.Product").First(&cart, id).Error
	if err != nil {
		return c.JSON(http.StatusNotFound, "Cart not found")
	}

	return c.JSON(http.StatusOK, cart)
}

func AddItemToCart(c echo.Context) error {
	cartID, _ := strconv.Atoi(c.Param("id"))
	var item struct {
		ProductID uint `json:"productID"`
		Quantity int `json:"quantity"`
	}

	if err := c.Bind(&item); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	cartItem := models.CartItem{
		CartID:	uint(cartID),
		ProductID:	item.ProductID,
		Quantity:	item.Quantity,
	}
	if err := db.DB.Create(&cartItem).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, cartItem)
}

func RemoveItemFromCart(c echo.Context) error {
	itemID, _ := strconv.Atoi(c.Param("itemId"))
	var item models.CartItem

	if err := db.DB.First(&item, itemID).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Item not found")
	}

	db.DB.Delete(&item)
	return c.NoContent(http.StatusNoContent)
}

func DeleteCart(c echo.Context) error {
	cartID, _ := strconv.Atoi(c.Param("id"))
	
	db.DB.Where("cart_id=?", cartID).Delete(&models.CartItem{})
	if err := db.DB.Delete(&models.Cart{}, cartID).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Cart not found")
	}
	
	return c.NoContent(http.StatusNoContent)
}
