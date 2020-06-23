// Code generated by go-swagger; DO NOT EDIT.

package exec_compat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateExecReader is a Reader for the CreateExec structure.
type CreateExecReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateExecReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateExecCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewCreateExecNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateExecConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateExecInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateExecCreated creates a CreateExecCreated with default headers values
func NewCreateExecCreated() *CreateExecCreated {
	return &CreateExecCreated{}
}

/*CreateExecCreated handles this case with default header values.

no error
*/
type CreateExecCreated struct {
}

func (o *CreateExecCreated) Error() string {
	return fmt.Sprintf("[POST /containers/{name}/exec][%d] createExecCreated ", 201)
}

func (o *CreateExecCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateExecNotFound creates a CreateExecNotFound with default headers values
func NewCreateExecNotFound() *CreateExecNotFound {
	return &CreateExecNotFound{}
}

/*CreateExecNotFound handles this case with default header values.

No such container
*/
type CreateExecNotFound struct {
	Payload *CreateExecNotFoundBody
}

func (o *CreateExecNotFound) Error() string {
	return fmt.Sprintf("[POST /containers/{name}/exec][%d] createExecNotFound  %+v", 404, o.Payload)
}

func (o *CreateExecNotFound) GetPayload() *CreateExecNotFoundBody {
	return o.Payload
}

func (o *CreateExecNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateExecNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateExecConflict creates a CreateExecConflict with default headers values
func NewCreateExecConflict() *CreateExecConflict {
	return &CreateExecConflict{}
}

/*CreateExecConflict handles this case with default header values.

container is paused
*/
type CreateExecConflict struct {
}

func (o *CreateExecConflict) Error() string {
	return fmt.Sprintf("[POST /containers/{name}/exec][%d] createExecConflict ", 409)
}

func (o *CreateExecConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateExecInternalServerError creates a CreateExecInternalServerError with default headers values
func NewCreateExecInternalServerError() *CreateExecInternalServerError {
	return &CreateExecInternalServerError{}
}

/*CreateExecInternalServerError handles this case with default header values.

Internal server error
*/
type CreateExecInternalServerError struct {
	Payload *CreateExecInternalServerErrorBody
}

func (o *CreateExecInternalServerError) Error() string {
	return fmt.Sprintf("[POST /containers/{name}/exec][%d] createExecInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateExecInternalServerError) GetPayload() *CreateExecInternalServerErrorBody {
	return o.Payload
}

func (o *CreateExecInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateExecInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateExecBody create exec body
swagger:model CreateExecBody
*/
type CreateExecBody struct {

	// Attach to stderr of the exec command
	AttachStderr bool `json:"AttachStderr,omitempty"`

	// Attach to stdin of the exec command
	AttachStdin bool `json:"AttachStdin,omitempty"`

	// Attach to stdout of the exec command
	AttachStdout bool `json:"AttachStdout,omitempty"`

	// Command to run, as a string or array of strings.
	Cmd []string `json:"Cmd"`

	// "Override the key sequence for detaching a container. Format is a single character [a-Z] or ctrl-<value> where <value> is one of: a-z, @, ^, [, , or _."
	//
	DetachKeys string `json:"DetachKeys,omitempty"`

	// A list of environment variables in the form ["VAR=value", ...]
	Env []string `json:"Env"`

	// Runs the exec process with extended privileges
	Privileged *bool `json:"Privileged,omitempty"`

	// Allocate a pseudo-TTY
	Tty bool `json:"Tty,omitempty"`

	// "The user, and optionally, group to run the exec process inside the container. Format is one of: user, user:group, uid, or uid:gid."
	//
	User string `json:"User,omitempty"`

	// The working directory for the exec process inside the container.
	WorkingDir string `json:"WorkingDir,omitempty"`
}

// Validate validates this create exec body
func (o *CreateExecBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateExecBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateExecBody) UnmarshalBinary(b []byte) error {
	var res CreateExecBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateExecInternalServerErrorBody create exec internal server error body
swagger:model CreateExecInternalServerErrorBody
*/
type CreateExecInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this create exec internal server error body
func (o *CreateExecInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateExecInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateExecInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res CreateExecInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateExecNotFoundBody create exec not found body
swagger:model CreateExecNotFoundBody
*/
type CreateExecNotFoundBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this create exec not found body
func (o *CreateExecNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateExecNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateExecNotFoundBody) UnmarshalBinary(b []byte) error {
	var res CreateExecNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
