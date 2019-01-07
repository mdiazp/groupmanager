package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mdiazp/gm/server/api"
	"github.com/mdiazp/gm/server/api/routes"
	"github.com/mdiazp/gm/server/conf"
	dbH "github.com/mdiazp/gm/server/db/handlers"
)

func main() {
	var (
		configPath  string
		environment string
		db          dbH.Handler
		logFile     *os.File
		apiBase     api.Base
		e           error
	)
	flag.StringVar(&configPath, "configpath", "/etc/gm-api", "Direccion del fichero de configuracion.")
	flag.StringVar(&environment, "env", "prod", "Entorno de ejecucion")
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
	defer db.Close()

	//JWT Handler
	jwth := api.NewJWTHandler(&config.JWTConfig)

	// LogFile
	logFile, e = os.Create(config.LogsDirectory + "/logs.log")
	if e != nil {
		log.Fatalf("Fail at create log file: %s", e.Error())
		panic(e)
	}
	logFile.Close()
	logFile, e = os.OpenFile(config.LogsDirectory+"/logs.log", os.O_RDWR|os.O_APPEND, 0660)
	if e != nil {
		log.Fatalf("Fail at open log file: %s", e.Error())
		panic(e)
	}
	defer logFile.Close()

	// ApiBase
	apiBase = api.NewBase(db, logFile, jwth, config.ADConfig, config.PublicFolderPath, environment)
	router := routes.Router(apiBase)

	fmt.Println("Running at " + config.Host + ":" + config.Port)

	// Run Server
	server := &http.Server{
		Addr:           config.Host + ":" + config.Port,
		Handler:        router,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: config.MaxHeaderBytes,
	}
	e = server.ListenAndServe()
	log.Fatalf("Server was down by: %s", e.Error())
}
