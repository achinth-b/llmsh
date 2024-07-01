package commands

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

var apiURL = "https://api.openai.com/v1/"
var availableChatModels = []string{"gpt-3.5-turbo-0125", "gpt-3.5-turbo"}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type RequestBody struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type ChatResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

func Chat(chatModel *string, chatInput *string) (string, error) {
	openAIAPIKey, exists := os.LookupEnv("OPENAI_API_KEY")

	if !(exists) {
		return "", errors.New("OPENAI_API_KEY environment variable not set")
	}

	if !IsOpenAIModelAvailable("chat", chatModel) {
		return "", errors.New("chosen model not available in the openai embeddings model family")
	}

	// Prepare the request body
	reqBody := RequestBody{
		Model: *chatModel,
		Messages: []Message{
			{Role: "user", Content: *chatInput},
		},
	}

	// Marshal the request body to JSON
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return "", err
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", apiURL+"chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+openAIAPIKey)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "", err
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return "", err
	}

	// Print the response
	var chatResponse ChatResponse
	err = json.Unmarshal(body, &chatResponse)
	if err != nil {
		fmt.Printf("Error parsing JSON response: %v\n", err)
		return "", err
	}
	return chatResponse.Choices[0].Message.Content, nil

}
