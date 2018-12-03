package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mdiazp/groupmanager/server/api"
	"github.com/mdiazp/groupmanager/server/api/routes"
	"github.com/mdiazp/groupmanager/server/conf"
	dbH "github.com/mdiazp/groupmanager/server/db/handlers"
	ldapH "github.com/mdiazp/groupmanager/server/ldap/handlers"
)

func main() {
	var (
		configPath  string
		environment string
		db          dbH.Handler
		ldap        ldapH.Handler
		logFile     *os.File
		apiBase     api.Base
		e           error
	)
	flag.StringVar(&configPath, "configpath", "/home/kino/my_configs/groupmanager", "Direccion del fichero de configuracion.")
	flag.StringVar(&environment, "env", "dev", "Entorno de ejecucion")
	flag.Parse()

	// Load Configuration
	config, e := conf.LoadConfiguration(configPath, environment)
	if e != nil {
		log.Fatalf("Fail at conf.LoadConfiguration: %s", e.Error())
		panic(e)
	}

	// Database Handler
	db, e = dbH.NewHandler(&config.DatabaseConfig)
	if e != nil {
		log.Fatalf("Fail at dbH.NewHandler: %s", e.Error())
		panic(e)
	}

	// Ldap Handler
	ldap, e = ldapH.NewHandler(&config.ADConfig)
	if e != nil {
		log.Fatalf("Fail at ldapH.NewHandler: %s", e.Error())
		panic(e)
	}

	// LogFile
	logFile, e = os.Create(config.LogFile)
	if e != nil {
		log.Fatalf("Fail at create log file: %s", e.Error())
		panic(e)
	}
	logFile.Close()
	logFile, e = os.OpenFile(config.LogFile, os.O_RDWR|os.O_APPEND, 0660)
	defer logFile.Close()
	if e != nil {
		log.Fatalf("Fail at open log file: %s", e.Error())
		panic(e)
	}

	// ApiBase
	apiBase = api.NewBase(db, ldap, logFile)
	router := routes.Router(apiBase)

	fmt.Println("Running at api.groupmanager.local:8080")

	// Run Server
	server := &http.Server{
		Addr:           "api.groupmanager.local:8080",
		Handler:        router,
		ReadTimeout:    time.Duration(1 * int64(time.Second)),
		WriteTimeout:   time.Duration(4 * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
