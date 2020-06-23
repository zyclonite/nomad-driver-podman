// Code generated by go-swagger; DO NOT EDIT.

package volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RemoveVolumeReader is a Reader for the RemoveVolume structure.
type RemoveVolumeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveVolumeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRemoveVolumeNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRemoveVolumeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewRemoveVolumeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRemoveVolumeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveVolumeNoContent creates a RemoveVolumeNoContent with default headers values
func NewRemoveVolumeNoContent() *RemoveVolumeNoContent {
	return &RemoveVolumeNoContent{}
}

/*RemoveVolumeNoContent handles this case with default header values.

no error
*/
type RemoveVolumeNoContent struct {
}

func (o *RemoveVolumeNoContent) Error() string {
	return fmt.Sprintf("[DELETE /libpod/volumes/{name}][%d] removeVolumeNoContent ", 204)
}

func (o *RemoveVolumeNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveVolumeNotFound creates a RemoveVolumeNotFound with default headers values
func NewRemoveVolumeNotFound() *RemoveVolumeNotFound {
	return &RemoveVolumeNotFound{}
}

/*RemoveVolumeNotFound handles this case with default header values.

No such volume
*/
type RemoveVolumeNotFound struct {
	Payload *RemoveVolumeNotFoundBody
}

func (o *RemoveVolumeNotFound) Error() string {
	return fmt.Sprintf("[DELETE /libpod/volumes/{name}][%d] removeVolumeNotFound  %+v", 404, o.Payload)
}

func (o *RemoveVolumeNotFound) GetPayload() *RemoveVolumeNotFoundBody {
	return o.Payload
}

func (o *RemoveVolumeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(RemoveVolumeNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveVolumeConflict creates a RemoveVolumeConflict with default headers values
func NewRemoveVolumeConflict() *RemoveVolumeConflict {
	return &RemoveVolumeConflict{}
}

/*RemoveVolumeConflict handles this case with default header values.

Volume is in use and cannot be removed
*/
type RemoveVolumeConflict struct {
}

func (o *RemoveVolumeConflict) Error() string {
	return fmt.Sprintf("[DELETE /libpod/volumes/{name}][%d] removeVolumeConflict ", 409)
}

func (o *RemoveVolumeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveVolumeInternalServerError creates a RemoveVolumeInternalServerError with default headers values
func NewRemoveVolumeInternalServerError() *RemoveVolumeInternalServerError {
	return &RemoveVolumeInternalServerError{}
}

/*RemoveVolumeInternalServerError handles this case with default header values.

Internal server error
*/
type RemoveVolumeInternalServerError struct {
	Payload *RemoveVolumeInternalServerErrorBody
}

func (o *RemoveVolumeInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /libpod/volumes/{name}][%d] removeVolumeInternalServerError  %+v", 500, o.Payload)
}

func (o *RemoveVolumeInternalServerError) GetPayload() *RemoveVolumeInternalServerErrorBody {
	return o.Payload
}

func (o *RemoveVolumeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(RemoveVolumeInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*RemoveVolumeInternalServerErrorBody remove volume internal server error body
swagger:model RemoveVolumeInternalServerErrorBody
*/
type RemoveVolumeInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this remove volume internal server error body
func (o *RemoveVolumeInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RemoveVolumeInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveVolumeInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res RemoveVolumeInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RemoveVolumeNotFoundBody remove volume not found body
swagger:model RemoveVolumeNotFoundBody
*/
type RemoveVolumeNotFoundBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this remove volume not found body
func (o *RemoveVolumeNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RemoveVolumeNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveVolumeNotFoundBody) UnmarshalBinary(b []byte) error {
	var res RemoveVolumeNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
