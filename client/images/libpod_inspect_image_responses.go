// Code generated by go-swagger; DO NOT EDIT.

package images

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
	"github.com/go-openapi/validate"

	"github.com/pascomnet/nomad-driver-podman/models"
)

// LibpodInspectImageReader is a Reader for the LibpodInspectImage structure.
type LibpodInspectImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LibpodInspectImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLibpodInspectImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewLibpodInspectImageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLibpodInspectImageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLibpodInspectImageOK creates a LibpodInspectImageOK with default headers values
func NewLibpodInspectImageOK() *LibpodInspectImageOK {
	return &LibpodInspectImageOK{}
}

/*LibpodInspectImageOK handles this case with default header values.

Inspect image
*/
type LibpodInspectImageOK struct {
	Payload *LibpodInspectImageOKBody
}

func (o *LibpodInspectImageOK) Error() string {
	return fmt.Sprintf("[GET /libpod/images/{name:.*}/json][%d] libpodInspectImageOK  %+v", 200, o.Payload)
}

func (o *LibpodInspectImageOK) GetPayload() *LibpodInspectImageOKBody {
	return o.Payload
}

func (o *LibpodInspectImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodInspectImageOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodInspectImageNotFound creates a LibpodInspectImageNotFound with default headers values
func NewLibpodInspectImageNotFound() *LibpodInspectImageNotFound {
	return &LibpodInspectImageNotFound{}
}

/*LibpodInspectImageNotFound handles this case with default header values.

No such image
*/
type LibpodInspectImageNotFound struct {
	Payload *LibpodInspectImageNotFoundBody
}

func (o *LibpodInspectImageNotFound) Error() string {
	return fmt.Sprintf("[GET /libpod/images/{name:.*}/json][%d] libpodInspectImageNotFound  %+v", 404, o.Payload)
}

func (o *LibpodInspectImageNotFound) GetPayload() *LibpodInspectImageNotFoundBody {
	return o.Payload
}

func (o *LibpodInspectImageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodInspectImageNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodInspectImageInternalServerError creates a LibpodInspectImageInternalServerError with default headers values
func NewLibpodInspectImageInternalServerError() *LibpodInspectImageInternalServerError {
	return &LibpodInspectImageInternalServerError{}
}

/*LibpodInspectImageInternalServerError handles this case with default header values.

Internal server error
*/
type LibpodInspectImageInternalServerError struct {
	Payload *LibpodInspectImageInternalServerErrorBody
}

func (o *LibpodInspectImageInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/images/{name:.*}/json][%d] libpodInspectImageInternalServerError  %+v", 500, o.Payload)
}

func (o *LibpodInspectImageInternalServerError) GetPayload() *LibpodInspectImageInternalServerErrorBody {
	return o.Payload
}

func (o *LibpodInspectImageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodInspectImageInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LibpodInspectImageInternalServerErrorBody libpod inspect image internal server error body
swagger:model LibpodInspectImageInternalServerErrorBody
*/
type LibpodInspectImageInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod inspect image internal server error body
func (o *LibpodInspectImageInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodInspectImageInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodInspectImageInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res LibpodInspectImageInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodInspectImageNotFoundBody libpod inspect image not found body
swagger:model LibpodInspectImageNotFoundBody
*/
type LibpodInspectImageNotFoundBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod inspect image not found body
func (o *LibpodInspectImageNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodInspectImageNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodInspectImageNotFoundBody) UnmarshalBinary(b []byte) error {
	var res LibpodInspectImageNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodInspectImageOKBody libpod inspect image o k body
swagger:model LibpodInspectImageOKBody
*/
type LibpodInspectImageOKBody struct {

	// annotations
	Annotations map[string]string `json:"Annotations,omitempty"`

	// architecture
	Architecture string `json:"Architecture,omitempty"`

	// author
	Author string `json:"Author,omitempty"`

	// comment
	Comment string `json:"Comment,omitempty"`

	// config
	Config *models.ImageConfig `json:"Config,omitempty"`

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"Created,omitempty"`

	// digest
	Digest models.Digest `json:"Digest,omitempty"`

	// graph driver
	GraphDriver *models.Data `json:"GraphDriver,omitempty"`

	// healthcheck
	Healthcheck *models.Schema2HealthConfig `json:"Healthcheck,omitempty"`

	// history
	History []*models.History `json:"History"`

	// ID
	ID string `json:"Id,omitempty"`

	// labels
	Labels map[string]string `json:"Labels,omitempty"`

	// manifest type
	ManifestType string `json:"ManifestType,omitempty"`

	// names history
	NamesHistory []string `json:"NamesHistory"`

	// os
	Os string `json:"Os,omitempty"`

	// parent
	Parent string `json:"Parent,omitempty"`

	// repo digests
	RepoDigests []string `json:"RepoDigests"`

	// repo tags
	RepoTags []string `json:"RepoTags"`

	// root f s
	RootFS *models.RootFS `json:"RootFS,omitempty"`

	// size
	Size int64 `json:"Size,omitempty"`

	// user
	User string `json:"User,omitempty"`

	// version
	Version string `json:"Version,omitempty"`

	// virtual size
	VirtualSize int64 `json:"VirtualSize,omitempty"`
}

// Validate validates this libpod inspect image o k body
func (o *LibpodInspectImageOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDigest(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateGraphDriver(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateHealthcheck(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateHistory(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRootFS(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LibpodInspectImageOKBody) validateConfig(formats strfmt.Registry) error {

	if swag.IsZero(o.Config) { // not required
		return nil
	}

	if o.Config != nil {
		if err := o.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("libpodInspectImageOK" + "." + "Config")
			}
			return err
		}
	}

	return nil
}

func (o *LibpodInspectImageOKBody) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(o.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("libpodInspectImageOK"+"."+"Created", "body", "date-time", o.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *LibpodInspectImageOKBody) validateDigest(formats strfmt.Registry) error {

	if swag.IsZero(o.Digest) { // not required
		return nil
	}

	if err := o.Digest.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("libpodInspectImageOK" + "." + "Digest")
		}
		return err
	}

	return nil
}

func (o *LibpodInspectImageOKBody) validateGraphDriver(formats strfmt.Registry) error {

	if swag.IsZero(o.GraphDriver) { // not required
		return nil
	}

	if o.GraphDriver != nil {
		if err := o.GraphDriver.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("libpodInspectImageOK" + "." + "GraphDriver")
			}
			return err
		}
	}

	return nil
}

func (o *LibpodInspectImageOKBody) validateHealthcheck(formats strfmt.Registry) error {

	if swag.IsZero(o.Healthcheck) { // not required
		return nil
	}

	if o.Healthcheck != nil {
		if err := o.Healthcheck.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("libpodInspectImageOK" + "." + "Healthcheck")
			}
			return err
		}
	}

	return nil
}

func (o *LibpodInspectImageOKBody) validateHistory(formats strfmt.Registry) error {

	if swag.IsZero(o.History) { // not required
		return nil
	}

	for i := 0; i < len(o.History); i++ {
		if swag.IsZero(o.History[i]) { // not required
			continue
		}

		if o.History[i] != nil {
			if err := o.History[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("libpodInspectImageOK" + "." + "History" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *LibpodInspectImageOKBody) validateRootFS(formats strfmt.Registry) error {

	if swag.IsZero(o.RootFS) { // not required
		return nil
	}

	if o.RootFS != nil {
		if err := o.RootFS.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("libpodInspectImageOK" + "." + "RootFS")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *LibpodInspectImageOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodInspectImageOKBody) UnmarshalBinary(b []byte) error {
	var res LibpodInspectImageOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
