package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"google.golang.org/genai"
)

func Chat(userInput string, history *ChatHistory) string {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: GetAPIKey(),
	})
	if err != nil {
		return "Oops! Couldn't reach Gemini."
	}

	// Append user input
	history.History = append(history.History, ChatMessage{
		Role:    "user",
		Message: userInput,
	})

	// Convert to Gemini format
	var contents []*genai.Content
	for _, m := range history.History {
		contents = append(contents, &genai.Content{
			Role: m.Role,
			Parts: []*genai.Part{
				{Text: m.Message},
			},
		})
	}

	sys := BuildSystemPrompt(history.User)

	stream := client.Models.GenerateContentStream(
		ctx,
		"gemini-1.5-flash",
		contents,
		&genai.GenerateContentConfig{
			SystemInstruction: &genai.Content{
				Role: "system",
				Parts: []*genai.Part{
					{Text: sys},
				},
			},
		},
	)

	var response strings.Builder
	for chunk := range stream {
		part := chunk.Candidates[0].Content.Parts[0]
		text := part.Text
		fmt.Print(text)
		response.WriteString(text)
	}

	reply := response.String()
	history.History = append(history.History, ChatMessage{
		Role:    "assistant",
		Message: reply,
	})

	return reply
}

func GetAPIKey() string {
	return os.Getenv("GEMINI_API_KEY")
}
