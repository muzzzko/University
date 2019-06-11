// Code generated by go-swagger; DO NOT EDIT.

package region

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "univers/pkg/models"
)

// GetRegionsOKCode is the HTTP code returned for type GetRegionsOK
const GetRegionsOKCode int = 200

/*GetRegionsOK успешное получение списка регионов

swagger:response getRegionsOK
*/
type GetRegionsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Region `json:"body,omitempty"`
}

// NewGetRegionsOK creates GetRegionsOK with default headers values
func NewGetRegionsOK() *GetRegionsOK {

	return &GetRegionsOK{}
}

// WithPayload adds the payload to the get regions o k response
func (o *GetRegionsOK) WithPayload(payload []*models.Region) *GetRegionsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get regions o k response
func (o *GetRegionsOK) SetPayload(payload []*models.Region) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRegionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Region, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetRegionsBadRequestCode is the HTTP code returned for type GetRegionsBadRequest
const GetRegionsBadRequestCode int = 400

/*GetRegionsBadRequest неверный запрос

swagger:response getRegionsBadRequest
*/
type GetRegionsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResult `json:"body,omitempty"`
}

// NewGetRegionsBadRequest creates GetRegionsBadRequest with default headers values
func NewGetRegionsBadRequest() *GetRegionsBadRequest {

	return &GetRegionsBadRequest{}
}

// WithPayload adds the payload to the get regions bad request response
func (o *GetRegionsBadRequest) WithPayload(payload *models.ErrorResult) *GetRegionsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get regions bad request response
func (o *GetRegionsBadRequest) SetPayload(payload *models.ErrorResult) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRegionsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}