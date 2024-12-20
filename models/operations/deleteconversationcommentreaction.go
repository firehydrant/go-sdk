// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type DeleteConversationCommentReactionRequest struct {
	ReactionID     string `pathParam:"style=simple,explode=false,name=reaction_id"`
	ConversationID string `pathParam:"style=simple,explode=false,name=conversation_id"`
	CommentID      string `pathParam:"style=simple,explode=false,name=comment_id"`
}

func (o *DeleteConversationCommentReactionRequest) GetReactionID() string {
	if o == nil {
		return ""
	}
	return o.ReactionID
}

func (o *DeleteConversationCommentReactionRequest) GetConversationID() string {
	if o == nil {
		return ""
	}
	return o.ConversationID
}

func (o *DeleteConversationCommentReactionRequest) GetCommentID() string {
	if o == nil {
		return ""
	}
	return o.CommentID
}

type DeleteConversationCommentReactionResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *DeleteConversationCommentReactionResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
