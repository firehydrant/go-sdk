// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type GetIncidentTaskRequest struct {
	TaskID     string `pathParam:"style=simple,explode=false,name=task_id"`
	IncidentID string `pathParam:"style=simple,explode=false,name=incident_id"`
}

func (o *GetIncidentTaskRequest) GetTaskID() string {
	if o == nil {
		return ""
	}
	return o.TaskID
}

func (o *GetIncidentTaskRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

type GetIncidentTaskResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *GetIncidentTaskResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}