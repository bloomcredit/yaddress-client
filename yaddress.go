package yaddress

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	baseURL        = "https://www.yaddress.net/api/address"
	defaultTimeout = 10 * time.Second
)

type Client struct {
	client  *http.Client
	userKey string
	baseUrl string
}

func NewClient(userKey string) (*Client, error) {
	return &Client{
		client: &http.Client{
			Timeout: defaultTimeout,
		},
		userKey: userKey,
		baseUrl: baseURL,
	}, nil
}

func (c *Client) ProcessAddress(addressLine1, addressLine2 string) (Address, error) {
	queryString := fmt.Sprintf("%s?AddressLine1=%s&AddressLine2=%s&UserKey=%s",
		c.baseUrl,
		url.QueryEscape(addressLine1),
		url.QueryEscape(addressLine2),
		url.QueryEscape(c.userKey))

	resp, err := c.client.Get(queryString)
	if err != nil {
		return Address{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Address{}, err
	}

	var r Address
	if err = json.Unmarshal(body, &r); err != nil {
		return Address{}, err
	}

	if r.ErrorCode != 0 {
		return Address{}, fmt.Errorf("API returned error %d: %s", r.ErrorCode, r.ErrorMessage)
	}

	return r, nil
}
