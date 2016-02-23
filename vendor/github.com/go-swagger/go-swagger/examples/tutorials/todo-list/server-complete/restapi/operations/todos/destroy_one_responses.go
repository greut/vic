package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit"

	"github.com/go-swagger/go-swagger/examples/tutorials/todo-list/server-complete/models"
)

/*DestroyOneNoContent Deleted

swagger:response destroyOneNoContent
*/
type DestroyOneNoContent struct {
}

// NewDestroyOneNoContent creates DestroyOneNoContent with default headers values
func NewDestroyOneNoContent() *DestroyOneNoContent {
	return &DestroyOneNoContent{}
}

// WriteResponse to the client
func (o *DestroyOneNoContent) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(204)
}

/*DestroyOneDefault error

swagger:response destroyOneDefault
*/
type DestroyOneDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewDestroyOneDefault creates DestroyOneDefault with default headers values
func NewDestroyOneDefault(code int) *DestroyOneDefault {
	if code <= 0 {
		code = 500
	}

	return &DestroyOneDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the destroy one default response
func (o *DestroyOneDefault) WithStatusCode(code int) *DestroyOneDefault {
	o._statusCode = code
	return o
}

// WithPayload adds the payload to the destroy one default response
func (o *DestroyOneDefault) WithPayload(payload *models.Error) *DestroyOneDefault {
	o.Payload = payload
	return o
}

// WriteResponse to the client
func (o *DestroyOneDefault) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
