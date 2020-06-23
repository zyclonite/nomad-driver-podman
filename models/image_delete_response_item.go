// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ImageDeleteResponseItem ImageDeleteResponseItem image delete response item
//
// swagger:model ImageDeleteResponseItem
type ImageDeleteResponseItem struct {

	// The image ID of an image that was deleted
	Deleted string `json:"Deleted,omitempty"`

	// The image ID of an image that was untagged
	Untagged string `json:"Untagged,omitempty"`
}

// Validate validates this image delete response item
func (m *ImageDeleteResponseItem) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ImageDeleteResponseItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageDeleteResponseItem) UnmarshalBinary(b []byte) error {
	var res ImageDeleteResponseItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
