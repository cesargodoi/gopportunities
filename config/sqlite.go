package config

import (
	"os"

	"github.com/cesargodoi/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	dbPath := "./db/main.db"

	// init logger
	logger := GetLogger("sqlite")

	// check if DB exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("DB file not found, creating...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	// create DB and connect
	db, err := gorm.Open(sqlite.Open(
		dbPath), &gorm.Config{},
	)
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	// migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return nil, err
	}

	// returning DB
	return db, nil
}
