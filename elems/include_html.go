//go:build !js && !wasm
// +build !js,!wasm

// Urls for normal build
package elems

import "errors"

type URLs []URL

func (u URLs) Len() int {
	return len(u)
}

func (u URLs) GetName(name string) (URL, error) {
	for _, v := range u {
		if v.Name == name {
			return v, nil
		}
	}
	return URL{}, errors.New("name not found")
}

func (u URLs) GetURL(url string) (URL, error) {
	for _, v := range u {
		if v.URL == url {
			return v, nil
		}
	}
	return URL{}, errors.New("url not found")
}

type URL struct {
	Name        string // Name of the URL
	URL         string // URL is the URL to the resource.
	Icon        string // Icon is the URL to the icon of the resource.
	LeftOrRight bool   // true = left, false = right
}
