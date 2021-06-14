package controllers

import (
	"github.com/akhidnukhlis/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"time"
)

var layoutFormat = "2006-01-02 15:04:05"

func FetchAllWarehouse(c echo.Context) error {
	result, err := models.FetchAllWarehouse()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreWarehouse(c echo.Context) error {
	WarehouseName 		:= c.FormValue("warehouseName")
	CreatedDate 		:= c.FormValue("createdDate")

	convCreatedDate, err := time.Parse(layoutFormat, CreatedDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.StoreWarehouse(WarehouseName, convCreatedDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateWarehouse(c echo.Context) error {
	WarehouseID 	:= c.FormValue("warehouseID")
	WarehouseName 	:= c.FormValue("warehouseName")
	ModifiedDate 	:= c.FormValue("modifiedDate")

	convWarehouseID, err := strconv.Atoi(WarehouseID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	convModifiedDate, err := time.Parse(layoutFormat, ModifiedDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateWarehouse(convWarehouseID, WarehouseName, convModifiedDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteWarehouse(c echo.Context) error {
	WarehouseID := c.FormValue("warehouseID")

	convWarehouseID, err := strconv.Atoi(WarehouseID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteWarehouse(convWarehouseID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}