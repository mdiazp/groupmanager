package controllers

import (
	"fmt"
	"net/http"

	"github.com/mdiazp/groupmanager/server/api"
)

// AccountController ...
type AccountController interface {
	api.Base
	Login(w http.ResponseWriter, r *http.Request)
	Profile(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

// NewAccountController ...
func NewAccountController(base api.Base) AccountController {
	return &accountController{
		Base: base,
	}
}

/////////////////////////////////////////////////////////////////////

type accountController struct {
	api.Base
}

func (c *accountController) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login")
	panic(fmt.Errorf("Not implemented yet"))
}

func (c *accountController) Profile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Profile")
	e := fmt.Errorf("Not implemented yet")
	c.WE(e, 501, "accountController.Profile")
}

func (c *accountController) Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Logout")
	// panic(fmt.Errorf("Not implemented yet"))
}
