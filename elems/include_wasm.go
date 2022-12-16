//go:build js && wasm
// +build js,wasm

// Urls for wasm and js
package elems

import (
	"errors"
	"net/url"
	"syscall/js"
)

type URLs []URL

func (u URLs) Len() int {
	return len(u)
}

func (u URLs) Less(i, j int) bool {
	return u[i].URL < u[j].URL
}

func (u URLs) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
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
	CallBack    func(body *Element, args []js.Value, url *url.URL)
}

func SimplePaginator(app *Element, urls URLs, onErr func(err error, app *Element)) func(this js.Value, args []js.Value, url *url.URL) interface{} {
	return func(this js.Value, args []js.Value, url *url.URL) interface{} {
		app.WasmClearInnerHTML()
		var appDiv = app.Div()
		if u, err := urls.GetURL(url.Path); err != nil {
			onErr(err, appDiv)
			return nil
		} else {
			u.CallBack(appDiv, append([]js.Value{this}, args...), url)
		}
		return nil
	}
}
