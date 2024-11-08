// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type UpdatePriorityRequest struct {
	PrioritySlug                  string                                   `pathParam:"style=simple,explode=false,name=priority_slug"`
	PatchV1PrioritiesPrioritySlug components.PatchV1PrioritiesPrioritySlug `request:"mediaType=application/json"`
}

func (o *UpdatePriorityRequest) GetPrioritySlug() string {
	if o == nil {
		return ""
	}
	return o.PrioritySlug
}

func (o *UpdatePriorityRequest) GetPatchV1PrioritiesPrioritySlug() components.PatchV1PrioritiesPrioritySlug {
	if o == nil {
		return components.PatchV1PrioritiesPrioritySlug{}
	}
	return o.PatchV1PrioritiesPrioritySlug
}

type UpdatePriorityResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Update a specific priority
	PriorityEntity *components.PriorityEntity
}

func (o *UpdatePriorityResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdatePriorityResponse) GetPriorityEntity() *components.PriorityEntity {
	if o == nil {
		return nil
	}
	return o.PriorityEntity
}
