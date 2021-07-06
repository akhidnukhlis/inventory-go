package models

import (
	"inventory-go/db"
	validator "github.com/go-playground/validator"
	"net/http"
	"time"
)

type Source struct {
	SourceId      	int    		`json:"sourceId"`
	SourceName    	string 		`json:"sourceName" validate:"required"`
	CreatedDate    	time.Time	`json:"createdDate"`
	ModifiedDate    time.Time	`json:"modifiedDate"`
}

func FetchAllSource() (Response, error) {
	var obj Source
	var arrows []Source
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM source"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.SourceId, &obj.SourceName, &obj.CreatedDate, &obj.ModifiedDate)
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

func StoreSource(sourceName string, createdDate time.Time) (Response, error) {
	var res Response

	v := validator.New()

	sou := Source{
		SourceName		: sourceName,
		CreatedDate		: createdDate,
	}

	err := v.Struct(sou)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "INSERT source (sourceName, createdDate) VALUES (?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(sourceName, createdDate)
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

func UpdateSource(sourceID int, sourceName string, modifiedDate time.Time) (Response, error) {
	var res Response

	v := validator.New()

	sou := Source{
		SourceId		: sourceID,
		SourceName		: sourceName,
		ModifiedDate	: modifiedDate,
	}

	err := v.Struct(sou)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "UPDATE source SET sourceName = ?, modifiedDate = ? WHERE sourceID = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(sourceName, modifiedDate, sourceID)
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

func DeleteSource(sourceID int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM source WHERE sourceID = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(sourceID)
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