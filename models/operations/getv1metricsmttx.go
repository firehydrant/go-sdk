// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"firehydrant/internal/utils"
	"firehydrant/models/components"
	"firehydrant/types"
	"fmt"
	"time"
)

// GetV1MetricsMttxQueryParamTagMatchStrategy - A matching strategy for the tags provided
type GetV1MetricsMttxQueryParamTagMatchStrategy string

const (
	GetV1MetricsMttxQueryParamTagMatchStrategyAny      GetV1MetricsMttxQueryParamTagMatchStrategy = "any"
	GetV1MetricsMttxQueryParamTagMatchStrategyMatchAll GetV1MetricsMttxQueryParamTagMatchStrategy = "match_all"
	GetV1MetricsMttxQueryParamTagMatchStrategyExclude  GetV1MetricsMttxQueryParamTagMatchStrategy = "exclude"
)

func (e GetV1MetricsMttxQueryParamTagMatchStrategy) ToPointer() *GetV1MetricsMttxQueryParamTagMatchStrategy {
	return &e
}
func (e *GetV1MetricsMttxQueryParamTagMatchStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "any":
		fallthrough
	case "match_all":
		fallthrough
	case "exclude":
		*e = GetV1MetricsMttxQueryParamTagMatchStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetV1MetricsMttxQueryParamTagMatchStrategy: %v", v)
	}
}

type SortBy string

const (
	SortByCountAsc        SortBy = "count_asc"
	SortByMttrAsc         SortBy = "mttr_asc"
	SortByMttaAsc         SortBy = "mtta_asc"
	SortByMttdAsc         SortBy = "mttd_asc"
	SortByMttmAsc         SortBy = "mttm_asc"
	SortByHealthinessAsc  SortBy = "healthiness_asc"
	SortByCountDesc       SortBy = "count_desc"
	SortByMttrDesc        SortBy = "mttr_desc"
	SortByMttaDesc        SortBy = "mtta_desc"
	SortByMttdDesc        SortBy = "mttd_desc"
	SortByMttmDesc        SortBy = "mttm_desc"
	SortByHealthinessDesc SortBy = "healthiness_desc"
)

func (e SortBy) ToPointer() *SortBy {
	return &e
}
func (e *SortBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "count_asc":
		fallthrough
	case "mttr_asc":
		fallthrough
	case "mtta_asc":
		fallthrough
	case "mttd_asc":
		fallthrough
	case "mttm_asc":
		fallthrough
	case "healthiness_asc":
		fallthrough
	case "count_desc":
		fallthrough
	case "mttr_desc":
		fallthrough
	case "mtta_desc":
		fallthrough
	case "mttd_desc":
		fallthrough
	case "mttm_desc":
		fallthrough
	case "healthiness_desc":
		*e = SortBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SortBy: %v", v)
	}
}

type GetV1MetricsMttxGroupBy string

const (
	GetV1MetricsMttxGroupByServices        GetV1MetricsMttxGroupBy = "services"
	GetV1MetricsMttxGroupByEnvironments    GetV1MetricsMttxGroupBy = "environments"
	GetV1MetricsMttxGroupByFunctionalities GetV1MetricsMttxGroupBy = "functionalities"
	GetV1MetricsMttxGroupByTeams           GetV1MetricsMttxGroupBy = "teams"
	GetV1MetricsMttxGroupBySeverities      GetV1MetricsMttxGroupBy = "severities"
	GetV1MetricsMttxGroupByUsers           GetV1MetricsMttxGroupBy = "users"
	GetV1MetricsMttxGroupByStartedDay      GetV1MetricsMttxGroupBy = "started_day"
	GetV1MetricsMttxGroupByStartedWeek     GetV1MetricsMttxGroupBy = "started_week"
	GetV1MetricsMttxGroupByStartedMonth    GetV1MetricsMttxGroupBy = "started_month"
	GetV1MetricsMttxGroupByCustomFields    GetV1MetricsMttxGroupBy = "custom_fields"
)

func (e GetV1MetricsMttxGroupBy) ToPointer() *GetV1MetricsMttxGroupBy {
	return &e
}
func (e *GetV1MetricsMttxGroupBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "services":
		fallthrough
	case "environments":
		fallthrough
	case "functionalities":
		fallthrough
	case "teams":
		fallthrough
	case "severities":
		fallthrough
	case "users":
		fallthrough
	case "started_day":
		fallthrough
	case "started_week":
		fallthrough
	case "started_month":
		fallthrough
	case "custom_fields":
		*e = GetV1MetricsMttxGroupBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetV1MetricsMttxGroupBy: %v", v)
	}
}

type GetV1MetricsMttxRequestBody struct {
	GroupBy []GetV1MetricsMttxGroupBy `multipartForm:"name=group_by"`
}

