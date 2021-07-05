package models

import (
	"github.com/akhidnukhlis/db"
	validator "github.com/go-playground/validator"
	"net/http"
	"time"
)

type Role struct {
	RoleCode      		string    	`json:"roleCode" validate:"required"`
	RoleName    		string 		`json:"roleName" validate:"required"`
	RoleDesc    		string 		`json:"roleDesc" validate:"required"`
	CreatedDate    		time.Time	`json:"createdDate"`
	ModifiedDate    	time.Time	`json:"modifiedDate"`
}

func FetchAllRole() (Response, error) {
	var obj Role
	var arrows []Role
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM role"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.RoleCode, &obj.RoleName, &obj.RoleDesc, &obj.CreatedDate, &obj.ModifiedDate)
		if err != nil {
			return res, err
		}

		arrows = append(arrows, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrows

	return res, nil
}

func StoreRole(roleCode string, roleName string, roleDesc string, createdDate time.Time) (Response, error) {
	var res Response

	v := validator.New()

	role := Role{
		RoleCode		: roleCode,
		RoleName		: roleName,
		RoleDesc		: roleDesc,
		CreatedDate		: createdDate,
	}

	err := v.Struct(role)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "INSERT role (roleCode, roleName, roleDesc, createdDate) VALUES (?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(roleCode, roleName, roleDesc, createdDate)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedId,
	}

	return res, nil
}

func UpdateRole(roleCode string, roleName string, roleDesc string, modifiedDate time.Time) (Response, error) {
	var res Response

	v := validator.New()

	role := Role{
		RoleCode		: roleCode,
		RoleName		: roleName,
		RoleDesc		: roleDesc,
		ModifiedDate	: modifiedDate,
	}

	err := v.Struct(role)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "UPDATE role SET roleName = ?, roleDesc = ?, modifiedDate = ? WHERE roleCode = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(roleName, roleDesc, modifiedDate, roleCode)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func DeleteRole(roleCode string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM role WHERE roleCode = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(roleCode)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
