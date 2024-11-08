// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type GetTeamRequest struct {
	TeamID string `pathParam:"style=simple,explode=false,name=team_id"`
	// Boolean to determine whether to return a slimified version of the teams object
	Lite *bool `queryParam:"style=form,explode=true,name=lite"`
}

func (o *GetTeamRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *GetTeamRequest) GetLite() *bool {
	if o == nil {
		return nil
	}
	return o.Lite
}

type GetTeamResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Retrieve a single team from its ID
	TeamEntity *components.TeamEntity
}

func (o *GetTeamResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetTeamResponse) GetTeamEntity() *components.TeamEntity {
	if o == nil {
		return nil
	}
	return o.TeamEntity
}