package http_client

import (
	"fmt"
	"net/http"
)

type HTTPClient struct {
	client *http.Client
}

func New() *HTTPClient {
	return &HTTPClient{
		client: &http.Client{},
	}
}

func (h HTTPClient) Get(uri string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed in creating requests: %w", err)
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %w", err)
	}

	return resp, nil
}
