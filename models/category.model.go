package models

import (
	validator "github.com/go-playground/validator"
	"inventory-go/db"
	"net/http"
	"time"
)

type Category struct {
	CategoryCode      	string    	`json:"categoryCode" validate:"required"`
	CategoryName    	string 		`json:"categoryName" validate:"required"`
	CreatedDate    		time.Time	`json:"createdDate"`
	ModifiedDate    	time.Time	`json:"modifiedDate"`
}

func FetchAllCategory() (Response, error) {
	var obj Category
	var arrows []Category
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM category"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.CategoryCode, &obj.CategoryName, &obj.CreatedDate, &obj.ModifiedDate)
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

func StoreCategory(categoryCode string, categoryName string, createdDate time.Time) (Response, error) {
	var res Response

	v := validator.New()

	cate := Category{
		CategoryCode		: categoryCode,
		CategoryName		: categoryName,
		CreatedDate			: createdDate,
	}

	err := v.Struct(cate)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "INSERT category (categoryCode, categoryName, createdDate) VALUES (?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(categoryCode, categoryName, createdDate)
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

func UpdateCategory(categoryCode string, categoryName string, modifiedDate time.Time) (Response, error) {
	var res Response

	v := validator.New()

	cate := Category{
		CategoryCode		: categoryCode,
		CategoryName		: categoryName,
		ModifiedDate		: modifiedDate,
	}

	err := v.Struct(cate)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "UPDATE category SET categoryName = ?, modifiedDate = ? WHERE categoryCode = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(categoryName, modifiedDate, categoryCode)
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

func DeleteCategory(categoryCode string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM category WHERE categoryCode = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(categoryCode)
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