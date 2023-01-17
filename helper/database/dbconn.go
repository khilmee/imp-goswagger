package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewConnectionDB(database string, host string, user string, password string, port int) (*gorm.DB, error) {
	var dialect gorm.Dialector

	//default singular
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		host,
		port,
		user,
		database,
		password,
		"disable",
	)

	dialect = postgres.Open(dsn)

	db, err := gorm.Open(dialect, gormConfig)
	if err != nil {
		return nil, err
	}

	return db, nil
}
