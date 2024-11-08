// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
)

// Meta - An object with additional error metadata
type Meta struct {
}

// BadRequest - ErrorEntity model
type BadRequest struct {
	Detail   *string  `json:"detail,omitempty"`
	Messages []string `json:"messages,omitempty"`
	// An object with additional error metadata
	Meta *Meta `json:"meta,omitempty"`
	// A stable code on which to match errors
	Code *string `json:"code,omitempty"`
}

var _ error = &BadRequest{}

func (e *BadRequest) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}