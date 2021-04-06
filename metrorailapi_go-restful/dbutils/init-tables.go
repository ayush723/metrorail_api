package dbutils

import (
	"database/sql"
	"log"
)

func Initialize(dbDriver *sql.DB) {
	statement, driverError := dbDriver.Prepare(train)
	if driverError != nil {
		log.Println(driverError)
	}
	//create train table
	_, statementError := statement.Exec()
	if statementError != nil {
		log.Println("table already exists!")
	}
	statement, _ = dbDriver.Prepare(station)
	statement.Exec()
	statement, _ = dbDriver.Prepare(schedule)
	statement.Exec()
	log.Println("all tables created/initialized successfully!")
}
