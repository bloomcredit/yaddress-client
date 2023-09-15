package yaddress

import "net/http"

type restClient interface {
	Get(string) (*http.Response, error)
}
