// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package tests

import (
	"context"
	"firehydrant"
	"firehydrant/models/components"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestWebhooks_PostV1Webhooks(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Webhooks.Create(ctx, components.PostV1Webhooks{
		URL: "https://probable-heating.com/",
	})
	require.NoError(t, err)
	assert.Equal(t, 201, res.HTTPMeta.Response.StatusCode)
}

func TestWebhooks_GetV1Webhooks(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Webhooks.List(ctx, nil, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestWebhooks_GetV1WebhooksWebhookIDDeliveries(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Webhooks.GetDeliveries(ctx, "<id>")
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestWebhooks_DeleteV1WebhooksWebhookID(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Webhooks.Delete(ctx, "<id>")
	require.NoError(t, err)
	assert.Equal(t, 204, res.HTTPMeta.Response.StatusCode)
}

func TestWebhooks_PatchV1WebhooksWebhookID(t *testing.T) {
	assert.Fail(t, "incomplete test found please make sure to address the following errors: [`missing request body values for required request body`]")
}

func TestWebhooks_GetV1WebhooksWebhookID(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Webhooks.Get(ctx, "<id>")
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}