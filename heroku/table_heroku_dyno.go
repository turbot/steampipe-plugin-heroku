package heroku

import (
	"context"

	heroku "github.com/heroku/heroku-go/v5"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableHerokuDyno(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "heroku_dyno",
		Description: "Dynos encapsulate running processes of an app on Heroku.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("app_name"),
			Hydrate:    listDyno,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id", "app_name"}),
			Hydrate:    getDyno,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier of this dyno."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of this process on this dyno."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of process."},
			// Other columns
			{Name: "app", Type: proto.ColumnType_JSON, Description: "App formation belongs to."},
			{Name: "app_name", Type: proto.ColumnType_STRING, Description: "App name formation belongs to.", Transform: transform.FromField("App.Name")},
			{Name: "attach_url", Type: proto.ColumnType_STRING, Description: "A URL to stream output from for attached processes or null for."},
			{Name: "command", Type: proto.ColumnType_STRING, Description: "Command used to start this process."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "When dyno was created."},
			{Name: "release", Type: proto.ColumnType_JSON, Description: "App release of the dyno."},
			{Name: "size", Type: proto.ColumnType_STRING, Description: "Dyno size (default: standard-1X)."},
			{Name: "state", Type: proto.ColumnType_STRING, Description: "Current status of process (either: crashed, down, idle, starting, or up)"},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "When process last changed state."},
		},
	}
}

func listDyno(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_dyno.listDyno", "connection_error", err)
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
	items, err := conn.DynoList(ctx, appName, &opts)
	if err != nil {
		if isNotFoundError(err) {
			return nil, nil
		}
		plugin.Logger(ctx).Error("heroku_dyno.listDyno", "query_error", err, "opts", opts)
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

func getDyno(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_dyno.getDyno", "connection_error", err)
		return nil, err
	}
	id := d.KeyColumnQuals["id"].GetStringValue()
	appName := d.KeyColumnQuals["app_name"].GetStringValue()
	item, err := conn.DynoInfo(ctx, appName, id)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_dyno.getDyno", "query_error", err, "id", id)
		return nil, err
	}
	return item, err
}