func (o *GetV1MetricsMttxRequestBody) GetGroupBy() []GetV1MetricsMttxGroupBy {
	if o == nil {
		return nil
	}
	return o.GroupBy
}

type GetV1MetricsMttxRequest struct {
	Page    *int `queryParam:"style=form,explode=true,name=page"`
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
	// A JSON string that defines 'logic' and 'user_data'
	Conditions *string `queryParam:"style=form,explode=true,name=conditions"`
	// A comma separated list of environment IDs or 'is_empty' to filter for incidents with no impacted environments
	Environments *string `queryParam:"style=form,explode=true,name=environments"`
	// A comma separated list of service IDs or 'is_empty' to filter for incidents with no impacted services
	Services *string `queryParam:"style=form,explode=true,name=services"`
	// A comma separated list of functionality IDs or 'is_empty' to filter for incidents with no impacted functionalities
	Functionalities *string `queryParam:"style=form,explode=true,name=functionalities"`
	// A comma separated list of infrastructure IDs. Returns incidents that do not have the following infrastructure ids associated with them.
	ExcludedInfrastructureIds *string `queryParam:"style=form,explode=true,name=excluded_infrastructure_ids"`
	// A comma separated list of team IDs
	Teams *string `queryParam:"style=form,explode=true,name=teams"`
	// A comma separated list of IDs for assigned teams or 'is_empty' to filter for incidents with no active team assignments
	AssignedTeams *string `queryParam:"style=form,explode=true,name=assigned_teams"`
	// Incident status
	Status *string `queryParam:"style=form,explode=true,name=status"`
	// Filters for incidents that started on or after this date
	StartDate types.Date `queryParam:"style=form,explode=true,name=start_date"`
	// Filters for incidents that started on or before this date
	EndDate types.Date `queryParam:"style=form,explode=true,name=end_date"`
	// Filters for incidents that were resolved at or after this time. Combine this with the `current_milestones` parameter if you wish to omit incidents that were re-opened and are still active.
	ResolvedAtOrAfter *time.Time `queryParam:"style=form,explode=true,name=resolved_at_or_after"`
	// Filters for incidents that were resolved at or before this time. Combine this with the `current_milestones` parameter if you wish to omit incidents that were re-opened and are still active.
	ResolvedAtOrBefore *time.Time `queryParam:"style=form,explode=true,name=resolved_at_or_before"`
	// Filters for incidents that were created at or after this time
	CreatedAtOrAfter *time.Time `queryParam:"style=form,explode=true,name=created_at_or_after"`
	// Filters for incidents that were created at or before this time
	CreatedAtOrBefore *time.Time `queryParam:"style=form,explode=true,name=created_at_or_before"`
	// A text query for an incident that searches on name, summary, and desciption
	Query *string `queryParam:"style=form,explode=true,name=query"`
	// A query to search incidents by their name
	Name *string `queryParam:"style=form,explode=true,name=name"`
	// The id of a previously saved search.
	SavedSearchID *string `queryParam:"style=form,explode=true,name=saved_search_id"`
	// A text value of priority
	Priorities *string `queryParam:"style=form,explode=true,name=priorities"`
	// Flag for including incidents where priority has not been set
	PriorityNotSet *bool `queryParam:"style=form,explode=true,name=priority_not_set"`
	// A text value of severity
	Severities *string `queryParam:"style=form,explode=true,name=severities"`
	// Flag for including incidents where severity has not been set
	SeverityNotSet *bool `queryParam:"style=form,explode=true,name=severity_not_set"`
	// A comma separated list of current milestones
	CurrentMilestones *string `queryParam:"style=form,explode=true,name=current_milestones"`
	// A comma separated list of tags
	Tags *string `queryParam:"style=form,explode=true,name=tags"`
	// A matching strategy for the tags provided
	TagMatchStrategy *GetV1MetricsMttxQueryParamTagMatchStrategy `queryParam:"style=form,explode=true,name=tag_match_strategy"`
	// Return archived incidents
	Archived *bool `queryParam:"style=form,explode=true,name=archived"`
	// Filters for incidents that were updated after this date
	UpdatedAfter *time.Time `queryParam:"style=form,explode=true,name=updated_after"`
	// Filters for incidents that were updated before this date
	UpdatedBefore *time.Time `queryParam:"style=form,explode=true,name=updated_before"`
	// A comma separated list of incident type IDs
	IncidentTypeID *string `queryParam:"style=form,explode=true,name=incident_type_id"`
	CustomFieldID  *string `queryParam:"style=form,explode=true,name=custom_field_id"`
	SortBy         *SortBy `queryParam:"style=form,explode=true,name=sort_by"`
	// Comma-separated list of measurements to include in the response
	Measurements *string                      `queryParam:"style=form,explode=true,name=measurements"`
	RequestBody  *GetV1MetricsMttxRequestBody `request:"mediaType=multipart/form-data"`
}

