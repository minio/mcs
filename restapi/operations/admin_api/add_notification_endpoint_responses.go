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

	"github.com/minio/console/models"
)

// AddNotificationEndpointCreatedCode is the HTTP code returned for type AddNotificationEndpointCreated
const AddNotificationEndpointCreatedCode int = 201

/*AddNotificationEndpointCreated A successful response.

swagger:response addNotificationEndpointCreated
*/
type AddNotificationEndpointCreated struct {

	/*
	  In: Body
	*/
	Payload *models.SetNotificationEndpointResponse `json:"body,omitempty"`
}

// NewAddNotificationEndpointCreated creates AddNotificationEndpointCreated with default headers values
func NewAddNotificationEndpointCreated() *AddNotificationEndpointCreated {

	return &AddNotificationEndpointCreated{}
}

// WithPayload adds the payload to the add notification endpoint created response
func (o *AddNotificationEndpointCreated) WithPayload(payload *models.SetNotificationEndpointResponse) *AddNotificationEndpointCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add notification endpoint created response
func (o *AddNotificationEndpointCreated) SetPayload(payload *models.SetNotificationEndpointResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddNotificationEndpointCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*AddNotificationEndpointDefault Generic error response.

swagger:response addNotificationEndpointDefault
*/
type AddNotificationEndpointDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddNotificationEndpointDefault creates AddNotificationEndpointDefault with default headers values
func NewAddNotificationEndpointDefault(code int) *AddNotificationEndpointDefault {
	if code <= 0 {
		code = 500
	}

	return &AddNotificationEndpointDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the add notification endpoint default response
func (o *AddNotificationEndpointDefault) WithStatusCode(code int) *AddNotificationEndpointDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the add notification endpoint default response
func (o *AddNotificationEndpointDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the add notification endpoint default response
func (o *AddNotificationEndpointDefault) WithPayload(payload *models.Error) *AddNotificationEndpointDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add notification endpoint default response
func (o *AddNotificationEndpointDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddNotificationEndpointDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
