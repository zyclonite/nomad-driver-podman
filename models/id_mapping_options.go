// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IDMappingOptions IDMappingOptions are used for specifying how ID mapping should be set up for
// a layer or container.
//
// swagger:model IDMappingOptions
type IDMappingOptions struct {

	// auto user ns
	AutoUserNs bool `json:"AutoUserNs,omitempty"`

	// auto user ns opts
	AutoUserNsOpts *AutoUserNsOptions `json:"AutoUserNsOpts,omitempty"`

	// g ID map
	GIDMap []*IDMap `json:"GIDMap"`

	// host g ID mapping
	HostGIDMapping bool `json:"HostGIDMapping,omitempty"`

	// UIDMap and GIDMap are used for setting up a layer's root filesystem
	// for use inside of a user namespace where ID mapping is being used.
	// If HostUIDMapping/HostGIDMapping is true, no mapping of the
	// respective type will be used.  Otherwise, if UIDMap and/or GIDMap
	// contain at least one mapping, one or both will be used.  By default,
	// if neither of those conditions apply, if the layer has a parent
	// layer, the parent layer's mapping will be used, and if it does not
	// have a parent layer, the mapping which was passed to the Store
	// object when it was initialized will be used.
	HostUIDMapping bool `json:"HostUIDMapping,omitempty"`

	// UID map
	UIDMap []*IDMap `json:"UIDMap"`
}

// Validate validates this ID mapping options
func (m *IDMappingOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoUserNsOpts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGIDMap(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUIDMap(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IDMappingOptions) validateAutoUserNsOpts(formats strfmt.Registry) error {

	if swag.IsZero(m.AutoUserNsOpts) { // not required
		return nil
	}

	if m.AutoUserNsOpts != nil {
		if err := m.AutoUserNsOpts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AutoUserNsOpts")
			}
			return err
		}
	}

	return nil
}

func (m *IDMappingOptions) validateGIDMap(formats strfmt.Registry) error {

	if swag.IsZero(m.GIDMap) { // not required
		return nil
	}

	for i := 0; i < len(m.GIDMap); i++ {
		if swag.IsZero(m.GIDMap[i]) { // not required
			continue
		}

		if m.GIDMap[i] != nil {
			if err := m.GIDMap[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GIDMap" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IDMappingOptions) validateUIDMap(formats strfmt.Registry) error {

	if swag.IsZero(m.UIDMap) { // not required
		return nil
	}

	for i := 0; i < len(m.UIDMap); i++ {
		if swag.IsZero(m.UIDMap[i]) { // not required
			continue
		}

		if m.UIDMap[i] != nil {
			if err := m.UIDMap[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("UIDMap" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IDMappingOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IDMappingOptions) UnmarshalBinary(b []byte) error {
	var res IDMappingOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
