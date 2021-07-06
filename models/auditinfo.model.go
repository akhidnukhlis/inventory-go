package models

import (
	"inventory-go/db"
	validator "github.com/go-playground/validator"
	"net/http"
)

type Auditing struct {
	AuditInfoID      	int    		`json:"auditInfoID" validate:"required"`
	AuditID    			int 		`json:"auditID" validate:"required"`
	GoodsCode    		string 		`json:"goodsCode" validate:"required"`
	SysStock    		int 		`json:"sysStock" validate:"required"`
	RealStock    		int 		`json:"realStock" validate:"required"`
	Note    			string 		`json:"note" validate:"required"`
}

func FetchAllAuditing() (Response, error) {
	var obj Auditing
	var arrows []Auditing
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM auditinfo"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.AuditInfoID, &obj.AuditID, &obj.GoodsCode, &obj.SysStock, &obj.RealStock, &obj.Note)
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

func StoreAuditing(auditInfoID int, auditID int, goodsCode string, sysStock int, realStock int, note string) (Response, error) {
	var res Response

	v := validator.New()

	audi := Auditing{
		AuditInfoID		: auditInfoID,
		AuditID			: auditID,
		GoodsCode		: goodsCode,
		SysStock		: sysStock,
		RealStock		: realStock,
		Note			: note,
	}

	err := v.Struct(audi)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "INSERT role (auditInfoID, auditID, goodsCode, sysStock, realStock, note) VALUES (?, ?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(auditInfoID, auditID, goodsCode, sysStock, realStock, note)
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