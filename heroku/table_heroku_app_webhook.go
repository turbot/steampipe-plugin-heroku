package heroku

import (
	"context"

	heroku "github.com/heroku/heroku-go/v5"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableHerokuAppWebhook(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "heroku_app_webhook",
		Description: "Webhooks define what web routes should be routed to an app on Heroku.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("app_name"),
			Hydrate:    listWebhook,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id", "app_name"}),
			Hydrate:    getWebhook,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The webhook's unique identifier."},
			{Name: "url", Type: proto.ColumnType_STRING, Description: "The URL where the webhook's notification requests are sent."},
			// Other columns
			{Name: "app", Type: proto.ColumnType_JSON, Description: "Identity of app. Only used for customer webhooks."},
			{Name: "app_name", Type: proto.ColumnType_STRING, Description: "The app name .", Transform: transform.FromField("App.Name")},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "When the webhook was created."},
			{Name: "include", Type: proto.ColumnType_JSON, Description: "The entities that the subscription provides notifications for."},
			{Name: "level", Type: proto.ColumnType_STRING, Description: "If `notify`, Heroku makes a single, fire-and-forget delivery attempt."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "When the webhook was updated."},
		},
	}
}

func listWebhook(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_app_webhook.listWebhook", "connection_error", err)
		return nil, err
	}
	appName := d.KeyColumnQuals["app_name"].GetStringValue()
	opts := heroku.ListRange{Field: "id", Max: 1000}
	limit := d.QueryContext.Limit
	if limit != nil {
		if *limit < int64(1000) {
			opts.Max = int(*limit)
		}
	}
	items, err := conn.AppWebhookList(ctx, appName, &opts)
	if err != nil {
		if isNotFoundError(err) {
			return nil, nil
		}
		plugin.Logger(ctx).Error("heroku_app_webhook.listWebhook", "query_error", err, "opts", opts)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
		// Context can be cancelled due to manual cancellation or the limit has been hit
		if plugin.IsCancelled(ctx) {
			return nil, nil
		}
	}
	return nil, nil
}

func getWebhook(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_app_webhook.getWebhook", "connection_error", err)
		return nil, err
	}
	id := d.KeyColumnQuals["id"].GetStringValue()
	appName := d.KeyColumnQuals["app_name"].GetStringValue()
	item, err := conn.AppWebhookInfo(ctx, appName, id)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_app_webhook.getApp", "query_error", err, "id", id)
		return nil, err
	}
	return item, err
}
