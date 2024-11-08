// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type ListStatusPageSubscribersRequest struct {
	NuncConnectionID string `pathParam:"style=simple,explode=false,name=nunc_connection_id"`
}

func (o *ListStatusPageSubscribersRequest) GetNuncConnectionID() string {
	if o == nil {
		return ""
	}
	return o.NuncConnectionID
}

type ListStatusPageSubscribersResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Retrieves the list of subscribers for a status page.
	NuncEmailSubscribersEntity *components.NuncEmailSubscribersEntity
}

func (o *ListStatusPageSubscribersResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListStatusPageSubscribersResponse) GetNuncEmailSubscribersEntity() *components.NuncEmailSubscribersEntity {
	if o == nil {
		return nil
	}
	return o.NuncEmailSubscribersEntity
}
