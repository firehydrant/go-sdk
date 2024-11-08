// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type DeleteTicketRequest struct {
	TicketID string `pathParam:"style=simple,explode=false,name=ticket_id"`
}

func (o *DeleteTicketRequest) GetTicketID() string {
	if o == nil {
		return ""
	}
	return o.TicketID
}

type DeleteTicketResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *DeleteTicketResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
