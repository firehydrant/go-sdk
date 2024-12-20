// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type ListPrioritiesRequest struct {
	Page    *int `queryParam:"style=form,explode=true,name=page"`
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
}

func (o *ListPrioritiesRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListPrioritiesRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type ListPrioritiesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Lists priorities
	PriorityEntity *components.PriorityEntity
}

func (o *ListPrioritiesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListPrioritiesResponse) GetPriorityEntity() *components.PriorityEntity {
	if o == nil {
		return nil
	}
	return o.PriorityEntity
}
