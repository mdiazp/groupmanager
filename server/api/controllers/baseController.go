package controllers

import (
	"net/http"
)

// BaseController ...
type BaseController interface {
	http.Handler
	GetMethods() []string
	GetRoute() string
	GetAccess() Permission
}
