// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type DeleteIncidentStatusPageRequest struct {
	StatusPageID string `pathParam:"style=simple,explode=false,name=status_page_id"`
	IncidentID   string `pathParam:"style=simple,explode=false,name=incident_id"`
}

func (o *DeleteIncidentStatusPageRequest) GetStatusPageID() string {
	if o == nil {
		return ""
	}
	return o.StatusPageID
}

func (o *DeleteIncidentStatusPageRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

type DeleteIncidentStatusPageResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *DeleteIncidentStatusPageResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
