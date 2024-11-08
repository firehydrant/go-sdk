// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type UpdateServiceRequest struct {
	ServiceID                string                              `pathParam:"style=simple,explode=false,name=service_id"`
	PatchV1ServicesServiceID components.PatchV1ServicesServiceID `request:"mediaType=application/json"`
}

func (o *UpdateServiceRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *UpdateServiceRequest) GetPatchV1ServicesServiceID() components.PatchV1ServicesServiceID {
	if o == nil {
		return components.PatchV1ServicesServiceID{}
	}
	return o.PatchV1ServicesServiceID
}

type UpdateServiceResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Update a services attributes, you may also add or remove functionalities from the service as well.
	// Note: You may not remove or add individual label key/value pairs. You must include the entire object to override label values.
	//
	ServiceEntity *components.ServiceEntity
}

func (o *UpdateServiceResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateServiceResponse) GetServiceEntity() *components.ServiceEntity {
	if o == nil {
		return nil
	}
	return o.ServiceEntity
}
