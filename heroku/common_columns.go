package heroku

import (
	"context"

	v5 "github.com/heroku/heroku-go/v5"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func commonColumns(c []*plugin.Column) []*plugin.Column {
	return append([]*plugin.Column{
		{
			Name:        "account_id",
			Description: "Unique identifier of an account.",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getAccountId,
			Transform:   transform.FromValue(),
		},
	}, c...)
}

// if the caching is required other than per connection, build a cache key for the call and use it in Memoize.
var getAccountIdMemoized = plugin.HydrateFunc(getAccountIdUncached).Memoize(memoize.WithCacheKeyFunction(getAccountIdCacheKey))

// declare a wrapper hydrate function to call the memoized function
// - this is required when a memoized function is used for a column definition
func getAccountId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	accountInfo, err := getAccountIdMemoized(ctx, d, h)
	if err != nil {
		return nil, err
	}
	return accountInfo.(*v5.Account).ID, nil
}

// Build a cache key for the call to getAccountIdCacheKey.
func getAccountIdCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := "getAccountId"
	return key, nil
}

func getAccountIdUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
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
	return item, nil
}
