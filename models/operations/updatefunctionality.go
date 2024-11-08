// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type UpdateFunctionalityRequest struct {
	FunctionalityID                       string                                           `pathParam:"style=simple,explode=false,name=functionality_id"`
	PatchV1FunctionalitiesFunctionalityID components.PatchV1FunctionalitiesFunctionalityID `request:"mediaType=application/json"`
}

func (o *UpdateFunctionalityRequest) GetFunctionalityID() string {
	if o == nil {
		return ""
	}
	return o.FunctionalityID
}

func (o *UpdateFunctionalityRequest) GetPatchV1FunctionalitiesFunctionalityID() components.PatchV1FunctionalitiesFunctionalityID {
	if o == nil {
		return components.PatchV1FunctionalitiesFunctionalityID{}
	}
	return o.PatchV1FunctionalitiesFunctionalityID
}

type UpdateFunctionalityResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Update a functionalities attributes
	FunctionalityEntity *components.FunctionalityEntity
}

func (o *UpdateFunctionalityResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateFunctionalityResponse) GetFunctionalityEntity() *components.FunctionalityEntity {
	if o == nil {
		return nil
	}
	return o.FunctionalityEntity
}
