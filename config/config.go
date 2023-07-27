package config

import (
	"fmt"

	"gorm.io/gorm"
)

// variavel 'global'
var (
	db     *gorm.DB
	logger *Logger

// logger *Logger
)

func Init() error {
	var err error

	db, err = InitializeSqlite()

	if err != nil {
		return fmt.Errorf("Error initializing sqlite: %v", err)
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}
