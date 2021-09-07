package heroku

import (
	"context"

	heroku "github.com/heroku/heroku-go/v5"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableHerokuTeamMember(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "heroku_team_member",
		Description: "A team member is an individual with access to a team.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("team_name"),
			Hydrate:       listTeamMember,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "email", Type: proto.ColumnType_STRING, Description: "Email address of the team member."},
			// Other columns
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "When the team was created."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier of the team member."},
			{Name: "identity_provider", Type: proto.ColumnType_JSON, Description: "Identity Provider information the member is federated with."},
			{Name: "is_federated", Transform: transform.FromField("Federated"), Type: proto.ColumnType_BOOL, Description: "Whether the user is federated and belongs to an Identity Provider."},
			{Name: "role", Type: proto.ColumnType_STRING, Description: "Role in the team."},
			{Name: "team_name", Type: proto.ColumnType_STRING, Description: "Role in the team.", Transform: transform.FromQual("team_name")},
			{Name: "two_factor_authentication", Type: proto.ColumnType_BOOL, Description: "Whether the enterprise team member has two factor authentication."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "When the membership record was updated."},
			{Name: "user", Type: proto.ColumnType_JSON, Description: "User information for the membership."},
		},
	}
}

func listTeamMember(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_team_member.listTeamMember", "connection_error", err)
		return nil, err
	}
	teamName := d.KeyColumnQuals["team_name"].GetStringValue()
	opts := heroku.ListRange{Field: "id", Max: 1000}
	items, err := conn.TeamMemberList(ctx, teamName, &opts)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_team_member.listTeamMember", "query_error", err, "opts", opts)
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
