// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type GetTicketingProjectRequest struct {
	TicketingProjectID string `pathParam:"style=simple,explode=false,name=ticketing_project_id"`
}

func (o *GetTicketingProjectRequest) GetTicketingProjectID() string {
	if o == nil {
		return ""
	}
	return o.TicketingProjectID
}

type GetTicketingProjectResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Retrieve a single ticketing project by ID
	TicketingProjectsProjectListItemEntity *components.TicketingProjectsProjectListItemEntity
}

func (o *GetTicketingProjectResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetTicketingProjectResponse) GetTicketingProjectsProjectListItemEntity() *components.TicketingProjectsProjectListItemEntity {
	if o == nil {
		return nil
	}
	return o.TicketingProjectsProjectListItemEntity
}