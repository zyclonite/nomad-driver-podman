// Code generated by go-swagger; DO NOT EDIT.

package pods

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

// NewKillPodParams creates a new KillPodParams object
// with the default values initialized.
func NewKillPodParams() *KillPodParams {
	var (
		signalDefault = string("SIGKILL")
	)
	return &KillPodParams{
		Signal: &signalDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewKillPodParamsWithTimeout creates a new KillPodParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewKillPodParamsWithTimeout(timeout time.Duration) *KillPodParams {
	var (
		signalDefault = string("SIGKILL")
	)
	return &KillPodParams{
		Signal: &signalDefault,

		timeout: timeout,
	}
}

// NewKillPodParamsWithContext creates a new KillPodParams object
// with the default values initialized, and the ability to set a context for a request
func NewKillPodParamsWithContext(ctx context.Context) *KillPodParams {
	var (
		signalDefault = string("SIGKILL")
	)
	return &KillPodParams{
		Signal: &signalDefault,

		Context: ctx,
	}
}

// NewKillPodParamsWithHTTPClient creates a new KillPodParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewKillPodParamsWithHTTPClient(client *http.Client) *KillPodParams {
	var (
		signalDefault = string("SIGKILL")
	)
	return &KillPodParams{
		Signal:     &signalDefault,
		HTTPClient: client,
	}
}

/*KillPodParams contains all the parameters to send to the API endpoint
for the kill pod operation typically these are written to a http.Request
*/
type KillPodParams struct {

	/*Name
	  the name or ID of the pod

	*/
	Name string
	/*Signal
	  signal to be sent to pod

	*/
	Signal *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the kill pod params
func (o *KillPodParams) WithTimeout(timeout time.Duration) *KillPodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kill pod params
func (o *KillPodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kill pod params
func (o *KillPodParams) WithContext(ctx context.Context) *KillPodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kill pod params
func (o *KillPodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kill pod params
func (o *KillPodParams) WithHTTPClient(client *http.Client) *KillPodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kill pod params
func (o *KillPodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the kill pod params
func (o *KillPodParams) WithName(name string) *KillPodParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the kill pod params
func (o *KillPodParams) SetName(name string) {
	o.Name = name
}

// WithSignal adds the signal to the kill pod params
func (o *KillPodParams) WithSignal(signal *string) *KillPodParams {
	o.SetSignal(signal)
	return o
}

// SetSignal adds the signal to the kill pod params
func (o *KillPodParams) SetSignal(signal *string) {
	o.Signal = signal
}

// WriteToRequest writes these params to a swagger request
func (o *KillPodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.Signal != nil {

		// query param signal
		var qrSignal string
		if o.Signal != nil {
			qrSignal = *o.Signal
		}
		qSignal := qrSignal
		if qSignal != "" {
			if err := r.SetQueryParam("signal", qSignal); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
