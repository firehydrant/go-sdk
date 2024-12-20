// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type GetIncidentRelationshipsRequest struct {
	IncidentID string `pathParam:"style=simple,explode=false,name=incident_id"`
}

func (o *GetIncidentRelationshipsRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

type GetIncidentRelationshipsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// List any parent/child relationships for an incident
	IncidentsRelationshipsEntity *components.IncidentsRelationshipsEntity
}

func (o *GetIncidentRelationshipsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetIncidentRelationshipsResponse) GetIncidentsRelationshipsEntity() *components.IncidentsRelationshipsEntity {
	if o == nil {
		return nil
	}
	return o.IncidentsRelationshipsEntity
}
