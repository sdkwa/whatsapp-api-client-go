# WhatsApp API Client Go

A Go client library for the SDKWA WhatsApp HTTP API.

## Installation

```bash
go get github.com/sdkwa/whatsapp-api-client-go
```

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sdkwa/whatsapp-api-client-go"
)

func main() {
	// Create client
	client, err := sdkwa.NewClient(sdkwa.Options{
		IDInstance:       "your_instance_id",
		APITokenInstance: "your_api_token",
	})
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	// Send a message
	response, err := client.SendMessage(ctx, sdkwa.SendMessageParams{
		ChatID:  "79999999999@c.us",
		Message: "Hello from Go SDK!",
	})
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	fmt.Printf("Message sent with ID: %s\n", response.IDMessage)
}
```

## Features

### Account Management
- Get/Set account settings
- Get account state
- Reboot/Logout account
- QR code authorization
- Phone number authorization

### Messaging
- Send text messages
- Send files (by upload or URL)
- Send contacts
- Send locations
- Upload files

### Receiving
- Receive notifications
- Get chat history
- Delete notifications

### Chat Management
- Get contacts and chats
- Set profile picture/name/status
- Get avatar
- Check WhatsApp availability
- Mark messages as read
- Archive/Unarchive chats
- Delete messages

### Group Management
- Create groups
- Update group name
- Get group data
- Add/Remove participants
- Set/Remove admin rights
- Leave group
- Set group picture

### Queue Management
- Show messages queue
- Clear messages queue

### Instance Management (with user credentials)
- Get instances
- Create instances
- Extend instances
- Delete instances
- Restore instances

### Webhooks & Real-time Events
- HTTP webhook handler
- WebSocket client
- Notification polling
- Event callbacks

## Configuration

The client accepts the following options:

```go
client, err := sdkwa.NewClient(sdkwa.Options{
	APIHost:            "https://api.sdkwa.pro", // Optional, defaults to official API
	IDInstance:         "your_instance_id",      // Required
	APITokenInstance:   "your_api_token",        // Required
	UserID:             "your_user_id",          // Optional, for instance management
	UserToken:          "your_user_token",       // Optional, for instance management
	Timeout:            30 * time.Second,        // Optional, HTTP timeout
	InsecureSkipVerify: false,                   // Optional, skip TLS verification
})
```

## Environment Variables

You can use environment variables for configuration:

- `SDKWA_API_HOST` - API host URL
- `SDKWA_ID_INSTANCE` - Instance ID
- `SDKWA_API_TOKEN` - API token
- `SDKWA_USER_ID` - User ID (for instance management)
- `SDKWA_USER_TOKEN` - User token (for instance management)

## Examples

### Send Message

```go
response, err := client.SendMessage(ctx, sdkwa.SendMessageParams{
	ChatID:  "79999999999@c.us",
	Message: "Hello World!",
})
```

### Send File by Upload

```go
file, err := os.Open("document.pdf")
if err != nil {
	log.Fatal(err)
}
defer file.Close()

response, err := client.SendFileByUpload(ctx, sdkwa.SendFileByUploadParams{
	ChatID:   "79999999999@c.us",
	File:     file,
	FileName: "document.pdf",
	Caption:  "Here's the document",
})
```

### Create Group

```go
response, err := client.CreateGroup(ctx, "My Group", []string{
	"79999999999@c.us",
	"79999999998@c.us",
})
```

### Handle Webhooks

```go
handler := sdkwa.NewWebhookHandler()

// Register message handler
handler.OnIncomingMessageText(func(data map[string]interface{}) error {
	fmt.Printf("Received message: %+v\n", data)
	return nil
})

// Use as HTTP handler
http.Handle("/webhook", handler)
http.ListenAndServe(":8080", nil)
```

### WebSocket Real-time Events

```go
handler := sdkwa.NewWebhookHandler()
handler.OnIncomingMessageText(func(data map[string]interface{}) error {
	fmt.Printf("Real-time message: %+v\n", data)
	return nil
})

wsClient := client.NewWebSocketClient(handler)
if err := wsClient.Connect(ctx); err != nil {
	log.Fatal(err)
}

// Listen for messages
if err := wsClient.Listen(ctx); err != nil {
	log.Fatal(err)
}
```

### Polling for Notifications

```go
handler := sdkwa.NewWebhookHandler()
handler.OnIncomingMessageText(func(data map[string]interface{}) error {
	fmt.Printf("Polled message: %+v\n", data)
	return nil
})

// Start polling (blocks until context is cancelled)
err := client.StartReceivingNotifications(ctx, handler)
```

## Error Handling

The client returns structured errors:

```go
response, err := client.SendMessage(ctx, params)
if err != nil {
	if apiError, ok := err.(*sdkwa.ErrorResponse); ok {
		fmt.Printf("API Error: %s (Status: %d)\n", apiError.Message, apiError.StatusCode)
	} else {
		fmt.Printf("Network Error: %v\n", err)
	}
}
```

## Testing

Run tests with:

```bash
go test ./...
```

For integration tests, set environment variables:

```bash
export SDKWA_API_HOST="https://api.sdkwa.pro"
export SDKWA_ID_INSTANCE="your_instance_id"
export SDKWA_API_TOKEN="your_api_token"
go test ./...
```

## Examples

See the `examples/` directory for complete working examples:

- `send_message/` - Basic message sending
- `get_qr/` - QR code authorization
- `webhook_server/` - HTTP webhook server
- `create_group/` - Group creation
- `instance_management/` - Instance management

## License

MIT License

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
