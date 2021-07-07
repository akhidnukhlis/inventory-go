package controllers

import (
	"github.com/labstack/echo"
	"inventory-go/helpers"
	"inventory-go/models"
	"net/http"
	"strconv"
	"time"
)

func FetchAllGoods(c echo.Context) error {
	result, err := models.FetchAllGoods()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreGoods(c echo.Context) error {
	GoodsCode 			:= c.FormValue("goodsCode")
	CategoryCode 		:= c.FormValue("categoryCode")
	GoodsName 			:= c.FormValue("goodsName")
	goodsLowestPrice 	:= c.FormValue("goodsLowestPrice")
	GoodsRetailPrice 	:= c.FormValue("goodsRetailPrice")
	WarehouseID 		:= c.FormValue("warehouseID")
	ProcStaffID 		:= c.FormValue("procStaffID")
	ProcMgrID 			:= c.FormValue("procMgrID")
	CreatedDate 		:= helpers.BuildTime()

	// convert string to int
	convgoodsLowestPrice, err := strconv.Atoi(goodsLowestPrice)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	convGoodsRetailPrice, err := strconv.Atoi(GoodsRetailPrice)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	convWarehouseID, err := strconv.Atoi(WarehouseID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	convProcStaffID, err := strconv.Atoi(ProcStaffID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	convProcMgrID, err := strconv.Atoi(ProcMgrID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// convert string to date
	convCreatedDate, err := time.Parse(helpers.LayoutFormat(), CreatedDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.StoreGoods(GoodsCode, CategoryCode, GoodsName, convgoodsLowestPrice, convGoodsRetailPrice, convWarehouseID, convProcStaffID, convProcMgrID, convCreatedDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateGoods(c echo.Context) error {
	GoodsCode 			:= c.FormValue("goodsCode")
	CategoryCode 		:= c.FormValue("categoryCode")
	GoodsName 			:= c.FormValue("goodsName")
	goodsLowestPrice 	:= c.FormValue("goodsLowestPrice")
	GoodsRetailPrice 	:= c.FormValue("goodsRetailPrice")
	WarehouseID 		:= c.FormValue("warehouseID")
	ProcStaffID 		:= c.FormValue("procStaffID")
	ProcMgrID 			:= c.FormValue("procMgrID")
	ModifiedDate 		:= helpers.BuildTime()

	// convert string to int
	convgoodsLowestPrice, err := strconv.Atoi(goodsLowestPrice)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	convGoodsRetailPrice, err := strconv.Atoi(GoodsRetailPrice)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	convWarehouseID, err := strconv.Atoi(WarehouseID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	convProcStaffID, err := strconv.Atoi(ProcStaffID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	convProcMgrID, err := strconv.Atoi(ProcMgrID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// convert string to date
	convModifiedDate, err := time.Parse(helpers.LayoutFormat(), ModifiedDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateGoods(GoodsCode, CategoryCode, GoodsName, convgoodsLowestPrice, convGoodsRetailPrice,
		convWarehouseID, convProcStaffID, convProcMgrID, convModifiedDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteGoods(c echo.Context) error {
	GoodsCode := c.FormValue("goodsCode")

	result, err := models.DeleteGoods(GoodsCode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}