package heroku

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type herokuConfig struct {
	Email  *string `hcl:"email"`
	APIKey *string `hcl:"api_key"`
}

func ConfigInstance() interface{} {
	return &herokuConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) herokuConfig {
	if connection == nil || connection.Config == nil {
		return herokuConfig{}
	}
	config, _ := connection.Config.(herokuConfig)
	return config
}
