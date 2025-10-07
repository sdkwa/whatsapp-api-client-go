package sdkwa

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestClient_NewClient tests the client creation
func TestClient_NewClient(t *testing.T) {
	tests := []struct {
		name    string
		opts    Options
		wantErr bool
	}{
		{
			name: "valid options",
			opts: Options{
				IDInstance:       "test-instance",
				APITokenInstance: "test-token",
			},
			wantErr: false,
		},
		{
			name: "missing id instance",
			opts: Options{
				APITokenInstance: "test-token",
			},
			wantErr: true,
		},
		{
			name: "missing api token instance",
			opts: Options{
				IDInstance: "test-instance",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewClient(tt.opts)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, client)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, client)
			}
		})
	}
}

// TestClient_DefaultValues tests default values in client creation
func TestClient_DefaultValues(t *testing.T) {
	client, err := NewClient(Options{
		IDInstance:       "test-instance",
		APITokenInstance: "test-token",
	})
	require.NoError(t, err)

	assert.Equal(t, "https://api.sdkwa.pro", client.apiHost)
	assert.Equal(t, "test-instance", client.idInstance)
	assert.Equal(t, "test-token", client.apiTokenInstance)
	assert.Equal(t, MessengerWhatsApp, client.messengerType)
	assert.Equal(t, "/whatsapp/test-instance", client.basePath)
}

// TestClient_TelegramMessengerType tests Telegram messenger type
func TestClient_TelegramMessengerType(t *testing.T) {
	client, err := NewClient(Options{
		IDInstance:       "test-instance",
		APITokenInstance: "test-token",
		MessengerType:    MessengerTelegram,
	})
	require.NoError(t, err)

	assert.Equal(t, MessengerTelegram, client.messengerType)
	assert.Equal(t, "/telegram/test-instance", client.basePath)
}

// TestClient_CustomAPIHost tests custom API host
func TestClient_CustomAPIHost(t *testing.T) {
	client, err := NewClient(Options{
		APIHost:          "https://custom.api.com/",
		IDInstance:       "test-instance",
		APITokenInstance: "test-token",
	})
	require.NoError(t, err)

	assert.Equal(t, "https://custom.api.com", client.apiHost)
}

// Integration test - only runs if environment variables are set
func TestClient_Integration(t *testing.T) {
	apiHost := os.Getenv("SDKWA_API_HOST")
	idInstance := os.Getenv("SDKWA_ID_INSTANCE")
	apiToken := os.Getenv("SDKWA_API_TOKEN")

	if apiHost == "" || idInstance == "" || apiToken == "" {
		t.Skip("Skipping integration test - environment variables not set")
	}

	client, err := NewClient(Options{
		APIHost:          apiHost,
		IDInstance:       idInstance,
		APITokenInstance: apiToken,
	})
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Test getting state instance
	state, err := client.GetStateInstance(ctx)
	require.NoError(t, err)
	assert.NotEmpty(t, state.StateInstance)

	// Test getting settings
	settings, err := client.GetSettings(ctx)
	require.NoError(t, err)
	assert.NotEmpty(t, settings)
}

// TestWebhookHandler tests webhook handler functionality
func TestWebhookHandler(t *testing.T) {
	handler := NewWebhookHandler()

	// Test callback registration
	called := false
	handler.OnIncomingMessageText(func(data map[string]interface{}) error {
		called = true
		return nil
	})

	// Test webhook handling
	webhookData := map[string]interface{}{
		"typeWebhook": "incomingMessageReceived",
		"messageData": map[string]interface{}{
			"typeMessage": "textMessage",
		},
	}

	err := handler.HandleWebhook(webhookData)
	assert.NoError(t, err)
	assert.True(t, called)
}

// TestSendMessageParams tests parameter structures
func TestSendMessageParams(t *testing.T) {
	params := SendMessageParams{
		ChatID:  "test@c.us",
		Message: "Hello, World!",
	}

	assert.Equal(t, "test@c.us", params.ChatID)
	assert.Equal(t, "Hello, World!", params.Message)
}
