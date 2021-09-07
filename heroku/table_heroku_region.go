package heroku

import (
	"context"

	heroku "github.com/heroku/heroku-go/v5"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableHerokuRegion(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "heroku_region",
		Description: "A region represents a geographic location in which your application may run.",
		List: &plugin.ListConfig{
			Hydrate: listRegion,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getRegion,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Unique name of region."},
			// Other columns
			{Name: "country", Type: proto.ColumnType_STRING, Description: "Country where the region exists."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "When region was created."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of region."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier of region."},
			{Name: "locale", Type: proto.ColumnType_STRING, Description: "Area in the country where the region exists."},
			{Name: "private_capable", Type: proto.ColumnType_BOOL, Description: "Whether or not region is available for creating a Private Space."},
			{Name: "provider", Type: proto.ColumnType_JSON, Description: "Provider of underlying substrate."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "When region was updated."},
		},
	}
}

func listRegion(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_region.listRegion", "connection_error", err)
		return nil, err
	}
	opts := heroku.ListRange{Field: "id", Max: 1000}
	items, err := conn.RegionList(ctx, &opts)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_region.listRegion", "query_error", err, "opts", opts)
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

func getRegion(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_region.getRegion", "connection_error", err)
		return nil, err
	}
	id := d.KeyColumnQuals["id"].GetStringValue()
	item, err := conn.RegionInfo(ctx, id)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_region.getRegion", "query_error", err, "id", id)
		return nil, err
	}
	return item, err
}
