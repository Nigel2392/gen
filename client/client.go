//go:build !js && !wasm
// +build !js,!wasm

// File might show errors depending on your GOOS and GOARCH

package client

import (
	"errors"
	"net/http"
)

// APIClient is a client that can be used to execute http requests.
// - Can be used to execute GET, POST, PUT, DELETE, PATCH requests.
type APIClient struct {
	// Client is the http client that will be used to execute the request.
	client *http.Client
	// Request is the request that will be executed.
	request *http.Request
	// Error func is used to handle errors that occur during the request.
	// - Can be specified, but is not required.
	// - If not set will panic.
	// - If set and returns true will panic
	// - Always recovers from panics when set
	errorFunc     func(err error) bool
	alwaysRecover bool
}

// Execute the request -> APIClient.exec
func (c *APIClient) Do(cb func(resp *http.Response)) {
	if c.request == nil {
		c.errorFunc(errors.New(ErrNoRequest))
	} else if cb == nil {
		c.errorFunc(errors.New(ErrNoCallback))
	}
	go func() {
		c.exec(cb)
	}()
}

// Execute the request
func (c *APIClient) exec(cb func(resp *http.Response)) {
	var resp, err = c.client.Do(c.request)
	c.clientErr(err)
	defer resp.Body.Close()
	c.clientErr(err)
	cb(resp)
}
