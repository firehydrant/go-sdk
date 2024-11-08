// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type ArchiveIncidentRequest struct {
	IncidentID string `pathParam:"style=simple,explode=false,name=incident_id"`
}

func (o *ArchiveIncidentRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

type ArchiveIncidentResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Archives an incident which will hide it from lists and metrics
	IncidentEntity *components.IncidentEntity
}

func (o *ArchiveIncidentResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ArchiveIncidentResponse) GetIncidentEntity() *components.IncidentEntity {
	if o == nil {
		return nil
	}
	return o.IncidentEntity
}
