package heroku

import (
	"context"

	heroku "github.com/heroku/heroku-go/v5"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func tableHerokuApp(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "heroku_app",
		Description: "An app represents the program that you would like to deploy and run on Heroku.",
		List: &plugin.ListConfig{
			Hydrate: listApp,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getApp,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Unique name of app."},
			// Other columns
			{Name: "acm", Type: proto.ColumnType_STRING, Description: "ACM status of this app."},
			{Name: "archived_at", Type: proto.ColumnType_TIMESTAMP, Description: "When app was archived."},
			{Name: "build_stack", Type: proto.ColumnType_JSON, Description: "Identity of the stack that will be used for new builds."},
			{Name: "buildpack_provided_description", Type: proto.ColumnType_STRING, Description: "Description from buildpack of app."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "When app was created."},
			{Name: "git_url", Type: proto.ColumnType_STRING, Description: "Git repo URL of app."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier of app."},
			{Name: "internal_routing", Type: proto.ColumnType_BOOL, Description: "Describes whether a Private Spaces app is externally routable or not."},
			{Name: "maintenance", Type: proto.ColumnType_BOOL, Description: "Maintenance status of app."},
			{Name: "organization", Type: proto.ColumnType_JSON, Description: "Identity of team."},
			{Name: "owner", Type: proto.ColumnType_JSON, Description: "Identity of app owner."},
			{Name: "region", Type: proto.ColumnType_JSON, Description: "Identity of app region."},
			{Name: "released_at", Type: proto.ColumnType_TIMESTAMP, Description: "When app was released."},
			{Name: "repo_size", Type: proto.ColumnType_INT, Description: "Git repo size in bytes of app."},
			{Name: "slug_size", Type: proto.ColumnType_INT, Description: "Slug size in bytes of app."},
			{Name: "space", Type: proto.ColumnType_JSON, Description: "Identity of space."},
			{Name: "stack", Type: proto.ColumnType_JSON, Description: "Identity of app stack."},
			{Name: "team", Type: proto.ColumnType_JSON, Description: "identity of team."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "When app was updated."},
			{Name: "web_url", Type: proto.ColumnType_STRING, Description: "Web URL of app."},
		},
	}
}

func listApp(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_app.listApp", "connection_error", err)
		return nil, err
	}
	opts := heroku.ListRange{Field: "id", Max: 1000}
	limit := d.QueryContext.Limit
	if limit != nil {
		if *limit < int64(1000) {
			opts.Max = int(*limit)
		}
	}
	items, err := conn.AppList(ctx, &opts)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_app.listApp", "query_error", err, "opts", opts)
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

func getApp(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_app.getApp", "connection_error", err)
		return nil, err
	}
	id := d.KeyColumnQuals["id"].GetStringValue()
	item, err := conn.AppInfo(ctx, id)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_app.getApp", "query_error", err, "id", id)
		return nil, err
	}
	return item, err
}
