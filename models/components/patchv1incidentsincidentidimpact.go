// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type PatchV1IncidentsIncidentIDImpactImpact struct {
	ID          string `json:"id"`
	ConditionID string `json:"condition_id"`
}

func (o *PatchV1IncidentsIncidentIDImpactImpact) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PatchV1IncidentsIncidentIDImpactImpact) GetConditionID() string {
	if o == nil {
		return ""
	}
	return o.ConditionID
}

type PatchV1IncidentsIncidentIDImpactStatusPages struct {
	ID              string `json:"id"`
	IntegrationSlug string `json:"integration_slug"`
}

func (o *PatchV1IncidentsIncidentIDImpactStatusPages) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PatchV1IncidentsIncidentIDImpactStatusPages) GetIntegrationSlug() string {
	if o == nil {
		return ""
	}
	return o.IntegrationSlug
}

// PatchV1IncidentsIncidentIDImpact - Allows updating an incident's impacted infrastructure, with the option to
// move the incident into a different milestone and provide a note to update
// the incident timeline and any attached status pages. If this method is
// requested with the PUT verb, impacts will be completely replaced with the
// information in the request body, even if not provided (effectively clearing
// all impacts). If this method is requested with the PATCH verb, the provided
// impacts will be added or updated, but no impacts will be removed.
type PatchV1IncidentsIncidentIDImpact struct {
	Note        *string                                       `json:"note,omitempty"`
	Milestone   *string                                       `json:"milestone,omitempty"`
	Impact      []PatchV1IncidentsIncidentIDImpactImpact      `json:"impact,omitempty"`
	StatusPages []PatchV1IncidentsIncidentIDImpactStatusPages `json:"status_pages,omitempty"`
}

func (o *PatchV1IncidentsIncidentIDImpact) GetNote() *string {
	if o == nil {
		return nil
	}
	return o.Note
}

func (o *PatchV1IncidentsIncidentIDImpact) GetMilestone() *string {
	if o == nil {
		return nil
	}
	return o.Milestone
}

func (o *PatchV1IncidentsIncidentIDImpact) GetImpact() []PatchV1IncidentsIncidentIDImpactImpact {
	if o == nil {
		return nil
	}
	return o.Impact
}

func (o *PatchV1IncidentsIncidentIDImpact) GetStatusPages() []PatchV1IncidentsIncidentIDImpactStatusPages {
	if o == nil {
		return nil
	}
	return o.StatusPages
}
