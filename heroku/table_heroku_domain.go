package heroku

import (
	"context"

	heroku "github.com/heroku/heroku-go/v5"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableHerokuDomain(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "heroku_domain",
		Description: "Domains define what web routes should be routed to an app on Heroku.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("app_name"),
			Hydrate:    listDomain,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id", "app_name"}),
			Hydrate:    getDomain,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier of this domain."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Current status of the release."},
			{Name: "kind", Type: proto.ColumnType_STRING, Description: "Type of domain name."},
			// Other columns

			{Name: "acm_status", Type: proto.ColumnType_STRING, Description: "status of this record's ACM.",},
			{Name: "acm_status_reason", Type: proto.ColumnType_STRING, Description: "Reason for the status of this record's ACM.",},
			{Name: "app", Type: proto.ColumnType_JSON, Description: "App that owns the domain."},
			{Name: "app_name", Type: proto.ColumnType_STRING, Description: "ACM status of this app.", Transform: transform.FromField("App.Name")},
			{Name: "cname", Type: proto.ColumnType_STRING, Description: "Canonical name record, the address to point a domain at."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "When release was created."},
			{Name: "hostname", Type: proto.ColumnType_STRING, Description: "Full hostname.",},
			{Name: "sni_endpoint", Type: proto.ColumnType_JSON, Description: "SNI endpoint the domain is associated with."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "When release was updated."},
		},
	}
}

func listDomain(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_domain.listDomain", "connection_error", err)
		return nil, err
	}
	appName := d.KeyColumnQuals["app_name"].GetStringValue()
	opts := heroku.ListRange{Field: "id"}
	items, err := conn.DomainList(ctx, appName, &opts)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_domain.listDomain", "query_error", err, "opts", opts)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getDomain(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_domain.getDomain", "connection_error", err)
		return nil, err
	}
	id := d.KeyColumnQuals["id"].GetStringValue()
	appName := d.KeyColumnQuals["app_name"].GetStringValue()
	item, err := conn.DomainInfo(ctx, appName, id)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_domain.getDomain", "query_error", err, "id", id)
		return nil, err
	}
	return item, err
}
