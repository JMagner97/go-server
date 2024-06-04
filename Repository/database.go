package repository

import "database/sql"

var instance *sql.DB

func init() {
	instance = nil
}

func NewDatabase(driverName, dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	if instance == nil {
		instance = db
	}
	return instance, nil
}

func GetDatabaseInstance() *sql.DB {
	return instance
}
