// Package sdkwa provides a Go client for the SDKWA WhatsApp HTTP API
package sdkwa

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
)

// Client represents the SDKWA API client
type Client struct {
	apiHost          string
	idInstance       string
	apiTokenInstance string
	messengerType    MessengerType
	userID           string
	userToken        string
	basePath         string
	httpClient       *http.Client
}

// RequestOptions contains options for individual API requests
type RequestOptions struct {
	MessengerType MessengerType // Override messenger type for this request
}

// MessengerType represents the messenger type
type MessengerType string

const (
	// MessengerWhatsApp represents WhatsApp messenger
	MessengerWhatsApp MessengerType = "whatsapp"
	// MessengerTelegram represents Telegram messenger
	MessengerTelegram MessengerType = "telegram"
)

// Options contains configuration options for the SDKWA client
type Options struct {
	APIHost            string        // API host URL, defaults to https://api.sdkwa.pro
	IDInstance         string        // Instance ID (required)
	APITokenInstance   string        // API token instance (required)
	MessengerType      MessengerType // Messenger type, defaults to whatsapp
	UserID             string        // User ID (optional, required for instance management)
	UserToken          string        // User token (optional, required for instance management)
	Timeout            time.Duration // HTTP client timeout, defaults to 30 seconds
	InsecureSkipVerify bool          // Skip TLS certificate verification
}

// NewClient creates a new SDKWA client with the provided options
func NewClient(opts Options) (*Client, error) {
	// Validate required options
	if opts.IDInstance == "" {
		return nil, errors.New("idInstance is required and must be non-empty")
	}
	if opts.APITokenInstance == "" {
		return nil, errors.New("apiTokenInstance is required and must be non-empty")
	}

	// Set defaults
	if opts.APIHost == "" {
		opts.APIHost = "https://api.sdkwa.pro"
	}
	if opts.MessengerType == "" {
		opts.MessengerType = MessengerWhatsApp
	}
	if opts.Timeout == 0 {
		opts.Timeout = 30 * time.Second
	}

	// Remove trailing slash from API host
	opts.APIHost = strings.TrimSuffix(opts.APIHost, "/")

	// Create HTTP client with custom transport
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: opts.InsecureSkipVerify,
		},
	}

	client := &Client{
		apiHost:          opts.APIHost,
		idInstance:       opts.IDInstance,
		apiTokenInstance: opts.APITokenInstance,
		messengerType:    opts.MessengerType,
		userID:           opts.UserID,
		userToken:        opts.UserToken,
		basePath:         fmt.Sprintf("/%s/%s", opts.MessengerType, opts.IDInstance),
		httpClient: &http.Client{
			Timeout:   opts.Timeout,
			Transport: transport,
		},
	}

	return client, nil
}

// ErrorResponse represents an error response from the API
type ErrorResponse struct {
	StatusCode int    `json:"statusCode,omitempty"`
	Timestamp  string `json:"timestamp,omitempty"`
	Path       string `json:"path,omitempty"`
	Message    string `json:"message"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("API error: %s", e.Message)
}

// request makes an HTTP request to the API
func (c *Client) request(ctx context.Context, method, path string, body interface{}, result interface{}, opts ...*RequestOptions) error {
	var bodyReader io.Reader
	contentType := "application/json"

	if body != nil {
		if formData, ok := body.(*bytes.Buffer); ok {
			bodyReader = formData
			contentType = "" // Will be set by multipart writer
		} else {
			jsonBody, err := json.Marshal(body)
			if err != nil {
				return fmt.Errorf("failed to marshal request body: %w", err)
			}
			bodyReader = bytes.NewReader(jsonBody)
		}
	}

	// Apply messenger type override if provided
	finalPath := path
	if len(opts) > 0 && opts[0] != nil && opts[0].MessengerType != "" {
		// Replace messenger type in path
		overrideBasePath := fmt.Sprintf("/%s/%s", opts[0].MessengerType, c.idInstance)
		finalPath = strings.Replace(path, c.basePath, overrideBasePath, 1)
	}

	fullURL := c.apiHost + finalPath
	req, err := http.NewRequestWithContext(ctx, method, fullURL, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+c.apiTokenInstance)
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		var apiError ErrorResponse
		if err := json.Unmarshal(respBody, &apiError); err != nil {
			return fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(respBody))
		}
		return &apiError
	}

	if result != nil {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("failed to unmarshal response: %w", err)
		}
	}

	return nil
}

// multipartRequest makes a multipart form request to the API
func (c *Client) multipartRequest(ctx context.Context, method, path string, fields map[string]string, files map[string]io.Reader, result interface{}, opts ...*RequestOptions) error {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// Add form fields
	for key, value := range fields {
		if err := writer.WriteField(key, value); err != nil {
			return fmt.Errorf("failed to write field %s: %w", key, err)
		}
	}

	// Add files
	for fieldName, fileReader := range files {
		part, err := writer.CreateFormFile(fieldName, "file")
		if err != nil {
			return fmt.Errorf("failed to create form file %s: %w", fieldName, err)
		}
		if _, err := io.Copy(part, fileReader); err != nil {
			return fmt.Errorf("failed to copy file data: %w", err)
		}
	}

	if err := writer.Close(); err != nil {
		return fmt.Errorf("failed to close multipart writer: %w", err)
	}

	// Apply messenger type override if provided
	finalPath := path
	if len(opts) > 0 && opts[0] != nil && opts[0].MessengerType != "" {
		// Replace messenger type in path
		overrideBasePath := fmt.Sprintf("/%s/%s", opts[0].MessengerType, c.idInstance)
		finalPath = strings.Replace(path, c.basePath, overrideBasePath, 1)
	}

	fullURL := c.apiHost + finalPath
	req, err := http.NewRequestWithContext(ctx, method, fullURL, &buf)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiTokenInstance)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		var apiError ErrorResponse
		if err := json.Unmarshal(respBody, &apiError); err != nil {
			return fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(respBody))
		}
		return &apiError
	}

	if result != nil {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("failed to unmarshal response: %w", err)
		}
	}

	return nil
}

// requestWithUserAuth makes a request with user authentication headers
func (c *Client) requestWithUserAuth(ctx context.Context, method, path string, body interface{}, result interface{}) error {
	if c.userID == "" || c.userToken == "" {
		return errors.New("userID and userToken are required for this operation")
	}

	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	fullURL := c.apiHost + path
	req, err := http.NewRequestWithContext(ctx, method, fullURL, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set user authentication headers
	req.Header.Set("x-user-id", c.userID)
	req.Header.Set("x-user-token", c.userToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		var apiError ErrorResponse
		if err := json.Unmarshal(respBody, &apiError); err != nil {
			return fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(respBody))
		}
		return &apiError
	}

	if result != nil {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("failed to unmarshal response: %w", err)
		}
	}

	return nil
}
