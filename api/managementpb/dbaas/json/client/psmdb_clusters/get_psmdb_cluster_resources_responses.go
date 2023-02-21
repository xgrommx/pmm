// Code generated by go-swagger; DO NOT EDIT.

package psmdb_clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetPSMDBClusterResourcesReader is a Reader for the GetPSMDBClusterResources structure.
type GetPSMDBClusterResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPSMDBClusterResourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPSMDBClusterResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetPSMDBClusterResourcesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPSMDBClusterResourcesOK creates a GetPSMDBClusterResourcesOK with default headers values
func NewGetPSMDBClusterResourcesOK() *GetPSMDBClusterResourcesOK {
	return &GetPSMDBClusterResourcesOK{}
}

/*
GetPSMDBClusterResourcesOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetPSMDBClusterResourcesOK struct {
	Payload *GetPSMDBClusterResourcesOKBody
}

func (o *GetPSMDBClusterResourcesOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/PSMDBCluster/Resources/Get][%d] getPsmdbClusterResourcesOk  %+v", 200, o.Payload)
}

func (o *GetPSMDBClusterResourcesOK) GetPayload() *GetPSMDBClusterResourcesOKBody {
	return o.Payload
}

func (o *GetPSMDBClusterResourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(GetPSMDBClusterResourcesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPSMDBClusterResourcesDefault creates a GetPSMDBClusterResourcesDefault with default headers values
func NewGetPSMDBClusterResourcesDefault(code int) *GetPSMDBClusterResourcesDefault {
	return &GetPSMDBClusterResourcesDefault{
		_statusCode: code,
	}
}

/*
GetPSMDBClusterResourcesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetPSMDBClusterResourcesDefault struct {
	_statusCode int

	Payload *GetPSMDBClusterResourcesDefaultBody
}

// Code gets the status code for the get PSMDB cluster resources default response
func (o *GetPSMDBClusterResourcesDefault) Code() int {
	return o._statusCode
}

func (o *GetPSMDBClusterResourcesDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/PSMDBCluster/Resources/Get][%d] GetPSMDBClusterResources default  %+v", o._statusCode, o.Payload)
}

func (o *GetPSMDBClusterResourcesDefault) GetPayload() *GetPSMDBClusterResourcesDefaultBody {
	return o.Payload
}

func (o *GetPSMDBClusterResourcesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(GetPSMDBClusterResourcesDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetPSMDBClusterResourcesBody get PSMDB cluster resources body
swagger:model GetPSMDBClusterResourcesBody
*/
type GetPSMDBClusterResourcesBody struct {
	// params
	Params *GetPSMDBClusterResourcesParamsBodyParams `json:"params,omitempty"`
}

