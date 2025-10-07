package sdkwa

import (
	"context"
	"io"
)

// Group methods

// UpdateGroupName changes the name of a group chat
func (c *Client) UpdateGroupName(ctx context.Context, groupID, groupName string, opts ...*RequestOptions) (*UpdateGroupNameResponse, error) {
	var result UpdateGroupNameResponse
	params := map[string]string{
		"groupId":   groupID,
		"groupName": groupName,
	}
	err := c.request(ctx, "POST", c.basePath+"/updateGroupName", params, &result, opts...)
	return &result, err
}

// GetGroupData retrieves information about a group chat
func (c *Client) GetGroupData(ctx context.Context, groupID string, opts ...*RequestOptions) (map[string]interface{}, error) {
	var result map[string]interface{}
	params := map[string]string{"groupId": groupID}
	err := c.request(ctx, "POST", c.basePath+"/getGroupData", params, &result, opts...)
	return result, err
}

// LeaveGroup allows the current account user to leave a specified group chat
func (c *Client) LeaveGroup(ctx context.Context, groupID string, opts ...*RequestOptions) (*LeaveGroupResponse, error) {
	var result LeaveGroupResponse
	params := map[string]string{"groupId": groupID}
	err := c.request(ctx, "POST", c.basePath+"/leaveGroup", params, &result, opts...)
	return &result, err
}

// SetGroupAdmin assigns administrator rights to a specified participant in a group chat
func (c *Client) SetGroupAdmin(ctx context.Context, groupID, participantChatID string, opts ...*RequestOptions) (*SetGroupAdminResponse, error) {
	var result SetGroupAdminResponse
	params := map[string]string{
		"groupId":           groupID,
		"participantChatId": participantChatID,
	}
	err := c.request(ctx, "POST", c.basePath+"/setGroupAdmin", params, &result, opts...)
	return &result, err
}

// RemoveGroupParticipant removes a specified participant from a group chat
func (c *Client) RemoveGroupParticipant(ctx context.Context, groupID, participantChatID string, opts ...*RequestOptions) (*RemoveGroupParticipantResponse, error) {
	var result RemoveGroupParticipantResponse
	params := map[string]string{
		"groupId":           groupID,
		"participantChatId": participantChatID,
	}
	err := c.request(ctx, "POST", c.basePath+"/removeGroupParticipant", params, &result, opts...)
	return &result, err
}

// RemoveAdmin revokes administrator rights from a specified participant in a group chat
func (c *Client) RemoveAdmin(ctx context.Context, groupID, participantChatID string, opts ...*RequestOptions) (*RemoveAdminResponse, error) {
	var result RemoveAdminResponse
	params := map[string]string{
		"groupId":           groupID,
		"participantChatId": participantChatID,
	}
	err := c.request(ctx, "POST", c.basePath+"/removeAdmin", params, &result, opts...)
	return &result, err
}

// CreateGroup creates a new group chat with the specified name and participants
func (c *Client) CreateGroup(ctx context.Context, groupName string, chatIDs []string, opts ...*RequestOptions) (*CreateGroupResponse, error) {
	var result CreateGroupResponse
	params := map[string]interface{}{
		"groupName": groupName,
		"chatIds":   chatIDs,
	}
	err := c.request(ctx, "POST", c.basePath+"/createGroup", params, &result, opts...)
	return &result, err
}

// AddGroupParticipant adds a specified participant to a group chat
func (c *Client) AddGroupParticipant(ctx context.Context, groupID, participantChatID string, opts ...*RequestOptions) (*AddGroupParticipantResponse, error) {
	var result AddGroupParticipantResponse
	params := map[string]string{
		"groupId":           groupID,
		"participantChatId": participantChatID,
	}
	err := c.request(ctx, "POST", c.basePath+"/addGroupParticipant", params, &result, opts...)
	return &result, err
}

// SetGroupPicture sets a new picture for a group chat
func (c *Client) SetGroupPicture(ctx context.Context, groupID string, file io.Reader, opts ...*RequestOptions) (*SetGroupPictureResponse, error) {
	fields := map[string]string{
		"groupId": groupID,
	}
	files := map[string]io.Reader{
		"file": file,
	}

	var result SetGroupPictureResponse
	err := c.multipartRequest(ctx, "POST", c.basePath+"/setGroupPicture", fields, files, &result, opts...)
	return &result, err
}

// Response types for group methods

// UpdateGroupNameResponse represents the response from updating group name
type UpdateGroupNameResponse struct {
	UpdateGroupName bool `json:"updateGroupName"`
}

// LeaveGroupResponse represents the response from leaving a group
type LeaveGroupResponse struct {
	LeaveGroup bool `json:"leaveGroup"`
}

// SetGroupAdminResponse represents the response from setting group admin
type SetGroupAdminResponse struct {
	SetGroupAdmin bool `json:"setGroupAdmin"`
}

// RemoveGroupParticipantResponse represents the response from removing group participant
type RemoveGroupParticipantResponse struct {
	RemoveParticipant bool `json:"removeParticipant"`
}

// RemoveAdminResponse represents the response from removing admin rights
type RemoveAdminResponse struct {
	RemoveAdmin bool `json:"removeAdmin"`
}

// CreateGroupResponse represents the response from creating a group
type CreateGroupResponse struct {
	Created         bool   `json:"created"`
	ChatID          string `json:"chatId"`
	GroupInviteLink string `json:"groupInviteLink"`
}

// AddGroupParticipantResponse represents the response from adding group participant
type AddGroupParticipantResponse struct {
	AddParticipant bool `json:"addParticipant"`
}

// SetGroupPictureResponse represents the response from setting group picture
type SetGroupPictureResponse struct {
	SetGroupPicture bool   `json:"setGroupPicture"`
	URLAvatar       string `json:"urlAvatar"`
	Reason          string `json:"reason"`
}
