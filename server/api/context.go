package api

import (
	"context"
	"net/http"

	"github.com/mdiazp/groupmanager/server/db/models"
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
func (b *base) ContextReadAuthor(r *http.Request) *models.User {
	x := ctxR(r, ContextVarResponse)
	if author, ok := x.(*models.User); ok {
		return author
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