// Validate validates this get PSMDB cluster resources body
func (o *GetPSMDBClusterResourcesBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPSMDBClusterResourcesBody) validateParams(formats strfmt.Registry) error {
	if swag.IsZero(o.Params) { // not required
		return nil
	}

	if o.Params != nil {
		if err := o.Params.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get PSMDB cluster resources body based on the context it is used
func (o *GetPSMDBClusterResourcesBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPSMDBClusterResourcesBody) contextValidateParams(ctx context.Context, formats strfmt.Registry) error {
	if o.Params != nil {
		if err := o.Params.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPSMDBClusterResourcesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPSMDBClusterResourcesBody) UnmarshalBinary(b []byte) error {
	var res GetPSMDBClusterResourcesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetPSMDBClusterResourcesDefaultBody get PSMDB cluster resources default body
swagger:model GetPSMDBClusterResourcesDefaultBody
*/
type GetPSMDBClusterResourcesDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*GetPSMDBClusterResourcesDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this get PSMDB cluster resources default body
func (o *GetPSMDBClusterResourcesDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPSMDBClusterResourcesDefaultBody) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GetPSMDBClusterResources default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GetPSMDBClusterResources default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get PSMDB cluster resources default body based on the context it is used
func (o *GetPSMDBClusterResourcesDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPSMDBClusterResourcesDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GetPSMDBClusterResources default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GetPSMDBClusterResources default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPSMDBClusterResourcesDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPSMDBClusterResourcesDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetPSMDBClusterResourcesDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetPSMDBClusterResourcesDefaultBodyDetailsItems0 get PSMDB cluster resources default body details items0
swagger:model GetPSMDBClusterResourcesDefaultBodyDetailsItems0
*/
type GetPSMDBClusterResourcesDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this get PSMDB cluster resources default body details items0
func (o *GetPSMDBClusterResourcesDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get PSMDB cluster resources default body details items0 based on context it is used
func (o *GetPSMDBClusterResourcesDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPSMDBClusterResourcesDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPSMDBClusterResourcesDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res GetPSMDBClusterResourcesDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetPSMDBClusterResourcesOKBody get PSMDB cluster resources OK body
swagger:model GetPSMDBClusterResourcesOKBody
*/
type GetPSMDBClusterResourcesOKBody struct {
	// expected
	Expected *GetPSMDBClusterResourcesOKBodyExpected `json:"expected,omitempty"`
}

// Validate validates this get PSMDB cluster resources OK body
func (o *GetPSMDBClusterResourcesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateExpected(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPSMDBClusterResourcesOKBody) validateExpected(formats strfmt.Registry) error {
	if swag.IsZero(o.Expected) { // not required
		return nil
	}

	if o.Expected != nil {
		if err := o.Expected.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getPsmdbClusterResourcesOk" + "." + "expected")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getPsmdbClusterResourcesOk" + "." + "expected")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get PSMDB cluster resources OK body based on the context it is used
func (o *GetPSMDBClusterResourcesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateExpected(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPSMDBClusterResourcesOKBody) contextValidateExpected(ctx context.Context, formats strfmt.Registry) error {
	if o.Expected != nil {
		if err := o.Expected.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getPsmdbClusterResourcesOk" + "." + "expected")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getPsmdbClusterResourcesOk" + "." + "expected")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPSMDBClusterResourcesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPSMDBClusterResourcesOKBody) UnmarshalBinary(b []byte) error {
	var res GetPSMDBClusterResourcesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetPSMDBClusterResourcesOKBodyExpected Resources contains Kubernetes cluster resources.
swagger:model GetPSMDBClusterResourcesOKBodyExpected
*/
type GetPSMDBClusterResourcesOKBodyExpected struct {
	// Memory in bytes.
	MemoryBytes string `json:"memory_bytes,omitempty"`

	// CPU in millicpus. For example 0.1 of CPU is equivalent to 100 millicpus.
	// See https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#meaning-of-cpu.
	CPUm string `json:"cpu_m,omitempty"`

	// Disk size in bytes.
	DiskSize string `json:"disk_size,omitempty"`
}

// Validate validates this get PSMDB cluster resources OK body expected
func (o *GetPSMDBClusterResourcesOKBodyExpected) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get PSMDB cluster resources OK body expected based on context it is used
func (o *GetPSMDBClusterResourcesOKBodyExpected) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPSMDBClusterResourcesOKBodyExpected) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPSMDBClusterResourcesOKBodyExpected) UnmarshalBinary(b []byte) error {
	var res GetPSMDBClusterResourcesOKBodyExpected
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetPSMDBClusterResourcesParamsBodyParams PSMDBClusterParams represents PSMDB cluster parameters that can be updated.
swagger:model GetPSMDBClusterResourcesParamsBodyParams
*/
type GetPSMDBClusterResourcesParamsBodyParams struct {
	// Cluster size.
	ClusterSize int32 `json:"cluster_size,omitempty"`

	// Docker image used for PSMDB.
	Image string `json:"image,omitempty"`

	// backup
	Backup *GetPSMDBClusterResourcesParamsBodyParamsBackup `json:"backup,omitempty"`

	// replicaset
	Replicaset *GetPSMDBClusterResourcesParamsBodyParamsReplicaset `json:"replicaset,omitempty"`

	// restore
	Restore *GetPSMDBClusterResourcesParamsBodyParamsRestore `json:"restore,omitempty"`
}

// Validate validates this get PSMDB cluster resources params body params
func (o *GetPSMDBClusterResourcesParamsBodyParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBackup(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateReplicaset(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRestore(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPSMDBClusterResourcesParamsBodyParams) validateBackup(formats strfmt.Registry) error {
	if swag.IsZero(o.Backup) { // not required
		return nil
	}

	if o.Backup != nil {
		if err := o.Backup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "backup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "backup")
			}
			return err
		}
	}

	return nil
}

func (o *GetPSMDBClusterResourcesParamsBodyParams) validateReplicaset(formats strfmt.Registry) error {
	if swag.IsZero(o.Replicaset) { // not required
		return nil
	}

	if o.Replicaset != nil {
		if err := o.Replicaset.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "replicaset")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "replicaset")
			}
			return err
		}
	}

	return nil
}

func (o *GetPSMDBClusterResourcesParamsBodyParams) validateRestore(formats strfmt.Registry) error {
	if swag.IsZero(o.Restore) { // not required
		return nil
	}

	if o.Restore != nil {
		if err := o.Restore.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "restore")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "restore")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get PSMDB cluster resources params body params based on the context it is used
func (o *GetPSMDBClusterResourcesParamsBodyParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateBackup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateReplicaset(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRestore(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPSMDBClusterResourcesParamsBodyParams) contextValidateBackup(ctx context.Context, formats strfmt.Registry) error {
	if o.Backup != nil {
		if err := o.Backup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "backup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "backup")
			}
			return err
		}
	}

	return nil
}

func (o *GetPSMDBClusterResourcesParamsBodyParams) contextValidateReplicaset(ctx context.Context, formats strfmt.Registry) error {
	if o.Replicaset != nil {
		if err := o.Replicaset.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "replicaset")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "replicaset")
			}
			return err
		}
	}

	return nil
}

func (o *GetPSMDBClusterResourcesParamsBodyParams) contextValidateRestore(ctx context.Context, formats strfmt.Registry) error {
	if o.Restore != nil {
		if err := o.Restore.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "restore")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "restore")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPSMDBClusterResourcesParamsBodyParams) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPSMDBClusterResourcesParamsBodyParams) UnmarshalBinary(b []byte) error {
	var res GetPSMDBClusterResourcesParamsBodyParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetPSMDBClusterResourcesParamsBodyParamsBackup Backup configuration for a database cluster
swagger:model GetPSMDBClusterResourcesParamsBodyParamsBackup
*/
type GetPSMDBClusterResourcesParamsBodyParamsBackup struct {
	// Backup Location id of stored backup location in PMM.
	LocationID string `json:"location_id,omitempty"`

	// Keep copies represents how many copies should retain.
	KeepCopies int32 `json:"keep_copies,omitempty"`

	// Cron expression represents cron expression
	CronExpression string `json:"cron_expression,omitempty"`

	// Service acccount used for backups
	ServiceAccount string `json:"service_account,omitempty"`
}

// Validate validates this get PSMDB cluster resources params body params backup
func (o *GetPSMDBClusterResourcesParamsBodyParamsBackup) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get PSMDB cluster resources params body params backup based on context it is used
func (o *GetPSMDBClusterResourcesParamsBodyParamsBackup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPSMDBClusterResourcesParamsBodyParamsBackup) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPSMDBClusterResourcesParamsBodyParamsBackup) UnmarshalBinary(b []byte) error {
	var res GetPSMDBClusterResourcesParamsBodyParamsBackup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetPSMDBClusterResourcesParamsBodyParamsReplicaset ReplicaSet container parameters.
// TODO Do not use inner messages in all public APIs (for consistency).
swagger:model GetPSMDBClusterResourcesParamsBodyParamsReplicaset
*/
type GetPSMDBClusterResourcesParamsBodyParamsReplicaset struct {
	// Disk size in bytes.
	DiskSize string `json:"disk_size,omitempty"`

	// Configuration for PSMDB cluster
	Configuration string `json:"configuration,omitempty"`

	// Storage Class for PSMDB cluster.
	StorageClass string `json:"storage_class,omitempty"`

	// compute resources
	ComputeResources *GetPSMDBClusterResourcesParamsBodyParamsReplicasetComputeResources `json:"compute_resources,omitempty"`
}

// Validate validates this get PSMDB cluster resources params body params replicaset
func (o *GetPSMDBClusterResourcesParamsBodyParamsReplicaset) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateComputeResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPSMDBClusterResourcesParamsBodyParamsReplicaset) validateComputeResources(formats strfmt.Registry) error {
	if swag.IsZero(o.ComputeResources) { // not required
		return nil
	}

	if o.ComputeResources != nil {
		if err := o.ComputeResources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "replicaset" + "." + "compute_resources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "replicaset" + "." + "compute_resources")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get PSMDB cluster resources params body params replicaset based on the context it is used
func (o *GetPSMDBClusterResourcesParamsBodyParamsReplicaset) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateComputeResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPSMDBClusterResourcesParamsBodyParamsReplicaset) contextValidateComputeResources(ctx context.Context, formats strfmt.Registry) error {
	if o.ComputeResources != nil {
		if err := o.ComputeResources.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "replicaset" + "." + "compute_resources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "params" + "." + "replicaset" + "." + "compute_resources")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPSMDBClusterResourcesParamsBodyParamsReplicaset) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPSMDBClusterResourcesParamsBodyParamsReplicaset) UnmarshalBinary(b []byte) error {
	var res GetPSMDBClusterResourcesParamsBodyParamsReplicaset
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetPSMDBClusterResourcesParamsBodyParamsReplicasetComputeResources ComputeResources represents container computer resources requests or limits.
swagger:model GetPSMDBClusterResourcesParamsBodyParamsReplicasetComputeResources
*/
type GetPSMDBClusterResourcesParamsBodyParamsReplicasetComputeResources struct {
	// CPUs in milliCPUs; 1000m = 1 vCPU.
	CPUm int32 `json:"cpu_m,omitempty"`

	// Memory in bytes.
	MemoryBytes string `json:"memory_bytes,omitempty"`
}

// Validate validates this get PSMDB cluster resources params body params replicaset compute resources
func (o *GetPSMDBClusterResourcesParamsBodyParamsReplicasetComputeResources) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get PSMDB cluster resources params body params replicaset compute resources based on context it is used
func (o *GetPSMDBClusterResourcesParamsBodyParamsReplicasetComputeResources) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPSMDBClusterResourcesParamsBodyParamsReplicasetComputeResources) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPSMDBClusterResourcesParamsBodyParamsReplicasetComputeResources) UnmarshalBinary(b []byte) error {
	var res GetPSMDBClusterResourcesParamsBodyParamsReplicasetComputeResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetPSMDBClusterResourcesParamsBodyParamsRestore Restore represents restoration payload to restore a database cluster from backup
swagger:model GetPSMDBClusterResourcesParamsBodyParamsRestore
*/
type GetPSMDBClusterResourcesParamsBodyParamsRestore struct {
	// Backup location in PMM.
	LocationID string `json:"location_id,omitempty"`

	// Destination filename.
	Destination string `json:"destination,omitempty"`

	// K8s Secrets name.
	SecretsName string `json:"secrets_name,omitempty"`
}

// Validate validates this get PSMDB cluster resources params body params restore
func (o *GetPSMDBClusterResourcesParamsBodyParamsRestore) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get PSMDB cluster resources params body params restore based on context it is used
func (o *GetPSMDBClusterResourcesParamsBodyParamsRestore) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPSMDBClusterResourcesParamsBodyParamsRestore) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPSMDBClusterResourcesParamsBodyParamsRestore) UnmarshalBinary(b []byte) error {
	var res GetPSMDBClusterResourcesParamsBodyParamsRestore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
