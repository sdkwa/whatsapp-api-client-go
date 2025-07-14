package sdkwa

import (
	"context"
)

// Queue methods

// ClearMessagesQueue clears all pending messages from the sending queue
func (c *Client) ClearMessagesQueue(ctx context.Context) (*ClearMessagesQueueResponse, error) {
	var result ClearMessagesQueueResponse
	err := c.request(ctx, "GET", c.basePath+"/clearMessagesQueue", nil, &result)
	return &result, err
}

// ShowMessagesQueue retrieves a list of messages currently in the sending queue
func (c *Client) ShowMessagesQueue(ctx context.Context) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := c.request(ctx, "GET", c.basePath+"/showMessagesQueue", nil, &result)
	return result, err
}

// Response types for queue methods

// ClearMessagesQueueResponse represents the response from clearing messages queue
type ClearMessagesQueueResponse struct {
	IsCleared bool `json:"isCleared"`
}
