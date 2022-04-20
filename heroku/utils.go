package heroku

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	heroku "github.com/heroku/heroku-go/v5"

	"github.com/pkg/errors"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func connect(_ context.Context, d *plugin.QueryData) (*heroku.Service, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "heroku"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*heroku.Service), nil
	}

	// Default to the env var settings
	email := os.Getenv("HEROKU_EMAIL")
	apiKey := os.Getenv("HEROKU_API_KEY")

	// Prefer config settings
	herokuConfig := GetConfig(d.Connection)
	if &herokuConfig != nil {
		if herokuConfig.Email != nil {
			email = *herokuConfig.Email
		}
		if herokuConfig.APIKey != nil {
			apiKey = *herokuConfig.APIKey
		}
	}

	// Error if the minimum config is not set
	if email == "" {
		return nil, errors.New("email must be configured")
	}
	if apiKey == "" {
		return nil, errors.New("api_key must be configured")
	}

	conn := heroku.NewService(&http.Client{
		Transport: &heroku.Transport{
			Username:  email,
			Password:  apiKey,
			UserAgent: fmt.Sprintf("%s steampipe-plugin-heroku", heroku.DefaultUserAgent),
			Transport: heroku.RoundTripWithRetryBackoff{
				// Configuration fields for ExponentialBackOff
				// InitialIntervalSeconds: 30,
				// RandomizationFactor:    0.25,
				// Multiplier:             2,
				// MaxIntervalSeconds:     900,
				// MaxElapsedTimeSeconds:  0,
			},
		},
	})

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, conn)

	return conn, nil
}

func isNotFoundError(err error) bool {
	return strings.Contains(err.Error(), "Couldn't find") || strings.Contains(err.Error(), "Not found") || strings.Contains(err.Error(), "Unable to process request with specified parameters")
}

/*
// TODO - This would be a better way to do the not found error check, but the
// error object doesn't seem to match?
func isNotFoundError(err error) bool {
	if herokuErr, ok := err.(heroku.Error); ok {
		return herokuErr.StatusCode == 404
	}
	return false
}
*/
