package controllers

import (
	"fmt"
	"github.com/akhidnukhlis/models"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

//var layoutFormat = "2006-01-02 15:04:05"

func getToday(format string) (todayString string){
	today := time.Now()
	todayString = today.Format(format);
	return
}

func FetchAllSupplier(c echo.Context) error {
	result, err := models.FetchAllSupplier()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreSupplier(c echo.Context) error {
	SupplierID 		:= time.Now()
	SupplierName 	:= c.FormValue("supplierName")

	convSupplierID, err := fmt.Println(SupplierID.Format("2006-01-02 15:04:05"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	cSID, err := fmt.Println("SUP", convSupplierID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.StoreSupplier(string(cSID), SupplierName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateSupplier(c echo.Context) error {
	supplierID := c.FormValue("suppierID")
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