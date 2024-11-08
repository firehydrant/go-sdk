// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type GetTeamSignalRuleRequest struct {
	TeamID string `pathParam:"style=simple,explode=false,name=team_id"`
	ID     string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetTeamSignalRuleRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *GetTeamSignalRuleRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetTeamSignalRuleResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *GetTeamSignalRuleResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
