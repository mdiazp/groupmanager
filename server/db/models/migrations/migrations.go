package migrations

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/mdiazp/gm/server/db/models"
)

// InitDB ...
func InitDB(db *gorm.DB) (e error) {
	gmModel := []models.GMModel{
		&models.GroupAdmin{},
		&models.GroupADUser{},
		&models.User{},
		&models.Group{},
	}

	model := make([]interface{}, 0)
	for _, x := range gmModel {
		model = append(model, x)
	}

	// db.DropTableIfExists(model...)
	// return
	db.SingularTable(true)
	e = db.AutoMigrate(model...).Error
	if e != nil {
		return fmt.Errorf("db.Automigrate: %s", e.Error())
	}

	for _, x := range gmModel {
		if e = x.AddSQLConstrainsts(db); e != nil {
			return e
		}
	}

	return
}
