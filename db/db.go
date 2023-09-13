package db

import (
	model "github.com/bytedeveloperr/hng-stage-2/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.Person{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
