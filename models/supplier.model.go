package models

import (
	"github.com/akhidnukhlis/db"
	validator "github.com/go-playground/validator"
	"net/http"
)

type Supplier struct {
	SupplierID      	string    		`json:"supplierID" validate:"required"`
	SupplierName    	string 			`json:"supplierName" validate:"required"`
}

func FetchAllSupplier() (Response, error) {
	var obj Supplier
	var arrobj []Supplier
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

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

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


