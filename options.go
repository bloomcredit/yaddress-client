package yaddress

import (
	"go.uber.org/zap"
	"net/http"
)

// Option represents a Client option
type Option func(*defaultClient)

// WithClient sets a custom client for making outgoing API calls
func WithClient(httpClient *http.Client) Option {
	return func(client *defaultClient) {
		client.httpClient = httpClient
	}
}

// WithLogger sets a logger for the client.
// Passing `nil` or not calling this Option will
// prevent the Client from logging anything
func WithLogger(logger *zap.SugaredLogger) Option {
	return func(client *defaultClient) {
		client.log = logger
	}
}
