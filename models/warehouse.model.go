package models

import (
	"github.com/akhidnukhlis/db"
	validator "github.com/go-playground/validator"
	"net/http"
	"time"
)

type Warehouse struct {
	WarehouseID      	int    			`json:"warehouseID"`
	WarehouseName    	string 			`json:"warehouseName" validate:"required"`
	CreatedDate    		time.Time		`json:"createdDate"`
	ModifiedDate    	time.Time		`json:"modifiedDate"`
}

func FetchAllWarehouse() (Response, error) {
	var obj Warehouse
	var arrobj []Warehouse
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM warehouse"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.WarehouseID, &obj.WarehouseName, &obj.CreatedDate, &obj.ModifiedDate)
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

func StoreWarehouse(warehouseName string, createdDate time.Time) (Response, error) {
	var res Response

	v := validator.New()

	ware := Warehouse{
		WarehouseName		: warehouseName,
		CreatedDate			: createdDate,
	}

	err := v.Struct(ware)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "INSERT warehouse (warehouseName, createdDate) VALUES (?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(warehouseName, createdDate)
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

func UpdateWarehouse(warehouseID int, warehouseName string, modifiedDate time.Time) (Response, error) {
	var res Response

	v := validator.New()

	ware := Warehouse{
		WarehouseID			: warehouseID,
		WarehouseName		: warehouseName,
		ModifiedDate		: modifiedDate,
	}

	err := v.Struct(ware)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "UPDATE warehouse SET warehouseName = ?, modifiedDate = ? WHERE warehouseID = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(warehouseName, modifiedDate, warehouseID)
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

func DeleteWarehouse(warehouseID int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM warehouse WHERE warehouseID = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(warehouseID)
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