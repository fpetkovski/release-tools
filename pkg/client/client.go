package client

import (
	"encoding/base64"
	"net/http"
	"time"
)

func NewBasicAuthClient(username, token string) *http.Client {
	client := http.Client{Timeout: 10 * time.Second}
	client.Transport = &transport{
		rt:       http.DefaultTransport,
		username: username,
		token:    token,
	}

	return &client
}

type transport struct {
	rt       http.RoundTripper
	username string
	token    string
}

func (t *transport) RoundTrip(request *http.Request) (*http.Response, error) {
	if request.Header.Get("Authorization") == "" {
		auth := t.username + ":" + t.token
		encodedValue := base64.StdEncoding.EncodeToString([]byte(auth))
		request.Header.Set("Authorization", "Basic "+encodedValue)
	}

	return t.rt.RoundTrip(request)
}
