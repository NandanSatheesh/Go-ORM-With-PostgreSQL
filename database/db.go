package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func GetDatabaseHandler() *gorm.DB {
	connectionString := fmt.Sprintf("%s://%s:%s@%s:%s/%s", DATABASE_VENDOR, USERNAME, PASSWORD, CONNECTION_URL, PORT_NUMBER, DATABASE_NAME)
	dbHandler, err := gorm.Open("postgres", connectionString)
	utils.CheckError(err)
	return dbHandler
}
