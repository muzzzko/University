// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetPingHandlerFunc turns a function with the right signature into a get ping handler
type GetPingHandlerFunc func(GetPingParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPingHandlerFunc) Handle(params GetPingParams) middleware.Responder {
	return fn(params)
}

// GetPingHandler interface for that can handle valid get ping params
type GetPingHandler interface {
	Handle(GetPingParams) middleware.Responder
}

// NewGetPing creates a new http.Handler for the get ping operation
func NewGetPing(ctx *middleware.Context, handler GetPingHandler) *GetPing {
	return &GetPing{Context: ctx, Handler: handler}
}

/*GetPing swagger:route GET /ping getPing

healthcheck

*/
type GetPing struct {
	Context *middleware.Context
	Handler GetPingHandler
}

func (o *GetPing) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetPingParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
