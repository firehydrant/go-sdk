// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type ListChangesRequest struct {
	Page    *int `queryParam:"style=form,explode=true,name=page"`
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
	// Filter changes by summary
	Query *string `queryParam:"style=form,explode=true,name=query"`
}

func (o *ListChangesRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListChangesRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *ListChangesRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

type ListChangesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *ListChangesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
