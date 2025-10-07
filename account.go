package sdkwa

import (
	"context"
)

// Account methods

// GetSettings retrieves the current account settings
func (c *Client) GetSettings(ctx context.Context, opts ...*RequestOptions) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := c.request(ctx, "GET", c.basePath+"/getSettings", nil, &result, opts...)
	return result, err
}

// SetSettings updates the account settings
func (c *Client) SetSettings(ctx context.Context, settings map[string]interface{}, opts ...*RequestOptions) (*SetSettingsResponse, error) {
	var result SetSettingsResponse
	err := c.request(ctx, "POST", c.basePath+"/setSettings", settings, &result, opts...)
	return &result, err
}

// GetStateInstance retrieves the current state of the account instance
func (c *Client) GetStateInstance(ctx context.Context, opts ...*RequestOptions) (*StateInstanceResponse, error) {
	var result StateInstanceResponse
	err := c.request(ctx, "GET", c.basePath+"/getStateInstance", nil, &result, opts...)
	return &result, err
}

// GetWarmingPhoneStatus gets the account warming state
func (c *Client) GetWarmingPhoneStatus(ctx context.Context, opts ...*RequestOptions) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := c.request(ctx, "GET", c.basePath+"/getWarmingPhoneStatus", nil, &result, opts...)
	return result, err
}

// Reboot reboots the specified account instance
func (c *Client) Reboot(ctx context.Context, opts ...*RequestOptions) (*RebootResponse, error) {
	var result RebootResponse
	err := c.request(ctx, "GET", c.basePath+"/reboot", nil, &result, opts...)
	return &result, err
}

// Logout logs out the specified account instance
func (c *Client) Logout(ctx context.Context, opts ...*RequestOptions) (*LogoutResponse, error) {
	var result LogoutResponse
	err := c.request(ctx, "GET", c.basePath+"/logout", nil, &result, opts...)
	return &result, err
}

// GetQR returns a QR code for authorizing your account
func (c *Client) GetQR(ctx context.Context, opts ...*RequestOptions) (*QRResponse, error) {
	var result QRResponse
	err := c.request(ctx, "GET", c.basePath+"/qr", nil, &result, opts...)
	return &result, err
}

// GetAuthorizationCode gets authorization code for account authorization
func (c *Client) GetAuthorizationCode(ctx context.Context, params GetAuthorizationCodeParams, opts ...*RequestOptions) (*GetAuthorizationCodeResponse, error) {
	var result GetAuthorizationCodeResponse
	err := c.request(ctx, "POST", c.basePath+"/getAuthorizationCode", params, &result, opts...)
	return &result, err
}

// RequestRegistrationCode requests a phone number registration code via SMS or call
func (c *Client) RequestRegistrationCode(ctx context.Context, params RequestRegistrationCodeParams, opts ...*RequestOptions) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := c.request(ctx, "POST", c.basePath+"/requestRegistrationCode", params, &result, opts...)
	return result, err
}

// SendRegistrationCode sends the phone number registration code
func (c *Client) SendRegistrationCode(ctx context.Context, params SendRegistrationCodeParams, opts ...*RequestOptions) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := c.request(ctx, "POST", c.basePath+"/sendRegistrationCode", params, &result, opts...)
	return result, err
}

// Response types for account methods

// SetSettingsResponse represents the response from setting account settings
type SetSettingsResponse struct {
	SaveSettings bool `json:"saveSettings"`
}

// StateInstanceResponse represents the response from getting account state
type StateInstanceResponse struct {
	StateInstance string `json:"stateInstance"`
}

// RebootResponse represents the response from rebooting an account
type RebootResponse struct {
	IsReboot bool `json:"isReboot"`
}

// LogoutResponse represents the response from logging out an account
type LogoutResponse struct {
	IsLogout bool `json:"isLogout"`
}

// QRResponse represents the response from getting QR code
type QRResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// GetAuthorizationCodeParams represents parameters for getting authorization code
type GetAuthorizationCodeParams struct {
	PhoneNumber int64 `json:"phoneNumber"`
}

// GetAuthorizationCodeResponse represents the response from getting authorization code
type GetAuthorizationCodeResponse struct {
	Status bool   `json:"status"`
	Code   string `json:"code"`
}

// RequestRegistrationCodeParams represents parameters for requesting registration code
type RequestRegistrationCodeParams struct {
	PhoneNumber int64  `json:"phoneNumber"`
	Method      string `json:"method"` // "sms" or "voice"
}

// SendRegistrationCodeParams represents parameters for sending registration code
type SendRegistrationCodeParams struct {
	Code string `json:"code"`
}
