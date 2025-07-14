package sdkwa

import (
	"context"
	"io"
)

// Sending methods

// SendMessage sends a text message to either a personal or group chat
func (c *Client) SendMessage(ctx context.Context, params SendMessageParams) (*SendMessageResponse, error) {
	var result SendMessageResponse
	err := c.request(ctx, "POST", c.basePath+"/sendMessage", params, &result)
	return &result, err
}

// SendContact sends a contact card message to a chat
func (c *Client) SendContact(ctx context.Context, params SendContactParams) (*SendContactResponse, error) {
	var result SendContactResponse
	err := c.request(ctx, "POST", c.basePath+"/sendContact", params, &result)
	return &result, err
}

// SendFileByUpload sends a file by uploading it using form-data
func (c *Client) SendFileByUpload(ctx context.Context, params SendFileByUploadParams) (*SendFileByUploadResponse, error) {
	fields := map[string]string{
		"chatId": params.ChatID,
	}

	if params.Caption != "" {
		fields["caption"] = params.Caption
	}
	if params.QuotedMessageID != "" {
		fields["quotedMessageId"] = params.QuotedMessageID
	}

	files := map[string]io.Reader{
		"file": params.File,
	}

	var result SendFileByUploadResponse
	err := c.multipartRequest(ctx, "POST", c.basePath+"/sendFileByUpload", fields, files, &result)
	return &result, err
}

// SendFileByURL sends a file by providing its URL
func (c *Client) SendFileByURL(ctx context.Context, params SendFileByURLParams) (*SendFileByURLResponse, error) {
	var result SendFileByURLResponse
	err := c.request(ctx, "POST", c.basePath+"/sendFileByUrl", params, &result)
	return &result, err
}

// SendLocation sends a location message to a chat
func (c *Client) SendLocation(ctx context.Context, params SendLocationParams) (*SendLocationResponse, error) {
	var result SendLocationResponse
	err := c.request(ctx, "POST", c.basePath+"/sendLocation", params, &result)
	return &result, err
}

// UploadFile uploads a file to storage for later sending
func (c *Client) UploadFile(ctx context.Context, file io.Reader) (*UploadFileResponse, error) {
	files := map[string]io.Reader{
		"file": file,
	}

	var result UploadFileResponse
	err := c.multipartRequest(ctx, "POST", c.basePath+"/uploadFile", nil, files, &result)
	return &result, err
}

// Parameter types for sending methods

// SendMessageParams represents parameters for sending a text message
type SendMessageParams struct {
	ChatID          string `json:"chatId"`
	Message         string `json:"message"`
	QuotedMessageID string `json:"quotedMessageId,omitempty"`
	ArchiveChat     bool   `json:"archiveChat,omitempty"`
	LinkPreview     bool   `json:"linkPreview,omitempty"`
}

// SendContactParams represents parameters for sending a contact
type SendContactParams struct {
	ChatID          string  `json:"chatId"`
	Contact         Contact `json:"contact"`
	QuotedMessageID string  `json:"quotedMessageId,omitempty"`
}

// Contact represents a contact information
type Contact struct {
	PhoneContact int64  `json:"phoneContact"`
	FirstName    string `json:"firstName,omitempty"`
	MiddleName   string `json:"middleName,omitempty"`
	LastName     string `json:"lastName,omitempty"`
	Company      string `json:"company,omitempty"`
}

// SendFileByUploadParams represents parameters for sending a file by upload
type SendFileByUploadParams struct {
	ChatID          string    `json:"chatId"`
	File            io.Reader `json:"-"` // File content
	FileName        string    `json:"fileName"`
	Caption         string    `json:"caption,omitempty"`
	QuotedMessageID string    `json:"quotedMessageId,omitempty"`
}

// SendFileByURLParams represents parameters for sending a file by URL
type SendFileByURLParams struct {
	ChatID          string `json:"chatId"`
	URLFile         string `json:"urlFile"`
	FileName        string `json:"fileName"`
	Caption         string `json:"caption,omitempty"`
	QuotedMessageID string `json:"quotedMessageId,omitempty"`
	ArchiveChat     bool   `json:"archiveChat,omitempty"`
}

// SendLocationParams represents parameters for sending a location
type SendLocationParams struct {
	ChatID          string  `json:"chatId"`
	NameLocation    string  `json:"nameLocation,omitempty"`
	Address         string  `json:"address,omitempty"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	QuotedMessageID string  `json:"quotedMessageId,omitempty"`
}

// Response types for sending methods

// SendMessageResponse represents the response from sending a message
type SendMessageResponse struct {
	IDMessage string `json:"idMessage"`
}

// SendContactResponse represents the response from sending a contact
type SendContactResponse struct {
	IDMessage string `json:"idMessage"`
}

// SendFileByUploadResponse represents the response from sending a file by upload
type SendFileByUploadResponse struct {
	IDMessage string `json:"idMessage"`
}

// SendFileByURLResponse represents the response from sending a file by URL
type SendFileByURLResponse struct {
	IDMessage string `json:"idMessage"`
}

// SendLocationResponse represents the response from sending a location
type SendLocationResponse struct {
	IDMessage string `json:"idMessage"`
}

// UploadFileResponse represents the response from uploading a file
type UploadFileResponse struct {
	URLFile string `json:"urlFile"`
}
