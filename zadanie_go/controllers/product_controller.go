package controllers

import (
	"zadanie_go/db"
	"zadanie_go/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetProducts(c echo.Context) error {
	var products []models.Product
	db.DB.Scopes(models.WithCategory()).Find(&products)
	return c.JSON(http.StatusOK, products)
}

func CreateProduct(c echo.Context) error {
	var product models.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	db.DB.Create(&product)
	return c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var product models.Product
	if err := db.DB.First(&product, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	db.DB.Save(&product)
	return c.JSON(http.StatusOK, product)
}

func DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var product models.Product
	if err := db.DB.First(&product, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}
	db.DB.Delete(&product)
	return c.NoContent(http.StatusNoContent)
	}

