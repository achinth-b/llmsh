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

type embeddingRequest struct {
	Input          string `json:"input"`
	Model          string `json:"model"`
	EncodingFormat string `json:"encoding_format"`
}

type embeddingResponse struct {
	Object string `json:"object"`
	Data   []struct {
		Object    string    `json:"object"`
		Embedding []float64 `json:"embedding"`
		Index     int       `json:"index"`
	} `json:"data"`
	Model string `json:"model"`
	Usage struct {
		PromptTokens int `json:"prompt_tokens"`
		TotalTokens  int `json:"total_tokens"`
	} `json:"usage"`
}

func Embedding(embModelName *string, embInput *string) ([]float64, error) {

	openAIAPIKey, exists := os.LookupEnv("OPENAI_API_KEY")

	if !(exists) {
		return nil, errors.New("OPENAI_API_KEY environment variable not set")
	}

	// Prepare the request payload

	if !IsOpenAIModelAvailable("embedding", embModelName) {
		return nil, errors.New("chosen model not available in the openai embeddings model family")
	}
	reqBody := embeddingRequest{
		Input:          *embInput,
		Model:          *embModelName,
		EncodingFormat: "float",
	}

	// Convert the request payload to JSON
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return nil, err
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", apiURL+"embeddings", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return nil, err
	}

	// Set the necessary headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", openAIAPIKey))

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return nil, err
	}

	// Parse the JSON response
	var embeddingResp embeddingResponse
	err = json.Unmarshal(body, &embeddingResp)

	if err != nil {
		fmt.Printf("Error parsing JSON response: %v\n", err)
		return nil, err
	}

	// Print the embedding
	resultEmbedding := embeddingResp.Data[0].Embedding
	fmt.Printf("Embedding: %v\n", resultEmbedding)
	return resultEmbedding, nil

}
