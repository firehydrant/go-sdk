// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type CreateMeasurementDefinitionRequestBody struct {
	Name                string  `json:"name"`
	Slug                *string `json:"slug,omitempty"`
	Description         *string `json:"description,omitempty"`
	StartsAtMilestoneID string  `json:"starts_at_milestone_id"`
	EndsAtMilestoneID   string  `json:"ends_at_milestone_id"`
}

func (o *CreateMeasurementDefinitionRequestBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateMeasurementDefinitionRequestBody) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *CreateMeasurementDefinitionRequestBody) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateMeasurementDefinitionRequestBody) GetStartsAtMilestoneID() string {
	if o == nil {
		return ""
	}
	return o.StartsAtMilestoneID
}

func (o *CreateMeasurementDefinitionRequestBody) GetEndsAtMilestoneID() string {
	if o == nil {
		return ""
	}
	return o.EndsAtMilestoneID
}

type CreateMeasurementDefinitionResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *CreateMeasurementDefinitionResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
