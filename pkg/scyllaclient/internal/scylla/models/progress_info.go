// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProgressInfo progress_info
//
// File transfer progress
//
// swagger:model progress_info
type ProgressInfo struct {

	// The current bytes
	CurrentBytes interface{} `json:"current_bytes,omitempty"`

	// direction
	Direction Direction `json:"direction,omitempty"`

	// The file name
	FileName string `json:"file_name,omitempty"`

	// The peer address
	Peer string `json:"peer,omitempty"`

	// The session index
	SessionIndex int32 `json:"session_index,omitempty"`

	// The total bytes
	TotalBytes interface{} `json:"total_bytes,omitempty"`
}

// Validate validates this progress info
func (m *ProgressInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProgressInfo) validateDirection(formats strfmt.Registry) error {

	if swag.IsZero(m.Direction) { // not required
		return nil
	}

	if err := m.Direction.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("direction")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProgressInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProgressInfo) UnmarshalBinary(b []byte) error {
	var res ProgressInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
