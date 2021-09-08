package heroku

import (
	"context"

	heroku "github.com/heroku/heroku-go/v5"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableHerokuAppRelease(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "heroku_app_release",
		Description: "A release represents a combination of code, config vars and add-ons for an app on Heroku.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("app_name"),
			Hydrate:    listAppRelease,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id", "app_name"}),
			Hydrate:    getAppRelease,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier of release."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Current status of the release."},
			{Name: "is_current", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Current"), Description: "Indicates this release as being the current one for the app."},
			// Other columns

			{Name: "addon_plan_names", Type: proto.ColumnType_JSON, Description: "Add-on plans installed on the app for this release."},
			{Name: "app_name", Type: proto.ColumnType_STRING, Description: "ACM status of this app.", Transform: transform.FromField("App.Name")},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "When release was created."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of changes in this release"},
			{Name: "output_stream_url", Type: proto.ColumnType_STRING, Description: "Release command output will be available from this URL as a stream."},
			{Name: "slug", Type: proto.ColumnType_JSON, Description: "Slug running in this release."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "When release was updated."},
			{Name: "user", Type: proto.ColumnType_JSON, Description: "User that created the release."},
		},
	}
}

func listAppRelease(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_app_release.listAppRelease", "connection_error", err)
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
	items, err := conn.ReleaseList(ctx, appName, &opts)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_app_release.listAppRelease", "query_error", err, "opts", opts)
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

func getAppRelease(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_app_release.getAppRelease", "connection_error", err)
		return nil, err
	}
	id := d.KeyColumnQuals["id"].GetStringValue()
	appName := d.KeyColumnQuals["app_name"].GetStringValue()
	item, err := conn.ReleaseInfo(ctx, appName, id)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_app_release.getAppRelease", "query_error", err, "id", id)
		return nil, err
	}
	return item, err
}
