package yaddress

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockClient struct {
	errCode    int
	errMessage string
	httpError  bool
}

func (mc *mockClient) Get(url string) (*http.Response, error) {
	w := httptest.NewRecorder()
	w.Write([]byte(fmt.Sprintf(`{"ErrorCode":%d, "ErrorMessage":"%s"}`, mc.errCode, mc.errMessage)))

	if mc.httpError {
		return nil, errors.New("mock http error")
	}
	return w.Result(), nil
}

func TestYaddress(t *testing.T) {
	type test struct {
		name       string
		errCode    int
		errMessage string
		addr1      string
		addr2      string
		shouldErr  bool
		errMsg     string
	}

	tests := []test{
		{
			name:      "HttpError",
			shouldErr: true,
			addr1:     "",
			addr2:     "",
		},
		{
			name:       "EmptyFields",
			errCode:    2,
			errMessage: "Invalid address: missing City-State-Zip line",
			addr1:      "",
			addr2:      "",
			shouldErr:  true,
			errMsg:     "Should give an error if both address lines are empty",
		},
		{
			name:    "PresentAddress1AndAddress2",
			errCode: 0,
			addr1:   "506 Fourth Avenue Unit 1",
			addr2:   "Asbury Prk, NJ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := DefaultLogger()
			yd := NewClient("", WithLogger(logger))

			yd.httpClient = &mockClient{errCode: tt.errCode, errMessage: tt.errMessage, httpError: tt.shouldErr}
			request := Request{AddressLine1: tt.addr1, AddressLine2: tt.addr2}

			_, err := yd.ProcessAddress(request)

			if tt.shouldErr {
				assert.Error(t, err, tt.errMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
