package models

import (
	validator "github.com/go-playground/validator"
	"inventory-go/db"
	"net/http"
	"time"
)

type So struct {
	SoCode			string		`json:"soCode"`
	GoodsCode		string		`json:"goodsCode" validate:"required"`
	UserCreated		int			`json:"userCreated" validate:"required"`
	SourceID		int			`json:"sourceID" validate:"required"`
	Customer		string		`json:"customer" validate:"required"`
	SoQty			int			`json:"soQty" validate:"required"`
	Currency		string		`json:"currency" validate:"required"`
	UnitPrice		int			`json:"unitPrice" validate:"required"`
	TotalPrice		int			`json:"totalPrice" validate:"required"`
	Courier			string		`json:"courier" validate:"required"`
	DeliveryCharge	int			`json:"deliveryCharge" validate:"required"`
	Status			string		`json:"status"`
	CreatedDate		time.Time	`json:"createdDate"`
	ModifiedDate	time.Time	`json:"modifiedDate"`
}

func FetchAllSo() (Response, error) {
	var obj So
	var arrows []So
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM so"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.SoCode, &obj.GoodsCode, &obj.UserCreated, &obj.SourceID, &obj.Customer, &obj.SoQty,
						&obj.Currency, &obj.UnitPrice, &obj.TotalPrice, &obj.Courier, &obj.DeliveryCharge, &obj.Status,
						&obj.CreatedDate, &obj.ModifiedDate)
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

func StoreSo(soCode string, goodsCode string, userCreated int, sourceID int, customer string, soQty int, currency string,
	unitPrice int, totalPrice int, courier string, deliveryCharge int, status string, createdDate time.Time) (Response, error) {
	var res Response

	v := validator.New()

	po := So{
		SoCode				: soCode,
		GoodsCode			: goodsCode,
		UserCreated			: userCreated,
		SourceID			: sourceID,
		Customer			: customer,
		SoQty				: soQty,
		Currency			: currency,
		UnitPrice			: unitPrice,
		TotalPrice			: totalPrice,
		Courier				: courier,
		DeliveryCharge		: deliveryCharge,
		Status				: status,
		CreatedDate			: createdDate,
	}

	err := v.Struct(po)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "INSERT so (soCode, goodsCode, userCreated, sourceID, customer, soQty, currency, unitPrice, " +
					"totalPrice, courier, deliveryCharge, status, createdDate) " +
					"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(soCode, goodsCode, userCreated, sourceID, customer, soQty, currency, unitPrice, totalPrice,
							 courier, deliveryCharge, status, createdDate)
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

func UpdateSo(goodsCode string, userCreated int, sourceID int, customer string, soQty int, currency string, unitPrice int,
			  totalPrice int, courier string, deliveryCharge int, status string, modifiedDate time.Time, soCode string) (Response, error) {
	var res Response

	v := validator.New()

	po := So{
		GoodsCode			: goodsCode,
		UserCreated			: userCreated,
		SourceID			: sourceID,
		Customer			: customer,
		SoQty				: soQty,
		Currency			: currency,
		UnitPrice			: unitPrice,
		TotalPrice			: totalPrice,
		Courier				: courier,
		DeliveryCharge		: deliveryCharge,
		Status				: status,
		ModifiedDate		: modifiedDate,
		SoCode				: soCode,
	}

	err := v.Struct(po)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "UPDATE so SET goodsCode = ?, userCreated = ?, sourceID = ?, customer = ?, soQty = ?, currency = ?," +
					"unitPrice = ?, totalPrice = ?, courier = ?, deliveryCharge = ?, status = ?, modifiedDate = ? WHERE soCode = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(goodsCode, userCreated, sourceID, customer, soQty, currency, unitPrice, totalPrice,
							 courier, deliveryCharge, status, modifiedDate, soCode)
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

func DeleteSo(poCode string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM so WHERE soCode = ?"

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