package config

import (
	"gorm.io/gorm"
)

// variavel 'global'
var (
	db     *gorm.DB
	logger *Logger

// logger *Logger
)

func Init() error {
	return nil
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}
