package controllers

import (
	"github.com/labstack/echo"
	"inventory-go/helpers"
	"inventory-go/models"
	"net/http"
	"strconv"
	"time"
)

func FetchAllWarehouse(c echo.Context) error {
	result, err := models.FetchAllWarehouse()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreWarehouse(c echo.Context) error {
	WarehouseName	:= c.FormValue("warehouseName")
	CreatedDate		:= helpers.BuildTime()

	convCreatedDate, err := time.Parse(helpers.LayoutFormat(), CreatedDate)
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
	ModifiedDate 	:= helpers.BuildTime()

	convWarehouseID, err := strconv.Atoi(WarehouseID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	convModifiedDate, err := time.Parse(helpers.LayoutFormat(), ModifiedDate)
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