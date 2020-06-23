// Code generated by go-swagger; DO NOT EDIT.

package containers_compat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RemoveContainerReader is a Reader for the RemoveContainer structure.
type RemoveContainerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveContainerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRemoveContainerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveContainerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRemoveContainerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewRemoveContainerConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRemoveContainerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveContainerNoContent creates a RemoveContainerNoContent with default headers values
func NewRemoveContainerNoContent() *RemoveContainerNoContent {
	return &RemoveContainerNoContent{}
}

/*RemoveContainerNoContent handles this case with default header values.

no error
*/
type RemoveContainerNoContent struct {
}

func (o *RemoveContainerNoContent) Error() string {
	return fmt.Sprintf("[DELETE /containers/{name}][%d] removeContainerNoContent ", 204)
}

func (o *RemoveContainerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveContainerBadRequest creates a RemoveContainerBadRequest with default headers values
func NewRemoveContainerBadRequest() *RemoveContainerBadRequest {
	return &RemoveContainerBadRequest{}
}

/*RemoveContainerBadRequest handles this case with default header values.

Bad parameter in request
*/
type RemoveContainerBadRequest struct {
	Payload *RemoveContainerBadRequestBody
}

func (o *RemoveContainerBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /containers/{name}][%d] removeContainerBadRequest  %+v", 400, o.Payload)
}

func (o *RemoveContainerBadRequest) GetPayload() *RemoveContainerBadRequestBody {
	return o.Payload
}

func (o *RemoveContainerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(RemoveContainerBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveContainerNotFound creates a RemoveContainerNotFound with default headers values
func NewRemoveContainerNotFound() *RemoveContainerNotFound {
	return &RemoveContainerNotFound{}
}

/*RemoveContainerNotFound handles this case with default header values.

No such container
*/
type RemoveContainerNotFound struct {
	Payload *RemoveContainerNotFoundBody
}

func (o *RemoveContainerNotFound) Error() string {
	return fmt.Sprintf("[DELETE /containers/{name}][%d] removeContainerNotFound  %+v", 404, o.Payload)
}

func (o *RemoveContainerNotFound) GetPayload() *RemoveContainerNotFoundBody {
	return o.Payload
}

func (o *RemoveContainerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(RemoveContainerNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveContainerConflict creates a RemoveContainerConflict with default headers values
func NewRemoveContainerConflict() *RemoveContainerConflict {
	return &RemoveContainerConflict{}
}

/*RemoveContainerConflict handles this case with default header values.

Conflict error in operation
*/
type RemoveContainerConflict struct {
	Payload *RemoveContainerConflictBody
}

func (o *RemoveContainerConflict) Error() string {
	return fmt.Sprintf("[DELETE /containers/{name}][%d] removeContainerConflict  %+v", 409, o.Payload)
}

func (o *RemoveContainerConflict) GetPayload() *RemoveContainerConflictBody {
	return o.Payload
}

func (o *RemoveContainerConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(RemoveContainerConflictBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveContainerInternalServerError creates a RemoveContainerInternalServerError with default headers values
func NewRemoveContainerInternalServerError() *RemoveContainerInternalServerError {
	return &RemoveContainerInternalServerError{}
}

/*RemoveContainerInternalServerError handles this case with default header values.

Internal server error
*/
type RemoveContainerInternalServerError struct {
	Payload *RemoveContainerInternalServerErrorBody
}

func (o *RemoveContainerInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /containers/{name}][%d] removeContainerInternalServerError  %+v", 500, o.Payload)
}

func (o *RemoveContainerInternalServerError) GetPayload() *RemoveContainerInternalServerErrorBody {
	return o.Payload
}

func (o *RemoveContainerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(RemoveContainerInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*RemoveContainerBadRequestBody remove container bad request body
swagger:model RemoveContainerBadRequestBody
*/
type RemoveContainerBadRequestBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this remove container bad request body
func (o *RemoveContainerBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RemoveContainerBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveContainerBadRequestBody) UnmarshalBinary(b []byte) error {
	var res RemoveContainerBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RemoveContainerConflictBody remove container conflict body
swagger:model RemoveContainerConflictBody
*/
type RemoveContainerConflictBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this remove container conflict body
func (o *RemoveContainerConflictBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RemoveContainerConflictBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveContainerConflictBody) UnmarshalBinary(b []byte) error {
	var res RemoveContainerConflictBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RemoveContainerInternalServerErrorBody remove container internal server error body
swagger:model RemoveContainerInternalServerErrorBody
*/
type RemoveContainerInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this remove container internal server error body
func (o *RemoveContainerInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RemoveContainerInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveContainerInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res RemoveContainerInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RemoveContainerNotFoundBody remove container not found body
swagger:model RemoveContainerNotFoundBody
*/
type RemoveContainerNotFoundBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this remove container not found body
func (o *RemoveContainerNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RemoveContainerNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveContainerNotFoundBody) UnmarshalBinary(b []byte) error {
	var res RemoveContainerNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
