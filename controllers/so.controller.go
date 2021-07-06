package controllers

import (
	"inventory-go/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"time"
)

func FetchAllSo(c echo.Context) error {
	result, err := models.FetchAllSo()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreSo(c echo.Context) error {
	SoCode 				:= c.FormValue("soCode")
	GoodsCode 			:= c.FormValue("goodsCode")
	UserCreated 		:= c.FormValue("userCreated")
	SourceID 			:= c.FormValue("sourceID")
	Customer 			:= c.FormValue("customer")
	SoQty 				:= c.FormValue("soQty")
	Currency 			:= c.FormValue("currency")
	UnitPrice 			:= c.FormValue("unitPrice")
	TotalPrice 			:= c.FormValue("totalPrice")
	Courier 			:= c.FormValue("courier")
	DeliveryCharge 		:= c.FormValue("deliveryCharge")
	Status 				:= c.FormValue("status")
	CreatedDate 		:= c.FormValue("createdDate")

	// convert string to int
	convUserCreated, err := strconv.Atoi(UserCreated)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	convSourceID, err := strconv.Atoi(SourceID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	convSoQty, err := strconv.Atoi(SoQty)
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

	convDeliveryCharge, err := strconv.Atoi(DeliveryCharge)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// convert string to date
	convCreatedDate, err := time.Parse(layoutFormat, CreatedDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.StoreSo(SoCode, GoodsCode, convUserCreated, convSourceID, Customer, convSoQty, Currency,
								  convUnitPrice, convTotalPrice, Courier, convDeliveryCharge, Status, convCreatedDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateSo(c echo.Context) error {
	SoCode 				:= c.FormValue("soCode")
	GoodsCode 			:= c.FormValue("goodsCode")
	UserCreated 		:= c.FormValue("userCreated")
	SourceID 			:= c.FormValue("sourceID")
	Customer 			:= c.FormValue("customer")
	SoQty 				:= c.FormValue("soQty")
	Currency 			:= c.FormValue("currency")
	UnitPrice 			:= c.FormValue("unitPrice")
	TotalPrice 			:= c.FormValue("totalPrice")
	Courier 			:= c.FormValue("courier")
	DeliveryCharge 		:= c.FormValue("deliveryCharge")
	Status 				:= c.FormValue("status")
	ModifiedDate 		:= c.FormValue("modifiedDate")

	// convert string to int
	convUserCreated, err := strconv.Atoi(UserCreated)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	convSourceID, err := strconv.Atoi(SourceID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	convSoQty, err := strconv.Atoi(SoQty)
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

	convDeliveryCharge, err := strconv.Atoi(DeliveryCharge)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// convert string to date
	convModifiedDate, err := time.Parse(layoutFormat, ModifiedDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateSo(GoodsCode, convUserCreated, convSourceID, Customer, convSoQty, Currency,
		convUnitPrice, convTotalPrice, Courier, convDeliveryCharge, Status, convModifiedDate, SoCode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteSo(c echo.Context) error {
	SoCode := c.FormValue("soCode")

	result, err := models.DeleteSo(SoCode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}