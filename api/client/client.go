package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudquery/cloudquery/plugins/destination/api/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/rs/zerolog"
)

type Client struct {
	plugin.UnimplementedSource
	httpClient *http.Client
	logger     zerolog.Logger
	spec       *spec.Spec
}

func New(_ context.Context, logger zerolog.Logger, s []byte, opts plugin.NewClientOptions) (plugin.Client, error) {
	c := &Client{
		logger: logger.With().Str("module", "api").Logger(),
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
	if opts.NoConnection {
		return c, nil
	}

	if err := json.Unmarshal(s, &c.spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal api spec: %w", err)
	}
	if err := c.spec.Validate(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) Close(ctx context.Context) error {
	// No resources to close in this example
	return nil
}
