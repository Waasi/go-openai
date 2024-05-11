package openai

import (
	"context"
	"fmt"
	"net/http"
)

type RunStreamResponse struct {
	ID             string             `json:"id"`
	Status         string             `json:"status"`
	RequiredAction *RunRequiredAction `json:"required_action,omitempty"`
	Object         string             `json:"object"`
	Created        int64              `json:"created"`
	Model          string             `json:"model"`
	Content        []MessageContent   `json:"content"`
}

// ChatCompletionStream
// Note: Perhaps it is more elegant to abstract Stream using generics.
type RunStream struct {
	*streamReader[RunStreamResponse]
}

// CreateChatCompletionStream — API call to create a chat completion w/ streaming
// support. It sets whether to stream back partial progress. If set, tokens will be
// sent as data-only server-sent events as they become available, with the
// stream terminated by a data: [DONE] message.
func (c *Client) CreateRunStream(
	ctx context.Context,
	threadID string,
	request RunRequest,
) (stream *RunStream, err error) {
	urlSuffix := fmt.Sprintf("/threads/%s/runs", threadID)
	if !checkEndpointSupportsModel(urlSuffix, request.Model) {
		err = ErrChatCompletionInvalidModel
		return
	}

	request.Stream = true
	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(urlSuffix, request.Model), withBody(request), withBetaAssistantV1())
	if err != nil {
		return nil, err
	}

	resp, err := sendRequestStream[RunStreamResponse](c, req)
	if err != nil {
		return
	}
	stream = &RunStream{
		streamReader: resp,
	}
	return
}
