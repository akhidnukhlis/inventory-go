package models

import (
	"github.com/akhidnukhlis/db"
	validator "github.com/go-playground/validator"
	"net/http"
	"time"
)

type Goods struct {
	GoodsCode      		string    	`json:"goodsCode" validate:"required"`
	CategoryCode    	string 		`json:"categoryCode" validate:"required"`
	GoodsName    		string 		`json:"goodsName" validate:"required"`
	GoodsLowesPrice    	int 		`json:"goodsLowesPrice" validate:"required"`
	GoodsRetailPrice    int 		`json:"goodsRetailPrice" validate:"required"`
	WarehouseID    		int 		`json:"warehouseID" validate:"required"`
	ProcStaffID    		int			`json:"procStaffID" validate:"required"`
	ProcMgrID    		int 		`json:"procMgrID" validate:"required"`
	CreatedDate    		time.Time	`json:"createdDate"`
	ModifiedDate    	time.Time	`json:"modifiedDate"`
}

func FetchAllGoods() (Response, error) {
	var obj Goods
	var arrows []Goods
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM goods"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.GoodsCode, &obj.CategoryCode, &obj.GoodsName, &obj.GoodsLowesPrice, &obj.GoodsRetailPrice,
			&obj.WarehouseID, &obj.ProcStaffID, &obj.ProcMgrID, &obj.CreatedDate, &obj.ModifiedDate)
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

func StoreGoods(goodsCode string, categoryCode string, goodsName string, goodsLowesPrice int,
				goodsRetailPrice int, warehouseID int, procStaffID int, procMgrID int,
				createdDate time.Time) (Response, error) {
	var res Response

	v := validator.New()

	good := Goods{
		GoodsCode				: goodsCode,
		CategoryCode			: categoryCode,
		GoodsName				: goodsName,
		GoodsLowesPrice			: goodsLowesPrice,
		GoodsRetailPrice		: goodsRetailPrice,
		WarehouseID				: warehouseID,
		ProcStaffID				: procStaffID,
		ProcMgrID				: procMgrID,
		CreatedDate				: createdDate,
	}

	err := v.Struct(good)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "INSERT goods (goodsCode, categoryCode, goodsName, goodsLowesPrice," +
					"goodsRetailPrice, warehouseID, procStaffID, procMgrID, createdDate" +
					") VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(goodsCode, categoryCode, goodsName, goodsLowesPrice, goodsRetailPrice,
							 warehouseID, procStaffID, procMgrID, createdDate)
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

func UpdateGoods(goodsCode string, categoryCode string, goodsName string, goodsLowesPrice int,
				 goodsRetailPrice int, warehouseID int, procStaffID int, procMgrID int,
				 modifiedDate time.Time) (Response, error) {
	var res Response

	v := validator.New()

	good := Goods{
		GoodsCode				: goodsCode,
		CategoryCode			: categoryCode,
		GoodsName				: goodsName,
		GoodsLowesPrice			: goodsLowesPrice,
		GoodsRetailPrice		: goodsRetailPrice,
		WarehouseID				: warehouseID,
		ProcStaffID				: procStaffID,
		ProcMgrID				: procMgrID,
		ModifiedDate			: modifiedDate,
	}

	err := v.Struct(good)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "UPDATE goods SET categoryCode = ?, goodsName = ?, goodsLowesPrice = ?," +
					"goodsRetailPrice = ?, warehouseID = ?, procStaffID = ?, procMgrID = ?," +
					"modifiedDate = ? WHERE goodsCode = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(categoryCode, goodsName, goodsLowesPrice, goodsRetailPrice,
							 warehouseID, procStaffID, procMgrID, modifiedDate, goodsCode)
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

func DeleteGoods(goodsCode string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM goods WHERE goodsCode = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(goodsCode)
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