package handlers

import (
	"fmt"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/mdiazp/gm/server/db/models/migrations"
)

// Handler ...
type Handler interface {
	UserHandler
	GroupHandler
	GroupAdminHandler
	GroupADUserHandler

	Close() error
}

// DatabaseConfig ...
type DatabaseConfig interface {
	GetDBHost() string
	GetDBPort() int
	GetDBName() string
	GetDBUser() string
	GetDBPassword() string
	GetDBDialect() string
}

// NewHandler ...
func NewHandler(conf DatabaseConfig) (Handler, error) {
	dbHost := conf.GetDBHost()
	dbPort := conf.GetDBPort()
	dbName := conf.GetDBName()
	dbUser := conf.GetDBUser()
	dbPassword := conf.GetDBPassword()
	dbDriver := conf.GetDBDialect()

	conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbName, dbPassword)

	db, e := gorm.Open(dbDriver, conn)
	if e != nil {
		return nil, fmt.Errorf("Fail opennig database: %s", e.Error())
	}

	e = migrations.InitDB(db)
	if e != nil {
		return nil, fmt.Errorf("Fail at InitDB migrations: %s", e.Error())
	}

	return &handler{
		DB: db,
	}, nil
}

//////////////////////////////////////////////////////////////////////////

type handler struct {
	*gorm.DB
}

func (h *handler) Close() error {
	return h.DB.Close()
}
