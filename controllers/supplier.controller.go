package controllers

import (
	"fmt"
	"github.com/labstack/echo"
	"inventory-go/helpers"
	"inventory-go/models"
	"net/http"
)
func FetchAllSupplier(c echo.Context) error {
	result, err := models.FetchAllSupplier()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreSupplier(c echo.Context) error {
	SupplierID := fmt.Sprint("SUP", helpers.BuildFileName())
	SupplierName := c.FormValue("supplierName")

	result, err := models.StoreSupplier(SupplierID, SupplierName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateSupplier(c echo.Context) error {
	supplierID := c.FormValue("supplierID")
	supplierName := c.FormValue("supplierName")

	result, err := models.UpdateSupplier(supplierID, supplierName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteSupplier(c echo.Context) error {
	supplierID := c.FormValue("supplierID")

	result, err := models.DeleteSupplier(supplierID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
