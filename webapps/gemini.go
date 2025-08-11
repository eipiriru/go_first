package webapps

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	genai "google.golang.org/genai"
)

func GetRandomKomentarByPrompt(prompt string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("GEMINI_API_KEY")

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		log.Fatalf("Failed to create GenAI client: %v", err)
	}

	result, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash", genai.Text(prompt), nil)
	if err != nil {
		log.Fatalf("API call error: %v", err)
	}

	return result.Text()
}
