// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UniversityUpdateDataFacultiesItemsConditionsItems university update data faculties items conditions items
// swagger:model universityUpdateDataFacultiesItemsConditionsItems
type UniversityUpdateDataFacultiesItemsConditionsItems struct {

	// admission condition
	// Required: true
	AdmissionCondition string `json:"admission_condition"`

	// id
	ID int64 `json:"id,omitempty"`
}

// Validate validates this university update data faculties items conditions items
func (m *UniversityUpdateDataFacultiesItemsConditionsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdmissionCondition(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UniversityUpdateDataFacultiesItemsConditionsItems) validateAdmissionCondition(formats strfmt.Registry) error {

	if err := validate.RequiredString("admission_condition", "body", string(m.AdmissionCondition)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UniversityUpdateDataFacultiesItemsConditionsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UniversityUpdateDataFacultiesItemsConditionsItems) UnmarshalBinary(b []byte) error {
	var res UniversityUpdateDataFacultiesItemsConditionsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
