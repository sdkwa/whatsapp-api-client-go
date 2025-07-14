package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	sdkwa "github.com/sdkwa/whatsapp-api-client-go"
)

func main() {
	// Example: Using the SDKWA Go library in a real application
	// This demonstrates common usage patterns

	// Create client with environment variables or defaults
	client, err := sdkwa.NewClient(sdkwa.Options{
		APIHost:          getEnv("SDKWA_API_HOST", "https://api.sdkwa.pro"),
		IDInstance:       getEnv("SDKWA_ID_INSTANCE", ""),
		APITokenInstance: getEnv("SDKWA_API_TOKEN", ""),
		Timeout:          30 * time.Second,
	})
	if err != nil {
		log.Fatalf("Failed to create SDKWA client: %v", err)
	}

	ctx := context.Background()

	// 1. Check account state
	fmt.Println("=== Checking Account State ===")
	state, err := client.GetStateInstance(ctx)
	if err != nil {
		log.Printf("Error getting state: %v", err)
	} else {
		fmt.Printf("Account state: %s\n", state.StateInstance)
	}

	// 2. Get account settings
	fmt.Println("\n=== Getting Account Settings ===")
	settings, err := client.GetSettings(ctx)
	if err != nil {
		log.Printf("Error getting settings: %v", err)
	} else {
		fmt.Printf("Settings: %+v\n", settings)
	}

	// 3. Send a simple message
	fmt.Println("\n=== Sending Message ===")
	messageResponse, err := client.SendMessage(ctx, sdkwa.SendMessageParams{
		ChatID:      "1234567890@c.us", // Replace with actual chat ID
		Message:     "Hello from SDKWA Go library! 🚀",
		LinkPreview: true,
	})
	if err != nil {
		log.Printf("Error sending message: %v", err)
	} else {
		fmt.Printf("Message sent with ID: %s\n", messageResponse.IDMessage)
	}

	// 4. Send a contact
	fmt.Println("\n=== Sending Contact ===")
	contactResponse, err := client.SendContact(ctx, sdkwa.SendContactParams{
		ChatID: "1234567890@c.us", // Replace with actual chat ID
		Contact: sdkwa.Contact{
			PhoneContact: 1234567890,
			FirstName:    "John",
			LastName:     "Doe",
			Company:      "SDKWA",
		},
	})
	if err != nil {
		log.Printf("Error sending contact: %v", err)
	} else {
		fmt.Printf("Contact sent with ID: %s\n", contactResponse.IDMessage)
	}

	// 5. Send location
	fmt.Println("\n=== Sending Location ===")
	locationResponse, err := client.SendLocation(ctx, sdkwa.SendLocationParams{
		ChatID:       "1234567890@c.us", // Replace with actual chat ID
		NameLocation: "SDKWA Headquarters",
		Address:      "Tech Street 123",
		Latitude:     40.7128,
		Longitude:    -74.0060,
	})
	if err != nil {
		log.Printf("Error sending location: %v", err)
	} else {
		fmt.Printf("Location sent with ID: %s\n", locationResponse.IDMessage)
	}

	// 6. Check WhatsApp availability
	fmt.Println("\n=== Checking WhatsApp Availability ===")
	availability, err := client.CheckWhatsApp(ctx, 1234567890) // Replace with actual phone number
	if err != nil {
		log.Printf("Error checking WhatsApp: %v", err)
	} else {
		fmt.Printf("WhatsApp exists: %t\n", availability.ExistsWhatsApp)
	}

	// 7. Get contacts
	fmt.Println("\n=== Getting Contacts ===")
	contacts, err := client.GetContacts(ctx)
	if err != nil {
		log.Printf("Error getting contacts: %v", err)
	} else {
		fmt.Printf("Found %d contacts\n", len(contacts))
	}

	// 8. Get chats
	fmt.Println("\n=== Getting Chats ===")
	chats, err := client.GetChats(ctx)
	if err != nil {
		log.Printf("Error getting chats: %v", err)
	} else {
		fmt.Printf("Found %d chats\n", len(chats))
	}

	// 9. Show message queue
	fmt.Println("\n=== Checking Message Queue ===")
	queue, err := client.ShowMessagesQueue(ctx)
	if err != nil {
		log.Printf("Error getting queue: %v", err)
	} else {
		fmt.Printf("Messages in queue: %d\n", len(queue))
	}

	// 10. Example of error handling
	fmt.Println("\n=== Error Handling Example ===")
	_, err = client.SendMessage(ctx, sdkwa.SendMessageParams{
		ChatID:  "invalid-chat-id",
		Message: "This will fail",
	})
	if err != nil {
		if apiError, ok := err.(*sdkwa.ErrorResponse); ok {
			fmt.Printf("API Error: %s (Status: %d)\n", apiError.Message, apiError.StatusCode)
		} else {
			fmt.Printf("Network/Other Error: %v\n", err)
		}
	}

	fmt.Println("\n=== Example Complete ===")
	fmt.Println("This demonstrates basic usage of the SDKWA Go library.")
	fmt.Println("Replace the phone numbers and chat IDs with real values to test.")
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
