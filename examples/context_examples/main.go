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
	// Create client
	client, err := sdkwa.NewClient(sdkwa.Options{
		APIHost:          getEnv("SDKWA_API_HOST", "https://api.sdkwa.pro"),
		IDInstance:       getEnv("SDKWA_ID_INSTANCE", ""),
		APITokenInstance: getEnv("SDKWA_API_TOKEN", ""),
	})
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Example 1: Basic context (simplest, never times out)
	fmt.Println("=== Example 1: Basic Context ===")
	basicContext()

	// Example 2: Context with timeout (recommended for production)
	fmt.Println("\n=== Example 2: Context with Timeout ===")
	contextWithTimeout(client)

	// Example 3: Context with cancellation
	fmt.Println("\n=== Example 3: Context with Cancellation ===")
	contextWithCancellation(client)

	// Example 4: Context with deadline
	fmt.Println("\n=== Example 4: Context with Deadline ===")
	contextWithDeadline(client)
}

func basicContext() {
	// This creates a context that never cancels or times out
	// Good for: Simple scripts, testing
	// Bad for: Production apps (no timeout protection)
	ctx := context.Background()

	fmt.Printf("Context type: %T\n", ctx)
	fmt.Println("This context will never timeout or cancel")
}

func contextWithTimeout(client *sdkwa.Client) {
	// This creates a context that automatically cancels after 10 seconds
	// Good for: Production apps, API calls that should not hang
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Always call cancel to free resources

	fmt.Println("Sending message with 10-second timeout...")
	response, err := client.SendMessage(ctx, sdkwa.SendMessageParams{
		ChatID:  "79999999999@c.us",
		Message: "Hello with timeout!",
	})

	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("❌ Request timed out after 10 seconds")
		} else {
			fmt.Printf("❌ Request failed: %v\n", err)
		}
	} else {
		fmt.Printf("✅ Message sent with ID: %s\n", response.IDMessage)
	}
}

func contextWithCancellation(client *sdkwa.Client) {
	// This creates a context that you can manually cancel
	// Good for: User-initiated cancellation, graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())

	// Simulate cancelling the request after 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("🛑 Manually cancelling the request...")
		cancel()
	}()

	fmt.Println("Sending message that will be cancelled...")
	response, err := client.SendMessage(ctx, sdkwa.SendMessageParams{
		ChatID:  "79999999999@c.us",
		Message: "This message might be cancelled!",
	})

	if err != nil {
		if ctx.Err() == context.Canceled {
			fmt.Println("❌ Request was cancelled")
		} else {
			fmt.Printf("❌ Request failed: %v\n", err)
		}
	} else {
		fmt.Printf("✅ Message sent with ID: %s\n", response.IDMessage)
	}
}

func contextWithDeadline(client *sdkwa.Client) {
	// This creates a context that cancels at a specific time
	// Good for: Operations that must complete by a certain time
	deadline := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	fmt.Printf("Sending message with deadline at %v...\n", deadline.Format("15:04:05"))
	response, err := client.SendMessage(ctx, sdkwa.SendMessageParams{
		ChatID:  "79999999999@c.us",
		Message: "Hello with deadline!",
	})

	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("❌ Request exceeded deadline")
		} else {
			fmt.Printf("❌ Request failed: %v\n", err)
		}
	} else {
		fmt.Printf("✅ Message sent with ID: %s\n", response.IDMessage)
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
