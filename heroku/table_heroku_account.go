package heroku

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func tableHerokuAccount(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "heroku_account",
		Description: "An account represents an individual signed up to use the Heroku platform.",
		List: &plugin.ListConfig{
			Hydrate: getAccount,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Full name of the account owner."},
			{Name: "email", Type: proto.ColumnType_STRING, Description: "Unique email address of account."},
			// Other columns
			{Name: "acknowledged_msa", Type: proto.ColumnType_BOOL, Description: "Whether account has acknowledged the MSA terms of service."},
			{Name: "acknowledged_msa_at", Type: proto.ColumnType_TIMESTAMP, Description: "When account has acknowledged the MSA terms of service."},
			{Name: "allow_tracking", Type: proto.ColumnType_BOOL, Description: "Whether to allow third party web activity tracking."},
			{Name: "beta", Type: proto.ColumnType_BOOL, Description: "Whether allowed to utilize beta Heroku features."},
			{Name: "country_of_residence", Type: proto.ColumnType_STRING, Description: "Country where account owner resides."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "When account was created."},
			{Name: "default_organization", Type: proto.ColumnType_JSON, Description: "Team selected by default."},
			{Name: "default_team", Type: proto.ColumnType_JSON, Description: "Team selected by default."},
			{Name: "delinquent_at", Type: proto.ColumnType_TIMESTAMP, Description: "When account became delinquent."},
			{Name: "federated", Type: proto.ColumnType_BOOL, Description: "Whether the user is federated and belongs to an Identity Provider."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier of an account."},
			{Name: "identity_provider", Type: proto.ColumnType_JSON, Description: "Identity Provider details for federated users."},
			{Name: "italian_customer_terms", Type: proto.ColumnType_STRING, Description: "Whether account has acknowledged the Italian customer terms of service."},
			{Name: "italian_partner_terms", Type: proto.ColumnType_STRING, Description: "Whether account has acknowledged the Italian provider terms of service."},
			{Name: "last_login", Type: proto.ColumnType_TIMESTAMP, Description: "When account last authorized with Heroku."},
			{Name: "sms_number", Type: proto.ColumnType_STRING, Description: "SMS number of account."},
			{Name: "suspended_at", Type: proto.ColumnType_TIMESTAMP, Description: "When account was suspended."},
			{Name: "two_factor_authentication", Type: proto.ColumnType_BOOL, Description: "Whether two-factor auth is enabled on the account."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "When account was updated."},
		},
	}
}

func getAccount(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_account.getAccount", "connection_error", err)
		return nil, err
	}
	item, err := conn.AccountInfo(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_account.getAccount", "query_error", err)
		return nil, err
	}
	d.StreamListItem(ctx, item)
	return nil, nil
}
