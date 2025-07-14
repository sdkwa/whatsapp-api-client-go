package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	sdkwa "github.com/sdkwa/whatsapp-api-client-go"
)

func main() {
	// Create client
	client, err := sdkwa.NewClient(sdkwa.Options{
		APIHost:          getEnv("SDKWA_API_HOST", "https://api.sdkwa.pro"),
		IDInstance:       getEnv("SDKWA_ID_INSTANCE", ""),
		APITokenInstance: getEnv("SDKWA_API_TOKEN", ""),
	})
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Create webhook handler
	handler := sdkwa.NewWebhookHandler()

	// Register handlers for different event types
	handler.OnIncomingMessageText(func(data map[string]interface{}) error {
		log.Printf("Received text message: %+v", data)
		return nil
	})

	handler.OnIncomingMessageFile(func(data map[string]interface{}) error {
		log.Printf("Received file message: %+v", data)
		return nil
	})

	handler.OnStateInstance(func(data map[string]interface{}) error {
		log.Printf("State instance changed: %+v", data)
		return nil
	})

	// Start webhook server
	http.Handle("/webhook", handler)
	server := &http.Server{Addr: ":8080"}

	go func() {
		log.Println("Starting webhook server on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Start receiving notifications via polling as fallback
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		log.Println("Starting notification polling...")
		if err := client.StartReceivingNotifications(ctx, handler); err != nil {
			log.Printf("Error in notification polling: %v", err)
		}
	}()

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("Shutting down...")
	server.Shutdown(context.Background())
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
