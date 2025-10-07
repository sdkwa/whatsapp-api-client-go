package main

import (
	"context"
	"fmt"
	"log"
	"os"

	sdkwa "github.com/sdkwa/whatsapp-api-client-go"
)

func main() {
	// Example 1: Create a Telegram client
	telegramClient, err := sdkwa.NewClient(sdkwa.Options{
		APIHost:          os.Getenv("SDKWA_API_HOST"),
		IDInstance:       os.Getenv("SDKWA_ID_INSTANCE"),
		APITokenInstance: os.Getenv("SDKWA_API_TOKEN"),
		MessengerType:    sdkwa.MessengerTelegram, // Set Telegram as default
	})
	if err != nil {
		log.Fatalf("Failed to create Telegram client: %v", err)
	}

	ctx := context.Background()

	// Use Telegram-specific methods
	fmt.Println("=== Telegram-Specific Methods ===")

	// Create a Telegram app
	createAppResp, err := telegramClient.CreateApp(ctx, sdkwa.CreateAppParams{
		Title:       "My Telegram Bot",
		ShortName:   "mybot",
		URL:         "https://example.com/bot",
		Description: "My awesome Telegram bot",
	})
	if err != nil {
		log.Printf("Failed to create app: %v", err)
	} else {
		fmt.Printf("App created: %+v\n", createAppResp)
	}

	// Send confirmation code
	codeResp, err := telegramClient.SendConfirmationCode(ctx, sdkwa.SendConfirmationCodeParams{
		PhoneNumber: 1234567890,
	})
	if err != nil {
		log.Printf("Failed to send confirmation code: %v", err)
	} else {
		fmt.Printf("Confirmation code sent: %+v\n", codeResp)
	}

	// Sign in with confirmation code
	signInResp, err := telegramClient.SignInWithConfirmationCode(ctx, sdkwa.SignInWithConfirmationCodeParams{
		Code: "123456",
	})
	if err != nil {
		log.Printf("Failed to sign in: %v", err)
	} else {
		fmt.Printf("Sign in response: %+v\n", signInResp)
	}

	// Example 2: Use WhatsApp client with Telegram override
	whatsappClient, err := sdkwa.NewClient(sdkwa.Options{
		APIHost:          os.Getenv("SDKWA_API_HOST"),
		IDInstance:       os.Getenv("SDKWA_ID_INSTANCE"),
		APITokenInstance: os.Getenv("SDKWA_API_TOKEN"),
		MessengerType:    sdkwa.MessengerWhatsApp, // Default to WhatsApp
	})
	if err != nil {
		log.Fatalf("Failed to create WhatsApp client: %v", err)
	}

	fmt.Println("\n=== Using Messenger Type Override ===")

	// Send a message using WhatsApp (default)
	_, err = whatsappClient.SendMessage(ctx, sdkwa.SendMessageParams{
		ChatID:  "1234567890@c.us",
		Message: "Hello from WhatsApp",
	})
	if err != nil {
		log.Printf("Failed to send WhatsApp message: %v", err)
	} else {
		fmt.Println("WhatsApp message sent successfully")
	}

	// Send a message using Telegram (override)
	_, err = whatsappClient.SendMessage(ctx, sdkwa.SendMessageParams{
		ChatID:  "123456789",
		Message: "Hello from Telegram",
	}, &sdkwa.RequestOptions{
		MessengerType: sdkwa.MessengerTelegram, // Override to Telegram
	})
	if err != nil {
		log.Printf("Failed to send Telegram message: %v", err)
	} else {
		fmt.Println("Telegram message sent successfully (via override)")
	}

	// Example 3: Get instance state for both messengers
	fmt.Println("\n=== Getting Instance State ===")

	// Get WhatsApp state
	whatsappState, err := whatsappClient.GetStateInstance(ctx)
	if err != nil {
		log.Printf("Failed to get WhatsApp state: %v", err)
	} else {
		fmt.Printf("WhatsApp state: %+v\n", whatsappState)
	}

	// Get Telegram state using override
	telegramState, err := whatsappClient.GetStateInstance(ctx, &sdkwa.RequestOptions{
		MessengerType: sdkwa.MessengerTelegram,
	})
	if err != nil {
		log.Printf("Failed to get Telegram state: %v", err)
	} else {
		fmt.Printf("Telegram state: %+v\n", telegramState)
	}

	// Example 4: Working with chats
	fmt.Println("\n=== Working with Chats ===")

	// Get WhatsApp chats
	whatsappChats, err := whatsappClient.GetChats(ctx)
	if err != nil {
		log.Printf("Failed to get WhatsApp chats: %v", err)
	} else {
		fmt.Printf("WhatsApp chats count: %d\n", len(whatsappChats))
	}

	// Get Telegram chats using override
	telegramChats, err := whatsappClient.GetChats(ctx, &sdkwa.RequestOptions{
		MessengerType: sdkwa.MessengerTelegram,
	})
	if err != nil {
		log.Printf("Failed to get Telegram chats: %v", err)
	} else {
		fmt.Printf("Telegram chats count: %d\n", len(telegramChats))
	}

	fmt.Println("\n=== Example Complete ===")
}
