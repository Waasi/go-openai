package openai

import (
	"context"
	"fmt"
	"net/http"
)

const (
	vectorStoreSuffix = "/vector_stores"
)

type VectorStore struct {
	ID           string            `json:"id"`
	Object       string            `json:"object"`
	Bytes        int64             `json:"bytes"`
	Name         string            `json:"name"`
	Status       string            `json:"status"`
	ExpiresAfter *ExpirationConfig `json:"expires_after,omitempty"`
	FileCounts   FileCountMap      `json:"file_counts"`

	CreatedAt    int64          `json:"created_at"`
	ExpiresAt    *int64         `json:"expires_at,omitempty"`
	LastActiveAt *int64         `json:"last_active_at,omitempty"`
	Metadata     map[string]any `json:"metadata,omitempty"`

	httpHeader
}

type ExpirationConfig struct {
	Anchor string `json:"anchor"`
	Days   int64  `json:"days"`
}

type FileCountMap struct {
	InProgress int64 `json:"in_progress"`
	Completed  int64 `json:"completed"`
	Failed     int64 `json:"failed"`
	Cancelled  int64 `json:"cancelled"`
	Total      int64 `json:"total"`
}

// VectorStoreRequest provides the vectorStore request parameters.
// When modifying the tools the API functions as the following:
type VectorStoreRequest struct {
	Name         *string           `json:"name,omitempty"`
	FileIDs      []string          `json:"file_ids,omitempty"`
	ExpiresAfter *ExpirationConfig `json:"expires_after,omitempty"`
	Metadata     map[string]any    `json:"metadata,omitempty"`
}

// CreateVectorStore creates a new vectorStore.
func (c *Client) CreateVectorStore(ctx context.Context, request VectorStoreRequest) (response VectorStore, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(vectorStoreSuffix), withBody(request))
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}

// ModifyVectorStore modifies an vectorStore.
func (c *Client) ModifyVectorStore(
	ctx context.Context,
	vectorStoreID string,
	request VectorStoreRequest,
) (response VectorStore, err error) {
	urlSuffix := fmt.Sprintf("%s/%s", vectorStoreSuffix, vectorStoreID)
	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(urlSuffix), withBody(request))
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}
