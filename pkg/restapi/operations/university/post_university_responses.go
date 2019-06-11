// Code generated by go-swagger; DO NOT EDIT.

package university

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "univers/pkg/models"
)

// PostUniversityOKCode is the HTTP code returned for type PostUniversityOK
const PostUniversityOKCode int = 200

/*PostUniversityOK университет успешно добавлен

swagger:response postUniversityOK
*/
type PostUniversityOK struct {

	/*
	  In: Body
	*/
	Payload *models.PostUniversityOKBody `json:"body,omitempty"`
}

// NewPostUniversityOK creates PostUniversityOK with default headers values
func NewPostUniversityOK() *PostUniversityOK {

	return &PostUniversityOK{}
}

// WithPayload adds the payload to the post university o k response
func (o *PostUniversityOK) WithPayload(payload *models.PostUniversityOKBody) *PostUniversityOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post university o k response
func (o *PostUniversityOK) SetPayload(payload *models.PostUniversityOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUniversityOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostUniversityUnprocessableEntityCode is the HTTP code returned for type PostUniversityUnprocessableEntity
const PostUniversityUnprocessableEntityCode int = 422

/*PostUniversityUnprocessableEntity неверный запрос

swagger:response postUniversityUnprocessableEntity
*/
type PostUniversityUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResult `json:"body,omitempty"`
}

// NewPostUniversityUnprocessableEntity creates PostUniversityUnprocessableEntity with default headers values
func NewPostUniversityUnprocessableEntity() *PostUniversityUnprocessableEntity {

	return &PostUniversityUnprocessableEntity{}
}

// WithPayload adds the payload to the post university unprocessable entity response
func (o *PostUniversityUnprocessableEntity) WithPayload(payload *models.ErrorResult) *PostUniversityUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post university unprocessable entity response
func (o *PostUniversityUnprocessableEntity) SetPayload(payload *models.ErrorResult) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUniversityUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}