func (g GetV1MetricsMttxRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetV1MetricsMttxRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetV1MetricsMttxRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetV1MetricsMttxRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *GetV1MetricsMttxRequest) GetConditions() *string {
	if o == nil {
		return nil
	}
	return o.Conditions
}

func (o *GetV1MetricsMttxRequest) GetEnvironments() *string {
	if o == nil {
		return nil
	}
	return o.Environments
}

func (o *GetV1MetricsMttxRequest) GetServices() *string {
	if o == nil {
		return nil
	}
	return o.Services
}

func (o *GetV1MetricsMttxRequest) GetFunctionalities() *string {
	if o == nil {
		return nil
	}
	return o.Functionalities
}

func (o *GetV1MetricsMttxRequest) GetExcludedInfrastructureIds() *string {
	if o == nil {
		return nil
	}
	return o.ExcludedInfrastructureIds
}

func (o *GetV1MetricsMttxRequest) GetTeams() *string {
	if o == nil {
		return nil
	}
	return o.Teams
}

func (o *GetV1MetricsMttxRequest) GetAssignedTeams() *string {
	if o == nil {
		return nil
	}
	return o.AssignedTeams
}

func (o *GetV1MetricsMttxRequest) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *GetV1MetricsMttxRequest) GetStartDate() types.Date {
	if o == nil {
		return types.Date{}
	}
	return o.StartDate
}

func (o *GetV1MetricsMttxRequest) GetEndDate() types.Date {
	if o == nil {
		return types.Date{}
	}
	return o.EndDate
}

func (o *GetV1MetricsMttxRequest) GetResolvedAtOrAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.ResolvedAtOrAfter
}

func (o *GetV1MetricsMttxRequest) GetResolvedAtOrBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.ResolvedAtOrBefore
}

func (o *GetV1MetricsMttxRequest) GetCreatedAtOrAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAtOrAfter
}

func (o *GetV1MetricsMttxRequest) GetCreatedAtOrBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAtOrBefore
}

func (o *GetV1MetricsMttxRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *GetV1MetricsMttxRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *GetV1MetricsMttxRequest) GetSavedSearchID() *string {
	if o == nil {
		return nil
	}
	return o.SavedSearchID
}

func (o *GetV1MetricsMttxRequest) GetPriorities() *string {
	if o == nil {
		return nil
	}
	return o.Priorities
}

func (o *GetV1MetricsMttxRequest) GetPriorityNotSet() *bool {
	if o == nil {
		return nil
	}
	return o.PriorityNotSet
}

func (o *GetV1MetricsMttxRequest) GetSeverities() *string {
	if o == nil {
		return nil
	}
	return o.Severities
}

func (o *GetV1MetricsMttxRequest) GetSeverityNotSet() *bool {
	if o == nil {
		return nil
	}
	return o.SeverityNotSet
}

func (o *GetV1MetricsMttxRequest) GetCurrentMilestones() *string {
	if o == nil {
		return nil
	}
	return o.CurrentMilestones
}

func (o *GetV1MetricsMttxRequest) GetTags() *string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *GetV1MetricsMttxRequest) GetTagMatchStrategy() *GetV1MetricsMttxQueryParamTagMatchStrategy {
	if o == nil {
		return nil
	}
	return o.TagMatchStrategy
}

func (o *GetV1MetricsMttxRequest) GetArchived() *bool {
	if o == nil {
		return nil
	}
	return o.Archived
}

func (o *GetV1MetricsMttxRequest) GetUpdatedAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAfter
}

func (o *GetV1MetricsMttxRequest) GetUpdatedBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedBefore
}

func (o *GetV1MetricsMttxRequest) GetIncidentTypeID() *string {
	if o == nil {
		return nil
	}
	return o.IncidentTypeID
}

func (o *GetV1MetricsMttxRequest) GetCustomFieldID() *string {
	if o == nil {
		return nil
	}
	return o.CustomFieldID
}

func (o *GetV1MetricsMttxRequest) GetSortBy() *SortBy {
	if o == nil {
		return nil
	}
	return o.SortBy
}

func (o *GetV1MetricsMttxRequest) GetMeasurements() *string {
	if o == nil {
		return nil
	}
	return o.Measurements
}

func (o *GetV1MetricsMttxRequest) GetRequestBody() *GetV1MetricsMttxRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type GetV1MetricsMttxResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Fetch infrastructure metrics based on custom query
	MetricsMttxDataEntity *components.MetricsMttxDataEntity
}

func (o *GetV1MetricsMttxResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetV1MetricsMttxResponse) GetMetricsMttxDataEntity() *components.MetricsMttxDataEntity {
	if o == nil {
		return nil
	}
	return o.MetricsMttxDataEntity
}
