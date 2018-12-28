package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mdiazp/gm/server/db/models"
)

func userRetrieve() {
	resp, e := http.Get(apiURL + "user/1")
	pe(e)

	var user models.User
	body, e := ioutil.ReadAll(resp.Body)
	pe(e)
	e = json.Unmarshal(body, &user)
	WR(resp, body)
	pe(e)
}
