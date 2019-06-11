// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetUniversitiesOKBody get universities o k body
// swagger:model getUniversitiesOKBody
type GetUniversitiesOKBody struct {

	// count university
	// Required: true
	CountUniversity int64 `json:"count_university"`

	// universities
	// Required: true
	Universities []*GetUniversitiesOKBodyUniversitiesItems `json:"universities"`
}

// Validate validates this get universities o k body
func (m *GetUniversitiesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCountUniversity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUniversities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetUniversitiesOKBody) validateCountUniversity(formats strfmt.Registry) error {

	if err := validate.Required("count_university", "body", int64(m.CountUniversity)); err != nil {
		return err
	}

	return nil
}

func (m *GetUniversitiesOKBody) validateUniversities(formats strfmt.Registry) error {

	if err := validate.Required("universities", "body", m.Universities); err != nil {
		return err
	}

	for i := 0; i < len(m.Universities); i++ {
		if swag.IsZero(m.Universities[i]) { // not required
			continue
		}

		if m.Universities[i] != nil {
			if err := m.Universities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("universities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetUniversitiesOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetUniversitiesOKBody) UnmarshalBinary(b []byte) error {
	var res GetUniversitiesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
