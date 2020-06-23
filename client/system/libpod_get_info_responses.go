// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/pascomnet/nomad-driver-podman/models"
)

// LibpodGetInfoReader is a Reader for the LibpodGetInfo structure.
type LibpodGetInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LibpodGetInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLibpodGetInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewLibpodGetInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLibpodGetInfoOK creates a LibpodGetInfoOK with default headers values
func NewLibpodGetInfoOK() *LibpodGetInfoOK {
	return &LibpodGetInfoOK{}
}

/*LibpodGetInfoOK handles this case with default header values.

Info
*/
type LibpodGetInfoOK struct {
	Payload *models.Info
}

func (o *LibpodGetInfoOK) Error() string {
	return fmt.Sprintf("[GET /libpod/info][%d] libpodGetInfoOK  %+v", 200, o.Payload)
}

func (o *LibpodGetInfoOK) GetPayload() *models.Info {
	return o.Payload
}

func (o *LibpodGetInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Info)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodGetInfoInternalServerError creates a LibpodGetInfoInternalServerError with default headers values
func NewLibpodGetInfoInternalServerError() *LibpodGetInfoInternalServerError {
	return &LibpodGetInfoInternalServerError{}
}

/*LibpodGetInfoInternalServerError handles this case with default header values.

Internal server error
*/
type LibpodGetInfoInternalServerError struct {
	Payload *LibpodGetInfoInternalServerErrorBody
}

func (o *LibpodGetInfoInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/info][%d] libpodGetInfoInternalServerError  %+v", 500, o.Payload)
}

func (o *LibpodGetInfoInternalServerError) GetPayload() *LibpodGetInfoInternalServerErrorBody {
	return o.Payload
}

func (o *LibpodGetInfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodGetInfoInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LibpodGetInfoInternalServerErrorBody libpod get info internal server error body
swagger:model LibpodGetInfoInternalServerErrorBody
*/
type LibpodGetInfoInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod get info internal server error body
func (o *LibpodGetInfoInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodGetInfoInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodGetInfoInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res LibpodGetInfoInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
