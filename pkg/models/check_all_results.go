// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CheckAllResults check all results
// swagger:model CheckAllResults
type CheckAllResults struct {

	// o k
	OK *bool `json:"OK,omitempty"`

	// hosts
	Hosts []*CheckAllResultsHostsItems0 `json:"hosts"`

	// hosts healthy
	HostsHealthy int32 `json:"hosts-healthy,omitempty"`

	// hosts number
	HostsNumber int32 `json:"hosts-number,omitempty"`

	// responses
	Responses map[string]CheckAllPodResult `json:"responses,omitempty"`
}

// Validate validates this check all results
func (m *CheckAllResults) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHosts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponses(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CheckAllResults) validateHosts(formats strfmt.Registry) error {

	if swag.IsZero(m.Hosts) { // not required
		return nil
	}

	for i := 0; i < len(m.Hosts); i++ {
		if swag.IsZero(m.Hosts[i]) { // not required
			continue
		}

		if m.Hosts[i] != nil {
			if err := m.Hosts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hosts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CheckAllResults) validateResponses(formats strfmt.Registry) error {

	if swag.IsZero(m.Responses) { // not required
		return nil
	}

	for k := range m.Responses {

		if err := validate.Required("responses"+"."+k, "body", m.Responses[k]); err != nil {
			return err
		}
		if val, ok := m.Responses[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CheckAllResults) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CheckAllResults) UnmarshalBinary(b []byte) error {
	var res CheckAllResults
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CheckAllResultsHostsItems0 check all results hosts items0
// swagger:model CheckAllResultsHostsItems0
type CheckAllResultsHostsItems0 struct {

	// host IP
	// Format: ipv4
	HostIP strfmt.IPv4 `json:"hostIP,omitempty"`

	// pod IP
	// Format: ipv4
	PodIP strfmt.IPv4 `json:"podIP,omitempty"`
}

// Validate validates this check all results hosts items0
func (m *CheckAllResultsHostsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePodIP(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CheckAllResultsHostsItems0) validateHostIP(formats strfmt.Registry) error {

	if swag.IsZero(m.HostIP) { // not required
		return nil
	}

	if err := validate.FormatOf("hostIP", "body", "ipv4", m.HostIP.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CheckAllResultsHostsItems0) validatePodIP(formats strfmt.Registry) error {

	if swag.IsZero(m.PodIP) { // not required
		return nil
	}

	if err := validate.FormatOf("podIP", "body", "ipv4", m.PodIP.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CheckAllResultsHostsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CheckAllResultsHostsItems0) UnmarshalBinary(b []byte) error {
	var res CheckAllResultsHostsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
