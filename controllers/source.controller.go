package controllers

import (
	"github.com/akhidnukhlis/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"time"
)

func FetchAllSource(c echo.Context) error {
	result, err := models.FetchAllSource()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreSource(c echo.Context) error {
	SourceName 		:= c.FormValue("sourceName")
	CreatedDate 	:= c.FormValue("createdDate")

	convCreatedDate, err := time.Parse(layoutFormat, CreatedDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.StoreSource(SourceName, convCreatedDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateSource(c echo.Context) error {
	SourceID 		:= c.FormValue("sourceID")
	SourceName 		:= c.FormValue("sourceName")
	ModifiedDate 	:= c.FormValue("modifiedDate")

	convSourceID, err := strconv.Atoi(SourceID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	convModifiedDate, err := time.Parse(layoutFormat, ModifiedDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateSource(convSourceID, SourceName, convModifiedDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteSource(c echo.Context) error {
	SourceID := c.FormValue("sourceID")

	convSourceID, err := strconv.Atoi(SourceID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteSource(convSourceID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}