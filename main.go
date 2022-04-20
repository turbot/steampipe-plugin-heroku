package main

import (
	"github.com/turbot/steampipe-plugin-heroku/heroku"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: heroku.Plugin})
}
