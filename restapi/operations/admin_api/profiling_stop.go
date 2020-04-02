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
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ProfilingStopHandlerFunc turns a function with the right signature into a profiling stop handler
type ProfilingStopHandlerFunc func(ProfilingStopParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ProfilingStopHandlerFunc) Handle(params ProfilingStopParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ProfilingStopHandler interface for that can handle valid profiling stop params
type ProfilingStopHandler interface {
	Handle(ProfilingStopParams, interface{}) middleware.Responder
}

// NewProfilingStop creates a new http.Handler for the profiling stop operation
func NewProfilingStop(ctx *middleware.Context, handler ProfilingStopHandler) *ProfilingStop {
	return &ProfilingStop{Context: ctx, Handler: handler}
}

/*ProfilingStop swagger:route POST /api/v1/profiling/stop AdminAPI profilingStop

Stop and download profile data

*/
type ProfilingStop struct {
	Context *middleware.Context
	Handler ProfilingStopHandler
}

func (o *ProfilingStop) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewProfilingStopParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
