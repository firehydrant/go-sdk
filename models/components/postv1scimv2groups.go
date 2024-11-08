// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Members struct {
	// String that represents the user's UUID to assign to the team
	Value string `json:"value"`
}

func (o *Members) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// PostV1ScimV2Groups - SCIM endpoint to create a new Team (Colloquial for Group in the SCIM protocol). Any members defined in the payload will be assigned to the team with no defined role.
type PostV1ScimV2Groups struct {
	// The name of the team being created
	DisplayName string    `json:"displayName"`
	Members     []Members `json:"members"`
}

func (o *PostV1ScimV2Groups) GetDisplayName() string {
	if o == nil {
		return ""
	}
	return o.DisplayName
}

func (o *PostV1ScimV2Groups) GetMembers() []Members {
	if o == nil {
		return []Members{}
	}
	return o.Members
}
