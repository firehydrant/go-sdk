// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type CreatePriorityResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Create a new priority
	PriorityEntity *components.PriorityEntity
}

func (o *CreatePriorityResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreatePriorityResponse) GetPriorityEntity() *components.PriorityEntity {
	if o == nil {
		return nil
	}
	return o.PriorityEntity
}
