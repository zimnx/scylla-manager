// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CollectdValue collectd_value
//
// Holds a collectd value
//
// swagger:model collectd_value
type CollectdValue struct {

	// An array of values
	Values []interface{} `json:"values"`
}

// Validate validates this collectd value
func (m *CollectdValue) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CollectdValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CollectdValue) UnmarshalBinary(b []byte) error {
	var res CollectdValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
