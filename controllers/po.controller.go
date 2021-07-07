package controllers

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo"
	"inventory-go/helpers"
	"inventory-go/models"
	"net/http"
	"strconv"
	"time"
)

func FetchAllPo(c echo.Context) error {
	result, err := models.FetchAllPo()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StorePo(c echo.Context) error {
	PoCode 				:= fmt.Sprint("PO-", helpers.BuildFileName())
	GoodsCode 			:= c.FormValue("goodsCode")
	SupplierCode 		:= c.FormValue("supplierCode")
	UserCreated 		:= c.FormValue("userCreated")
	UserApproved 		:= c.FormValue("userApproved")
	PoQty 				:= c.FormValue("poQty")
	Currency 			:= c.FormValue("currency")
	UnitPrice 			:= c.FormValue("unitPrice")
	TotalPrice 			:= c.FormValue("totalPrice")
	Status 				:= c.FormValue("status")
	CreatedDate 		:= helpers.BuildTime()

	// convert string to int
	convUserCreated, err := strconv.Atoi(UserCreated)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	convUserApproved, err := strconv.ParseInt(UserApproved, 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	convPoQty, err := strconv.Atoi(PoQty)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	convUnitPrice, err := strconv.Atoi(UnitPrice)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	convTotalPrice, err := strconv.Atoi(TotalPrice)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// convert string to date
	convCreatedDate, err := time.Parse(helpers.LayoutFormat(), CreatedDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.StorePo(PoCode, GoodsCode, SupplierCode, convUserCreated, sql.NullInt64{Int64: convUserApproved, Valid: true},
								  convPoQty, Currency, convUnitPrice, convTotalPrice, Status, convCreatedDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdatePo(c echo.Context) error {
	PoCode 				:= c.FormValue("poCode")
	GoodsCode 			:= c.FormValue("goodsCode")
	SupplierCode 		:= c.FormValue("supplierCode")
	UserCreated 		:= c.FormValue("userCreated")
	UserApproved 		:= c.FormValue("userApproved")
	PoQty 				:= c.FormValue("poQty")
	Currency 			:= c.FormValue("currency")
	UnitPrice 			:= c.FormValue("unitPrice")
	TotalPrice 			:= c.FormValue("totalPrice")
	Status 				:= c.FormValue("status")
	ModifiedDate 		:= helpers.BuildTime()

	// convert string to int
	convUserCreated, err := strconv.Atoi(UserCreated)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	convUserApproved, err := strconv.ParseInt(UserApproved, 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	convPoQty, err := strconv.Atoi(PoQty)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	convUnitPrice, err := strconv.Atoi(UnitPrice)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	convTotalPrice, err := strconv.Atoi(TotalPrice)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// convert string to date
	convModifiedDate, err := time.Parse(helpers.LayoutFormat(), ModifiedDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdatePo(PoCode, GoodsCode, SupplierCode, convUserCreated, sql.NullInt64{Int64: convUserApproved, Valid: true},
								   convPoQty, Currency, convUnitPrice, convTotalPrice, Status, convModifiedDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeletePo(c echo.Context) error {
	PoCode := c.FormValue("poCode")

	result, err := models.DeletePo(PoCode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}