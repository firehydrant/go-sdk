// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/internal/utils"
	"firehydrant/models/components"
	"time"
)

type ListRetrospectiveReportsRequest struct {
	Page    *int `queryParam:"style=form,explode=true,name=page"`
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
	// Filter the reports by an incident ID
	IncidentID *string `queryParam:"style=form,explode=true,name=incident_id"`
	// Filter for reports updated after the given ISO8601 timestamp
	UpdatedSince *time.Time `queryParam:"style=form,explode=true,name=updated_since"`
}

func (l ListRetrospectiveReportsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListRetrospectiveReportsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListRetrospectiveReportsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListRetrospectiveReportsRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *ListRetrospectiveReportsRequest) GetIncidentID() *string {
	if o == nil {
		return nil
	}
	return o.IncidentID
}

func (o *ListRetrospectiveReportsRequest) GetUpdatedSince() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedSince
}

type ListRetrospectiveReportsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// List all reports
	PostMortemsPostMortemReportEntityPaginated *components.PostMortemsPostMortemReportEntityPaginated
}

func (o *ListRetrospectiveReportsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListRetrospectiveReportsResponse) GetPostMortemsPostMortemReportEntityPaginated() *components.PostMortemsPostMortemReportEntityPaginated {
	if o == nil {
		return nil
	}
	return o.PostMortemsPostMortemReportEntityPaginated
}