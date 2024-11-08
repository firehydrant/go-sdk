// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type GetAwsCloudTrailBatchRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetAwsCloudTrailBatchRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetAwsCloudTrailBatchResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Retrieve a single CloudTrail batch.
	IntegrationsAwsCloudtrailBatchEntity *components.IntegrationsAwsCloudtrailBatchEntity
}

func (o *GetAwsCloudTrailBatchResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetAwsCloudTrailBatchResponse) GetIntegrationsAwsCloudtrailBatchEntity() *components.IntegrationsAwsCloudtrailBatchEntity {
	if o == nil {
		return nil
	}
	return o.IntegrationsAwsCloudtrailBatchEntity
}
