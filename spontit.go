// Package spontit provides a client for the Spontit API
package spontit

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const (
	baseURL = "https://api.spontit.com/v3"
)

// Client represents a Spontit API client
type Client struct {
	config     *Config
	httpClient *http.Client
}

// NewClient returns a new Spontit client
func NewClient() (*Client, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Error loading .env file")
	}

	config := &Config{
		userID: os.Getenv("SPONTIT_USERID"),
		key:    os.Getenv("SPONTIT_KEY"),
	}

	return &Client{
		config:     config,
		httpClient: http.DefaultClient,
	}, nil
}
