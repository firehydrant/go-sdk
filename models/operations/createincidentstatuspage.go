// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type CreateIncidentStatusPageRequest struct {
	IncidentID                           string                                          `pathParam:"style=simple,explode=false,name=incident_id"`
	PostV1IncidentsIncidentIDStatusPages components.PostV1IncidentsIncidentIDStatusPages `request:"mediaType=application/json"`
}

func (o *CreateIncidentStatusPageRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *CreateIncidentStatusPageRequest) GetPostV1IncidentsIncidentIDStatusPages() components.PostV1IncidentsIncidentIDStatusPages {
	if o == nil {
		return components.PostV1IncidentsIncidentIDStatusPages{}
	}
	return o.PostV1IncidentsIncidentIDStatusPages
}

type CreateIncidentStatusPageResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Add a status page to an incident.
	IncidentsStatusPageEntity *components.IncidentsStatusPageEntity
}

func (o *CreateIncidentStatusPageResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateIncidentStatusPageResponse) GetIncidentsStatusPageEntity() *components.IncidentsStatusPageEntity {
	if o == nil {
		return nil
	}
	return o.IncidentsStatusPageEntity
}