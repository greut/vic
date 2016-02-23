package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit/middleware"
)

// UpdateOneHandlerFunc turns a function with the right signature into a update one handler
type UpdateOneHandlerFunc func(UpdateOneParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateOneHandlerFunc) Handle(params UpdateOneParams) middleware.Responder {
	return fn(params)
}

// UpdateOneHandler interface for that can handle valid update one params
type UpdateOneHandler interface {
	Handle(UpdateOneParams) middleware.Responder
}

// NewUpdateOne creates a new http.Handler for the update one operation
func NewUpdateOne(ctx *middleware.Context, handler UpdateOneHandler) *UpdateOne {
	return &UpdateOne{Context: ctx, Handler: handler}
}

/*UpdateOne swagger:route PUT /{id} todos updateOne

UpdateOne update one API

*/
type UpdateOne struct {
	Context *middleware.Context
	Handler UpdateOneHandler
}

func (o *UpdateOne) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewUpdateOneParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
