// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type ListConversationCommentReactionsRequest struct {
	ConversationID string `pathParam:"style=simple,explode=false,name=conversation_id"`
	CommentID      string `pathParam:"style=simple,explode=false,name=comment_id"`
}

func (o *ListConversationCommentReactionsRequest) GetConversationID() string {
	if o == nil {
		return ""
	}
	return o.ConversationID
}

func (o *ListConversationCommentReactionsRequest) GetCommentID() string {
	if o == nil {
		return ""
	}
	return o.CommentID
}

type ListConversationCommentReactionsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *ListConversationCommentReactionsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
