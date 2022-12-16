//go:build js && wasm
// +build js,wasm

package client

// File might show errors depending on your GOOS and GOARCH

import (
	"errors"
	"net/http"

	"github.com/Nigel2392/gen/cookies"
	"github.com/Nigel2392/gen/predef/loaders"
)

type APIClient struct {
	// Client is the http client that will be used to execute the request.
	client *http.Client
	// Request is the request that will be executed.
	request *http.Request
	// Loader element is used to show a loader when the request is being executed.
	// This is only used in the WASM version of the client.
	ld *loaders.Loader
	// Error func is used to handle errors that occur during the request.
	// - Can be specified, but is not required.
	// - If not set will panic.
	// - If set and returns true will panic
	// - Always recovers from panics when set
	errorFunc     func(err error) bool
	alwaysRecover bool
}

// WithLoader sets the loader that will be used to show a loader when the request is being executed.
func (c *APIClient) WithLoader(ld *loaders.Loader) *APIClient {
	c.ld = ld
	return c
}

// Execute the request
func (c *APIClient) Do(cb func(resp *http.Response)) {
	if c.request == nil {
		c.errorFunc(errors.New(ErrNoRequest))
	} else if cb == nil {
		c.errorFunc(errors.New(ErrNoCallback))
	}
	if c.ld != nil {
		c.ld.Run(func() {
			c.exec(cb)
		})
	} else {
		go func() {
			c.exec(cb)
		}()
	}
}

// Add specified cookies from cookies.GlobalJar to the request
func (ac *APIClient) WithCookies(c ...string) *APIClient {
	if ac.request == nil {
		ac.errorFunc(errors.New(ErrNoRequest))
	}
	cookies.GlobalJar.Dump(ac.request, c...)
	return ac
}

// Execute the request -> APIClient.exec
func (c *APIClient) exec(cb func(resp *http.Response)) {
	var resp, err = c.client.Do(c.request)
	c.clientErr(err)
	defer resp.Body.Close()
	if err == nil {
		err = cookies.GlobalJar.SetHTTPCookies(resp.Cookies())
		c.clientErr(err)
	}
	c.clientErr(err)
	cb(resp)
}
