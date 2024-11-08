// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"firehydrant/models/components"
	"fmt"
)

// QueryParamType - The type of the relation to the incident
type QueryParamType string

const (
	QueryParamTypeCaused    QueryParamType = "caused"
	QueryParamTypeFixed     QueryParamType = "fixed"
	QueryParamTypeSuspect   QueryParamType = "suspect"
	QueryParamTypeDismissed QueryParamType = "dismissed"
)

func (e QueryParamType) ToPointer() *QueryParamType {
	return &e
}
func (e *QueryParamType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "caused":
		fallthrough
	case "fixed":
		fallthrough
	case "suspect":
		fallthrough
	case "dismissed":
		*e = QueryParamType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for QueryParamType: %v", v)
	}
}

type ListIncidentRelatedChangesRequest struct {
	Page    *int `queryParam:"style=form,explode=true,name=page"`
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
	// The type of the relation to the incident
	Type       *QueryParamType `queryParam:"style=form,explode=true,name=type"`
	IncidentID string          `pathParam:"style=simple,explode=false,name=incident_id"`
}

func (o *ListIncidentRelatedChangesRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListIncidentRelatedChangesRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *ListIncidentRelatedChangesRequest) GetType() *QueryParamType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ListIncidentRelatedChangesRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

type ListIncidentRelatedChangesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// List related changes that have been attached to an incident
	IncidentsRelatedChangeEventEntityPaginated *components.IncidentsRelatedChangeEventEntityPaginated
}

func (o *ListIncidentRelatedChangesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListIncidentRelatedChangesResponse) GetIncidentsRelatedChangeEventEntityPaginated() *components.IncidentsRelatedChangeEventEntityPaginated {
	if o == nil {
		return nil
	}
	return o.IncidentsRelatedChangeEventEntityPaginated
}