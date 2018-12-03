package tests

import (
	"fmt"
	"testing"

	"github.com/mdiazp/groupmanager/server/conf"
	dbhandlers "github.com/mdiazp/groupmanager/server/db/handlers"
)

// CONFIG ...
var CONFIG *conf.Configuration

func init() {
	var e error
	configPath := "/home/kino/my_configs/groupmanager"
	CONFIG, e = conf.LoadConfiguration(configPath)
	if e != nil {
		panic(fmt.Errorf("Fail loading configuration: %s", e.Error()))
	}
}

// GetDBHandler ...
func GetDBHandler(t *testing.T) dbhandlers.ModelHandler {
	db, e := dbhandlers.NewModelHandler(CONFIG)
	if e != nil {
		t.Fatalf("Fail newModelHandler: %s", e.Error())
	}

	return db
}
