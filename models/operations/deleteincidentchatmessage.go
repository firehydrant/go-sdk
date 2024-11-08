// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type DeleteIncidentChatMessageRequest struct {
	MessageID  string `pathParam:"style=simple,explode=false,name=message_id"`
	IncidentID string `pathParam:"style=simple,explode=false,name=incident_id"`
}

func (o *DeleteIncidentChatMessageRequest) GetMessageID() string {
	if o == nil {
		return ""
	}
	return o.MessageID
}

func (o *DeleteIncidentChatMessageRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

type DeleteIncidentChatMessageResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Delete an existing generic chat message on an incident.
	EventGenericChatMessageEntity *components.EventGenericChatMessageEntity
}

func (o *DeleteIncidentChatMessageResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteIncidentChatMessageResponse) GetEventGenericChatMessageEntity() *components.EventGenericChatMessageEntity {
	if o == nil {
		return nil
	}
	return o.EventGenericChatMessageEntity
}
