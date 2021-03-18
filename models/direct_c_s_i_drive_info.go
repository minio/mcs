// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Console Server
// Copyright (c) 2021 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DirectCSIDriveInfo direct c s i drive info
//
// swagger:model directCSIDriveInfo
type DirectCSIDriveInfo struct {

	// allocated
	Allocated int64 `json:"allocated,omitempty"`

	// capacity
	Capacity int64 `json:"capacity,omitempty"`

	// drive
	Drive string `json:"drive,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// node
	Node string `json:"node,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// volumes
	Volumes int64 `json:"volumes,omitempty"`
}

// Validate validates this direct c s i drive info
func (m *DirectCSIDriveInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DirectCSIDriveInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectCSIDriveInfo) UnmarshalBinary(b []byte) error {
	var res DirectCSIDriveInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
