// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"firehydrant/internal/utils"
	"time"
)

// ChecklistTemplateEntity model
type ChecklistTemplateEntity struct {
	ID          *string                `json:"id,omitempty"`
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	CreatedAt   *string                `json:"created_at,omitempty"`
	UpdatedAt   *time.Time             `json:"updated_at,omitempty"`
	Checks      []ChecklistCheckEntity `json:"checks,omitempty"`
	Owner       *TeamEntityLite        `json:"owner,omitempty"`
	// List of services that use this checklist
	ConnectedServices []ServiceEntityChecklist `json:"connected_services,omitempty"`
}

func (c ChecklistTemplateEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ChecklistTemplateEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ChecklistTemplateEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ChecklistTemplateEntity) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ChecklistTemplateEntity) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *ChecklistTemplateEntity) GetCreatedAt() *string {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *ChecklistTemplateEntity) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *ChecklistTemplateEntity) GetChecks() []ChecklistCheckEntity {
	if o == nil {
		return nil
	}
	return o.Checks
}

func (o *ChecklistTemplateEntity) GetOwner() *TeamEntityLite {
	if o == nil {
		return nil
	}
	return o.Owner
}

func (o *ChecklistTemplateEntity) GetConnectedServices() []ServiceEntityChecklist {
	if o == nil {
		return nil
	}
	return o.ConnectedServices
}
