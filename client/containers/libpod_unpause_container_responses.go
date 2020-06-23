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

// LibpodUnpauseContainerReader is a Reader for the LibpodUnpauseContainer structure.
type LibpodUnpauseContainerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LibpodUnpauseContainerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewLibpodUnpauseContainerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewLibpodUnpauseContainerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLibpodUnpauseContainerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLibpodUnpauseContainerNoContent creates a LibpodUnpauseContainerNoContent with default headers values
func NewLibpodUnpauseContainerNoContent() *LibpodUnpauseContainerNoContent {
	return &LibpodUnpauseContainerNoContent{}
}

/*LibpodUnpauseContainerNoContent handles this case with default header values.

no error
*/
type LibpodUnpauseContainerNoContent struct {
}

func (o *LibpodUnpauseContainerNoContent) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/unpause][%d] libpodUnpauseContainerNoContent ", 204)
}

func (o *LibpodUnpauseContainerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLibpodUnpauseContainerNotFound creates a LibpodUnpauseContainerNotFound with default headers values
func NewLibpodUnpauseContainerNotFound() *LibpodUnpauseContainerNotFound {
	return &LibpodUnpauseContainerNotFound{}
}

/*LibpodUnpauseContainerNotFound handles this case with default header values.

No such container
*/
type LibpodUnpauseContainerNotFound struct {
	Payload *LibpodUnpauseContainerNotFoundBody
}

func (o *LibpodUnpauseContainerNotFound) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/unpause][%d] libpodUnpauseContainerNotFound  %+v", 404, o.Payload)
}

func (o *LibpodUnpauseContainerNotFound) GetPayload() *LibpodUnpauseContainerNotFoundBody {
	return o.Payload
}

func (o *LibpodUnpauseContainerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodUnpauseContainerNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodUnpauseContainerInternalServerError creates a LibpodUnpauseContainerInternalServerError with default headers values
func NewLibpodUnpauseContainerInternalServerError() *LibpodUnpauseContainerInternalServerError {
	return &LibpodUnpauseContainerInternalServerError{}
}

/*LibpodUnpauseContainerInternalServerError handles this case with default header values.

Internal server error
*/
type LibpodUnpauseContainerInternalServerError struct {
	Payload *LibpodUnpauseContainerInternalServerErrorBody
}

func (o *LibpodUnpauseContainerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/unpause][%d] libpodUnpauseContainerInternalServerError  %+v", 500, o.Payload)
}

func (o *LibpodUnpauseContainerInternalServerError) GetPayload() *LibpodUnpauseContainerInternalServerErrorBody {
	return o.Payload
}

func (o *LibpodUnpauseContainerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodUnpauseContainerInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LibpodUnpauseContainerInternalServerErrorBody libpod unpause container internal server error body
swagger:model LibpodUnpauseContainerInternalServerErrorBody
*/
type LibpodUnpauseContainerInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod unpause container internal server error body
func (o *LibpodUnpauseContainerInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodUnpauseContainerInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodUnpauseContainerInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res LibpodUnpauseContainerInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodUnpauseContainerNotFoundBody libpod unpause container not found body
swagger:model LibpodUnpauseContainerNotFoundBody
*/
type LibpodUnpauseContainerNotFoundBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod unpause container not found body
func (o *LibpodUnpauseContainerNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodUnpauseContainerNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodUnpauseContainerNotFoundBody) UnmarshalBinary(b []byte) error {
	var res LibpodUnpauseContainerNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
