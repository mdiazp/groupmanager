package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mdiazp/gm/server/db/models"
)

const (
	registerURL = "account/register"
)

// Register ...
func Register() {
	user := &models.User{
		Username: "manuel.diaz",
		Password: "123",
		Rol:      models.RolAdmin,
		PersonalInfo: models.PersonalInfo{
			Name:        "Manuel Alejandro Diaz Perez",
			Address:     "Montequin",
			Phone:       "48766081",
			Description: "Developer",
		},
	}

	body, e := json.Marshal(user)
	pe(e)

	resp, e := http.Post(apiURL+registerURL, ct, bytes.NewReader(body))
	pe(e)

	bodyR, e := ioutil.ReadAll(resp.Body)
	WR(resp, bodyR)
	pe(e)
}
