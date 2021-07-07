package controllers

import (
	"github.com/labstack/echo"
	"inventory-go/models"
	"net/http"
	"strconv"
)

func FetchAllAuditing(c echo.Context) error {
	result, err := models.FetchAllAuditing()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreAuditing(c echo.Context) error {
	AuditID 			:= c.FormValue("auditID")
	GoodsCode 			:= c.FormValue("goodsCode")
	SysStock 			:= c.FormValue("sysStock")
	RealStock 			:= c.FormValue("realStock")
	Note 				:= c.FormValue("note")

	// convert string to int
	convAuditID, err := strconv.Atoi(AuditID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	convSysStock, err := strconv.Atoi(SysStock)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	convRealStock, err := strconv.Atoi(RealStock)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.StoreAuditing(convAuditID, GoodsCode, convSysStock, convRealStock, Note)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}