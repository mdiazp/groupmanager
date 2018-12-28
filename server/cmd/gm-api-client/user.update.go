package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mdiazp/gm/server/db/models"
)

// UserUpdate ...
func userUpdate() {
	user := &models.User{
		ID:       1,
		Username: "Kino",
		Password: "123",
		Rol:      models.RolAdmin,
		PersonalInfo: models.PersonalInfo{
			Name:        "Manuel Alejandro Diaz Perez (El Kino)",
			Address:     "Virtudes",
			Phone:       "48766081",
			Description: "Software Developer",
			OwnerID:     1,
			OwnerType:   "sm_user",
		},
	}

	body, e := json.Marshal(user)
	pe(e)

	r, e := http.NewRequest("PATCH", apiURL+"user/1", bytes.NewReader(body))
	//r.Header.Set("Content-Type", ct)
	pe(e)

	resp, e := http.DefaultClient.Do(r)

	bodyR, e := ioutil.ReadAll(resp.Body)
	WR(resp, bodyR)
	pe(e)
}
