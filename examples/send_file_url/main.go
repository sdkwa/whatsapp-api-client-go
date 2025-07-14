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

	// Example 1: Send an image by URL
	fmt.Println("=== Sending Image by URL ===")
	imageResponse, err := client.SendFileByURL(ctx, sdkwa.SendFileByURLParams{
		ChatID:   "79999999999@c.us", // Replace with actual chat ID
		URLFile:  "https://via.placeholder.com/300x200.png?text=Sample+Image",
		FileName: "sample_image.png",
		Caption:  "This is a sample image sent via URL! 📸",
	})
	if err != nil {
		log.Printf("Failed to send image: %v", err)
	} else {
		fmt.Printf("✅ Image sent with ID: %s\n", imageResponse.IDMessage)
	}

	// Example 2: Send a PDF document by URL
	fmt.Println("\n=== Sending PDF Document by URL ===")
	pdfResponse, err := client.SendFileByURL(ctx, sdkwa.SendFileByURLParams{
		ChatID:   "79999999999@c.us", // Replace with actual chat ID
		URLFile:  "https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf",
		FileName: "sample_document.pdf",
		Caption:  "Here's a PDF document! 📄",
	})
	if err != nil {
		log.Printf("Failed to send PDF: %v", err)
	} else {
		fmt.Printf("✅ PDF sent with ID: %s\n", pdfResponse.IDMessage)
	}

	// Example 3: Send a video by URL
	fmt.Println("\n=== Sending Video by URL ===")
	videoResponse, err := client.SendFileByURL(ctx, sdkwa.SendFileByURLParams{
		ChatID:   "79999999999@c.us", // Replace with actual chat ID
		URLFile:  "https://sample-videos.com/zip/10/mp4/SampleVideo_1280x720_1mb.mp4",
		FileName: "sample_video.mp4",
		Caption:  "Check out this video! 🎥",
	})
	if err != nil {
		log.Printf("Failed to send video: %v", err)
	} else {
		fmt.Printf("✅ Video sent with ID: %s\n", videoResponse.IDMessage)
	}

	// Example 4: Send file with quoted message and archive chat
	fmt.Println("\n=== Sending File with Advanced Options ===")
	advancedResponse, err := client.SendFileByURL(ctx, sdkwa.SendFileByURLParams{
		ChatID:          "79999999999@c.us", // Replace with actual chat ID
		URLFile:         "https://jsonplaceholder.typicode.com/posts/1",
		FileName:        "data.json",
		Caption:         "Here's some JSON data 📊",
		QuotedMessageID: "", // Add message ID to quote if needed
		ArchiveChat:     false,
	})
	if err != nil {
		log.Printf("Failed to send file with advanced options: %v", err)
	} else {
		fmt.Printf("✅ File with advanced options sent with ID: %s\n", advancedResponse.IDMessage)
	}

	// Example 5: Upload file first, then send by URL
	fmt.Println("\n=== Upload File Then Send by URL ===")
	
	// Note: This would require having a local file to upload first
	// Here's how you would do it:
	/*
	file, err := os.Open("local_file.txt")
	if err != nil {
		log.Printf("Failed to open local file: %v", err)
		return
	}
	defer file.Close()

	// Upload the file
	uploadResponse, err := client.UploadFile(ctx, file)
	if err != nil {
		log.Printf("Failed to upload file: %v", err)
		return
	}

	fmt.Printf("File uploaded to: %s\n", uploadResponse.URLFile)

	// Now send the uploaded file by URL
	uploadedFileResponse, err := client.SendFileByURL(ctx, sdkwa.SendFileByURLParams{
		ChatID:   "79999999999@c.us",
		URLFile:  uploadResponse.URLFile,
		FileName: "uploaded_file.txt",
		Caption:  "This file was uploaded first, then sent! 📁",
	})
	if err != nil {
		log.Printf("Failed to send uploaded file: %v", err)
	} else {
		fmt.Printf("✅ Uploaded file sent with ID: %s\n", uploadedFileResponse.IDMessage)
	}
	*/

	fmt.Println("\n=== SendFileByURL Examples Complete ===")
	fmt.Println("Replace the phone numbers and URLs with real values to test.")
	fmt.Println("\nSupported file types:")
	fmt.Println("📸 Images: JPG, PNG, GIF, WEBP")
	fmt.Println("🎥 Videos: MP4, AVI, MOV, 3GP")
	fmt.Println("🎵 Audio: MP3, WAV, OGG, AAC")
	fmt.Println("📄 Documents: PDF, DOC, DOCX, XLS, XLSX, PPT, PPTX, TXT")
	fmt.Println("📁 Archives: ZIP, RAR, 7Z")
	fmt.Println("💾 Maximum file size: 100MB")
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
