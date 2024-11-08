// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type GetSignalsWebhookTargetRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetSignalsWebhookTargetRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetSignalsWebhookTargetResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *GetSignalsWebhookTargetResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
