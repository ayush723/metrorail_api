package dbutils

import (
	"database/sql"
	"log"
)
package dbutils

func Initialize(dbDriver *sql.DB) {
	
const train = `
	CREATE TABLE IF NOT EXISTS train(
		ID INTEGER PRIAMRY KEY ,
		DRIVER_NAME VARCHAR(64) NULL,
		OPERATING_STATUS BOOLEAN
	)
`
const station = `
	CREATE TABLE IF NOT EXISTS station(
		ID INTEGER PRIMARY KEY ,
		NAME VARCHAR(64) NULL,
		OPENING_TIME TIME NULL,
		CLOSING_TIME TIME NULL
	)
`

const schedule = `
	CREATE TABLE IF NOT EXISTS schedule(
		ID INTEGER PRIMARY KEY ,
		TRAIN_ID INT,
		STATION_ID INT,
		ARRIVAL_TIME TIME,
		FOREIGN KEY (TRAIN_ID) REFERENCES train(ID),
		FOREIGN KEY (STATION_ID) REFERENCES station(ID)
	)
`
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
