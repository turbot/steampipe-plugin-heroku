package heroku

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableHerokuPipeline(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "heroku_pipeline",
		Description: "A pipeline allows grouping of apps into different stages.",
		/*
			List: &plugin.ListConfig{
				Hydrate: listPipeline,
			},
		*/
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getPipeline,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of pipeline."},
			// Other columns
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "When pipeline was created."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier of pipeline."},
			{Name: "owner", Type: proto.ColumnType_JSON, Description: "Owner of a pipeline."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "When pipeline was updated."},
		},
	}
}

/*
// TODO - The Heroku go SDK doesn't support Pipeline list? Seems like a bug?
func listPipeline(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_pipeline.listPipeline", "connection_error", err)
		return nil, err
	}
	opts := heroku.ListRange{Field: "id"}
	items, err := conn.PipelineList(ctx, &opts)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_pipeline.listPipeline", "query_error", err, "opts", opts)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}
*/

func getPipeline(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_pipeline.getPipeline", "connection_error", err)
		return nil, err
	}
	id := d.KeyColumnQuals["id"].GetStringValue()
	item, err := conn.PipelineInfo(ctx, id)
	if err != nil {
		plugin.Logger(ctx).Error("heroku_pipeline.getPipeline", "query_error", err, "id", id)
		return nil, err
	}
	return item, err
}
