package yaddress

import (
	"encoding/json"
	"fmt"
	"github.com/bloomcredit/prettyzap"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"time"
)

const (
	baseURL        = "https://www.yaddress.net/api/address"
	defaultTimeout = 10 * time.Second
	errorAPI       = "API returned error %d: %s"
)

type defaultClient struct {
	httpClient restClient
	userKey    string
	baseUrl    string
	log        *zap.SugaredLogger
}

// DefaultLogger creates the standard default logger to be used for the Client
func DefaultLogger() *zap.SugaredLogger {
	log, _ := prettyzap.ConfigureZapLogger("")
	return log.Named("yaddress-client")
}

func NewClient(userKey string, opts ...Option) *defaultClient {
	client := &defaultClient{
		userKey: userKey,
		baseUrl: baseURL,
	}

	for _, opt := range opts {
		opt(client)
	}

	// Setup httpClient
	if client.httpClient == nil {

		cl := &http.Client{
			Timeout: defaultTimeout,
		}
		client.httpClient = cl
	}

	// Use noop logger if none provided
	if client.log == nil {
		client.log = zap.NewNop().Sugar()
	}

	return client
}

func (c *defaultClient) generateQueryString(req Request) string {
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

	c.log.Infof("formated query is: %s", baseQuery)

	return baseQuery
}

func (c *defaultClient) ProcessAddress(req Request) (YaddressResult, error) {
	queryString := c.generateQueryString(req)

	resp, err := c.httpClient.Get(queryString)
	if err != nil {
		return YaddressResult{}, err
	}

	defer resp.Body.Close()

	result := YaddressResult{}

	var r Address
	if err = json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return YaddressResult{}, err
	}

	if r.ErrorCode != 0 {
		err = fmt.Errorf(errorAPI, r.ErrorCode, r.ErrorMessage)
		c.log.Error(err)
		result.Debug.ErrorMessage = r.ErrorMessage
		result.Debug.ErrorCode = r.ErrorCode
		return result, err
	}

	c.log.Infof("received data: %+v\n", &r)
	result.Result = r
	return result, nil
}
