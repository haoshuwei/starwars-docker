package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PutExhaustportHandlerFunc turns a function with the right signature into a put exhaustport handler
type PutExhaustportHandlerFunc func(PutExhaustportParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutExhaustportHandlerFunc) Handle(params PutExhaustportParams) middleware.Responder {
	return fn(params)
}

// PutExhaustportHandler interface for that can handle valid put exhaustport params
type PutExhaustportHandler interface {
	Handle(PutExhaustportParams) middleware.Responder
}

// NewPutExhaustport creates a new http.Handler for the put exhaustport operation
func NewPutExhaustport(ctx *middleware.Context, handler PutExhaustportHandler) *PutExhaustport {
	return &PutExhaustport{Context: ctx, Handler: handler}
}

/*PutExhaustport swagger:route PUT /exhaustport putExhaustport

Put something into the thermal exhaust port of the deathstar

*/
type PutExhaustport struct {
	Context *middleware.Context
	Handler PutExhaustportHandler
}

func (o *PutExhaustport) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPutExhaustportParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
