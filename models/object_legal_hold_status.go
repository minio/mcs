// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Console Server
// Copyright (c) 2020 MinIO, Inc.
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
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ObjectLegalHoldStatus object legal hold status
//
// swagger:model objectLegalHoldStatus
type ObjectLegalHoldStatus string

const (

	// ObjectLegalHoldStatusEnabled captures enum value "enabled"
	ObjectLegalHoldStatusEnabled ObjectLegalHoldStatus = "enabled"

	// ObjectLegalHoldStatusDisabled captures enum value "disabled"
	ObjectLegalHoldStatusDisabled ObjectLegalHoldStatus = "disabled"
)

// for schema
var objectLegalHoldStatusEnum []interface{}

func init() {
	var res []ObjectLegalHoldStatus
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		objectLegalHoldStatusEnum = append(objectLegalHoldStatusEnum, v)
	}
}

func (m ObjectLegalHoldStatus) validateObjectLegalHoldStatusEnum(path, location string, value ObjectLegalHoldStatus) error {
	if err := validate.EnumCase(path, location, value, objectLegalHoldStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this object legal hold status
func (m ObjectLegalHoldStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateObjectLegalHoldStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
