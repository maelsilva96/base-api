package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

func GetDbTest() *gorm.DB {
	LoadEnvFile()
	strConn := os.Getenv("STR_CONN_TEST")
	db, err := gorm.Open(mysql.Open(strConn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
