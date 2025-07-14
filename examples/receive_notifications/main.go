import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	sdkwa "github.com/sdkwa/whatsapp-api-client-go"
)

func main() {
	client, err := sdkwa.NewClient(sdkwa.Options{
		IDInstance:       os.Getenv("SDKWA_ID_INSTANCE"),
		APITokenInstance: os.Getenv("SDKWA_API_TOKEN"),
	})
	if err != nil {
		log.Fatal(err)
	}

	// Receive notification
	notification, err := client.ReceiveNotification(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	if len(notification) == 0 {
		fmt.Println("📭 No notifications")
		return
	}

	// Print and delete notification
	jsonData, _ := json.MarshalIndent(notification, "", "  ")
	fmt.Printf("📬 Notification:\n%s\n", jsonData)

	if receiptID, ok := notification["receiptId"].(float64); ok {
		client.DeleteNotification(context.Background(), int64(receiptID))
		fmt.Println("✅ Deleted")
	}
}