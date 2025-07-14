package main

import (
	"context"
	"fmt"
	"log"
	"os"

	sdkwa "github.com/sdkwa/whatsapp-api-client-go"
)

func main() {
	// Create client with user credentials for instance management
	client, err := sdkwa.NewClient(sdkwa.Options{
		APIHost:          getEnv("SDKWA_API_HOST", "https://api.sdkwa.pro"),
		IDInstance:       getEnv("SDKWA_ID_INSTANCE", ""),
		APITokenInstance: getEnv("SDKWA_API_TOKEN", ""),
		UserID:           getEnv("SDKWA_USER_ID", ""),
		UserToken:        getEnv("SDKWA_USER_TOKEN", ""),
	})
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	// Get all instances
	instances, err := client.GetInstances(ctx)
	if err != nil {
		log.Fatalf("Failed to get instances: %v", err)
	}

	fmt.Printf("User instances: %+v\n", instances)

	// Create new instance (uncomment to test)
	/*
		newInstance, err := client.CreateInstance(ctx, sdkwa.CreateInstanceParams{
			Tariff: "developer",
			Period: "month1",
		})
		if err != nil {
			log.Fatalf("Failed to create instance: %v", err)
		}

		fmt.Printf("New instance created: %+v\n", newInstance)
	*/
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
