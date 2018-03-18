// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

/*PetGetParams contains all the parameters to send to the API endpoint
for the pet get operation typically these are written to a http.Request
*/
type PetGetParams struct {

	/*PetID
	  ID of pet to return

	*/
	PetID int64

	Timeout time.Duration
}

// WriteToRequest writes these params to a swagger request
func (o *PetGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.Timeout); err != nil {
		return err
	}
	var res []error

	// path param petId
	if err := r.SetPathParam("petId", swag.FormatInt64(o.PetID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}