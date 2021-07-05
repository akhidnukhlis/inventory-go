package models

import (
	"database/sql"
	"github.com/akhidnukhlis/db"
	validator "github.com/go-playground/validator"
	"net/http"
	"time"
)

type Po struct {
	PoCode      		string    			`json:"poCode" validate:"required"`
	GoodsCode    		string 				`json:"goodsCode" validate:"required"`
	SupplierCode    	string 				`json:"supplierCode" validate:"required"`
	UserCreated    		int 				`json:"userCreated" validate:"required"`
	UserApproved    	sql.NullInt64 		`json:"userApproved" validate:"required"`
	PoQty    			int 				`json:"poQty" validate:"required"`
	Currency    		string 				`json:"currency" validate:"required"`
	UnitPrice    		int 				`json:"unitPrice" validate:"required"`
	TotalPrice    		int 				`json:"totalPrice" validate:"required"`
	Status    			string 				`json:"status" validate:"required"`
	CreatedDate    		time.Time			`json:"createdDate"`
	ModifiedDate    	time.Time			`json:"modifiedDate"`
}

func FetchAllPo() (Response, error) {
	var obj Po
	var arrows []Po
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM po"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.PoCode, &obj.GoodsCode, &obj.SupplierCode, &obj.UserCreated, &obj.UserApproved,
						&obj.PoQty, &obj.Currency, &obj.UnitPrice, &obj.TotalPrice, &obj.Status, &obj.CreatedDate,
						&obj.ModifiedDate)
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

func StorePo(poCode string, goodsCode string, supplierCode string, userCreated int, userApproved sql.NullInt64,
			 poQty int, currency string, unitPrice int, totalPrice int, status string,
			 createdDate time.Time) (Response, error) {
	var res Response

	v := validator.New()

	po := Po{
		PoCode				: poCode,
		GoodsCode			: goodsCode,
		SupplierCode		: supplierCode,
		UserCreated			: userCreated,
		UserApproved		: userApproved,
		PoQty				: poQty,
		Currency			: currency,
		UnitPrice			: unitPrice,
		TotalPrice			: totalPrice,
		Status				: status,
		CreatedDate			: createdDate,
	}

	err := v.Struct(po)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "INSERT po (poCode, goodsCode, supplierCode, userCreated," +
					"userApproved, poQty, currency, unitPrice, totalPrice, status," +
					"createdDate) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(poCode, goodsCode, supplierCode, userCreated, userApproved,
							 poQty, currency, unitPrice, totalPrice, status, createdDate)
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

func UpdatePo(poCode string, goodsCode string, supplierCode string, userCreated int, userApproved sql.NullInt64,
			  poQty int, currency string, unitPrice int, totalPrice int, status string,
			  modifiedDate time.Time) (Response, error) {
	var res Response

	v := validator.New()

	po := Po{
		PoCode				: poCode,
		GoodsCode			: goodsCode,
		SupplierCode		: supplierCode,
		UserCreated			: userCreated,
		UserApproved		: userApproved,
		PoQty				: poQty,
		Currency			: currency,
		UnitPrice			: unitPrice,
		TotalPrice			: totalPrice,
		Status				: status,
		ModifiedDate		: modifiedDate,
	}

	err := v.Struct(po)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "UPDATE po SET goodsCode = ?, supplierCode = ?, userCreated = ?," +
					"userApproved = ?, poQty = ?, currency = ?, unitPrice = ?, totalPrice = ?," +
					"status = ?,  modifiedDate = ? WHERE poCode = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(goodsCode, supplierCode, userCreated, userApproved, poQty,
							 currency, unitPrice, totalPrice, status, modifiedDate, poCode)
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

func DeletePo(poCode string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM po WHERE poCode = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(poCode)
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