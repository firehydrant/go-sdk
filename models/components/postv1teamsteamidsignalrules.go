// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// TargetType - The type of target that the rule will notify when matched.
type TargetType string

const (
	TargetTypeEscalationPolicy TargetType = "EscalationPolicy"
	TargetTypeOnCallSchedule   TargetType = "OnCallSchedule"
	TargetTypeUser             TargetType = "User"
	TargetTypeWebhook          TargetType = "Webhook"
)

func (e TargetType) ToPointer() *TargetType {
	return &e
}
func (e *TargetType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "EscalationPolicy":
		fallthrough
	case "OnCallSchedule":
		fallthrough
	case "User":
		fallthrough
	case "Webhook":
		*e = TargetType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TargetType: %v", v)
	}
}

// PostV1TeamsTeamIDSignalRulesNotificationPriorityOverride - A notification priority that will be set on the resulting alert (default: HIGH)
type PostV1TeamsTeamIDSignalRulesNotificationPriorityOverride string

const (
	PostV1TeamsTeamIDSignalRulesNotificationPriorityOverrideHigh    PostV1TeamsTeamIDSignalRulesNotificationPriorityOverride = "HIGH"
	PostV1TeamsTeamIDSignalRulesNotificationPriorityOverrideMedium  PostV1TeamsTeamIDSignalRulesNotificationPriorityOverride = "MEDIUM"
	PostV1TeamsTeamIDSignalRulesNotificationPriorityOverrideLow     PostV1TeamsTeamIDSignalRulesNotificationPriorityOverride = "LOW"
	PostV1TeamsTeamIDSignalRulesNotificationPriorityOverrideUnknown PostV1TeamsTeamIDSignalRulesNotificationPriorityOverride = ""
)

func (e PostV1TeamsTeamIDSignalRulesNotificationPriorityOverride) ToPointer() *PostV1TeamsTeamIDSignalRulesNotificationPriorityOverride {
	return &e
}
func (e *PostV1TeamsTeamIDSignalRulesNotificationPriorityOverride) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "HIGH":
		fallthrough
	case "MEDIUM":
		fallthrough
	case "LOW":
		fallthrough
	case "":
		*e = PostV1TeamsTeamIDSignalRulesNotificationPriorityOverride(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostV1TeamsTeamIDSignalRulesNotificationPriorityOverride: %v", v)
	}
}

// PostV1TeamsTeamIDSignalRules - Create a Signals rule for a team.
type PostV1TeamsTeamIDSignalRules struct {
	// The rule's name.
	Name string `json:"name"`
	// The CEL expression that defines the rule.
	Expression string `json:"expression"`
	// The type of target that the rule will notify when matched.
	TargetType TargetType `json:"target_type"`
	// The ID of the target that the rule will notify when matched.
	TargetID string `json:"target_id"`
	// The ID of an incident type that should be used when an alert is promoted to an incident
	IncidentTypeID *string `json:"incident_type_id,omitempty"`
	// A notification priority that will be set on the resulting alert (default: HIGH)
	NotificationPriorityOverride *PostV1TeamsTeamIDSignalRulesNotificationPriorityOverride `json:"notification_priority_override,omitempty"`
}

func (o *PostV1TeamsTeamIDSignalRules) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PostV1TeamsTeamIDSignalRules) GetExpression() string {
	if o == nil {
		return ""
	}
	return o.Expression
}

func (o *PostV1TeamsTeamIDSignalRules) GetTargetType() TargetType {
	if o == nil {
		return TargetType("")
	}
	return o.TargetType
}

func (o *PostV1TeamsTeamIDSignalRules) GetTargetID() string {
	if o == nil {
		return ""
	}
	return o.TargetID
}

func (o *PostV1TeamsTeamIDSignalRules) GetIncidentTypeID() *string {
	if o == nil {
		return nil
	}
	return o.IncidentTypeID
}

func (o *PostV1TeamsTeamIDSignalRules) GetNotificationPriorityOverride() *PostV1TeamsTeamIDSignalRulesNotificationPriorityOverride {
	if o == nil {
		return nil
	}
	return o.NotificationPriorityOverride
}
