package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// DatabaseInfo is the struct holding the  general information needed to access a database
type DatabaseInfo struct {
	Host         string
	Port         string
	User         string
	Password     string
	DatabaseName string
}

var DBConnection *sql.DB

// Connect connects to a database with the information provided in the DatabaseInfo struct
func (databaseInfo *DatabaseInfo) Connect() error {

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		databaseInfo.Host, databaseInfo.Port, databaseInfo.User, databaseInfo.Password, databaseInfo.DatabaseName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	DBConnection = db

	err = db.Ping()
	if err != nil {
		return err
	}

	return nil
}

/*
Disconnect disconnects the api from the database connection
*/
func (databaseInfo *DatabaseInfo) Disconnect() error {
	err := DBConnection.Close()
	if err != nil {
		return err
	}
	return nil
}
