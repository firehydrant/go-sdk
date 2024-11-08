// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type CreateIncidentNoteRequest struct {
	IncidentID                     string                                    `pathParam:"style=simple,explode=false,name=incident_id"`
	PostV1IncidentsIncidentIDNotes components.PostV1IncidentsIncidentIDNotes `request:"mediaType=application/json"`
}

func (o *CreateIncidentNoteRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *CreateIncidentNoteRequest) GetPostV1IncidentsIncidentIDNotes() components.PostV1IncidentsIncidentIDNotes {
	if o == nil {
		return components.PostV1IncidentsIncidentIDNotes{}
	}
	return o.PostV1IncidentsIncidentIDNotes
}

type CreateIncidentNoteResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Create a new note on for an incident. The visibility field on a note determines where it gets posted.
	EventNoteEntity *components.EventNoteEntity
}

func (o *CreateIncidentNoteResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateIncidentNoteResponse) GetEventNoteEntity() *components.EventNoteEntity {
	if o == nil {
		return nil
	}
	return o.EventNoteEntity
}
