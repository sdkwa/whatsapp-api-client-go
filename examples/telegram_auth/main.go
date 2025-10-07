package main

import (
	"context"
	"fmt"
	"log"
	"os"

	sdkwa "github.com/sdkwa/whatsapp-api-client-go"
)

func main() {
	// Create a Telegram client
	client, err := sdkwa.NewClient(sdkwa.Options{
		IDInstance:       os.Getenv("SDKWA_ID_INSTANCE"),
		APITokenInstance: os.Getenv("SDKWA_API_TOKEN"),
		MessengerType:    sdkwa.MessengerTelegram,
	})
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	// Step 1: Create a Telegram app
	fmt.Println("Creating Telegram app...")
	appResp, err := client.CreateApp(ctx, sdkwa.CreateAppParams{
		Title:       "My Awesome Bot",
		ShortName:   "mybot",
		URL:         "https://example.com/bot",
		Description: "A helpful Telegram bot",
	})
	if err != nil {
		log.Fatalf("Failed to create app: %v", err)
	}
	fmt.Printf("✅ App created successfully!\n")
	fmt.Printf("   App ID: %s\n", appResp.Data.AppID)
	fmt.Printf("   Result: %v\n\n", appResp.Result)

	// Step 2: Send confirmation code
	fmt.Println("Sending confirmation code...")
	phoneNumber := int64(1234567890) // Replace with actual phone number
	codeResp, err := client.SendConfirmationCode(ctx, sdkwa.SendConfirmationCodeParams{
		PhoneNumber: phoneNumber,
	})
	if err != nil {
		log.Fatalf("Failed to send confirmation code: %v", err)
	}
	fmt.Printf("✅ Confirmation code sent!\n")
	fmt.Printf("   Result: %v\n", codeResp.Result)
	fmt.Printf("   Message: %s\n\n", codeResp.Message)

	// Step 3: Sign in with confirmation code
	// Note: In a real application, you would wait for the user to receive
	// and enter the confirmation code
	fmt.Println("To complete sign-in, enter the confirmation code you received:")
	var code string
	fmt.Print("Code: ")
	fmt.Scanln(&code)

	signInResp, err := client.SignInWithConfirmationCode(ctx, sdkwa.SignInWithConfirmationCodeParams{
		Code: code,
	})
	if err != nil {
		log.Fatalf("Failed to sign in: %v", err)
	}
	fmt.Printf("✅ Successfully signed in!\n")
	fmt.Printf("   Result: %v\n", signInResp.Result)
	fmt.Printf("   Message: %s\n\n", signInResp.Message)

	// Step 4: Check instance state
	fmt.Println("Checking instance state...")
	state, err := client.GetStateInstance(ctx)
	if err != nil {
		log.Fatalf("Failed to get state: %v", err)
	}
	fmt.Printf("✅ Instance state: %s\n", state.StateInstance)
}
