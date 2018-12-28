package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"

	dbH "github.com/mdiazp/gm/server/db/handlers"
	"github.com/mdiazp/gm/server/db/models"
)

// Base ...
type Base interface {
	DB() dbH.Handler
	Logger() *log.Logger
	JWTHandler() JWTHandler
	PublicFolderPath() string
	GetEnv() string

	ReadJSON(w http.ResponseWriter, r *http.Request, objs ...interface{})

	GetPInt(w http.ResponseWriter, r *http.Request, vname string) int
	GetPString(w http.ResponseWriter, r *http.Request, vname string) string

	GetQInt(w http.ResponseWriter, r *http.Request, vname string, required ...bool) *int
	GetQBool(w http.ResponseWriter, r *http.Request, vname string, required ...bool) *bool
	GetQString(w http.ResponseWriter, r *http.Request, vname string, required ...bool) *string
	GetQPaginator(w http.ResponseWriter, r *http.Request) *dbH.Paginator
	GetQOrderBy(w http.ResponseWriter, r *http.Request) *dbH.OrderBy

	Validate(w http.ResponseWriter, objs ...Valid)

	ContextWriteAuthor(r *http.Request, author *models.User)
	ContextReadAuthor(w http.ResponseWriter, r *http.Request, required ...bool) *models.User

	WE(w http.ResponseWriter, e error, status int)
	WR(w http.ResponseWriter, status int, body interface{})
}

// NewBase ...
func NewBase(db dbH.Handler, logFile *os.File, jwth JWTHandler, publicFolderPath string, env string) Base {
	return &base{
		db:               db,
		logger:           NewLogger(logFile),
		publicFolderPath: publicFolderPath,
		jwth:             jwth,
		env:              env,
	}
}

type Valid interface {
	Valid() error
}

///////////////////////////////////////////////////////////////////////////

type base struct {
	db               dbH.Handler
	logger           *log.Logger
	jwth             JWTHandler
	publicFolderPath string
	env              string
}

func (b *base) DB() dbH.Handler {
	return b.db
}

func (b *base) Logger() *log.Logger {
	return b.logger
}

func (b *base) JWTHandler() JWTHandler {
	return b.jwth
}

func (b *base) PublicFolderPath() string {
	return b.publicFolderPath
}

func (b *base) GetEnv() string {
	return b.env
}

func (b *base) ReadJSON(w http.ResponseWriter, r *http.Request, objs ...interface{}) {
	decoder := json.NewDecoder(r.Body)
	for _, obj := range objs {
		e := decoder.Decode(obj)
		if e != nil {
			e = fmt.Errorf("Request's body has wrong format: %s", e.Error())
			b.WE(w, e, 400)
		}
	}
}

func (b *base) Validate(w http.ResponseWriter, objs ...Valid) {
	for _, obj := range objs {
		e := obj.Valid()
		b.WE(w, e, 400)
	}
}

func (b *base) GetPInt(w http.ResponseWriter, r *http.Request, vname string) int {
	vs := b.GetPString(w, r, vname)
	v, e := strconv.Atoi(vs)
	if e != nil {
		b.WE(w, fmt.Errorf("%s from url's path must be an integer: %s", vname, e.Error()), 400)
	}
	return v
}

func (b *base) GetPString(w http.ResponseWriter, r *http.Request, vname string) string {
	vars := mux.Vars(r)
	v, ok := vars[vname]
	if !ok {
		b.WE(w, fmt.Errorf("Required %s from url's path is missed", vname), 400)
	}
	return v
}

func (b *base) GetQInt(w http.ResponseWriter, r *http.Request, vname string, required ...bool) *int {
	vs := b.GetQString(w, r, vname, required...)
	if vs == nil {
		return nil
	}

	v, e := strconv.Atoi(*vs)
	if e != nil {
		b.WE(w, fmt.Errorf("Required %s from URL Query must be an integer: %s", vname, e.Error()), 400)
	}
	return &v
}

func (b *base) GetQBool(w http.ResponseWriter, r *http.Request, vname string, required ...bool) *bool {
	vs := b.GetQString(w, r, vname, required...)
	if vs == nil {
		return nil
	}

	v, e := strconv.ParseBool(*vs)
	if e != nil {
		b.WE(w, fmt.Errorf("Required %s from URL Query must be a bool: %s", vname, e.Error()), 400)
	}
	return &v
}

func (b *base) GetQString(w http.ResponseWriter, r *http.Request, vname string, required ...bool) *string {
	v := r.URL.Query().Get(vname)
	req := false
	if len(required) > 0 {
		req = required[0]
	}
	if v == "" && req {
		b.WE(w, fmt.Errorf("Required %s from URL Query is missed", vname), 400)
	}
	if v == "" {
		return nil
	}
	return &v
}

func (b *base) GetQPaginator(w http.ResponseWriter, r *http.Request) *dbH.Paginator {
	p := dbH.Paginator{}
	limit := b.GetQInt(w, r, "limit", false)
	offset := b.GetQInt(w, r, "offset", false)
	if limit == nil || offset == nil {
		return nil
	}
	p.Limit = *limit
	p.Offset = *offset
	return &p
}

func (b *base) GetQOrderBy(w http.ResponseWriter, r *http.Request) *dbH.OrderBy {
	ob := dbH.OrderBy{}
	by := b.GetQString(w, r, "orderby", false)
	desc := b.GetQBool(w, r, "desc", false)
	if by == nil {
		return nil
	}
	if desc == nil {
		tmp := false
		desc = &tmp
	}
	ob.By = *by
	ob.DESC = *desc
	return &ob
}

func (b *base) WE(w http.ResponseWriter, e error, status int) {
	if e == nil {
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	body, _ := json.Marshal(fmt.Sprintf("%s", e.Error()))
	w.Write(body)

	panic(
		Error{
			Status:   status,
			Location: WAI(2),
			error:    e,
		},
	)
}

func (b *base) WR(w http.ResponseWriter, status int, body interface{}) {
	bod, e := json.Marshal(body)
	if e != nil {
		b.WE(w, e, 500)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	w.Write(bod)
}
