// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type GetIncidentChannelRequest struct {
	IncidentID string `pathParam:"style=simple,explode=false,name=incident_id"`
}

func (o *GetIncidentChannelRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

type GetIncidentChannelResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Gives chat channel information for the specified incident
	IncidentsChannelEntity *components.IncidentsChannelEntity
}

func (o *GetIncidentChannelResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetIncidentChannelResponse) GetIncidentsChannelEntity() *components.IncidentsChannelEntity {
	if o == nil {
		return nil
	}
	return o.IncidentsChannelEntity
}
