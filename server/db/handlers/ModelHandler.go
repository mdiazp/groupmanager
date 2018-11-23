package handlers

import (
	"database/sql"
	"fmt"
)

// ModelHandler ...
type ModelHandler interface {
	UserDBHandler
	GroupAdminDBHandler
	GroupUserDBHandler
}

// DatabaseConfig ...
type DatabaseConfig interface {
	GetDBName() string
	GetDBUser() string
	GetDBPassword() string
	GetDBDriver() string
}

// NewModelHandler ...
func NewModelHandler(conf DatabaseConfig) ModelHandler {
	dbName := conf.GetDBName()
	dbUser := conf.GetDBUser()
	dbPassword := conf.GetDBPassword()
	dbDriver := conf.GetDBDriver()

	conn := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable",
		dbUser, dbName, dbPassword)

	db, e := sql.Open(dbDriver, conn)
	if e != nil {
		panic(e)
	}

	return &modelHandler{
		DB: db,
	}
}

type modelHandler struct {
	DB *sql.DB
}
