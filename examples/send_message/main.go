package main

import (
	"context"
	"fmt"
	"log"
	"os"

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

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
