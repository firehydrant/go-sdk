// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"firehydrant/internal/utils"
	"time"
)

type TeamEntityLite struct {
	ID             *string       `json:"id,omitempty"`
	Name           *string       `json:"name,omitempty"`
	Description    *string       `json:"description,omitempty"`
	Slug           *string       `json:"slug,omitempty"`
	CreatedAt      *time.Time    `json:"created_at,omitempty"`
	UpdatedAt      *time.Time    `json:"updated_at,omitempty"`
	SignalsIcalURL *string       `json:"signals_ical_url,omitempty"`
	CreatedBy      *AuthorEntity `json:"created_by,omitempty"`
	InSupportHours *bool         `json:"in_support_hours,omitempty"`
}

func (t TeamEntityLite) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TeamEntityLite) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TeamEntityLite) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TeamEntityLite) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TeamEntityLite) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *TeamEntityLite) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *TeamEntityLite) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *TeamEntityLite) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *TeamEntityLite) GetSignalsIcalURL() *string {
	if o == nil {
		return nil
	}
	return o.SignalsIcalURL
}

func (o *TeamEntityLite) GetCreatedBy() *AuthorEntity {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}

func (o *TeamEntityLite) GetInSupportHours() *bool {
	if o == nil {
		return nil
	}
	return o.InSupportHours
}
