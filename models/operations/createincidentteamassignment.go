// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type CreateIncidentTeamAssignmentRequest struct {
	IncidentID                               string                                              `pathParam:"style=simple,explode=false,name=incident_id"`
	PostV1IncidentsIncidentIDTeamAssignments components.PostV1IncidentsIncidentIDTeamAssignments `request:"mediaType=application/json"`
}

func (o *CreateIncidentTeamAssignmentRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *CreateIncidentTeamAssignmentRequest) GetPostV1IncidentsIncidentIDTeamAssignments() components.PostV1IncidentsIncidentIDTeamAssignments {
	if o == nil {
		return components.PostV1IncidentsIncidentIDTeamAssignments{}
	}
	return o.PostV1IncidentsIncidentIDTeamAssignments
}

type CreateIncidentTeamAssignmentResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *CreateIncidentTeamAssignmentResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}