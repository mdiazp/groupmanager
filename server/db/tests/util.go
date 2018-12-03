package tests

import (
	"math/rand"

	"github.com/mdiazp/groupmanager/server/db/models"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// GetRandString ...
func GetRandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// GetRandomUserRol ...
func GetRandomUserRol() models.Rol {
	x := rand.Int() % 2
	if x == 0 {
		return models.RolSuperadmin
	} else {
		return models.RolGroupadmin
	}
}
