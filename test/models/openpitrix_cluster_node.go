// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixClusterNode openpitrix cluster node
// swagger:model openpitrixClusterNode
type OpenpitrixClusterNode struct {

	// auto backup
	AutoBackup bool `json:"auto_backup,omitempty"`

	// cluster common
	ClusterCommon *OpenpitrixClusterCommon `json:"cluster_common,omitempty"`

	// cluster id
	ClusterID string `json:"cluster_id,omitempty"`

	// cluster role
	ClusterRole *OpenpitrixClusterRole `json:"cluster_role,omitempty"`

	// create time
	CreateTime strfmt.DateTime `json:"create_time,omitempty"`

	// custom metadata
	CustomMetadata string `json:"custom_metadata,omitempty"`

	// device
	Device string `json:"device,omitempty"`

	// eip
	Eip string `json:"eip,omitempty"`

	// global server id
	GlobalServerID *ProtobufUint32Value `json:"global_server_id,omitempty"`

	// group id
	GroupID *ProtobufUint32Value `json:"group_id,omitempty"`

	// health status
	HealthStatus string `json:"health_status,omitempty"`

	// instance id
	InstanceID string `json:"instance_id,omitempty"`

	// is backup
	IsBackup bool `json:"is_backup,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// node id
	NodeID string `json:"node_id,omitempty"`

	// owner
	Owner string `json:"owner,omitempty"`

	// private ip
	PrivateIP string `json:"private_ip,omitempty"`

	// pub key
	PubKey string `json:"pub_key,omitempty"`

	// role
	Role string `json:"role,omitempty"`

	// server id
	ServerID *ProtobufUint32Value `json:"server_id,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// status time
	StatusTime strfmt.DateTime `json:"status_time,omitempty"`

	// subnet id
	SubnetID string `json:"subnet_id,omitempty"`

	// transition status
	TransitionStatus string `json:"transition_status,omitempty"`

	// volume id
	VolumeID string `json:"volume_id,omitempty"`
}

// Validate validates this openpitrix cluster node
func (m *OpenpitrixClusterNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterCommon(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateClusterRole(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGlobalServerID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGroupID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateServerID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixClusterNode) validateClusterCommon(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterCommon) { // not required
		return nil
	}

	if m.ClusterCommon != nil {

		if err := m.ClusterCommon.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster_common")
			}
			return err
		}
	}

	return nil
}

func (m *OpenpitrixClusterNode) validateClusterRole(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterRole) { // not required
		return nil
	}

	if m.ClusterRole != nil {

		if err := m.ClusterRole.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster_role")
			}
			return err
		}
	}

	return nil
}

func (m *OpenpitrixClusterNode) validateGlobalServerID(formats strfmt.Registry) error {

	if swag.IsZero(m.GlobalServerID) { // not required
		return nil
	}

	if m.GlobalServerID != nil {

		if err := m.GlobalServerID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("global_server_id")
			}
			return err
		}
	}

	return nil
}

func (m *OpenpitrixClusterNode) validateGroupID(formats strfmt.Registry) error {

	if swag.IsZero(m.GroupID) { // not required
		return nil
	}

	if m.GroupID != nil {

		if err := m.GroupID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group_id")
			}
			return err
		}
	}

	return nil
}

func (m *OpenpitrixClusterNode) validateServerID(formats strfmt.Registry) error {

	if swag.IsZero(m.ServerID) { // not required
		return nil
	}

	if m.ServerID != nil {

		if err := m.ServerID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("server_id")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixClusterNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixClusterNode) UnmarshalBinary(b []byte) error {
	var res OpenpitrixClusterNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
