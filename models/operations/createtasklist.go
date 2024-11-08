// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type CreateTaskListResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Creates a new task list
	TaskListEntity *components.TaskListEntity
}

func (o *CreateTaskListResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateTaskListResponse) GetTaskListEntity() *components.TaskListEntity {
	if o == nil {
		return nil
	}
	return o.TaskListEntity
}