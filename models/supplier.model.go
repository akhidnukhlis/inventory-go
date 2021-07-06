package models

import (
	validator "github.com/go-playground/validator"
	"inventory-go/db"
	"net/http"
)

type Supplier struct {
	SupplierID      	string    		`json:"supplierID" validate:"required"`
	SupplierName    	string 			`json:"supplierName" validate:"required"`
}

func FetchAllSupplier() (Response, error) {
	var obj Supplier
	var arrows []Supplier
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM supplier"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.SupplierID, &obj.SupplierName)
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

func StoreSupplier(supplierID string, supplierName string) (Response, error) {
	var res Response

	v := validator.New()

	ware := Supplier{
		SupplierID		: supplierID,
		SupplierName	: supplierName,
	}

	err := v.Struct(ware)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "INSERT supplier (supplierID, supplierName) VALUES (?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(supplierID, supplierName)
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

func UpdateSupplier(supplierID string, supplierName string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE supplier SET supplierName = ? WHERE supplierID = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(supplierName, supplierID)
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

func DeleteSupplier(supplierID string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM supplier WHERE supplierID = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(supplierID)
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