package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

// Helper function to send data to API
func (c *Client) sendToAPI(endpoint string, data []byte) error {
	req, err := http.NewRequest("POST", c.spec.BaseURL+endpoint, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Set additional headers from the config
	for key, value := range c.spec.Headers {
		req.Header.Set(key, value)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("received error response: %s", resp.Status)
	}

	return nil
}

// WriteTable sends records to the /append endpoint
func (c *Client) WriteTable(ctx context.Context, msg *message.WriteInsert) error {
	data, err := json.Marshal(msg.Record)
	if err != nil {
		return fmt.Errorf("failed to marshal record: %w", err)
	}

	if err := c.sendToAPI("/append", data); err != nil {
		return err
	}

	return nil
}

// MigrateTable sends table schema to the /migrate endpoint
func (c *Client) MigrateTable(ctx context.Context, table *schema.Table) error {
	schemaData, err := json.Marshal(table)
	if err != nil {
		return fmt.Errorf("failed to marshal schema: %w", err)
	}

	if err := c.sendToAPI("/migrate", schemaData); err != nil {
		return err
	}

	return nil
}

// Write processes messages from the channel and writes them to the API
func (c *Client) Write(ctx context.Context, msgs <-chan message.WriteMessage) error {
	for msg := range msgs {
		switch m := msg.(type) {
		case *message.WriteInsert:
			// Migrate the table schema before writing data
			if err := c.MigrateTable(ctx, m.GetTable()); err != nil {
				return err
			}

			if err := c.WriteTable(ctx, m); err != nil {
				return err
			}
		case *message.WriteMigrateTable:
			if err := c.MigrateTable(ctx, m.GetTable()); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported message type: %T", msg)
		}
	}
	return nil
}
