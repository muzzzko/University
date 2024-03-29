// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// University данные университета
// swagger:model university
type University struct {

	// адресс университета
	// Required: true
	Address string `json:"address"`

	// faculties
	// Required: true
	Faculties []*UniversityFacultiesItems `json:"faculties"`

	// идентификатор университета
	// Required: true
	ID int64 `json:"id"`

	// название университета
	// Required: true
	Name string `json:"name"`

	// rector
	// Required: true
	Rector *UniversityRector `json:"rector"`

	// region
	// Required: true
	Region *UniversityRegion `json:"region"`

	// профиль вуза
	// Required: true
	Shape string `json:"shape"`

	// статус вуза
	// Required: true
	// Enum: [state commercial]
	Status string `json:"status"`

	// кличество студентов
	// Required: true
	StudentNumber int64 `json:"student_number"`
}

// Validate validates this university
func (m *University) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFaculties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShape(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStudentNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *University) validateAddress(formats strfmt.Registry) error {

	if err := validate.RequiredString("address", "body", string(m.Address)); err != nil {
		return err
	}

	return nil
}

func (m *University) validateFaculties(formats strfmt.Registry) error {

	if err := validate.Required("faculties", "body", m.Faculties); err != nil {
		return err
	}

	for i := 0; i < len(m.Faculties); i++ {
		if swag.IsZero(m.Faculties[i]) { // not required
			continue
		}

		if m.Faculties[i] != nil {
			if err := m.Faculties[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("faculties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *University) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *University) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *University) validateRector(formats strfmt.Registry) error {

	if err := validate.Required("rector", "body", m.Rector); err != nil {
		return err
	}

	if m.Rector != nil {
		if err := m.Rector.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rector")
			}
			return err
		}
	}

	return nil
}

func (m *University) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	if m.Region != nil {
		if err := m.Region.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("region")
			}
			return err
		}
	}

	return nil
}

func (m *University) validateShape(formats strfmt.Registry) error {

	if err := validate.RequiredString("shape", "body", string(m.Shape)); err != nil {
		return err
	}

	return nil
}

var universityTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["state","commercial"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		universityTypeStatusPropEnum = append(universityTypeStatusPropEnum, v)
	}
}

const (

	// UniversityStatusState captures enum value "state"
	UniversityStatusState string = "state"

	// UniversityStatusCommercial captures enum value "commercial"
	UniversityStatusCommercial string = "commercial"
)

// prop value enum
func (m *University) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, universityTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *University) validateStatus(formats strfmt.Registry) error {

	if err := validate.RequiredString("status", "body", string(m.Status)); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *University) validateStudentNumber(formats strfmt.Registry) error {

	if err := validate.Required("student_number", "body", int64(m.StudentNumber)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *University) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *University) UnmarshalBinary(b []byte) error {
	var res University
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
