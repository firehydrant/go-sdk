// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package tests

import (
	"context"
	"firehydrant"
	"firehydrant/models/components"
	"firehydrant/models/operations"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTeams_ListSchedules(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Teams.ListSchedules(ctx, nil, nil, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.Equal(t, &components.ScheduleEntityPaginated{}, res.ScheduleEntityPaginated)
}

func TestTeams_ListTeams(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Teams.List(ctx, operations.ListTeamsRequest{})
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.Equal(t, &components.TeamEntityPaginated{}, res.TeamEntityPaginated)
}

func TestTeams_CreateTeam(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Teams.Create(ctx, components.PostV1Teams{
		Name: "<value>",
	})
	require.NoError(t, err)
	assert.Equal(t, 201, res.HTTPMeta.Response.StatusCode)
	assert.Equal(t, &components.TeamEntity{}, res.TeamEntity)
}

func TestTeams_GetTeam(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Teams.Get(ctx, "<id>", nil)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.Equal(t, &components.TeamEntity{}, res.TeamEntity)
}

func TestTeams_ArchiveTeam(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Teams.Archive(ctx, "<id>")
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.Equal(t, &components.TeamEntity{}, res.TeamEntity)
}

func TestTeams_UpdateTeam(t *testing.T) {
	assert.Fail(t, "incomplete test found please make sure to address the following errors: [`missing request body values for required request body`]")
}

func TestTeams_ListTeamOnCallSchedules(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Teams.ListOnCallSchedules(ctx, "<id>", nil, nil, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestTeams_CreateTeamOnCallSchedule(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Teams.CreateOnCallSchedule(ctx, "<id>", components.PostV1TeamsTeamIDOnCallSchedules{
		Name:     "<value>",
		TimeZone: "America/Porto_Velho",
		Strategy: components.Strategy{
			Type: components.PostV1TeamsTeamIDOnCallSchedulesTypeCustom,
		},
	})
	require.NoError(t, err)
	assert.Equal(t, 201, res.HTTPMeta.Response.StatusCode)
}

func TestTeams_GetTeamScheduleShift(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Teams.GetScheduleShift(ctx, "<id>", "<id>", "<id>")
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestTeams_DeleteTeamScheduleShift(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Teams.DeleteScheduleShift(ctx, "<id>", "<id>", "<id>")
	require.NoError(t, err)
	assert.Equal(t, 204, res.HTTPMeta.Response.StatusCode)
}

func TestTeams_UpdateTeamScheduleShift(t *testing.T) {
	assert.Fail(t, "incomplete test found please make sure to address the following errors: [`missing request body values for required request body`]")
}

func TestTeams_CreateTeamEscalationPolicy(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Teams.CreateEscalationPolicy(ctx, "<id>", components.PostV1TeamsTeamIDEscalationPolicies{
		Name: "<value>",
		Steps: []components.PostV1TeamsTeamIDEscalationPoliciesSteps{
			components.PostV1TeamsTeamIDEscalationPoliciesSteps{
				Targets: []components.Targets{
					components.Targets{
						Type: components.PostV1TeamsTeamIDEscalationPoliciesTypeWebhook,
						ID:   "<id>",
					},
					components.Targets{
						Type: components.PostV1TeamsTeamIDEscalationPoliciesTypeEntireTeam,
						ID:   "<id>",
					},
					components.Targets{
						Type: components.PostV1TeamsTeamIDEscalationPoliciesTypeEntireTeam,
						ID:   "<id>",
					},
				},
				Timeout: "<value>",
			},
		},
	})
	require.NoError(t, err)
	assert.Equal(t, 201, res.HTTPMeta.Response.StatusCode)
}

func TestTeams_UpdateTeamEscalationPolicy(t *testing.T) {
	assert.Fail(t, "incomplete test found please make sure to address the following errors: [`missing request body values for required request body`]")
}

func TestTeams_CreateTeamOnCallScheduleShift(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Teams.CreateShift(ctx, "<id>", "<id>", components.PostV1TeamsTeamIDOnCallSchedulesScheduleIDShifts{
		StartTime: "<value>",
		EndTime:   "<value>",
	})
	require.NoError(t, err)
	assert.Equal(t, 201, res.HTTPMeta.Response.StatusCode)
}