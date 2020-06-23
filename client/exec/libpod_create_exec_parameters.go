// Code generated by go-swagger; DO NOT EDIT.

package exec

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewLibpodCreateExecParams creates a new LibpodCreateExecParams object
// with the default values initialized.
func NewLibpodCreateExecParams() *LibpodCreateExecParams {
	var ()
	return &LibpodCreateExecParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLibpodCreateExecParamsWithTimeout creates a new LibpodCreateExecParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLibpodCreateExecParamsWithTimeout(timeout time.Duration) *LibpodCreateExecParams {
	var ()
	return &LibpodCreateExecParams{

		timeout: timeout,
	}
}

// NewLibpodCreateExecParamsWithContext creates a new LibpodCreateExecParams object
// with the default values initialized, and the ability to set a context for a request
func NewLibpodCreateExecParamsWithContext(ctx context.Context) *LibpodCreateExecParams {
	var ()
	return &LibpodCreateExecParams{

		Context: ctx,
	}
}

// NewLibpodCreateExecParamsWithHTTPClient creates a new LibpodCreateExecParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLibpodCreateExecParamsWithHTTPClient(client *http.Client) *LibpodCreateExecParams {
	var ()
	return &LibpodCreateExecParams{
		HTTPClient: client,
	}
}

/*LibpodCreateExecParams contains all the parameters to send to the API endpoint
for the libpod create exec operation typically these are written to a http.Request
*/
type LibpodCreateExecParams struct {

	/*Control
	  Attributes for create

	*/
	Control LibpodCreateExecBody
	/*Name
	  name of container

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the libpod create exec params
func (o *LibpodCreateExecParams) WithTimeout(timeout time.Duration) *LibpodCreateExecParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the libpod create exec params
func (o *LibpodCreateExecParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the libpod create exec params
func (o *LibpodCreateExecParams) WithContext(ctx context.Context) *LibpodCreateExecParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the libpod create exec params
func (o *LibpodCreateExecParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the libpod create exec params
func (o *LibpodCreateExecParams) WithHTTPClient(client *http.Client) *LibpodCreateExecParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the libpod create exec params
func (o *LibpodCreateExecParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithControl adds the control to the libpod create exec params
func (o *LibpodCreateExecParams) WithControl(control LibpodCreateExecBody) *LibpodCreateExecParams {
	o.SetControl(control)
	return o
}

// SetControl adds the control to the libpod create exec params
func (o *LibpodCreateExecParams) SetControl(control LibpodCreateExecBody) {
	o.Control = control
}

// WithName adds the name to the libpod create exec params
func (o *LibpodCreateExecParams) WithName(name string) *LibpodCreateExecParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the libpod create exec params
func (o *LibpodCreateExecParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *LibpodCreateExecParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Control); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
