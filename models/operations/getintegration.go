// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type GetIntegrationRequest struct {
	// Integration UUID
	IntegrationID string `pathParam:"style=simple,explode=false,name=integration_id"`
}

func (o *GetIntegrationRequest) GetIntegrationID() string {
	if o == nil {
		return ""
	}
	return o.IntegrationID
}

type GetIntegrationResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Retrieve a single integration
	IntegrationsIntegrationEntity *components.IntegrationsIntegrationEntity
}

func (o *GetIntegrationResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetIntegrationResponse) GetIntegrationsIntegrationEntity() *components.IntegrationsIntegrationEntity {
	if o == nil {
		return nil
	}
	return o.IntegrationsIntegrationEntity
}
