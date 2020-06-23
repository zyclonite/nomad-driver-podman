// Code generated by go-swagger; DO NOT EDIT.

package system_compat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetInfoReader is a Reader for the GetInfo structure.
type GetInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInfoOK creates a GetInfoOK with default headers values
func NewGetInfoOK() *GetInfoOK {
	return &GetInfoOK{}
}

/*GetInfoOK handles this case with default header values.

to be determined
*/
type GetInfoOK struct {
}

func (o *GetInfoOK) Error() string {
	return fmt.Sprintf("[GET /info][%d] getInfoOK ", 200)
}

func (o *GetInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInfoInternalServerError creates a GetInfoInternalServerError with default headers values
func NewGetInfoInternalServerError() *GetInfoInternalServerError {
	return &GetInfoInternalServerError{}
}

/*GetInfoInternalServerError handles this case with default header values.

Internal server error
*/
type GetInfoInternalServerError struct {
	Payload *GetInfoInternalServerErrorBody
}

func (o *GetInfoInternalServerError) Error() string {
	return fmt.Sprintf("[GET /info][%d] getInfoInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInfoInternalServerError) GetPayload() *GetInfoInternalServerErrorBody {
	return o.Payload
}

func (o *GetInfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetInfoInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetInfoInternalServerErrorBody get info internal server error body
swagger:model GetInfoInternalServerErrorBody
*/
type GetInfoInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this get info internal server error body
func (o *GetInfoInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetInfoInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetInfoInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetInfoInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
