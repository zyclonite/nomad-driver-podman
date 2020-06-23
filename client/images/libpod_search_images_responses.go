// Code generated by go-swagger; DO NOT EDIT.

package images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LibpodSearchImagesReader is a Reader for the LibpodSearchImages structure.
type LibpodSearchImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LibpodSearchImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLibpodSearchImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewLibpodSearchImagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLibpodSearchImagesOK creates a LibpodSearchImagesOK with default headers values
func NewLibpodSearchImagesOK() *LibpodSearchImagesOK {
	return &LibpodSearchImagesOK{}
}

/*LibpodSearchImagesOK handles this case with default header values.

Search results
*/
type LibpodSearchImagesOK struct {
	Payload *LibpodSearchImagesOKBody
}

func (o *LibpodSearchImagesOK) Error() string {
	return fmt.Sprintf("[GET /libpod/images/search][%d] libpodSearchImagesOK  %+v", 200, o.Payload)
}

func (o *LibpodSearchImagesOK) GetPayload() *LibpodSearchImagesOKBody {
	return o.Payload
}

func (o *LibpodSearchImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodSearchImagesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodSearchImagesInternalServerError creates a LibpodSearchImagesInternalServerError with default headers values
func NewLibpodSearchImagesInternalServerError() *LibpodSearchImagesInternalServerError {
	return &LibpodSearchImagesInternalServerError{}
}

/*LibpodSearchImagesInternalServerError handles this case with default header values.

Internal server error
*/
type LibpodSearchImagesInternalServerError struct {
	Payload *LibpodSearchImagesInternalServerErrorBody
}

func (o *LibpodSearchImagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/images/search][%d] libpodSearchImagesInternalServerError  %+v", 500, o.Payload)
}

func (o *LibpodSearchImagesInternalServerError) GetPayload() *LibpodSearchImagesInternalServerErrorBody {
	return o.Payload
}

func (o *LibpodSearchImagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodSearchImagesInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LibpodSearchImagesInternalServerErrorBody libpod search images internal server error body
swagger:model LibpodSearchImagesInternalServerErrorBody
*/
type LibpodSearchImagesInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod search images internal server error body
func (o *LibpodSearchImagesInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodSearchImagesInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodSearchImagesInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res LibpodSearchImagesInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodSearchImagesOKBody libpod search images o k body
swagger:model LibpodSearchImagesOKBody
*/
type LibpodSearchImagesOKBody struct {

	// Automated indicates if the image was created by an automated build.
	Automated string `json:"Automated,omitempty"`

	// Description of the image.
	Description string `json:"Description,omitempty"`

	// Index is the image index (e.g., "docker.io" or "quay.io")
	Index string `json:"Index,omitempty"`

	// Name is the canoncical name of the image (e.g., "docker.io/library/alpine").
	Name string `json:"Name,omitempty"`

	// Official indicates if it's an official image.
	Official string `json:"Official,omitempty"`

	// Stars is the number of stars of the image.
	Stars int64 `json:"Stars,omitempty"`
}

// Validate validates this libpod search images o k body
func (o *LibpodSearchImagesOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodSearchImagesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodSearchImagesOKBody) UnmarshalBinary(b []byte) error {
	var res LibpodSearchImagesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
