// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type CreateChecklistTemplateResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Creates a checklist template for the organization
	ChecklistTemplateEntity *components.ChecklistTemplateEntity
}

func (o *CreateChecklistTemplateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateChecklistTemplateResponse) GetChecklistTemplateEntity() *components.ChecklistTemplateEntity {
	if o == nil {
		return nil
	}
	return o.ChecklistTemplateEntity
}
