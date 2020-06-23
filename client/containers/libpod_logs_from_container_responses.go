// Code generated by go-swagger; DO NOT EDIT.

package containers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LibpodLogsFromContainerReader is a Reader for the LibpodLogsFromContainer structure.
type LibpodLogsFromContainerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LibpodLogsFromContainerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLibpodLogsFromContainerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewLibpodLogsFromContainerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLibpodLogsFromContainerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLibpodLogsFromContainerOK creates a LibpodLogsFromContainerOK with default headers values
func NewLibpodLogsFromContainerOK() *LibpodLogsFromContainerOK {
	return &LibpodLogsFromContainerOK{}
}

/*LibpodLogsFromContainerOK handles this case with default header values.

logs returned as a stream in response body.
*/
type LibpodLogsFromContainerOK struct {
}

func (o *LibpodLogsFromContainerOK) Error() string {
	return fmt.Sprintf("[GET /libpod/containers/{name}/logs][%d] libpodLogsFromContainerOK ", 200)
}

func (o *LibpodLogsFromContainerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLibpodLogsFromContainerNotFound creates a LibpodLogsFromContainerNotFound with default headers values
func NewLibpodLogsFromContainerNotFound() *LibpodLogsFromContainerNotFound {
	return &LibpodLogsFromContainerNotFound{}
}

/*LibpodLogsFromContainerNotFound handles this case with default header values.

No such container
*/
type LibpodLogsFromContainerNotFound struct {
	Payload *LibpodLogsFromContainerNotFoundBody
}

func (o *LibpodLogsFromContainerNotFound) Error() string {
	return fmt.Sprintf("[GET /libpod/containers/{name}/logs][%d] libpodLogsFromContainerNotFound  %+v", 404, o.Payload)
}

func (o *LibpodLogsFromContainerNotFound) GetPayload() *LibpodLogsFromContainerNotFoundBody {
	return o.Payload
}

func (o *LibpodLogsFromContainerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodLogsFromContainerNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodLogsFromContainerInternalServerError creates a LibpodLogsFromContainerInternalServerError with default headers values
func NewLibpodLogsFromContainerInternalServerError() *LibpodLogsFromContainerInternalServerError {
	return &LibpodLogsFromContainerInternalServerError{}
}

/*LibpodLogsFromContainerInternalServerError handles this case with default header values.

Internal server error
*/
type LibpodLogsFromContainerInternalServerError struct {
	Payload *LibpodLogsFromContainerInternalServerErrorBody
}

func (o *LibpodLogsFromContainerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/containers/{name}/logs][%d] libpodLogsFromContainerInternalServerError  %+v", 500, o.Payload)
}

func (o *LibpodLogsFromContainerInternalServerError) GetPayload() *LibpodLogsFromContainerInternalServerErrorBody {
	return o.Payload
}

func (o *LibpodLogsFromContainerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodLogsFromContainerInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LibpodLogsFromContainerInternalServerErrorBody libpod logs from container internal server error body
swagger:model LibpodLogsFromContainerInternalServerErrorBody
*/
type LibpodLogsFromContainerInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod logs from container internal server error body
func (o *LibpodLogsFromContainerInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodLogsFromContainerInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodLogsFromContainerInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res LibpodLogsFromContainerInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodLogsFromContainerNotFoundBody libpod logs from container not found body
swagger:model LibpodLogsFromContainerNotFoundBody
*/
type LibpodLogsFromContainerNotFoundBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod logs from container not found body
func (o *LibpodLogsFromContainerNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodLogsFromContainerNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodLogsFromContainerNotFoundBody) UnmarshalBinary(b []byte) error {
	var res LibpodLogsFromContainerNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
