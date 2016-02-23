package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit/middleware"
)

// GetPetByIDHandlerFunc turns a function with the right signature into a get pet by id handler
type GetPetByIDHandlerFunc func(GetPetByIDParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPetByIDHandlerFunc) Handle(params GetPetByIDParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetPetByIDHandler interface for that can handle valid get pet by id params
type GetPetByIDHandler interface {
	Handle(GetPetByIDParams, interface{}) middleware.Responder
}

// NewGetPetByID creates a new http.Handler for the get pet by id operation
func NewGetPetByID(ctx *middleware.Context, handler GetPetByIDHandler) *GetPetByID {
	return &GetPetByID{Context: ctx, Handler: handler}
}

/*GetPetByID swagger:route GET /pet/{petId} pet getPetById

Find pet by ID

Returns a single pet

*/
type GetPetByID struct {
	Context *middleware.Context
	Handler GetPetByIDHandler
}

func (o *GetPetByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetPetByIDParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
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
