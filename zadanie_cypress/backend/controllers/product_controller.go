package controllers

import (
	"backend/db"
	"backend/models"
	"net/http"
	"strconv"
	"fmt"
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
        fmt.Println("Error binding product:", err)
        return c.JSON(http.StatusBadRequest, err)
    }

	fmt.Println("Product data:", product)
    if product.Name == "" || product.Price == 0 || product.CategoryID == 0 {
        return c.JSON(http.StatusBadRequest, "All fields (name, price, category_id) are required")
    }

    if err := db.DB.Create(&product).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }
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
		fmt.Println("Error finding product:", err)
		fmt.Println("Product ID:", id)
		return c.JSON(http.StatusNotFound, "Product not found")
	}
	err :=db.DB.Delete(&product).Error
	if err != nil {
		fmt.Println("Error deleting product:", err)
		return c.JSON(http.StatusInternalServerError, "Failed to delete product")
	}
	fmt.Println("Product deleted successfully")
	return c.NoContent(http.StatusNoContent)
	}

