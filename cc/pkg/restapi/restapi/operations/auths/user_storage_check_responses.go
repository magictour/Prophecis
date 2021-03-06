// Code generated by go-swagger; DO NOT EDIT.

package auths

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "mlss-controlcenter-go/pkg/models"
)

// UserStorageCheckOKCode is the HTTP code returned for type UserStorageCheckOK
const UserStorageCheckOKCode int = 200

/*UserStorageCheckOK auth by username and path.

swagger:response userStorageCheckOK
*/
type UserStorageCheckOK struct {

	/*
	  In: Body
	*/
	Payload *models.Result `json:"body,omitempty"`
}

// NewUserStorageCheckOK creates UserStorageCheckOK with default headers values
func NewUserStorageCheckOK() *UserStorageCheckOK {

	return &UserStorageCheckOK{}
}

// WithPayload adds the payload to the user storage check o k response
func (o *UserStorageCheckOK) WithPayload(payload *models.Result) *UserStorageCheckOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user storage check o k response
func (o *UserStorageCheckOK) SetPayload(payload *models.Result) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserStorageCheckOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UserStorageCheckUnauthorizedCode is the HTTP code returned for type UserStorageCheckUnauthorized
const UserStorageCheckUnauthorizedCode int = 401

/*UserStorageCheckUnauthorized Unauthorized

swagger:response userStorageCheckUnauthorized
*/
type UserStorageCheckUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUserStorageCheckUnauthorized creates UserStorageCheckUnauthorized with default headers values
func NewUserStorageCheckUnauthorized() *UserStorageCheckUnauthorized {

	return &UserStorageCheckUnauthorized{}
}

// WithPayload adds the payload to the user storage check unauthorized response
func (o *UserStorageCheckUnauthorized) WithPayload(payload *models.Error) *UserStorageCheckUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user storage check unauthorized response
func (o *UserStorageCheckUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserStorageCheckUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UserStorageCheckNotFoundCode is the HTTP code returned for type UserStorageCheckNotFound
const UserStorageCheckNotFoundCode int = 404

/*UserStorageCheckNotFound url to add namespace not found.

swagger:response userStorageCheckNotFound
*/
type UserStorageCheckNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUserStorageCheckNotFound creates UserStorageCheckNotFound with default headers values
func NewUserStorageCheckNotFound() *UserStorageCheckNotFound {

	return &UserStorageCheckNotFound{}
}

// WithPayload adds the payload to the user storage check not found response
func (o *UserStorageCheckNotFound) WithPayload(payload *models.Error) *UserStorageCheckNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user storage check not found response
func (o *UserStorageCheckNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserStorageCheckNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
