package heroku

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-heroku",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: isNotFoundError,
		},
		TableMap: map[string]*plugin.Table{
			"heroku_account":     tableHerokuAccount(ctx),
			"heroku_addon":      tableHerokuAddOn(ctx),
			"heroku_app":         tableHerokuApp(ctx),
			"heroku_app_release": tableHerokuAppRelease(ctx),
			"heroku_app_webhook": tableHerokuAppWebhook(ctx),
			"heroku_domain":      tableHerokuDomain(ctx),
			"heroku_dyno":        tableHerokuDyno(ctx),
			"heroku_key":         tableHerokuKey(ctx),
			"heroku_pipeline":    tableHerokuPipeline(ctx),
			"heroku_region":      tableHerokuRegion(ctx),
			"heroku_team":        tableHerokuTeam(ctx),
			"heroku_team_member": tableHerokuTeamMember(ctx),
		},
	}
	return p
}
