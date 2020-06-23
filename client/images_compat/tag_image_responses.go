// Code generated by go-swagger; DO NOT EDIT.

package images_compat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TagImageReader is a Reader for the TagImage structure.
type TagImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TagImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewTagImageCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTagImageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTagImageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewTagImageConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTagImageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTagImageCreated creates a TagImageCreated with default headers values
func NewTagImageCreated() *TagImageCreated {
	return &TagImageCreated{}
}

/*TagImageCreated handles this case with default header values.

no error
*/
type TagImageCreated struct {
}

func (o *TagImageCreated) Error() string {
	return fmt.Sprintf("[POST /images/{name:.*}/tag][%d] tagImageCreated ", 201)
}

func (o *TagImageCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTagImageBadRequest creates a TagImageBadRequest with default headers values
func NewTagImageBadRequest() *TagImageBadRequest {
	return &TagImageBadRequest{}
}

/*TagImageBadRequest handles this case with default header values.

Bad parameter in request
*/
type TagImageBadRequest struct {
	Payload *TagImageBadRequestBody
}

func (o *TagImageBadRequest) Error() string {
	return fmt.Sprintf("[POST /images/{name:.*}/tag][%d] tagImageBadRequest  %+v", 400, o.Payload)
}

func (o *TagImageBadRequest) GetPayload() *TagImageBadRequestBody {
	return o.Payload
}

func (o *TagImageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(TagImageBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTagImageNotFound creates a TagImageNotFound with default headers values
func NewTagImageNotFound() *TagImageNotFound {
	return &TagImageNotFound{}
}

/*TagImageNotFound handles this case with default header values.

No such image
*/
type TagImageNotFound struct {
	Payload *TagImageNotFoundBody
}

func (o *TagImageNotFound) Error() string {
	return fmt.Sprintf("[POST /images/{name:.*}/tag][%d] tagImageNotFound  %+v", 404, o.Payload)
}

func (o *TagImageNotFound) GetPayload() *TagImageNotFoundBody {
	return o.Payload
}

func (o *TagImageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(TagImageNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTagImageConflict creates a TagImageConflict with default headers values
func NewTagImageConflict() *TagImageConflict {
	return &TagImageConflict{}
}

/*TagImageConflict handles this case with default header values.

Conflict error in operation
*/
type TagImageConflict struct {
	Payload *TagImageConflictBody
}

func (o *TagImageConflict) Error() string {
	return fmt.Sprintf("[POST /images/{name:.*}/tag][%d] tagImageConflict  %+v", 409, o.Payload)
}

func (o *TagImageConflict) GetPayload() *TagImageConflictBody {
	return o.Payload
}

func (o *TagImageConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(TagImageConflictBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTagImageInternalServerError creates a TagImageInternalServerError with default headers values
func NewTagImageInternalServerError() *TagImageInternalServerError {
	return &TagImageInternalServerError{}
}

/*TagImageInternalServerError handles this case with default header values.

Internal server error
*/
type TagImageInternalServerError struct {
	Payload *TagImageInternalServerErrorBody
}

func (o *TagImageInternalServerError) Error() string {
	return fmt.Sprintf("[POST /images/{name:.*}/tag][%d] tagImageInternalServerError  %+v", 500, o.Payload)
}

func (o *TagImageInternalServerError) GetPayload() *TagImageInternalServerErrorBody {
	return o.Payload
}

func (o *TagImageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(TagImageInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*TagImageBadRequestBody tag image bad request body
swagger:model TagImageBadRequestBody
*/
type TagImageBadRequestBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this tag image bad request body
func (o *TagImageBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TagImageBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TagImageBadRequestBody) UnmarshalBinary(b []byte) error {
	var res TagImageBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*TagImageConflictBody tag image conflict body
swagger:model TagImageConflictBody
*/
type TagImageConflictBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this tag image conflict body
func (o *TagImageConflictBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TagImageConflictBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TagImageConflictBody) UnmarshalBinary(b []byte) error {
	var res TagImageConflictBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*TagImageInternalServerErrorBody tag image internal server error body
swagger:model TagImageInternalServerErrorBody
*/
type TagImageInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this tag image internal server error body
func (o *TagImageInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TagImageInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TagImageInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res TagImageInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*TagImageNotFoundBody tag image not found body
swagger:model TagImageNotFoundBody
*/
type TagImageNotFoundBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this tag image not found body
func (o *TagImageNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TagImageNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TagImageNotFoundBody) UnmarshalBinary(b []byte) error {
	var res TagImageNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
