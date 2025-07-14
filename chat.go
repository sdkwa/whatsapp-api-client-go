package sdkwa

import (
	"context"
	"io"
)

// Chat/Contact methods

// GetContacts retrieves a list of contacts for the current account
func (c *Client) GetContacts(ctx context.Context) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := c.request(ctx, "GET", c.basePath+"/getContacts", nil, &result)
	return result, err
}

// GetChats retrieves a list of all chats for the current account
func (c *Client) GetChats(ctx context.Context) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := c.request(ctx, "GET", c.basePath+"/getChats", nil, &result)
	return result, err
}

// GetContactInfo retrieves detailed information about a contact
func (c *Client) GetContactInfo(ctx context.Context, chatID string) (map[string]interface{}, error) {
	var result map[string]interface{}
	params := map[string]string{"chatId": chatID}
	err := c.request(ctx, "GET", c.basePath+"/getContactInfo", params, &result)
	return result, err
}

// SetProfilePicture sets a new profile picture for the account
func (c *Client) SetProfilePicture(ctx context.Context, file io.Reader) (*SetProfilePictureResponse, error) {
	files := map[string]io.Reader{
		"file": file,
	}

	var result SetProfilePictureResponse
	err := c.multipartRequest(ctx, "POST", c.basePath+"/setProfilePicture", nil, files, &result)
	return &result, err
}

// SetProfileName sets a new profile name for the account
func (c *Client) SetProfileName(ctx context.Context, name string) error {
	params := map[string]string{"name": name}
	return c.request(ctx, "POST", c.basePath+"/setProfileName", params, nil)
}

// SetProfileStatus sets a new status message for the account
func (c *Client) SetProfileStatus(ctx context.Context, status string) error {
	params := map[string]string{"status": status}
	return c.request(ctx, "POST", c.basePath+"/setProfileStatus", params, nil)
}

// GetAvatar returns the avatar URL for a user or group chat
func (c *Client) GetAvatar(ctx context.Context, chatID string) (map[string]interface{}, error) {
	var result map[string]interface{}
	params := map[string]string{"chatId": chatID}
	err := c.request(ctx, "POST", c.basePath+"/getAvatar", params, &result)
	return result, err
}

// CheckWhatsApp checks if a WhatsApp account exists for the specified phone number
func (c *Client) CheckWhatsApp(ctx context.Context, phoneNumber int64) (*CheckWhatsAppResponse, error) {
	var result CheckWhatsAppResponse
	params := map[string]int64{"phoneNumber": phoneNumber}
	err := c.request(ctx, "POST", c.basePath+"/checkWhatsapp", params, &result)
	return &result, err
}

// ReadChat marks messages in a chat as read
func (c *Client) ReadChat(ctx context.Context, params ReadChatParams) (*ReadChatResponse, error) {
	var result ReadChatResponse
	err := c.request(ctx, "POST", c.basePath+"/readChat", params, &result)
	return &result, err
}

// ArchiveChat archives a chat
func (c *Client) ArchiveChat(ctx context.Context, chatID string) error {
	params := map[string]string{"chatId": chatID}
	return c.request(ctx, "POST", c.basePath+"/archiveChat", params, nil)
}

// UnarchiveChat unarchives a chat
func (c *Client) UnarchiveChat(ctx context.Context, chatID string) error {
	params := map[string]string{"chatId": chatID}
	return c.request(ctx, "POST", c.basePath+"/unarchiveChat", params, nil)
}

// DeleteMessage deletes a message from a chat
func (c *Client) DeleteMessage(ctx context.Context, chatID, messageID string) error {
	params := map[string]string{
		"chatId":    chatID,
		"idMessage": messageID,
	}
	return c.request(ctx, "POST", c.basePath+"/deleteMessage", params, nil)
}

// Parameter types for chat/contact methods

// ReadChatParams represents parameters for marking chat messages as read
type ReadChatParams struct {
	ChatID    string `json:"chatId"`
	IDMessage string `json:"idMessage,omitempty"`
}

// Response types for chat/contact methods

// SetProfilePictureResponse represents the response from setting profile picture
type SetProfilePictureResponse struct {
	SetProfilePicture bool   `json:"setProfilePicture"`
	URLAvatar         string `json:"urlAvatar"`
	Reason            string `json:"reason"`
}

// CheckWhatsAppResponse represents the response from checking WhatsApp availability
type CheckWhatsAppResponse struct {
	ExistsWhatsApp bool `json:"existsWhatsapp"`
}

// ReadChatResponse represents the response from marking chat messages as read
type ReadChatResponse struct {
	SetRead bool `json:"setRead"`
}
