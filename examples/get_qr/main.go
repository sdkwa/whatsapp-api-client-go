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

	// Get QR code for authorization
	qr, err := client.GetQR(ctx)
	if err != nil {
		log.Fatalf("Failed to get QR code: %v", err)
	}

	fmt.Printf("QR Code Type: %s\n", qr.Type)
	fmt.Printf("QR Code Message: %s\n", qr.Message)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
