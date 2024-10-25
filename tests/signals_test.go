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

func TestSignals_GetV1SignalsAnalyticsTimeseries(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Signals.GetAnalyticsTimeseries(ctx, operations.GetV1SignalsAnalyticsTimeseriesRequest{})
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestSignals_GetV1SignalsAnalyticsGroupedMetrics(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Signals.ReportGroupedMetrics(ctx, operations.GetV1SignalsAnalyticsGroupedMetricsRequest{})
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestSignals_GetV1SignalsAnalyticsMttx(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Signals.GetMttxAnalytics(ctx, operations.GetV1SignalsAnalyticsMttxRequest{})
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestSignals_GetV1SignalsEventSources(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Signals.ListEventSources(ctx, nil, nil, nil, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestSignals_PostV1SignalsEmailTargets(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Signals.CreateEmailTarget(ctx, components.PostV1SignalsEmailTargets{
		Name: "<value>",
	})
	require.NoError(t, err)
	assert.Equal(t, 201, res.HTTPMeta.Response.StatusCode)
}

func TestSignals_GetV1SignalsEmailTargets(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Signals.ListEmailTargets(ctx, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestSignals_DeleteV1SignalsEmailTargetsID(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Signals.DeleteEmailTarget(ctx, "<id>")
	require.NoError(t, err)
	assert.Equal(t, 204, res.HTTPMeta.Response.StatusCode)
}

func TestSignals_PatchV1SignalsEmailTargetsID(t *testing.T) {
	assert.Fail(t, "incomplete test found please make sure to address the following errors: [`missing request body values for required request body`]")
}

func TestSignals_PostV1SignalsWebhookTargets(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Signals.CreateWebhookTarget(ctx, components.PostV1SignalsWebhookTargets{
		Name: "<value>",
		URL:  "https://critical-prohibition.biz",
	})
	require.NoError(t, err)
	assert.Equal(t, 201, res.HTTPMeta.Response.StatusCode)
}

func TestSignals_GetV1SignalsWebhookTargets(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Signals.GetWebhookTargets(ctx, nil, nil, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestSignals_DeleteV1SignalsWebhookTargetsID(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Signals.DeleteWebhookTarget(ctx, "<id>")
	require.NoError(t, err)
	assert.Equal(t, 204, res.HTTPMeta.Response.StatusCode)
}

func TestSignals_PatchV1SignalsWebhookTargetsID(t *testing.T) {
	assert.Fail(t, "incomplete test found please make sure to address the following errors: [`missing request body values for required request body`]")
}

func TestSignals_GetV1SignalsWebhookTargetsID(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Signals.GetWebhookTarget(ctx, "<id>")
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestSignals_GetV1SignalsTransposers(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Signals.GetTransposers(ctx)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestSignals_GetV1SignalsIngestURL(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Signals.GetIngestURL(ctx, nil, nil, nil, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestSignals_PostV1SignalsDebugger(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Signals.PostDebugger(ctx, components.PostV1SignalsDebugger{
		Expression: "<value>",
		Signals: []components.Signals{
			components.Signals{},
		},
	})
	require.NoError(t, err)
	assert.Equal(t, 201, res.HTTPMeta.Response.StatusCode)
}