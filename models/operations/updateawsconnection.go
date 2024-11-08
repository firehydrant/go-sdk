// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type UpdateAwsConnectionRequest struct {
	ID                                  string                                         `pathParam:"style=simple,explode=false,name=id"`
	PatchV1IntegrationsAwsConnectionsID components.PatchV1IntegrationsAwsConnectionsID `request:"mediaType=application/json"`
}

func (o *UpdateAwsConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateAwsConnectionRequest) GetPatchV1IntegrationsAwsConnectionsID() components.PatchV1IntegrationsAwsConnectionsID {
	if o == nil {
		return components.PatchV1IntegrationsAwsConnectionsID{}
	}
	return o.PatchV1IntegrationsAwsConnectionsID
}

type UpdateAwsConnectionResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Update the AWS connection with the provided data.
	IntegrationsAwsConnectionEntity *components.IntegrationsAwsConnectionEntity
}

func (o *UpdateAwsConnectionResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateAwsConnectionResponse) GetIntegrationsAwsConnectionEntity() *components.IntegrationsAwsConnectionEntity {
	if o == nil {
		return nil
	}
	return o.IntegrationsAwsConnectionEntity
}
