package db

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

var dbConnection *gorm.DB

func Connect() (*gorm.DB, error) {

	if dbConnection != nil {
		return dbConnection, nil
	}

	var err error
	dbConnection, err = gorm.Open(sqlite.Open(os.Getenv("DB")), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		return nil, fmt.Errorf("Faild to initialize a new db connection : %v", err)

	}

	return dbConnection, nil
}

func GetDB() (*gorm.DB, error) {
	return dbConnection, nil
}
