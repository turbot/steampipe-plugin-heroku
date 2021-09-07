package heroku

import (
	"context"

	heroku "github.com/heroku/heroku-go/v5"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableHerokuTeam(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "heroku_team",
		Description: "Teams allow you to manage access to a shared group of applications and other resources.",
		List: &plugin.ListConfig{
			Hydrate: listTeam,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getTeam,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Unique name of team."},
			// Other columns
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "When the team was created."},
			{Name: "credit_card_collections", Type: proto.ColumnType_BOOL, Description: "Whether charges incurred by the team are paid by credit card."},
			{Name: "is_default", Transform: transform.FromField("Default"), Type: proto.ColumnType_BOOL, Description: "Whether to use this team when none is specified."},
			{Name: "enterprise_account", Type: proto.ColumnType_JSON, Description: "Enterprise Account associated with the team."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier of team."},
			{Name: "identity_provider", Type: proto.ColumnType_JSON, Description: "Identity Provider associated with the team."},
			{Name: "membership_limit", Type: proto.ColumnType_DOUBLE, Description: "Upper limit of members allowed in a team."},
			{Name: "provisioned_licenses", Type: proto.ColumnType_BOOL, Description: "Whether the team is provisioned licenses by salesforce."},
			{Name: "role", Type: proto.ColumnType_STRING, Description: "Role in the team."},
			{Name: "team_type", Transform: transform.FromField("Type"), Type: proto.ColumnType_STRING, Description: "Type of team."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "When the team was updated."},
		},
	}
}

func listTeam(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_team.listTeam", "connection_error", err)
		return nil, err
	}
	opts := heroku.ListRange{Field: "id", Max: 1000}
	items, err := conn.TeamList(ctx, &opts)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_team.listTeam", "query_error", err, "opts", opts)
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

func getTeam(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_team.getTeam", "connection_error", err)
		return nil, err
	}
	id := d.KeyColumnQuals["id"].GetStringValue()
	item, err := conn.TeamInfo(ctx, id)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_team.getTeam", "query_error", err, "id", id)
		return nil, err
	}
	return item, err
}
