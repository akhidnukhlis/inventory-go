package controllers

import (
	"inventory-go/models"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func FetchAllCategory(c echo.Context) error {
	result, err := models.FetchAllCategory()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreCategory(c echo.Context) error {
	CategoryCode 		:= c.FormValue("categoryCode")
	CategoryName 		:= c.FormValue("categoryName")
	CreatedDate 		:= c.FormValue("createdDate")

	convCreatedDate, err := time.Parse(layoutFormat, CreatedDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.StoreCategory(CategoryCode, CategoryName, convCreatedDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateCategory(c echo.Context) error {
	CategoryCode 		:= c.FormValue("categoryCode")
	CategoryName 		:= c.FormValue("categoryName")
	ModifiedDate 		:= c.FormValue("modifiedDate")

	convModifiedDate, err := time.Parse(layoutFormat, ModifiedDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateCategory(CategoryCode, CategoryName, convModifiedDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteCategory(c echo.Context) error {
	CategoryCode := c.FormValue("categoryCode")

	result, err := models.DeleteCategory(CategoryCode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
