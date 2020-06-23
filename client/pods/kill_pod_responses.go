// Code generated by go-swagger; DO NOT EDIT.

package pods

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

// KillPodReader is a Reader for the KillPod structure.
type KillPodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KillPodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKillPodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKillPodBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKillPodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewKillPodConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKillPodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewKillPodOK creates a KillPodOK with default headers values
func NewKillPodOK() *KillPodOK {
	return &KillPodOK{}
}

/*KillPodOK handles this case with default header values.

Kill Pod
*/
type KillPodOK struct {
	Payload *models.PodKillReport
}

func (o *KillPodOK) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/{name}/kill][%d] killPodOK  %+v", 200, o.Payload)
}

func (o *KillPodOK) GetPayload() *models.PodKillReport {
	return o.Payload
}

func (o *KillPodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PodKillReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKillPodBadRequest creates a KillPodBadRequest with default headers values
func NewKillPodBadRequest() *KillPodBadRequest {
	return &KillPodBadRequest{}
}

/*KillPodBadRequest handles this case with default header values.

Bad parameter in request
*/
type KillPodBadRequest struct {
	Payload *KillPodBadRequestBody
}

func (o *KillPodBadRequest) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/{name}/kill][%d] killPodBadRequest  %+v", 400, o.Payload)
}

func (o *KillPodBadRequest) GetPayload() *KillPodBadRequestBody {
	return o.Payload
}

func (o *KillPodBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(KillPodBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKillPodNotFound creates a KillPodNotFound with default headers values
func NewKillPodNotFound() *KillPodNotFound {
	return &KillPodNotFound{}
}

/*KillPodNotFound handles this case with default header values.

No such pod
*/
type KillPodNotFound struct {
	Payload *KillPodNotFoundBody
}

func (o *KillPodNotFound) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/{name}/kill][%d] killPodNotFound  %+v", 404, o.Payload)
}

func (o *KillPodNotFound) GetPayload() *KillPodNotFoundBody {
	return o.Payload
}

func (o *KillPodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(KillPodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKillPodConflict creates a KillPodConflict with default headers values
func NewKillPodConflict() *KillPodConflict {
	return &KillPodConflict{}
}

/*KillPodConflict handles this case with default header values.

Conflict error in operation
*/
type KillPodConflict struct {
	Payload *KillPodConflictBody
}

func (o *KillPodConflict) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/{name}/kill][%d] killPodConflict  %+v", 409, o.Payload)
}

func (o *KillPodConflict) GetPayload() *KillPodConflictBody {
	return o.Payload
}

func (o *KillPodConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(KillPodConflictBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKillPodInternalServerError creates a KillPodInternalServerError with default headers values
func NewKillPodInternalServerError() *KillPodInternalServerError {
	return &KillPodInternalServerError{}
}

/*KillPodInternalServerError handles this case with default header values.

Internal server error
*/
type KillPodInternalServerError struct {
	Payload *KillPodInternalServerErrorBody
}

func (o *KillPodInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/{name}/kill][%d] killPodInternalServerError  %+v", 500, o.Payload)
}

func (o *KillPodInternalServerError) GetPayload() *KillPodInternalServerErrorBody {
	return o.Payload
}

func (o *KillPodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(KillPodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*KillPodBadRequestBody kill pod bad request body
swagger:model KillPodBadRequestBody
*/
type KillPodBadRequestBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this kill pod bad request body
func (o *KillPodBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *KillPodBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KillPodBadRequestBody) UnmarshalBinary(b []byte) error {
	var res KillPodBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*KillPodConflictBody kill pod conflict body
swagger:model KillPodConflictBody
*/
type KillPodConflictBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this kill pod conflict body
func (o *KillPodConflictBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *KillPodConflictBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KillPodConflictBody) UnmarshalBinary(b []byte) error {
	var res KillPodConflictBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*KillPodInternalServerErrorBody kill pod internal server error body
swagger:model KillPodInternalServerErrorBody
*/
type KillPodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this kill pod internal server error body
func (o *KillPodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *KillPodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KillPodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res KillPodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*KillPodNotFoundBody kill pod not found body
swagger:model KillPodNotFoundBody
*/
type KillPodNotFoundBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this kill pod not found body
func (o *KillPodNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *KillPodNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KillPodNotFoundBody) UnmarshalBinary(b []byte) error {
	var res KillPodNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
