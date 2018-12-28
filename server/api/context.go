package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mdiazp/gm/server/db/models"
)

// ContextVar ...
type ContextVar string

const (
	// ContextVarAuthor ...
	ContextVarAuthor ContextVar = "Author"
	// ContextVarResponse ...
	ContextVarResponse ContextVar = "Response"
)

// ContextWriteResponse ...
func (b *base) ContextWriteResponse(r *http.Request, res *Response) {
	ctxW(r, ContextVarResponse, res)
}

// ContextReadResponse ...
func (b *base) ContextReadResponse(r *http.Request) *Response {
	x := ctxR(r, ContextVarResponse)
	if res, ok := x.(*Response); ok {
		return res
	}
	return nil
}

// ContextWriteAuthor ...
func (b *base) ContextWriteAuthor(r *http.Request, author *models.User) {
	ctxW(r, ContextVarAuthor, author)
}

// ContextReadAuthor ...
func (b *base) ContextReadAuthor(w http.ResponseWriter, r *http.Request, required ...bool) *models.User {
	x := ctxR(r, ContextVarAuthor)
	if author, ok := x.(*models.User); ok {
		return author
	}
	if len(required) > 0 && required[0] {
		b.WE(w, fmt.Errorf("User was not found in context"), 500)
	}
	return nil
}

func ctxR(r *http.Request, cv ContextVar) interface{} {
	return r.Context().Value(cv)
}

func ctxW(r *http.Request, cv ContextVar, value interface{}) {
	ctx := context.WithValue(r.Context(), cv, value)
	*r = *(r.WithContext(ctx))
}
