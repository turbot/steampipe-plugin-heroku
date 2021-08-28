package heroku

import (
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/schema"
)

type herokuConfig struct {
	Email  *string `cty:"email"`
	APIKey *string `cty:"api_key"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"email": {
		Type: schema.TypeString,
	},
	"api_key": {
		Type: schema.TypeString,
	},
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
