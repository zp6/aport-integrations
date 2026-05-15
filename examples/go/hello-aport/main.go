package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// APort Hello World Example in Go
// This example demonstrates basic integration with the APort API.

const defaultBaseURL = "https://api.aport.io/v1"

type APortClient struct {
	APIKey  string
	BaseURL string
	Client  *http.Client
}

type HealthResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

func NewAPortClient(apiKey string) *APortClient {
	return &APortClient{
		APIKey:  apiKey,
		BaseURL: defaultBaseURL,
		Client:  &http.Client{},
	}
}

func (c *APortClient) doRequest(method, path string) (*http.Response, error) {
	req, err := http.NewRequest(method, c.BaseURL+path, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")
	return c.Client.Do(req)
}

// CheckHealth verifies the APort API is reachable
func (c *APortClient) CheckHealth() (*HealthResponse, error) {
	resp, err := c.doRequest("GET", "/health")
	if err != nil {
		return nil, fmt.Errorf("health check failed: %w", err)
	}
	defer resp.Body.Close()

	var health HealthResponse
	if err := json.NewDecoder(resp.Body).Decode(&health); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}
	return &health, nil
}

func main() {
	apiKey := os.Getenv("APORT_API_KEY")
	if apiKey == "" {
		log.Println("Warning: APORT_API_KEY not set, using demo mode")
		apiKey = "demo-key"
	}

	client := NewAPortClient(apiKey)

	// Step 1: Check API health
	fmt.Println("Hello, APort! 👋")
	fmt.Println("Checking API health...")

	health, err := client.CheckHealth()
	if err != nil {
		log.Printf("Health check error: %v", err)
		fmt.Println("API is not reachable. Please check your configuration.")
		return
	}

	fmt.Printf("APort API Status: %s (version %s)\n", health.Status, health.Version)
	fmt.Println("Integration successful! ✅")
}
