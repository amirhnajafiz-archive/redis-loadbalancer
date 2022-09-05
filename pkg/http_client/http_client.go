package http_client

import "net/http"

type HTTPClient struct {
	client *http.Client
}

func New() *HTTPClient {
	return &HTTPClient{
		client: &http.Client{},
	}
}
