package exa

import (
	"context"
	"errors"
	"fmt"

	"github.com/builtbyrobben/exa-cli/internal/api"
)

var (
	errQueryRequired = errors.New("query is required")
	errURLRequired   = errors.New("url is required")
)

const defaultBaseURL = "https://api.exa.ai"

// Client wraps the API client with Exa-specific methods.
type Client struct {
	*api.Client
}

// NewClient creates a new Exa API client.
func NewClient(apiKey string) *Client {
	return &Client{
		Client: api.NewClient(apiKey,
			api.WithBaseURL(defaultBaseURL),
			api.WithUserAgent("exa-cli/1.0"),
			api.WithAuthHeader("x-api-key", apiKey),
		),
	}
}

// Search performs a web search.
func (c *Client) Search(ctx context.Context, req SearchRequest) (*SearchResponse, error) {
	if req.Query == "" {
		return nil, errQueryRequired
	}

	var result SearchResponse
	if err := c.Post(ctx, "/search", req, &result); err != nil {
		return nil, fmt.Errorf("search: %w", err)
	}

	return &result, nil
}

// SearchAndContents performs a search and returns page contents.
func (c *Client) SearchAndContents(ctx context.Context, req SearchRequest) (*SearchResponse, error) {
	if req.Query == "" {
		return nil, errQueryRequired
	}

	if req.Contents == nil {
		req.Contents = &ContentOptions{
			Text: &TextOptions{},
		}
	}

	var result SearchResponse
	if err := c.Post(ctx, "/search", req, &result); err != nil {
		return nil, fmt.Errorf("search and contents: %w", err)
	}

	return &result, nil
}

// FindSimilar finds pages similar to a given URL.
func (c *Client) FindSimilar(ctx context.Context, req FindSimilarRequest) (*SearchResponse, error) {
	if req.URL == "" {
		return nil, errURLRequired
	}

	var result SearchResponse
	if err := c.Post(ctx, "/findSimilar", req, &result); err != nil {
		return nil, fmt.Errorf("find similar: %w", err)
	}

	return &result, nil
}

// Answer gets an AI-generated answer with citations.
func (c *Client) Answer(ctx context.Context, req AnswerRequest) (*AnswerResponse, error) {
	if req.Query == "" {
		return nil, errQueryRequired
	}

	var result AnswerResponse
	if err := c.Post(ctx, "/answer", req, &result); err != nil {
		return nil, fmt.Errorf("answer: %w", err)
	}

	return &result, nil
}
