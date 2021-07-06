package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"inventory-go/config"
	"log"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()

	// mysql connection setting
	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME + "?parseTime=true"
	db, err = sql.Open("mysql", connectionString)

	// postgresql connection setting
	// connectionPsql := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", conf.DB_HOST, conf.DB_PORT, conf.DB_USERNAME, conf.DB_PASSWORD, conf.DB_NAME)
	// db, err := sql.Open("postgres", connectionPsql)
	// defer db.Close()

	if err != nil {
		log.Println(err)
		panic("connectionString error..")
	}

	err = db.Ping()
	if err != nil {
		log.Println(err)
		panic("DSN Invalid")
	}
}

func CreateCon() *sql.DB {
	return db
}
