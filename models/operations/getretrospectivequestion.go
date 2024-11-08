// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type GetRetrospectiveQuestionRequest struct {
	QuestionID string `pathParam:"style=simple,explode=false,name=question_id"`
}

func (o *GetRetrospectiveQuestionRequest) GetQuestionID() string {
	if o == nil {
		return ""
	}
	return o.QuestionID
}

type GetRetrospectiveQuestionResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *GetRetrospectiveQuestionResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}