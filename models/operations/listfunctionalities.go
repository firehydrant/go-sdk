// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type ListFunctionalitiesRequest struct {
	// A query to search functionalities by their name or description
	Query *string `queryParam:"style=form,explode=true,name=query"`
	// A query to search functionalities by their name
	Name *string `queryParam:"style=form,explode=true,name=name"`
	// A query to search services by if they are impacted with active incidents
	Impacted *string `queryParam:"style=form,explode=true,name=impacted"`
	// A comma separated list of label key / values in the format of 'key=value,key2=value2'. To filter change events that have a key (with no specific value), omit the value
	Labels *string `queryParam:"style=form,explode=true,name=labels"`
	// A query to search functionalities by their owning team ID
	Owner *string `queryParam:"style=form,explode=true,name=owner"`
	// Boolean to determine whether to return a slimified version of the functionalities object
	Lite    *bool `queryParam:"style=form,explode=true,name=lite"`
	Page    *int  `queryParam:"style=form,explode=true,name=page"`
	PerPage *int  `queryParam:"style=form,explode=true,name=per_page"`
}

func (o *ListFunctionalitiesRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *ListFunctionalitiesRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ListFunctionalitiesRequest) GetImpacted() *string {
	if o == nil {
		return nil
	}
	return o.Impacted
}

func (o *ListFunctionalitiesRequest) GetLabels() *string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *ListFunctionalitiesRequest) GetOwner() *string {
	if o == nil {
		return nil
	}
	return o.Owner
}

func (o *ListFunctionalitiesRequest) GetLite() *bool {
	if o == nil {
		return nil
	}
	return o.Lite
}

func (o *ListFunctionalitiesRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListFunctionalitiesRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type ListFunctionalitiesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// List all of the functionalities that have been added to the organiation
	FunctionalityEntityPaginated *components.FunctionalityEntityPaginated
}

func (o *ListFunctionalitiesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListFunctionalitiesResponse) GetFunctionalityEntityPaginated() *components.FunctionalityEntityPaginated {
	if o == nil {
		return nil
	}
	return o.FunctionalityEntityPaginated
}
