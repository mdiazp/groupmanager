package main

import (
	"fmt"

	"github.com/mdiazp/gm/server/conf"
	dbhandlers "github.com/mdiazp/gm/server/db/handlers"
	"github.com/mdiazp/gm/server/db/models"
	"golang.org/x/exp/rand"
)

var (
	config *conf.Configuration
	db     dbhandlers.Handler
)

func init() {
	var e error
	configPath := "/home/kino/my_configs/gm"
	config, e = conf.LoadConfiguration(configPath, "dev")
	if e != nil {
		panic(fmt.Errorf("Fail loading configuration: %s", e.Error()))
	}
	db, e = dbhandlers.NewHandler(config)

	if e != nil {
		panic(fmt.Errorf("Fail at dbhandlers.NewHandler: %s", e.Error()))
	}

}

func randString(len int) string {
	abc := "abcdefghijklmnopqrstuvwxyz"

	s := ""
	for i := 0; i < len; i++ {
		s += string(abc[rand.Int()%26])
	}

	return s
}

func main() {
	for i := 0; i < 100; i++ {
		act := true
		if rand.Int()%2 == 0 {
			act = false
		}
		g := models.Group{
			Name:        randString(10),
			Description: "Some description",
			Actived:     act,
		}

		e := db.CreateGroup(&g)
		if e != nil {
			panic(e)
		}
	}
}
