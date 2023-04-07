package heroku

import (
	"context"

	heroku "github.com/heroku/heroku-go/v5"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableHerokuAddOn(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "heroku_addon",
		Description: "Add-ons are components, services, or pieces of infrastructure that are fully maintained for you, either by a third-party provider or by Heroku.",
		List: &plugin.ListConfig{
			Hydrate: listAddOn,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getAddOn,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Globally unique name of the add-on."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier of add-on."},
			// Other columns
			{Name: "addon_service", Type: proto.ColumnType_JSON, Description: "Identity of add-on service."},
			{Name: "actions", Type: proto.ColumnType_JSON, Description: "Provider actions for this specific add-on."},
			{Name: "provider_id", Type: proto.ColumnType_STRING, Description: "Id of this add-on with its provider."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "When add-on was created."},
			{Name: "state", Type: proto.ColumnType_STRING, Description: "State in the add-on's lifecycle."},
			{Name: "app", Type: proto.ColumnType_JSON, Description: "Billing application associated with this add-on."},
			{Name: "billed_price", Type: proto.ColumnType_JSON, Description: "Billed price."},
			{Name: "billing_entity", Type: proto.ColumnType_JSON, Description: "Billing entity associated with this add-on."},
			{Name: "config_vars", Type: proto.ColumnType_JSON, Description: "Config vars exposed to the owning app by this add-on."},
			{Name: "plan", Type: proto.ColumnType_JSON, Description: "Identity of add-on plan."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "When add-on was updated."},
			{Name: "web_url", Type: proto.ColumnType_STRING, Description: "URL for logging into web interface of add-on (e.g. a dashboard)."},
		},
	}
}

func listAddOn(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_addon.listAddOn", "connection_error", err)
		return nil, err
	}
	opts := heroku.ListRange{Field: "id", Max: 1000}
	limit := d.QueryContext.Limit
	if limit != nil {
		if *limit < int64(1000) {
			opts.Max = int(*limit)
		}
	}
	items, err := conn.AddOnList(ctx, &opts)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_addon.listAddOn", "query_error", err, "opts", opts)
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

func getAddOn(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_addon.getAddOn", "connection_error", err)
		return nil, err
	}
	id := d.EqualsQuals["id"].GetStringValue()
	item, err := conn.AddOnInfo(ctx, id)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_addon.getAddOn", "query_error", err, "id", id)
		return nil, err
	}
	return item, err
}
