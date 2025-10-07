package sdkwa

import (
	"context"
	"fmt"
)

// Receiving methods

// ReceiveNotification retrieves a single incoming notification from the notifications queue
func (c *Client) ReceiveNotification(ctx context.Context, opts ...*RequestOptions) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := c.request(ctx, "GET", c.basePath+"/receiveNotification", nil, &result, opts...)
	return result, err
}

// DeleteNotification deletes an incoming notification from the notification queue
func (c *Client) DeleteNotification(ctx context.Context, receiptID int64, opts ...*RequestOptions) (*DeleteNotificationResponse, error) {
	var result DeleteNotificationResponse
	err := c.request(ctx, "DELETE", fmt.Sprintf("%s/deleteNotification/%d", c.basePath, receiptID), nil, &result, opts...)
	return &result, err
}

// GetChatHistory returns the message history for a specified chat
func (c *Client) GetChatHistory(ctx context.Context, params GetChatHistoryParams, opts ...*RequestOptions) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := c.request(ctx, "POST", c.basePath+"/getChatHistory", params, &result, opts...)
	return result, err
}

// Parameter types for receiving methods

// GetChatHistoryParams represents parameters for getting chat history
type GetChatHistoryParams struct {
	ChatID string `json:"chatId"`
	Count  int    `json:"count,omitempty"`
}

// Response types for receiving methods

// DeleteNotificationResponse represents the response from deleting a notification
type DeleteNotificationResponse struct {
	Result bool `json:"result"`
}
