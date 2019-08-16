// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// BackupProgress backup progress
// swagger:model backupProgress
type BackupProgress struct {

	// dcs
	Dcs []string `json:"dcs"`

	// failed
	Failed int64 `json:"failed,omitempty"`

	// keyspaces
	Keyspaces []*KeyspaceProgress `json:"keyspaces"`

	// size
	Size int64 `json:"size,omitempty"`

	// skipped
	Skipped int64 `json:"skipped,omitempty"`

	// uploaded
	Uploaded int64 `json:"uploaded,omitempty"`
}

// Validate validates this backup progress
func (m *BackupProgress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKeyspaces(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupProgress) validateKeyspaces(formats strfmt.Registry) error {

	if swag.IsZero(m.Keyspaces) { // not required
		return nil
	}

	for i := 0; i < len(m.Keyspaces); i++ {
		if swag.IsZero(m.Keyspaces[i]) { // not required
			continue
		}

		if m.Keyspaces[i] != nil {
			if err := m.Keyspaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("keyspaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackupProgress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupProgress) UnmarshalBinary(b []byte) error {
	var res BackupProgress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}