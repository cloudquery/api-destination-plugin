package client

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func (c *Client) Read(_ context.Context, table *schema.Table, res chan<- arrow.Record) error {
	return nil
}
