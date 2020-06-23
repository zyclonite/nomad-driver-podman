// Code generated by go-swagger; DO NOT EDIT.

package networks

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

	"github.com/pascomnet/nomad-driver-podman/models"
)

// NewLibpodCreateNetworkParams creates a new LibpodCreateNetworkParams object
// with the default values initialized.
func NewLibpodCreateNetworkParams() *LibpodCreateNetworkParams {
	var ()
	return &LibpodCreateNetworkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLibpodCreateNetworkParamsWithTimeout creates a new LibpodCreateNetworkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLibpodCreateNetworkParamsWithTimeout(timeout time.Duration) *LibpodCreateNetworkParams {
	var ()
	return &LibpodCreateNetworkParams{

		timeout: timeout,
	}
}

// NewLibpodCreateNetworkParamsWithContext creates a new LibpodCreateNetworkParams object
// with the default values initialized, and the ability to set a context for a request
func NewLibpodCreateNetworkParamsWithContext(ctx context.Context) *LibpodCreateNetworkParams {
	var ()
	return &LibpodCreateNetworkParams{

		Context: ctx,
	}
}

// NewLibpodCreateNetworkParamsWithHTTPClient creates a new LibpodCreateNetworkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLibpodCreateNetworkParamsWithHTTPClient(client *http.Client) *LibpodCreateNetworkParams {
	var ()
	return &LibpodCreateNetworkParams{
		HTTPClient: client,
	}
}

/*LibpodCreateNetworkParams contains all the parameters to send to the API endpoint
for the libpod create network operation typically these are written to a http.Request
*/
type LibpodCreateNetworkParams struct {

	/*Create
	  attributes for creating a container

	*/
	Create *models.NetworkCreateOptions
	/*Name
	  optional name for new network

	*/
	Name *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the libpod create network params
func (o *LibpodCreateNetworkParams) WithTimeout(timeout time.Duration) *LibpodCreateNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the libpod create network params
func (o *LibpodCreateNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the libpod create network params
func (o *LibpodCreateNetworkParams) WithContext(ctx context.Context) *LibpodCreateNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the libpod create network params
func (o *LibpodCreateNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the libpod create network params
func (o *LibpodCreateNetworkParams) WithHTTPClient(client *http.Client) *LibpodCreateNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the libpod create network params
func (o *LibpodCreateNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreate adds the create to the libpod create network params
func (o *LibpodCreateNetworkParams) WithCreate(create *models.NetworkCreateOptions) *LibpodCreateNetworkParams {
	o.SetCreate(create)
	return o
}

// SetCreate adds the create to the libpod create network params
func (o *LibpodCreateNetworkParams) SetCreate(create *models.NetworkCreateOptions) {
	o.Create = create
}

// WithName adds the name to the libpod create network params
func (o *LibpodCreateNetworkParams) WithName(name *string) *LibpodCreateNetworkParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the libpod create network params
func (o *LibpodCreateNetworkParams) SetName(name *string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *LibpodCreateNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Create != nil {
		if err := r.SetBodyParam(o.Create); err != nil {
			return err
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
