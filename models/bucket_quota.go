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
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BucketQuota bucket quota
//
// swagger:model bucketQuota
type BucketQuota struct {

	// quota
	Quota int64 `json:"quota,omitempty"`

	// type
	// Enum: [hard fifo]
	Type string `json:"type,omitempty"`
}

// Validate validates this bucket quota
func (m *BucketQuota) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var bucketQuotaTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["hard","fifo"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bucketQuotaTypeTypePropEnum = append(bucketQuotaTypeTypePropEnum, v)
	}
}

const (

	// BucketQuotaTypeHard captures enum value "hard"
	BucketQuotaTypeHard string = "hard"

	// BucketQuotaTypeFifo captures enum value "fifo"
	BucketQuotaTypeFifo string = "fifo"
)

// prop value enum
func (m *BucketQuota) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, bucketQuotaTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BucketQuota) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this bucket quota based on context it is used
func (m *BucketQuota) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BucketQuota) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BucketQuota) UnmarshalBinary(b []byte) error {
	var res BucketQuota
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
