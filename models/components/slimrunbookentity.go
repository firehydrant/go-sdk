// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"firehydrant/internal/utils"
	"time"
)

type SlimRunbookEntity struct {
	ID             *string          `json:"id,omitempty"`
	Name           *string          `json:"name,omitempty"`
	Summary        *string          `json:"summary,omitempty"`
	Description    *string          `json:"description,omitempty"`
	Type           *string          `json:"type,omitempty"`
	CreatedAt      *time.Time       `json:"created_at,omitempty"`
	UpdatedAt      *time.Time       `json:"updated_at,omitempty"`
	AttachmentRule *RulesRuleEntity `json:"attachment_rule,omitempty"`
	Owner          *TeamEntityLite  `json:"owner,omitempty"`
	// categories the runbook applies to
	Categories *string `json:"categories,omitempty"`
}

func (s SlimRunbookEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SlimRunbookEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SlimRunbookEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SlimRunbookEntity) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *SlimRunbookEntity) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *SlimRunbookEntity) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *SlimRunbookEntity) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *SlimRunbookEntity) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *SlimRunbookEntity) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *SlimRunbookEntity) GetAttachmentRule() *RulesRuleEntity {
	if o == nil {
		return nil
	}
	return o.AttachmentRule
}

func (o *SlimRunbookEntity) GetOwner() *TeamEntityLite {
	if o == nil {
		return nil
	}
	return o.Owner
}

func (o *SlimRunbookEntity) GetCategories() *string {
	if o == nil {
		return nil
	}
	return o.Categories
}
