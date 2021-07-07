package models

import (
	validator "github.com/go-playground/validator"
	"inventory-go/db"
	"net/http"
	"time"
)

type Audit struct {
	AuditID      		int    		`json:"auditID"`
	CreatedDate    		time.Time	`json:"createdDate"`
	Auditor    			int 		`json:"auditor" validate:"required"`
}

func FetchAllAudit() (Response, error) {
	var obj Audit
	var arrows []Audit
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM audit"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.AuditID, &obj.CreatedDate, &obj.Auditor)
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

func StoreAudit(createdDate time.Time, auditor int) (Response, error) {
	var res Response

	v := validator.New()

	audit := Audit{
		CreatedDate		: createdDate,
		Auditor			: auditor,
	}

	err := v.Struct(audit)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "INSERT audit (createdDate, auditor) VALUES (?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(createdDate, auditor)
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
