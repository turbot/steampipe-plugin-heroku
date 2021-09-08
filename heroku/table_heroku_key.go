package heroku

import (
	"context"

	heroku "github.com/heroku/heroku-go/v5"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableHerokuKey(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "heroku_key",
		Description: "Keys represent public SSH keys associated with an account and are used to authorize accounts as they are performing git operations.",
		List: &plugin.ListConfig{
			Hydrate: listKey,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getKey,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "comment", Type: proto.ColumnType_STRING, Description: "Comment on the key."},
			// Other columns
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "When key was created."},
			{Name: "email", Type: proto.ColumnType_STRING, Description: "Deprecated. Please refer to 'comment' instead."},
			{Name: "fingerprint", Type: proto.ColumnType_STRING, Description: "A unique identifying string based on contents."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier of this key."},
			{Name: "public_key", Type: proto.ColumnType_STRING, Description: "Full public_key as uploaded."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "When key was updated."},
		},
	}
}

func listKey(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_key.listKey", "connection_error", err)
		return nil, err
	}
	opts := heroku.ListRange{Field: "id", Max: 1000}
	limit := d.QueryContext.Limit
	if limit != nil {
		if *limit < int64(1000) {
			opts.Max = int(*limit)
		}
	}
	items, err := conn.KeyList(ctx, &opts)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_key.listKey", "query_error", err, "opts", opts)
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

func getKey(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_key.getKey", "connection_error", err)
		return nil, err
	}
	id := d.KeyColumnQuals["id"].GetStringValue()
	item, err := conn.KeyInfo(ctx, id)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_key.getKey", "query_error", err, "id", id)
		return nil, err
	}
	return item, err
}
