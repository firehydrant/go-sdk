// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type UpdateRetrospectiveReasonRequest struct {
	ReportID                                         string                                                      `pathParam:"style=simple,explode=false,name=report_id"`
	ReasonID                                         string                                                      `pathParam:"style=simple,explode=false,name=reason_id"`
	PatchV1PostMortemsReportsReportIDReasonsReasonID components.PatchV1PostMortemsReportsReportIDReasonsReasonID `request:"mediaType=application/json"`
}

func (o *UpdateRetrospectiveReasonRequest) GetReportID() string {
	if o == nil {
		return ""
	}
	return o.ReportID
}

func (o *UpdateRetrospectiveReasonRequest) GetReasonID() string {
	if o == nil {
		return ""
	}
	return o.ReasonID
}

func (o *UpdateRetrospectiveReasonRequest) GetPatchV1PostMortemsReportsReportIDReasonsReasonID() components.PatchV1PostMortemsReportsReportIDReasonsReasonID {
	if o == nil {
		return components.PatchV1PostMortemsReportsReportIDReasonsReasonID{}
	}
	return o.PatchV1PostMortemsReportsReportIDReasonsReasonID
}

type UpdateRetrospectiveReasonResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Update a contributing factor
	PostMortemsReasonEntity *components.PostMortemsReasonEntity
}

func (o *UpdateRetrospectiveReasonResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateRetrospectiveReasonResponse) GetPostMortemsReasonEntity() *components.PostMortemsReasonEntity {
	if o == nil {
		return nil
	}
	return o.PostMortemsReasonEntity
}
