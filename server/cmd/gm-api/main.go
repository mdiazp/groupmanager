package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
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

	///////////////////////////////////////////////////////////////////////////////////
	//JWT Handler
	keyFile, e := os.OpenFile(configPath+"/key.perm", os.O_RDWR|os.O_CREATE, 0660)
	if e != nil {
		log.Fatalf("Fail opening key.pem file")
		panic(e)
	}
	defer keyFile.Close()
	keyBytes := FileToBytes(keyFile)
	keyBlock, _ := pem.Decode(keyBytes)
	var pKey *rsa.PrivateKey
	alreadyDefined := (keyBlock != nil)
	if !alreadyDefined {
		pKey, e = rsa.GenerateKey(rand.Reader, 512)
		if e != nil {
			log.Fatalf("Fail at rsa.GenerateKey")
			panic(e)
		}
	} else {
		pKey, e = x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
		if e != nil {
			log.Fatalf("Fail at parse privateKey")
			panic(e)
		}
	}

	jwth := api.NewJWTHandler(pKey)

	// Save the pkey if it was not defined
	if !alreadyDefined {
		keyBlock = &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(pKey),
		}
		if err := pem.Encode(keyFile, keyBlock); err != nil {
			log.Fatalf("Failed to write data to key.pem: %s", err)
			panic(e)
		}
	}
	///////////////////////////////////////////////////////////////////////////////////

	// LogFile
	tim := time.Now()
	pln := fmt.Sprintf("%d-%02d-%02dT%02d-%02d-%02d",
		tim.Year(), tim.Month(), tim.Day(),
		tim.Hour(), tim.Minute(), tim.Second())
	logFile, e = os.OpenFile(config.LogsDirectory+"/"+pln+"-gm-logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0660)
	if e != nil {
		log.Fatalf("Fail at open log file: %s", e.Error())
		panic(e)
	}
	defer logFile.Close()

	// ApiBase
	apiBase = api.NewBase(
		db, logFile, jwth, config.ADConfig,
		config.PublicFolderPath, environment,
		config.UserRootPassword,
	)
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
