package infrastructure

import (
	"time"

	"github.com/sergiovirahonda/inventory-manager/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	SQL *gorm.DB
}

var Database *DB

var dbConn = &DB{}

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDbLifeTime = 5 * time.Minute

func NewDBInstance() *DB {
	dsn := config.GetPostgresConnectionString()
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(maxIdleDbConn)
	sqlDB.SetMaxOpenConns(maxOpenDbConn)
	sqlDB.SetConnMaxLifetime(maxDbLifeTime)
	dbConn.SQL = db
	return dbConn
}

func DbManager() *gorm.DB {
	return dbConn.SQL
}
