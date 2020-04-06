// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Console Server
// Copyright (c) 2020 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package admin_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/minio/mcs/models"
)

// ListPoliciesOKCode is the HTTP code returned for type ListPoliciesOK
const ListPoliciesOKCode int = 200

/*ListPoliciesOK A successful response.

swagger:response listPoliciesOK
*/
type ListPoliciesOK struct {

	/*
	  In: Body
	*/
	Payload *models.ListPoliciesResponse `json:"body,omitempty"`
}

// NewListPoliciesOK creates ListPoliciesOK with default headers values
func NewListPoliciesOK() *ListPoliciesOK {

	return &ListPoliciesOK{}
}

// WithPayload adds the payload to the list policies o k response
func (o *ListPoliciesOK) WithPayload(payload *models.ListPoliciesResponse) *ListPoliciesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list policies o k response
func (o *ListPoliciesOK) SetPayload(payload *models.ListPoliciesResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListPoliciesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ListPoliciesDefault Generic error response.

swagger:response listPoliciesDefault
*/
type ListPoliciesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListPoliciesDefault creates ListPoliciesDefault with default headers values
func NewListPoliciesDefault(code int) *ListPoliciesDefault {
	if code <= 0 {
		code = 500
	}

	return &ListPoliciesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list policies default response
func (o *ListPoliciesDefault) WithStatusCode(code int) *ListPoliciesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list policies default response
func (o *ListPoliciesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list policies default response
func (o *ListPoliciesDefault) WithPayload(payload *models.Error) *ListPoliciesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list policies default response
func (o *ListPoliciesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListPoliciesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
