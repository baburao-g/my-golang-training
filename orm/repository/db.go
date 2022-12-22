package repository

import (
	"log"
	"orm/model"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	dsn := "sqlserver://SA:BrU@u2@@9StepvApp@206.189.129.40?database=demodb"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error, connecting a db")
	}
	db.AutoMigrate(&model.User{}, &model.Product{})
	return db
}
