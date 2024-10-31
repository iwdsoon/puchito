package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func InitDb(SqlString string) {
	fmt.Println(SqlString)
	conn, err := gorm.Open(mysql.Open(SqlString), &gorm.Config{
		SkipDefaultTransaction: false,
		PrepareStmt:            false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	db = conn
	db.Logger = db.Logger.LogMode(logger.Info)

	fmt.Println("conn")
}
func GetDb() *gorm.DB {
	return db
}