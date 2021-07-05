package controllers

import (
	"github.com/akhidnukhlis/models"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func FetchAllRole(c echo.Context) error {
	result, err := models.FetchAllRole()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreRole(c echo.Context) error {
	RoleCode 		:= c.FormValue("roleCode")
	RoleName 		:= c.FormValue("roleName")
	RoleDesc 		:= c.FormValue("roleDesc")
	CreatedDate 	:= c.FormValue("createdDate")

	convCreatedDate, err := time.Parse(layoutFormat, CreatedDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.StoreRole(RoleCode, RoleName, RoleDesc, convCreatedDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateRole(c echo.Context) error {
	RoleCode 		:= c.FormValue("roleCode")
	RoleName 		:= c.FormValue("roleName")
	RoleDesc 		:= c.FormValue("roleDesc")
	ModifiedDate 	:= c.FormValue("modifiedDate")

	convModifiedDate, err := time.Parse(layoutFormat, ModifiedDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateRole(RoleCode, RoleName, RoleDesc, convModifiedDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteRole(c echo.Context) error {
	RoleCode := c.FormValue("roleCode")

	result, err := models.DeleteRole(RoleCode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}