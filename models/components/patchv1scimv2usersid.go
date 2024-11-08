// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Operations struct {
	// The operation to perform on the user. Options are add, remove, replace
	Op string `json:"op"`
	// The path to the attribute to be modified
	Path string `json:"path"`
}

func (o *Operations) GetOp() string {
	if o == nil {
		return ""
	}
	return o.Op
}

func (o *Operations) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}

// PatchV1ScimV2UsersID - PATCH SCIM endpoint to update a User. This endpoint is used to update a resource's attributes.
type PatchV1ScimV2UsersID struct {
	// An optional trail to log the request
	Trail *string `json:"trail,omitempty"`
	// An array of operations to perform on the user
	Operations []Operations `json:"Operations"`
}

func (o *PatchV1ScimV2UsersID) GetTrail() *string {
	if o == nil {
		return nil
	}
	return o.Trail
}

func (o *PatchV1ScimV2UsersID) GetOperations() []Operations {
	if o == nil {
		return []Operations{}
	}
	return o.Operations
}
