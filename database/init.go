package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB

// InitDB creates (if not created) and connects to the database
func InitDB() error {
	if err := os.MkdirAll("./lightcom", os.ModePerm); err != nil {
		return err
	}

	var err error
	if db, err = gorm.Open(sqlite.Open("./lightcom/database.sql")); err != nil {
		return err
	}

	db.AutoMigrate(&UserModel{}, &AccessTokenModel{}, &KeysModel{}, &MessageModel{})
	if err = GenerateKeys(); err != nil {
		return err
	}

	return nil
}
