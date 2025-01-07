package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ollama/ollama/api"
)

var (
	FALSE = false
	TRUE  = true
)


func main() {

	ctx := context.Background()

	ollamaUrl := os.Getenv("OLLAMA_HOST")
	model := os.Getenv("LLM")

	fmt.Println("🌍", ollamaUrl, "📕", model)

	client, errCli := api.ClientFromEnvironment()
	if errCli != nil {
		log.Fatal("😡:", errCli)
	}

	// Prompt construction
	messages := []api.Message{
		{Role: "user", Content: "Who is Jean-Luc Picard?"},
	}

	req := &api.ChatRequest{
		Model:    model,
		Messages: messages,
		Options: map[string]interface{}{
			"temperature":    0.9,
			"repeat_last_n":  2,
			"repeat_penalty": 1.8,
			"top_k":          10,
			"top_p":          0.5,
		},
		Stream:    &TRUE,
	}

	respFunc := func(resp api.ChatResponse) error {
		fmt.Print(resp.Message.Content)
		return nil
	}

	// Start the chat completion
	errChat := client.Chat(ctx, req, respFunc)
	if errChat != nil {
		log.Fatal("😡:", errChat)
	}

}
