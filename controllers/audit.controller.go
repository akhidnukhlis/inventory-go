package controllers

import (
	"github.com/labstack/echo"
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
	AuditID 			:= c.FormValue("auditID")
	CreatedDate 		:= c.FormValue("createdDate")
	Auditor 			:= c.FormValue("auditor")

	// convert string to int
	convAuditID, err := strconv.Atoi(AuditID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	convAuditor, err := strconv.Atoi(Auditor)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// convert string to date
	convCreatedDate, err := time.Parse(layoutFormat, CreatedDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.StoreAudit(convAuditID, convCreatedDate, convAuditor)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
