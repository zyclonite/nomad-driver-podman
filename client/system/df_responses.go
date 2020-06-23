// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/pascomnet/nomad-driver-podman/models"
)

// DfReader is a Reader for the Df structure.
type DfReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DfReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDfOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewDfInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDfOK creates a DfOK with default headers values
func NewDfOK() *DfOK {
	return &DfOK{}
}

/*DfOK handles this case with default header values.

Disk usage
*/
type DfOK struct {
	Payload *DfOKBody
}

func (o *DfOK) Error() string {
	return fmt.Sprintf("[GET /libpod/system/df][%d] dfOK  %+v", 200, o.Payload)
}

func (o *DfOK) GetPayload() *DfOKBody {
	return o.Payload
}

func (o *DfOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DfOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDfInternalServerError creates a DfInternalServerError with default headers values
func NewDfInternalServerError() *DfInternalServerError {
	return &DfInternalServerError{}
}

/*DfInternalServerError handles this case with default header values.

Internal server error
*/
type DfInternalServerError struct {
	Payload *DfInternalServerErrorBody
}

func (o *DfInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/system/df][%d] dfInternalServerError  %+v", 500, o.Payload)
}

func (o *DfInternalServerError) GetPayload() *DfInternalServerErrorBody {
	return o.Payload
}

func (o *DfInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DfInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DfInternalServerErrorBody df internal server error body
swagger:model DfInternalServerErrorBody
*/
type DfInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this df internal server error body
func (o *DfInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DfInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DfInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res DfInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DfOKBody df o k body
swagger:model DfOKBody
*/
type DfOKBody struct {

	// containers
	Containers []*models.SystemDfContainerReport `json:"Containers"`

	// images
	Images []*models.SystemDfImageReport `json:"Images"`

	// volumes
	Volumes []*models.SystemDfVolumeReport `json:"Volumes"`
}

// Validate validates this df o k body
func (o *DfOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateContainers(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateVolumes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DfOKBody) validateContainers(formats strfmt.Registry) error {

	if swag.IsZero(o.Containers) { // not required
		return nil
	}

	for i := 0; i < len(o.Containers); i++ {
		if swag.IsZero(o.Containers[i]) { // not required
			continue
		}

		if o.Containers[i] != nil {
			if err := o.Containers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dfOK" + "." + "Containers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *DfOKBody) validateImages(formats strfmt.Registry) error {

	if swag.IsZero(o.Images) { // not required
		return nil
	}

	for i := 0; i < len(o.Images); i++ {
		if swag.IsZero(o.Images[i]) { // not required
			continue
		}

		if o.Images[i] != nil {
			if err := o.Images[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dfOK" + "." + "Images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *DfOKBody) validateVolumes(formats strfmt.Registry) error {

	if swag.IsZero(o.Volumes) { // not required
		return nil
	}

	for i := 0; i < len(o.Volumes); i++ {
		if swag.IsZero(o.Volumes[i]) { // not required
			continue
		}

		if o.Volumes[i] != nil {
			if err := o.Volumes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dfOK" + "." + "Volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DfOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DfOKBody) UnmarshalBinary(b []byte) error {
	var res DfOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
