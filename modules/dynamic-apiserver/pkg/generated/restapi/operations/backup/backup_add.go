// Code generated by go-swagger; DO NOT EDIT.

package backup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/pkg/generated/models"
)

// BackupAddHandlerFunc turns a function with the right signature into a backup add handler
type BackupAddHandlerFunc func(BackupAddParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn BackupAddHandlerFunc) Handle(params BackupAddParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// BackupAddHandler interface for that can handle valid backup add params
type BackupAddHandler interface {
	Handle(BackupAddParams, *models.Principal) middleware.Responder
}

// NewBackupAdd creates a new http.Handler for the backup add operation
func NewBackupAdd(ctx *middleware.Context, handler BackupAddHandler) *BackupAdd {
	return &BackupAdd{Context: ctx, Handler: handler}
}

/*
	BackupAdd swagger:route POST /backups/ backup backupAdd

create backup object

Create backup object
*/
type BackupAdd struct {
	Context *middleware.Context
	Handler BackupAddHandler
}

func (o *BackupAdd) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewBackupAddParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
