package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	dbH "github.com/mdiazp/groupmanager/server/db/handlers"
	"github.com/mdiazp/groupmanager/server/db/models"
	ldapH "github.com/mdiazp/groupmanager/server/ldap/handlers"
)

// Base ...
type Base interface {
	DB() dbH.Handler
	Ldap() ldapH.Handler
	Logger() *log.Logger
	ContextWriteResponse(r *http.Request, res *Response)
	ContextReadResponse(r *http.Request) *Response
	ContextWriteAuthor(r *http.Request, author *models.User)
	ContextReadAuthor(r *http.Request) *models.User
	WE(e error, status int, s ...string)
	WR(r *http.Request, status int, body interface{})
}

// NewBase ...
func NewBase(db dbH.Handler, ldap ldapH.Handler, logFile *os.File) Base {
	return &base{
		db:     db,
		ldap:   ldap,
		logger: NewLogger(logFile),
	}
}

///////////////////////////////////////////////////////////////////////////

type base struct {
	db     dbH.Handler
	ldap   ldapH.Handler
	logger *log.Logger
}

func (b *base) DB() dbH.Handler {
	return b.db
}

func (b *base) Ldap() ldapH.Handler {
	return b.ldap
}

func (b *base) Logger() *log.Logger {
	return b.logger
}

func (b *base) WE(e error, status int, s ...string) {
	if e == nil {
		return
	}
	if len(s) > 0 {
		e = fmt.Errorf("%s: %s", s[0], e.Error())
	}
	panic(
		Error{
			Status: status,
			error:  e,
		},
	)
}

func (b *base) WR(r *http.Request, status int, body interface{}) {
	b.ContextWriteResponse(
		r,
		&Response{
			Status: status,
			Body:   body,
		},
	)
}
