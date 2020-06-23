// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NetworkCreateOptions NetworkCreateOptions describes options to create a network
//
// swagger:model NetworkCreateOptions
type NetworkCreateOptions struct {

	// disable DNS
	DisableDNS bool `json:"DisableDNS,omitempty"`

	// driver
	Driver string `json:"Driver,omitempty"`

	// gateway
	Gateway IP `json:"Gateway,omitempty"`

	// internal
	Internal bool `json:"Internal,omitempty"`

	// mac v l a n
	MacVLAN string `json:"MacVLAN,omitempty"`

	// range
	Range *IPNet `json:"Range,omitempty"`

	// subnet
	Subnet *IPNet `json:"Subnet,omitempty"`
}

// Validate validates this network create options
func (m *NetworkCreateOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGateway(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkCreateOptions) validateGateway(formats strfmt.Registry) error {

	if swag.IsZero(m.Gateway) { // not required
		return nil
	}

	if err := m.Gateway.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Gateway")
		}
		return err
	}

	return nil
}

func (m *NetworkCreateOptions) validateRange(formats strfmt.Registry) error {

	if swag.IsZero(m.Range) { // not required
		return nil
	}

	if m.Range != nil {
		if err := m.Range.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Range")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkCreateOptions) validateSubnet(formats strfmt.Registry) error {

	if swag.IsZero(m.Subnet) { // not required
		return nil
	}

	if m.Subnet != nil {
		if err := m.Subnet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Subnet")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkCreateOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkCreateOptions) UnmarshalBinary(b []byte) error {
	var res NetworkCreateOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
