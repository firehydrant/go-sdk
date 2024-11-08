// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type UpdateScimUserRequest struct {
	ID                   string                          `pathParam:"style=simple,explode=false,name=id"`
	PatchV1ScimV2UsersID components.PatchV1ScimV2UsersID `request:"mediaType=application/scim+json"`
}

func (o *UpdateScimUserRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateScimUserRequest) GetPatchV1ScimV2UsersID() components.PatchV1ScimV2UsersID {
	if o == nil {
		return components.PatchV1ScimV2UsersID{}
	}
	return o.PatchV1ScimV2UsersID
}

type UpdateScimUserResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *UpdateScimUserResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
