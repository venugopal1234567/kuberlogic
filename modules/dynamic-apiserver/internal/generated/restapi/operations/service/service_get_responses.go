// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/internal/generated/models"
)

// ServiceGetOKCode is the HTTP code returned for type ServiceGetOK
const ServiceGetOKCode int = 200

/*ServiceGetOK item edited

swagger:response serviceGetOK
*/
type ServiceGetOK struct {

	/*
	  In: Body
	*/
	Payload *models.Service `json:"body,omitempty"`
}

// NewServiceGetOK creates ServiceGetOK with default headers values
func NewServiceGetOK() *ServiceGetOK {

	return &ServiceGetOK{}
}

// WithPayload adds the payload to the service get o k response
func (o *ServiceGetOK) WithPayload(payload *models.Service) *ServiceGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service get o k response
func (o *ServiceGetOK) SetPayload(payload *models.Service) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceGetBadRequestCode is the HTTP code returned for type ServiceGetBadRequest
const ServiceGetBadRequestCode int = 400

/*ServiceGetBadRequest invalid input, object invalid

swagger:response serviceGetBadRequest
*/
type ServiceGetBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceGetBadRequest creates ServiceGetBadRequest with default headers values
func NewServiceGetBadRequest() *ServiceGetBadRequest {

	return &ServiceGetBadRequest{}
}

// WithPayload adds the payload to the service get bad request response
func (o *ServiceGetBadRequest) WithPayload(payload *models.Error) *ServiceGetBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service get bad request response
func (o *ServiceGetBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceGetBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceGetUnauthorizedCode is the HTTP code returned for type ServiceGetUnauthorized
const ServiceGetUnauthorizedCode int = 401

/*ServiceGetUnauthorized bad authentication

swagger:response serviceGetUnauthorized
*/
type ServiceGetUnauthorized struct {
}

// NewServiceGetUnauthorized creates ServiceGetUnauthorized with default headers values
func NewServiceGetUnauthorized() *ServiceGetUnauthorized {

	return &ServiceGetUnauthorized{}
}

// WriteResponse to the client
func (o *ServiceGetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ServiceGetForbiddenCode is the HTTP code returned for type ServiceGetForbidden
const ServiceGetForbiddenCode int = 403

/*ServiceGetForbidden bad permissions

swagger:response serviceGetForbidden
*/
type ServiceGetForbidden struct {
}

// NewServiceGetForbidden creates ServiceGetForbidden with default headers values
func NewServiceGetForbidden() *ServiceGetForbidden {

	return &ServiceGetForbidden{}
}

// WriteResponse to the client
func (o *ServiceGetForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// ServiceGetNotFoundCode is the HTTP code returned for type ServiceGetNotFound
const ServiceGetNotFoundCode int = 404

/*ServiceGetNotFound item not found

swagger:response serviceGetNotFound
*/
type ServiceGetNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceGetNotFound creates ServiceGetNotFound with default headers values
func NewServiceGetNotFound() *ServiceGetNotFound {

	return &ServiceGetNotFound{}
}

// WithPayload adds the payload to the service get not found response
func (o *ServiceGetNotFound) WithPayload(payload *models.Error) *ServiceGetNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service get not found response
func (o *ServiceGetNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceGetNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceGetUnprocessableEntityCode is the HTTP code returned for type ServiceGetUnprocessableEntity
const ServiceGetUnprocessableEntityCode int = 422

/*ServiceGetUnprocessableEntity bad validation

swagger:response serviceGetUnprocessableEntity
*/
type ServiceGetUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceGetUnprocessableEntity creates ServiceGetUnprocessableEntity with default headers values
func NewServiceGetUnprocessableEntity() *ServiceGetUnprocessableEntity {

	return &ServiceGetUnprocessableEntity{}
}

// WithPayload adds the payload to the service get unprocessable entity response
func (o *ServiceGetUnprocessableEntity) WithPayload(payload *models.Error) *ServiceGetUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service get unprocessable entity response
func (o *ServiceGetUnprocessableEntity) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceGetUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceGetServiceUnavailableCode is the HTTP code returned for type ServiceGetServiceUnavailable
const ServiceGetServiceUnavailableCode int = 503

/*ServiceGetServiceUnavailable internal server error

swagger:response serviceGetServiceUnavailable
*/
type ServiceGetServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceGetServiceUnavailable creates ServiceGetServiceUnavailable with default headers values
func NewServiceGetServiceUnavailable() *ServiceGetServiceUnavailable {

	return &ServiceGetServiceUnavailable{}
}

// WithPayload adds the payload to the service get service unavailable response
func (o *ServiceGetServiceUnavailable) WithPayload(payload *models.Error) *ServiceGetServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service get service unavailable response
func (o *ServiceGetServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceGetServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}