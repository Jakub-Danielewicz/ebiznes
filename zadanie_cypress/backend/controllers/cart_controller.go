package controllers

import (
	"backend/db"
	"backend/models"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type CartRequest struct {
	UserID uint `json:"userID"`
	Items  []struct {
		ProductID uint `json:"productID"`
		Quantity  int  `json:"quantity"`
	} `json:"items"`
}

func GetOrCreateCart(c echo.Context) error {
	sess, err := session.Get("session", c)
	cartID, ok := sess.Values["cart_id"].(uint)

	if err != nil {
		return err
	}
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
		Secure:   false,
	}
	var cart models.Cart
	if ok {
		err := db.DB.Preload("CartItems.Product").First(&cart, cartID).Error
		if err != nil {
			return c.JSON(http.StatusNotFound, "Cart not found")
		} else {
			return c.JSON(http.StatusOK, cart)
		}
	}
	cart = models.Cart{}
	db.DB.Create(&cart)
	sess.Values["cart_id"] = cart.ID
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, cart)
}

func GetCart(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var cart models.Cart
	if err := db.DB.First(&cart, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Cart not found")
	}
	err := db.DB.Preload("CartItems.Product").First(&cart, id).Error
	if err != nil {
		return c.JSON(http.StatusNotFound, "Cart not found")
	}

	return c.JSON(http.StatusOK, cart)
}

func AddItemToCart(c echo.Context) error {
	cartID, _ := strconv.Atoi(c.Param("id"))
	var item struct {
		ProductID uint `json:"productId"`
		Quantity  int  `json:"quantity"`
	}

	if err := c.Bind(&item); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if item.Quantity <= 0 {
		return c.JSON(http.StatusBadRequest, "Quantity must be greater than 0")
	}

	var product models.Product
	if err := db.DB.First(&product, item.ProductID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, "Product does not exist")
	}

	cartItem := models.CartItem{
		CartID:    uint(cartID),
		ProductID: item.ProductID,
		Quantity:  item.Quantity,
	}
	if err := db.DB.Create(&cartItem).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if err := db.DB.Preload("Product").First(&cartItem, cartItem.ID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Nie udało się załadować produktu")
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
	db.DB.Where("cart_id=?", cartID).Delete(models.CartItem{})
	var cart models.Cart
	if err := db.DB.First(&cart, cartID).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Cart not found")
	}
	if err := db.DB.Delete(&cart).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete cart")
	}

	return c.NoContent(http.StatusNoContent)
}
