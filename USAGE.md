# SDKWA WhatsApp API Client for Go

A comprehensive Go wrapper library for the SDKWA WhatsApp HTTP API, providing easy access to WhatsApp messaging functionality.

## Quick Setup Guide

### 1. Installation

```bash
go get github.com/sdkwa/whatsapp-api-client-go
```

### 2. Basic Usage

```go
package main

import (
    "context"
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
        log.Fatal(err)
    }

    // Send a message
    response, err := client.SendMessage(context.Background(), sdkwa.SendMessageParams{
        ChatID:  "1234567890@c.us",
        Message: "Hello from Go!",
    })
    if err != nil {
        log.Fatal(err)
    }
    
    log.Printf("Message sent with ID: %s", response.IDMessage)
}
```

### 3. Environment Variables

Set these environment variables for easy configuration:

```bash
export SDKWA_API_HOST="https://api.sdkwa.pro"
export SDKWA_ID_INSTANCE="your_instance_id"
export SDKWA_API_TOKEN="your_api_token"
export SDKWA_USER_ID="your_user_id"        # Optional: for instance management
export SDKWA_USER_TOKEN="your_user_token"  # Optional: for instance management
```

### 4. Available Features

#### Account Management
- ✅ Get/Set account settings
- ✅ Get account state
- ✅ Reboot/Logout account
- ✅ QR code authorization
- ✅ Phone number authorization

#### Messaging
- ✅ Send text messages
- ✅ Send files (upload & URL)
- ✅ Send contacts
- ✅ Send locations
- ✅ Upload files

#### Receiving
- ✅ Receive notifications
- ✅ Get chat history
- ✅ Delete notifications
- ✅ Notification polling

#### Chat Management
- ✅ Get contacts and chats
- ✅ Set profile picture/name/status
- ✅ Check WhatsApp availability
- ✅ Mark messages as read
- ✅ Archive/Unarchive chats
- ✅ Delete messages

#### Group Management
- ✅ Create groups
- ✅ Update group name/picture
- ✅ Add/Remove participants
- ✅ Set/Remove admin rights
- ✅ Leave group

#### Queue Management
- ✅ Show/Clear message queue

#### Instance Management
- ✅ Get instances
- ✅ Create instances
- ✅ Extend instances
- ✅ Delete instances
- ✅ Restore instances

#### Real-time Events
- ✅ HTTP webhook handler
- ✅ WebSocket client
- ✅ Event callbacks

### 5. Examples

Run the example programs:

```bash
# Send a message
./bin/send_message.exe

# Get QR code for authorization
./bin/get_qr.exe

# Start webhook server
./bin/webhook_server.exe

# Create a group
./bin/create_group.exe

# Manage instances
./bin/instance_management.exe
```

### 6. Package Structure

```
github.com/sdkwa/whatsapp-api-client-go/
├── client.go         # Main client and HTTP handling
├── account.go        # Account management methods
├── sending.go        # Message sending methods
├── receiving.go      # Message receiving methods
├── chat.go          # Chat management methods
├── group.go         # Group management methods
├── queue.go         # Queue management methods
├── instance.go      # Instance management methods
├── webhook.go       # Webhook and WebSocket handling
├── client_test.go   # Tests
└── examples/        # Example applications
```

### 7. Use in Your Project

```go
// go.mod
module your-project

go 1.19

require github.com/sdkwa/whatsapp-api-client-go v1.0.0
```

```go
// main.go
package main

import (
    "context"
    "log"
    
    "github.com/sdkwa/whatsapp-api-client-go"
)

func main() {
    client, err := sdkwa.NewClient(sdkwa.Options{
        IDInstance:       "your_instance_id",
        APITokenInstance: "your_api_token",
    })
    if err != nil {
        log.Fatal(err)
    }
    
    // Your WhatsApp automation code here
    // ... use client methods
}
```

### 8. Testing

```bash
# Run unit tests
go test -v ./...

# Run with coverage
go test -v -cover ./...

# Run integration tests (with environment variables set)
go test -v ./... -tags=integration
```

### 9. Build & Distribution

```bash
# Build for current platform
go build -o myapp ./

# Cross-compile for different platforms
GOOS=linux GOARCH=amd64 go build -o myapp-linux ./
GOOS=windows GOARCH=amd64 go build -o myapp.exe ./
GOOS=darwin GOARCH=amd64 go build -o myapp-mac ./
```

This Go library provides a complete 1:1 mapping of the JavaScript/TypeScript SDKWA library functionality, making it easy to integrate WhatsApp messaging into your Go applications.
