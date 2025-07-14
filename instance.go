package sdkwa

import (
	"context"
)

// Instance Management methods (user-level)

// GetInstances retrieves all account instances created by the user
func (c *Client) GetInstances(ctx context.Context) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := c.requestWithUserAuth(ctx, "POST", "/api/v1/instance/user/instances/list", nil, &result)
	return result, err
}

// CreateInstance creates a new user instance with the specified tariff and period
func (c *Client) CreateInstance(ctx context.Context, params CreateInstanceParams) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := c.requestWithUserAuth(ctx, "POST", "/api/v1/instance/user/instance/createByOrder", params, &result)
	return result, err
}

// ExtendInstance renews a paid user instance for the specified period and tariff
func (c *Client) ExtendInstance(ctx context.Context, params ExtendInstanceParams) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := c.requestWithUserAuth(ctx, "POST", "/api/v1/instance/user/instance/extendByOrder", params, &result)
	return result, err
}

// DeleteInstance deletes a user instance by its ID
func (c *Client) DeleteInstance(ctx context.Context, idInstance int64) (map[string]interface{}, error) {
	var result map[string]interface{}
	params := map[string]int64{"idInstance": idInstance}
	err := c.requestWithUserAuth(ctx, "POST", "/api/v1/instance/user/instance/delete", params, &result)
	return result, err
}

// RestoreInstance restores a user instance by its ID
func (c *Client) RestoreInstance(ctx context.Context, idInstance int64) (map[string]interface{}, error) {
	var result map[string]interface{}
	params := map[string]int64{"idInstance": idInstance}
	err := c.requestWithUserAuth(ctx, "POST", "/api/v1/instance/user/instance/restore", params, &result)
	return result, err
}

// Parameter types for instance management methods

// CreateInstanceParams represents parameters for creating an instance
type CreateInstanceParams struct {
	Tariff      string `json:"tariff"`
	Period      string `json:"period"`
	PaymentType string `json:"paymentType,omitempty"`
}

// ExtendInstanceParams represents parameters for extending an instance
type ExtendInstanceParams struct {
	IDInstance  int64  `json:"idInstance"`
	Tariff      string `json:"tariff"`
	Period      string `json:"period"`
	PaymentType string `json:"paymentType,omitempty"`
}
