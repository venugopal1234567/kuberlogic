// Code generated by go-swagger; DO NOT EDIT.

package restore

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/pkg/generated/models"
)

// RestoreDeleteHandlerFunc turns a function with the right signature into a restore delete handler
type RestoreDeleteHandlerFunc func(RestoreDeleteParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn RestoreDeleteHandlerFunc) Handle(params RestoreDeleteParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// RestoreDeleteHandler interface for that can handle valid restore delete params
type RestoreDeleteHandler interface {
	Handle(RestoreDeleteParams, *models.Principal) middleware.Responder
}

// NewRestoreDelete creates a new http.Handler for the restore delete operation
func NewRestoreDelete(ctx *middleware.Context, handler RestoreDeleteHandler) *RestoreDelete {
	return &RestoreDelete{Context: ctx, Handler: handler}
}

/*
	RestoreDelete swagger:route DELETE /restores/{RestoreID}/ restore restoreDelete

deletes a restore item

Deletes a restore object
*/
type RestoreDelete struct {
	Context *middleware.Context
	Handler RestoreDeleteHandler
}

func (o *RestoreDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewRestoreDeleteParams()
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
