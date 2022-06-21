// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	"github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/internal/generated/models"
)

// NewServiceEditParams creates a new ServiceEditParams object
//
// There are no default values defined in the spec.
func NewServiceEditParams() ServiceEditParams {

	return ServiceEditParams{}
}

// ServiceEditParams contains all the bound params for the service edit operation
// typically these are obtained from a http.Request
//
// swagger:parameters serviceEdit
type ServiceEditParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*service Resource ID
	  Required: true
	  Max Length: 120
	  Min Length: 3
	  Pattern: [a-z0-9]([-a-z0-9]*[a-z0-9])?
	  In: path
	*/
	ServiceID string
	/*service item
	  Required: true
	  In: body
	*/
	ServiceItem *models.Service
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewServiceEditParams() beforehand.
func (o *ServiceEditParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rServiceID, rhkServiceID, _ := route.Params.GetOK("ServiceID")
	if err := o.bindServiceID(rServiceID, rhkServiceID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Service
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("serviceItem", "body", ""))
			} else {
				res = append(res, errors.NewParseError("serviceItem", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.ServiceItem = &body
			}
		}
	} else {
		res = append(res, errors.Required("serviceItem", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindServiceID binds and validates parameter ServiceID from path.
func (o *ServiceEditParams) bindServiceID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ServiceID = raw

	if err := o.validateServiceID(formats); err != nil {
		return err
	}

	return nil
}

// validateServiceID carries on validations for parameter ServiceID
func (o *ServiceEditParams) validateServiceID(formats strfmt.Registry) error {

	if err := validate.MinLength("ServiceID", "path", o.ServiceID, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("ServiceID", "path", o.ServiceID, 120); err != nil {
		return err
	}

	if err := validate.Pattern("ServiceID", "path", o.ServiceID, `[a-z0-9]([-a-z0-9]*[a-z0-9])?`); err != nil {
		return err
	}

	return nil
}