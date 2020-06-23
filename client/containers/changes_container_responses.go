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

// ChangesContainerReader is a Reader for the ChangesContainer structure.
type ChangesContainerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangesContainerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangesContainerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewChangesContainerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewChangesContainerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangesContainerOK creates a ChangesContainerOK with default headers values
func NewChangesContainerOK() *ChangesContainerOK {
	return &ChangesContainerOK{}
}

/*ChangesContainerOK handles this case with default header values.

Array of Changes
*/
type ChangesContainerOK struct {
}

func (o *ChangesContainerOK) Error() string {
	return fmt.Sprintf("[GET /libpod/containers/{name}/changes][%d] changesContainerOK ", 200)
}

func (o *ChangesContainerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangesContainerNotFound creates a ChangesContainerNotFound with default headers values
func NewChangesContainerNotFound() *ChangesContainerNotFound {
	return &ChangesContainerNotFound{}
}

/*ChangesContainerNotFound handles this case with default header values.

No such container
*/
type ChangesContainerNotFound struct {
	Payload *ChangesContainerNotFoundBody
}

func (o *ChangesContainerNotFound) Error() string {
	return fmt.Sprintf("[GET /libpod/containers/{name}/changes][%d] changesContainerNotFound  %+v", 404, o.Payload)
}

func (o *ChangesContainerNotFound) GetPayload() *ChangesContainerNotFoundBody {
	return o.Payload
}

func (o *ChangesContainerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangesContainerNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangesContainerInternalServerError creates a ChangesContainerInternalServerError with default headers values
func NewChangesContainerInternalServerError() *ChangesContainerInternalServerError {
	return &ChangesContainerInternalServerError{}
}

/*ChangesContainerInternalServerError handles this case with default header values.

Internal server error
*/
type ChangesContainerInternalServerError struct {
	Payload *ChangesContainerInternalServerErrorBody
}

func (o *ChangesContainerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/containers/{name}/changes][%d] changesContainerInternalServerError  %+v", 500, o.Payload)
}

func (o *ChangesContainerInternalServerError) GetPayload() *ChangesContainerInternalServerErrorBody {
	return o.Payload
}

func (o *ChangesContainerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangesContainerInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ChangesContainerInternalServerErrorBody changes container internal server error body
swagger:model ChangesContainerInternalServerErrorBody
*/
type ChangesContainerInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this changes container internal server error body
func (o *ChangesContainerInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangesContainerInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangesContainerInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ChangesContainerInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangesContainerNotFoundBody changes container not found body
swagger:model ChangesContainerNotFoundBody
*/
type ChangesContainerNotFoundBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this changes container not found body
func (o *ChangesContainerNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangesContainerNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangesContainerNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ChangesContainerNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
