package controllers

import (
	"inventory-go/models"
	"github.com/labstack/echo"
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
	AuditInfoID 		:= c.FormValue("auditInfoID")
	AuditID 			:= c.FormValue("auditID")
	GoodsCode 			:= c.FormValue("goodsCode")
	SysStock 			:= c.FormValue("sysStock")
	RealStock 			:= c.FormValue("realStock")
	Note 				:= c.FormValue("note")

	// convert string to int
	convAuditInfoID, err := strconv.Atoi(AuditInfoID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

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

	result, err := models.StoreAuditing(convAuditInfoID, convAuditID, GoodsCode, convSysStock, convRealStock, Note)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}