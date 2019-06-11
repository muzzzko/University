// Code generated by go-swagger; DO NOT EDIT.

package university

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetUniversityParams creates a new GetUniversityParams object
// no default values defined in spec.
func NewGetUniversityParams() GetUniversityParams {

	return GetUniversityParams{}
}

// GetUniversityParams contains all the bound params for the get university operation
// typically these are obtained from a http.Request
//
// swagger:parameters getUniversity
type GetUniversityParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*идентификатор университета
	  Required: true
	  In: path
	*/
	UniversityID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetUniversityParams() beforehand.
func (o *GetUniversityParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rUniversityID, rhkUniversityID, _ := route.Params.GetOK("university_id")
	if err := o.bindUniversityID(rUniversityID, rhkUniversityID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindUniversityID binds and validates parameter UniversityID from path.
func (o *GetUniversityParams) bindUniversityID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("university_id", "path", "int64", raw)
	}
	o.UniversityID = value

	return nil
}