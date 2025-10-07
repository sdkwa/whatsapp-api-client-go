package sdkwa

import "context"

// Telegram-specific methods

// CreateApp creates a new Telegram application
func (c *Client) CreateApp(ctx context.Context, params CreateAppParams, opts ...*RequestOptions) (*CreateAppResponse, error) {
	var result CreateAppResponse
	err := c.request(ctx, "POST", c.basePath+"/createApp", params, &result, opts...)
	return &result, err
}

// SendConfirmationCode sends a confirmation code for Telegram account authorization
func (c *Client) SendConfirmationCode(ctx context.Context, params SendConfirmationCodeParams, opts ...*RequestOptions) (*SendConfirmationCodeResponse, error) {
	var result SendConfirmationCodeResponse
	err := c.request(ctx, "POST", c.basePath+"/sendConfirmationCode", params, &result, opts...)
	return &result, err
}

// SignInWithConfirmationCode signs in using a confirmation code for Telegram
func (c *Client) SignInWithConfirmationCode(ctx context.Context, params SignInWithConfirmationCodeParams, opts ...*RequestOptions) (*SignInWithConfirmationCodeResponse, error) {
	var result SignInWithConfirmationCodeResponse
	err := c.request(ctx, "POST", c.basePath+"/signInWithConfirmationCode", params, &result, opts...)
	return &result, err
}

// Parameter types for Telegram methods

// CreateAppParams represents parameters for creating a Telegram app
type CreateAppParams struct {
	Title       string `json:"title"`
	ShortName   string `json:"shortName"`
	URL         string `json:"url"`
	Description string `json:"description,omitempty"`
}

// SendConfirmationCodeParams represents parameters for sending confirmation code
type SendConfirmationCodeParams struct {
	PhoneNumber int64 `json:"phoneNumber"`
}

// SignInWithConfirmationCodeParams represents parameters for signing in with confirmation code
type SignInWithConfirmationCodeParams struct {
	Code string `json:"code"`
}

// Response types for Telegram methods

// CreateAppResponse represents the response from creating a Telegram app
type CreateAppResponse struct {
	Result bool `json:"result"`
	Data   struct {
		AppID string `json:"appId"`
	} `json:"data"`
}

// SendConfirmationCodeResponse represents the response from sending confirmation code
type SendConfirmationCodeResponse struct {
	Result  bool   `json:"result"`
	Message string `json:"message"`
}

// SignInWithConfirmationCodeResponse represents the response from signing in with confirmation code
type SignInWithConfirmationCodeResponse struct {
	Result  bool   `json:"result"`
	Message string `json:"message"`
}
