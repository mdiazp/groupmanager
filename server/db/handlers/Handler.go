package handlers

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Handler ...
type Handler interface {
	UserHandler
	GroupHandler
	GroupAdminHandler
	GroupADUserHandler
}

// DatabaseConfig ...
type DatabaseConfig interface {
	GetDBName() string
	GetDBUser() string
	GetDBPassword() string
	GetDBDriver() string
}

// NewHandler ...
func NewHandler(conf DatabaseConfig) (Handler, error) {
	dbName := conf.GetDBName()
	dbUser := conf.GetDBUser()
	dbPassword := conf.GetDBPassword()
	dbDriver := conf.GetDBDriver()

	conn := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable",
		dbUser, dbName, dbPassword)

	db, e := sql.Open(dbDriver, conn)
	if e != nil {
		return nil, fmt.Errorf("Fail opennig database: %s", e.Error())
	}

	return &handler{
		DB: db,
	}, nil
}

type handler struct {
	DB *sql.DB
}
