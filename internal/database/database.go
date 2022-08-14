package database

import (
	"context"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database interface {
	ConnectDatabase() gorm.DB
}

type database struct {
	ctx context.Context
}

func NewDatabase(ctx context.Context) Database {
	return &database{
		ctx,
	}
}

func (r *database) ConnectDatabase() gorm.DB {
	db, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database")
	}
	return *db
}
