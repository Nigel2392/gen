//go:build js && wasm
// +build js,wasm

package loaders

import (
	"encoding/json"
	"net/http"
	"syscall/js"

	"github.com/Nigel2392/gen/elems"
)

const (
	ID_LOADER_CONTAINER = "gohtml-main-loader-container"
	ID_LOADER           = "gohtml-main-loader"
)

type Loader struct {
	appendTo       string
	Activated      bool
	jsVal          js.Value
	created        bool
	deleteOnFinish bool
	className      string
	getLoaderElem  func(idContainer, idLoader string) *elems.Element
}

func NewLoader(appendTo string, className string, deleteOnFinish bool, f ...func(idContainer, idLoader string) *elems.Element) *Loader {
	var loadFunc = LoaderRing
	if len(f) > 0 {
		loadFunc = f[0]
	}
	return &Loader{
		appendTo:       appendTo,
		className:      className,
		getLoaderElem:  loadFunc,
		deleteOnFinish: deleteOnFinish,
	}
}

func (l *Loader) Stop() {
	l.Delete()
}

func (l *Loader) Run(f func()) {
	if !l.created {
		l.create()
	}
	l.activate()
	go func() {
		f()
		l.finalize()
	}()
}

func (l *Loader) finalize() {
	if l.deleteOnFinish {
		l.Delete()
	} else {
		l.Deactivate()
	}
}

func (l *Loader) RunHTTP(fetchURL string, method string, requestChanger func(rq *http.Request), cb func(resp *http.Response) error) {
	if !l.created {
		l.create()
	}
	l.activate()
	var onErr = func(err error) {
		l.finalize()
		var div = elems.Div(err.Error()).Class("alert alert-danger").Style("width:100%")
		div.WasmGenerate(l.appendTo)
	}
	go func() {
		elems.HTTPRequest(fetchURL, method, requestChanger, cb, onErr)
		l.finalize()
	}()
}

func (l *Loader) RunJsonToStruct(fetchURL string, method string, strct any, requestChanger func(rq *http.Request), cb func(resp *http.Response, s any) error) {
	// Activate or create the loader
	if !l.created {
		l.create()
	}
	l.activate()
	var onErr = func(err error) {
		l.finalize()
		var div = elems.Div(err.Error()).Class("alert alert-danger").Style("width:100%")
		div.WasmGenerate(l.appendTo)
	}
	// Create a new callback that will decode the json into the struct
	var newCallback = func(resp *http.Response) error {
		if err := json.NewDecoder(resp.Body).Decode(strct); err != nil {
			return err
		}
		// Call the original callback
		return cb(resp, strct)
	}
	// Run the async http request
	go func() {
		elems.HTTPRequest(fetchURL, method, requestChanger, newCallback, onErr)
		l.finalize()
	}()

}

func (l *Loader) create() elems.WasmValue {
	var loader_container = l.getLoaderElem(ID_LOADER_CONTAINER, ID_LOADER)
	var loader_val = loader_container.WasmGenerate(l.appendTo)
	l.jsVal = js.Value(loader_val)
	l.created = true
	return loader_val
}

func (l *Loader) activate() {
	if !l.Activated {
		l.Activated = true
		l.jsVal.Get("style").Set("display", "block")
	}
}

func (l *Loader) Deactivate() {
	if l.Activated {
		l.Activated = false
		l.jsVal.Get("style").Set("display", "none")
	}
}

func (l *Loader) Delete() {
	l.created = false
	l.jsVal.Call("remove")
}
