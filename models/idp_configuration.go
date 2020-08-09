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
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IdpConfiguration idp configuration
//
// swagger:model idpConfiguration
type IdpConfiguration struct {

	// active directory
	ActiveDirectory *IdpConfigurationActiveDirectory `json:"active_directory,omitempty"`

	// oidc
	Oidc *IdpConfigurationOidc `json:"oidc,omitempty"`
}

// Validate validates this idp configuration
func (m *IdpConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActiveDirectory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOidc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdpConfiguration) validateActiveDirectory(formats strfmt.Registry) error {

	if swag.IsZero(m.ActiveDirectory) { // not required
		return nil
	}

	if m.ActiveDirectory != nil {
		if err := m.ActiveDirectory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("active_directory")
			}
			return err
		}
	}

	return nil
}

func (m *IdpConfiguration) validateOidc(formats strfmt.Registry) error {

	if swag.IsZero(m.Oidc) { // not required
		return nil
	}

	if m.Oidc != nil {
		if err := m.Oidc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oidc")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IdpConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdpConfiguration) UnmarshalBinary(b []byte) error {
	var res IdpConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IdpConfigurationActiveDirectory idp configuration active directory
//
// swagger:model IdpConfigurationActiveDirectory
type IdpConfigurationActiveDirectory struct {

	// group name attribute
	GroupNameAttribute string `json:"group_name_attribute,omitempty"`

	// group search base dn
	GroupSearchBaseDn string `json:"group_search_base_dn,omitempty"`

	// group search filter
	GroupSearchFilter string `json:"group_search_filter,omitempty"`

	// server insecure
	ServerInsecure bool `json:"server_insecure,omitempty"`

	// skip tls verification
	SkipTLSVerification bool `json:"skip_tls_verification,omitempty"`

	// url
	// Required: true
	URL *string `json:"url"`

	// user search filter
	// Required: true
	UserSearchFilter *string `json:"user_search_filter"`

	// username format
	// Required: true
	UsernameFormat *string `json:"username_format"`
}

// Validate validates this idp configuration active directory
func (m *IdpConfigurationActiveDirectory) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserSearchFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsernameFormat(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdpConfigurationActiveDirectory) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("active_directory"+"."+"url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

func (m *IdpConfigurationActiveDirectory) validateUserSearchFilter(formats strfmt.Registry) error {

	if err := validate.Required("active_directory"+"."+"user_search_filter", "body", m.UserSearchFilter); err != nil {
		return err
	}

	return nil
}

func (m *IdpConfigurationActiveDirectory) validateUsernameFormat(formats strfmt.Registry) error {

	if err := validate.Required("active_directory"+"."+"username_format", "body", m.UsernameFormat); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IdpConfigurationActiveDirectory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdpConfigurationActiveDirectory) UnmarshalBinary(b []byte) error {
	var res IdpConfigurationActiveDirectory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IdpConfigurationOidc idp configuration oidc
//
// swagger:model IdpConfigurationOidc
type IdpConfigurationOidc struct {

	// client id
	// Required: true
	ClientID *string `json:"client_id"`

	// secret id
	// Required: true
	SecretID *string `json:"secret_id"`

	// url
	// Required: true
	URL *string `json:"url"`
}

// Validate validates this idp configuration oidc
func (m *IdpConfigurationOidc) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecretID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdpConfigurationOidc) validateClientID(formats strfmt.Registry) error {

	if err := validate.Required("oidc"+"."+"client_id", "body", m.ClientID); err != nil {
		return err
	}

	return nil
}

func (m *IdpConfigurationOidc) validateSecretID(formats strfmt.Registry) error {

	if err := validate.Required("oidc"+"."+"secret_id", "body", m.SecretID); err != nil {
		return err
	}

	return nil
}

func (m *IdpConfigurationOidc) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("oidc"+"."+"url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IdpConfigurationOidc) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdpConfigurationOidc) UnmarshalBinary(b []byte) error {
	var res IdpConfigurationOidc
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
