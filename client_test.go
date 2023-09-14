package yaddress

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockClient struct {
	errCode    int
	errMessage string
}

func (mc *mockClient) Get(url string) (*http.Response, error) {
	w := httptest.NewRecorder()
	w.Write([]byte(fmt.Sprintf(`{"ErrorCode":%d, "ErrorMessage":"%s"}`, mc.errCode, mc.errMessage)))
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
			name:       "EmptyFields",
			errCode:    2,
			errMessage: "Invalid address: missing City-State-Zip line",
			addr1:      "",
			addr2:      "",
			shouldErr:  true,
			errMsg:     "Should give an error if both address lines are empty",
		},
		{
			name:    "EmptyAddress1ButPresentAddress2",
			errCode: 0,
			addr1:   "",
			addr2:   "Chicago, IL",
		},
		{
			name:       "PresentAddress1ButAbsentAddress2",
			errCode:    2,
			errMessage: "Invalid address: missing City-State-Zip line",
			addr1:      "1009 S Oakley",
			addr2:      "",
			shouldErr:  true,
			errMsg:     "Should give an error, address2 is required",
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
			yd, err := NewClient("")
			assert.NoError(t, err)

			yd.client = &mockClient{errCode: tt.errCode, errMessage: tt.errMessage}
			request := Request{AddressLine1: tt.addr1, AddressLine2: tt.addr2}

			_, err = yd.ProcessAddress(request)

			if tt.shouldErr {
				assert.Error(t, err, tt.errMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
