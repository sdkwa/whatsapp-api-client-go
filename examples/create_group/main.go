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

	// Create a group
	response, err := client.CreateGroup(ctx, "My Go Group", []string{
		"79999999999@c.us", // Add participant phone numbers
	})
	if err != nil {
		log.Fatalf("Failed to create group: %v", err)
	}

	fmt.Printf("Group created successfully!\n")
	fmt.Printf("Group ID: %s\n", response.ChatID)
	fmt.Printf("Invite Link: %s\n", response.GroupInviteLink)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
