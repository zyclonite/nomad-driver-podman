// Code generated by go-swagger; DO NOT EDIT.

package containers

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

// NewLibpodUnpauseContainerParams creates a new LibpodUnpauseContainerParams object
// with the default values initialized.
func NewLibpodUnpauseContainerParams() *LibpodUnpauseContainerParams {
	var ()
	return &LibpodUnpauseContainerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLibpodUnpauseContainerParamsWithTimeout creates a new LibpodUnpauseContainerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLibpodUnpauseContainerParamsWithTimeout(timeout time.Duration) *LibpodUnpauseContainerParams {
	var ()
	return &LibpodUnpauseContainerParams{

		timeout: timeout,
	}
}

// NewLibpodUnpauseContainerParamsWithContext creates a new LibpodUnpauseContainerParams object
// with the default values initialized, and the ability to set a context for a request
func NewLibpodUnpauseContainerParamsWithContext(ctx context.Context) *LibpodUnpauseContainerParams {
	var ()
	return &LibpodUnpauseContainerParams{

		Context: ctx,
	}
}

// NewLibpodUnpauseContainerParamsWithHTTPClient creates a new LibpodUnpauseContainerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLibpodUnpauseContainerParamsWithHTTPClient(client *http.Client) *LibpodUnpauseContainerParams {
	var ()
	return &LibpodUnpauseContainerParams{
		HTTPClient: client,
	}
}

/*LibpodUnpauseContainerParams contains all the parameters to send to the API endpoint
for the libpod unpause container operation typically these are written to a http.Request
*/
type LibpodUnpauseContainerParams struct {

	/*Name
	  the name or ID of the container

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the libpod unpause container params
func (o *LibpodUnpauseContainerParams) WithTimeout(timeout time.Duration) *LibpodUnpauseContainerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the libpod unpause container params
func (o *LibpodUnpauseContainerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the libpod unpause container params
func (o *LibpodUnpauseContainerParams) WithContext(ctx context.Context) *LibpodUnpauseContainerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the libpod unpause container params
func (o *LibpodUnpauseContainerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the libpod unpause container params
func (o *LibpodUnpauseContainerParams) WithHTTPClient(client *http.Client) *LibpodUnpauseContainerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the libpod unpause container params
func (o *LibpodUnpauseContainerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the libpod unpause container params
func (o *LibpodUnpauseContainerParams) WithName(name string) *LibpodUnpauseContainerParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the libpod unpause container params
func (o *LibpodUnpauseContainerParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *LibpodUnpauseContainerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
