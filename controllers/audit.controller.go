package controllers

import (
	"github.com/labstack/echo"
	"inventory-go/helpers"
	"inventory-go/models"
	"net/http"
	"strconv"
	"time"
)

func FetchAllAudit(c echo.Context) error {
	result, err := models.FetchAllAudit()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreAudit(c echo.Context) error {
	CreatedDate 		:= helpers.BuildTime()
	Auditor 			:= c.FormValue("auditor")

	// convert string to date
	convCreatedDate, err := time.Parse(helpers.LayoutFormat(), CreatedDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// convert string to int
	convAuditor, err := strconv.Atoi(Auditor)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.StoreAudit(convCreatedDate, convAuditor)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
