package db

import (
	"Test_MNC_2/lib/env"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "gorm.io/gorm/logger"
)

func Postgresql() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn()), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
}

func dsn() string {
	host := "host=" + env.String("DATABASE_HOST", "localhost")
	port := "port=" + env.String("DATABASE_PORT", "5432")
	dbname := "dbname=" + env.String("DATABASE_NAME", "bank-api")
	user := "user=" + env.String("DATABASE_USER", "postgres")
	password := "password=" + env.String("DATABASE_PASSWORD", "postgreaff")
	return fmt.Sprintln(host, port, dbname, user, password)
}
