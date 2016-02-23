package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

// NewAddCommentToTaskParams creates a new AddCommentToTaskParams object
// with the default values initialized.
func NewAddCommentToTaskParams() *AddCommentToTaskParams {
	var ()
	return &AddCommentToTaskParams{}
}

/*AddCommentToTaskParams contains all the parameters to send to the API endpoint
for the add comment to task operation typically these are written to a http.Request
*/
type AddCommentToTaskParams struct {

	/*Body
	  The comment to add

	*/
	Body AddCommentToTaskBody
	/*ID
	  The id of the item

	*/
	ID int64
}

// WithBody adds the body to the add comment to task params
func (o *AddCommentToTaskParams) WithBody(body AddCommentToTaskBody) *AddCommentToTaskParams {
	o.Body = body
	return o
}

// WithID adds the id to the add comment to task params
func (o *AddCommentToTaskParams) WithID(id int64) *AddCommentToTaskParams {
	o.ID = id
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *AddCommentToTaskParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
