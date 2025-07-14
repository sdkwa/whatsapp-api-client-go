package sdkwa

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

// WebhookType represents the type of webhook event
type WebhookType string

const (
	WebhookTypeStateInstanceChanged                       WebhookType = "stateInstanceChanged"
	WebhookTypeOutgoingMessageStatus                      WebhookType = "outgoingMessageStatus"
	WebhookTypeIncomingMessageReceivedTextMessage         WebhookType = "incomingMessageReceived_textMessage"
	WebhookTypeIncomingMessageReceivedImageMessage        WebhookType = "incomingMessageReceived_imageMessage"
	WebhookTypeIncomingMessageReceivedLocationMessage     WebhookType = "incomingMessageReceived_locationMessage"
	WebhookTypeIncomingMessageReceivedContactMessage      WebhookType = "incomingMessageReceived_contactMessage"
	WebhookTypeIncomingMessageReceivedExtendedTextMessage WebhookType = "incomingMessageReceived_extendedTextMessage"
	WebhookTypeDeviceInfo                                 WebhookType = "deviceInfo"
)

// WebhookCallback represents a callback function for webhook events
type WebhookCallback func(data map[string]interface{}) error

// WebhookHandler handles webhook events from the SDKWA API
type WebhookHandler struct {
	callbacks map[WebhookType]WebhookCallback
}

// NewWebhookHandler creates a new webhook handler
func NewWebhookHandler() *WebhookHandler {
	return &WebhookHandler{
		callbacks: make(map[WebhookType]WebhookCallback),
	}
}

// OnStateInstance registers a callback for state instance changed events
func (w *WebhookHandler) OnStateInstance(callback WebhookCallback) {
	w.callbacks[WebhookTypeStateInstanceChanged] = callback
}

// OnOutgoingMessageStatus registers a callback for outgoing message status events
func (w *WebhookHandler) OnOutgoingMessageStatus(callback WebhookCallback) {
	w.callbacks[WebhookTypeOutgoingMessageStatus] = callback
}

// OnIncomingMessageText registers a callback for incoming text message events
func (w *WebhookHandler) OnIncomingMessageText(callback WebhookCallback) {
	w.callbacks[WebhookTypeIncomingMessageReceivedTextMessage] = callback
}

// OnIncomingMessageFile registers a callback for incoming file message events
func (w *WebhookHandler) OnIncomingMessageFile(callback WebhookCallback) {
	w.callbacks[WebhookTypeIncomingMessageReceivedImageMessage] = callback
}

// OnIncomingMessageLocation registers a callback for incoming location message events
func (w *WebhookHandler) OnIncomingMessageLocation(callback WebhookCallback) {
	w.callbacks[WebhookTypeIncomingMessageReceivedLocationMessage] = callback
}

// OnIncomingMessageContact registers a callback for incoming contact message events
func (w *WebhookHandler) OnIncomingMessageContact(callback WebhookCallback) {
	w.callbacks[WebhookTypeIncomingMessageReceivedContactMessage] = callback
}

// OnIncomingMessageExtendedText registers a callback for incoming extended text message events
func (w *WebhookHandler) OnIncomingMessageExtendedText(callback WebhookCallback) {
	w.callbacks[WebhookTypeIncomingMessageReceivedExtendedTextMessage] = callback
}

// OnDeviceInfo registers a callback for device info events
func (w *WebhookHandler) OnDeviceInfo(callback WebhookCallback) {
	w.callbacks[WebhookTypeDeviceInfo] = callback
}

// HandleWebhook processes a webhook request
func (w *WebhookHandler) HandleWebhook(data map[string]interface{}) error {
	var webhookType WebhookType

	// Determine webhook type
	if typeWebhook, ok := data["typeWebhook"].(string); ok {
		if messageData, ok := data["messageData"].(map[string]interface{}); ok {
			if typeMessage, ok := messageData["typeMessage"].(string); ok {
				webhookType = WebhookType(fmt.Sprintf("%s_%s", typeWebhook, typeMessage))
			}
		} else {
			webhookType = WebhookType(typeWebhook)
		}
	}

	// Execute callback if registered
	if callback, exists := w.callbacks[webhookType]; exists {
		return callback(data)
	}

	return nil
}

// ServeHTTP implements http.Handler interface for webhook handling
func (w *WebhookHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(rw, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := w.HandleWebhook(data); err != nil {
		log.Printf("Error handling webhook: %v", err)
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

// WebSocketClient handles WebSocket connections for real-time events
type WebSocketClient struct {
	conn     *websocket.Conn
	client   *Client
	handler  *WebhookHandler
	stopChan chan struct{}
}

// NewWebSocketClient creates a new WebSocket client
func (c *Client) NewWebSocketClient(handler *WebhookHandler) *WebSocketClient {
	return &WebSocketClient{
		client:   c,
		handler:  handler,
		stopChan: make(chan struct{}),
	}
}

// Connect establishes a WebSocket connection
func (ws *WebSocketClient) Connect(ctx context.Context) error {
	// Convert HTTP(S) URL to WebSocket URL
	wsURL := strings.Replace(ws.client.apiHost, "http://", "ws://", 1)
	wsURL = strings.Replace(wsURL, "https://", "wss://", 1)
	wsURL = fmt.Sprintf("%s/ws/%s", wsURL, ws.client.idInstance)

	// Add authorization as query parameter
	u, err := url.Parse(wsURL)
	if err != nil {
		return fmt.Errorf("invalid WebSocket URL: %w", err)
	}

	q := u.Query()
	q.Set("token", ws.client.apiTokenInstance)
	u.RawQuery = q.Encode()

	// Create WebSocket connection
	dialer := websocket.Dialer{
		HandshakeTimeout: 10 * time.Second,
	}

	conn, _, err := dialer.DialContext(ctx, u.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to connect to WebSocket: %w", err)
	}

	ws.conn = conn
	return nil
}

// Listen starts listening for WebSocket messages
func (ws *WebSocketClient) Listen(ctx context.Context) error {
	if ws.conn == nil {
		return fmt.Errorf("WebSocket connection not established")
	}

	defer ws.conn.Close()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ws.stopChan:
			return nil
		default:
			var message map[string]interface{}
			if err := ws.conn.ReadJSON(&message); err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					return fmt.Errorf("WebSocket error: %w", err)
				}
				return nil
			}

			// Handle the message using the webhook handler
			if ws.handler != nil {
				if err := ws.handler.HandleWebhook(message); err != nil {
					log.Printf("Error handling WebSocket message: %v", err)
				}
			}
		}
	}
}

// Close closes the WebSocket connection
func (ws *WebSocketClient) Close() error {
	close(ws.stopChan)
	if ws.conn != nil {
		return ws.conn.Close()
	}
	return nil
}

// StartReceivingNotifications starts receiving notifications via polling
func (c *Client) StartReceivingNotifications(ctx context.Context, handler *WebhookHandler) error {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			notification, err := c.ReceiveNotification(ctx)
			if err != nil {
				log.Printf("Error receiving notification: %v", err)
				continue
			}

			// If notification is empty, continue
			if len(notification) == 0 {
				continue
			}

			// Handle the notification
			if handler != nil {
				if err := handler.HandleWebhook(notification); err != nil {
					log.Printf("Error handling notification: %v", err)
				}
			}

			// Delete the notification if it has a receiptId
			if receiptIDFloat, ok := notification["receiptId"].(float64); ok {
				receiptID := int64(receiptIDFloat)
				if _, err := c.DeleteNotification(ctx, receiptID); err != nil {
					log.Printf("Error deleting notification: %v", err)
				}
			}
		}
	}
}
