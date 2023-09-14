package yaddress

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	baseURL        = "https://www.yaddress.net/api/address"
	defaultTimeout = 10 * time.Second
	errorAPI       = "API returned error %d: %s"
)

func NewClient(userKey string) (*client, error) {
	return &client{
		client: &http.Client{
			Timeout: defaultTimeout,
		},
		userKey: userKey,
		baseUrl: baseURL,
	}, nil
}

func (c *client) generateQueryString(req Request) string {
	// Base URL
	baseQuery := fmt.Sprintf("%s?UserKey=%s", c.baseUrl, c.userKey)

	// Add AddressLine1 if present
	if req.AddressLine1 != "" {
		baseQuery += fmt.Sprintf("&AddressLine1=%s", url.QueryEscape(req.AddressLine1))
	}

	// Add AddressLine2 if present
	if req.AddressLine2 != "" {
		baseQuery += fmt.Sprintf("&AddressLine2=%s", url.QueryEscape(req.AddressLine2))
	}

	fmt.Println(baseQuery)

	return baseQuery
}

func (c *client) ProcessAddress(req Request) (*Address, error) {
	queryString := c.generateQueryString(req)

	resp, err := c.client.Get(queryString)
	if err != nil {
		return &Address{}, err
	}

	defer resp.Body.Close()

	var r Address
	if err = json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return &Address{}, err
	}

	if r.ErrorCode != 0 {
		return &Address{}, fmt.Errorf(errorAPI, r.ErrorCode, r.ErrorMessage)
	}

	return &r, nil
}
