//go:build js && wasm
// +build js,wasm

package elems

import (
	"net/http"
	"net/url"
	"strings"
	"syscall/js"

	"github.com/Nigel2392/gen/cookies"
	"github.com/Nigel2392/gen/genstring"
)

type Ready chan bool

func (r Ready) Finalize() {
	r <- true
	close(r)
}

type WasmValue js.Value

func (v WasmValue) Parent() WasmValue {
	return WasmValue((js.Value)(v).Get("parentNode"))
}

func (v WasmValue) AddEventListener(eventName string, listener func(this js.Value, args []js.Value) interface{}) WasmValue {
	js.Value(v).Call("addEventListener", eventName, listener)
	return v
}

func (v WasmValue) RemoveEventListener(eventName string) WasmValue {
	js.Value(v).Call("removeEventListener", eventName)
	return v
}

func (v WasmValue) Remove() WasmValue {
	js.Value(v).Call("remove")
	return v
}

func (v WasmValue) SetInnerHTML(html string) WasmValue {
	js.Value(v).Set("innerHTML", html)
	return v
}

func (v WasmValue) Set(name string, value any) WasmValue {
	js.Value(v).Set(name, value)
	return v
}

func (v WasmValue) Get(name string) WasmValue {
	return WasmValue(js.Value(v).Get(name))
}

func (v WasmValue) GetID() string {
	return js.Value(v).Get("id").String()
}

func (v WasmValue) GetClasses() []string {
	var classes = js.Value(v).Get("classList")
	var classList = make([]string, classes.Length())
	for i := 0; i < classes.Length(); i++ {
		classList[i] = classes.Index(i).String()
	}
	return classList
}

func (v WasmValue) AppendChild(child WasmValue) WasmValue {
	js.Value(v).Call("appendChild", js.Value(child))
	return v
}

func (v WasmValue) InsertBefore(child WasmValue, before WasmValue) WasmValue {
	js.Value(v).Call("insertBefore", js.Value(child), js.Value(before))
	return v
}

func (v WasmValue) InsertAfter(child WasmValue, after WasmValue) WasmValue {
	js.Value(v).Call("insertAfter", js.Value(child), js.Value(after))
	return v
}

func (v WasmValue) RemoveChild(child WasmValue) WasmValue {
	js.Value(v).Call("removeChild", js.Value(child))
	return v
}

func (v WasmValue) ReplaceChild(child WasmValue, before WasmValue) WasmValue {
	js.Value(v).Call("replaceChild", js.Value(child), js.Value(before))
	return v
}

func (v WasmValue) GetChildNodes() []WasmValue {
	var nodes = js.Value(v).Get("childNodes")
	var nodeList = make([]WasmValue, nodes.Length())
	for i := 0; i < nodes.Length(); i++ {
		nodeList[i] = WasmValue(nodes.Index(i))
	}
	return nodeList
}

func (v WasmValue) GetChildNode(index int) WasmValue {
	return WasmValue(js.Value(v).Get("childNodes").Index(index))
}

func (v WasmValue) GetChildNodeCount() int {
	return js.Value(v).Get("childNodes").Length()
}

// Clear the inner html of an element.
func (e *Element) WasmClearInnerHTML() {
	e.text = ""
	e.inner = make(Elements, 0)
	if e.GetID() != "" {
		js.Global().Get("document").Call("getElementById", e.GetID()).Set("innerHTML", "")
	}
	//var classes = e.GetClasses()
	//if len(classes) > 0 {
	//	for _, class := range classes {
	//		var jsval = js.Global().Get("document").Call("querySelectorAll", "."+class)
	//		if jsval.Length() > 1 {
	//			continue
	//		} else {
	//			jsval.Set("innerHTML", "")
	//			break
	//		}
	//	}
	//}
}

// Add an eventlistener, needs to be activated with AddEventListener().Finalize() after rendering to document.
func (e *Element) AddEventListener(eventName string, listener func(this js.Value, args []js.Value) interface{}) (string, Ready) {
	var className = genstring.RandStringBytesMaskImprSrcUnsafe(16)
	className = "gohtml-" + className
	e.Class(className) // fix 1.1.7.4
	var ready = addListener(eventName, className, listener)
	return className, ready
}

// Add an eventlistener for all elements, needs to be activated with AddEventListeners().Finalize() after rendering to document.
func (e Elements) AddEventListeners(eventName string, listener func(this js.Value, args []js.Value) interface{}) (string, Ready) {
	var className = genstring.RandStringBytesMaskImprSrcUnsafe(16)
	var rdy Ready
	className = "gohtml-" + className
	e.ForEach(func(element *Element) bool {
		element.Class(className)
		return true
	})
	rdy = addListener(eventName, className, listener)
	return className, rdy
}

// Add an eventlistener, needs to be activated with addListener().Finalize() after rendering to document.
func addListener(eventName, className string, listener func(this js.Value, args []js.Value) interface{}) Ready {
	var ready = make(Ready)
	go func() {
		<-ready
		println("Adding listener " + eventName + " for " + className)
		var document = js.Global().Get("document")
		var elements = document.Call("getElementsByClassName", className)
		for i := 0; i < elements.Length(); i++ {
			elements.Index(i).Call("addEventListener", eventName, js.FuncOf(listener))
		}
	}()
	return ready
}

// Add an eventlistener to toggle gohtml-active class from a element class.
func (elements *Elements) ActiveToggleListener(fn ...func(this js.Value, args []js.Value, url *url.URL) any) Ready {
	var _, rdy = elements.AddEventListeners("click", func(this js.Value, args []js.Value) any {
		// Stop the default action, if needed
		var event = args[0]
		event.Call("preventDefault")
		// Toggle active class
		var className = this.Get("className").String()
		var elements = js.Global().Get("document").Call("getElementsByClassName", className)
		for i := 0; i < elements.Length(); i++ {
			// Disable active class on all elements
			elements.Index(i).Get("classList").Call("remove", "gohtml-active")
		}
		// Enable active class on clicked element
		this.Get("classList").Call("add", "gohtml-active")
		if len(fn) > 0 {
			url, err := url.Parse(this.Get("href").String())
			if err != nil {
				panic(err)
			}
			return fn[0](this, args, url)
		}
		return nil
	})
	return rdy
}

func (e *Element) WasmFormSubmit(f func(data map[string]string, elements []js.Value)) (string, Ready) {
	if !strings.EqualFold(e.Type, "form") {
		panic("WasmFormSubmit() can only be used on form elements")
	}
	var className, rdy = e.AddEventListener("submit", func(this js.Value, args []js.Value) interface{} {
		var event = args[0]
		event.Call("preventDefault")
		var data = make(map[string]string)
		var elements = this.Get("elements")
		var length = elements.Get("length").Int()
		var elemList = make([]js.Value, length)
		for i := 0; i < length; i++ {
			var element = elements.Index(i)
			var name = element.Get("name").String()
			var value = element.Get("value").String()
			data[name] = value
			elemList[i] = element
		}
		f(data, elemList)
		return nil
	})
	return className, rdy
}

func (e *Element) WasmClearForm() {
	var id = e.GetID()
	var form = js.Global().Get("document").Call("getElementById", id)
	if !form.IsUndefined() && !form.IsNull() {
		form.Call("reset")
	}
}

func (e *Element) WasmFormSubmitToStruct(strct any, f func(strct any, elements []js.Value)) (string, Ready) {
	if !strings.EqualFold(e.Type, "form") {
		panic("WasmFormSubmit() can only be used on form elements")
	}
	var className, rdy = e.WasmFormSubmit(func(data map[string]string, elements []js.Value) {
		var err = FormDataToStruct(data, strct)
		if err != nil {
			panic(err)
		}
		f(strct, elements)
	})
	return className, rdy
}

// Returns root of generated HTML
func (e *Element) WasmGenerate(appendTo string, rdys ...Ready) WasmValue {
	var body js.Value
	// Append to body
	if appendTo != "" {
		body = js.Global().Get("document").Call("getElementById", appendTo)
	} else {
		body = js.Global().Get("document").Get("body")
	}
	var root = e.wasmGen(body)
	if len(rdys) > 0 {
		for _, rdy := range rdys {
			if rdy != nil {
				rdy.Finalize()
			}
		}
	}
	return WasmValue(root)
}

// Generate HTML with JS calls inside of WASM.
func (e *Element) wasmGen(last js.Value) js.Value {
	var current = js.Global().Get("document").Call("createElement", e.Type)

	last.Call("appendChild", current)

	// Set attributes
	for k, attr := range e.attrs {
		// current.Set(k, strings.Join(attr, " "))
		current.Call("setAttribute", k, strings.Join(attr, " "))
	}
	for k, attr := range e.semicolattrs {
		// current.Set(k, strings.Join(attr, ";"))
		current.Call("setAttribute", k, strings.Join(attr, ";"))
	}
	for k, attr := range e.boolattrs {
		// current.Set(k, attr)
		current.Call("setAttribute", k, attr)
	}
	if !e.textAfter {
		if e.text != "" {
			current.Call("insertAdjacentHTML", "afterbegin", e.text)
		}
	}
	for _, child := range e.inner {
		child.wasmGen(current)
	}
	if e.textAfter {
		if e.text != "" {
			current.Call("insertAdjacentHTML", "beforeend", e.text)
		}
	}
	return current
}

// Append a script tag with src attribute to the head of the document
func (e *Element) WasmScript(src string, typ ...string) js.Value {
	return WasmScript(src, typ...)
}

// Append a script tag from source code to the head of the document
func (e *Element) WasmScriptInline(sourceCode string, typ ...string) js.Value {
	return WasmScriptInline(sourceCode, typ...)
}

// Append a script tag with src attribute to the head of the document
func WasmScript(src string, typ ...string) js.Value {
	var script = js.Global().Get("document").Call("createElement", "script")
	script.Set("src", src)
	if len(typ) > 0 {
		script.Set("type", typ[0])
	}
	js.Global().Get("document").Get("head").Call("appendChild", script)
	return script
}

// Append a script tag from source code to the head of the document
func WasmScriptInline(sourceCode string, typ ...string) js.Value {
	var script = js.Global().Get("document").Call("createElement", "script")
	script.Set("innerHTML", sourceCode)
	if len(typ) > 0 {
		script.Set("type", typ[0])
	}
	js.Global().Get("document").Get("head").Call("appendChild", script)
	return script
}

// Append a style tag with src attribute to the head of the document
func (e *Element) WasmStyle(src string, typ ...string) js.Value {
	return WasmStyle(src, typ...)
}

// Append a style tag from source code to the head of the document
func (e *Element) WasmStyleInline(sourceCode string, typ ...string) js.Value {
	return WasmStyleInline(sourceCode, typ...)
}

// Append a style tag with src attribute to the head of the document
func WasmStyle(src string, typ ...string) js.Value {
	var style = js.Global().Get("document").Call("createElement", "link")
	style.Set("rel", "stylesheet")
	style.Set("href", src)
	if len(typ) > 0 {
		style.Set("type", typ[0])
	}
	js.Global().Get("document").Get("head").Call("appendChild", style)
	return style
}

// Append a style tag from source code to the head of the document
func WasmStyleInline(sourceCode string, typ ...string) js.Value {
	var style = js.Global().Get("document").Call("createElement", "style")
	style.Set("innerHTML", sourceCode)
	if len(typ) > 0 {
		style.Set("type", typ[0])
	}
	js.Global().Get("document").Get("head").Call("appendChild", style)
	return style
}

// Used to asynchronously create a http request for javascript with a callback function.
// Most useful for parsing data, since the url and method is gained from the javascript call.
func JSAsyncHTTPRequest(cb func(r *http.Request) (js.Value, error)) func() js.Func {
	return func() js.Func {
		return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			var requestMethod string
			var requestUrl = args[0].String() // Get the URL from the arguments
			if len(args) > 1 {
				requestMethod = args[1].String() // Get the method from the arguments
			} else {
				requestMethod = "GET" // Default to GET
			}
			var handler = asyncHttpHandler(requestUrl, requestMethod, cb) // Create promise handler
			var promiseConstructor = js.Global().Get("Promise")           // Create promise
			return promiseConstructor.New(handler)                        // Return promise
		})
	}
}

// Asynchronously fetch from a predefined URL/method with a predefined callback function.
func JSAsyncHTTPRequestTo(fetchURL, method string, cb func(r *http.Request) (js.Value, error)) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var handler = asyncHttpHandler(fetchURL, method, cb) // Create promise handler
		var promiseConstructor = js.Global().Get("Promise")  // Create promise
		return promiseConstructor.New(handler)               // Return promise
	})
}

// func GenericHTTPRequest(fetchURL, method, encoding string, )

func HTTPRequest(fetchURL, method string, requestChanger func(rq *http.Request), cb func(resp *http.Response) error, onError func(err error)) { //chan struct{} {
	// var done = make(chan struct{})
	var client = http.Client{}
	var req, err = http.NewRequest(method, fetchURL, nil)
	if err != nil {
		onError(err)
		return //nil
	}
	if requestChanger != nil {
		requestChanger(req)
	}
	resp, err := client.Do(req)
	if err != nil {
		onError(err)
		return
	}
	defer resp.Body.Close()

	err = cookies.GlobalJar.SetHTTPCookies(resp.Cookies())

	if err != nil {
		onError(err)
		return
	}

	err = cb(resp)
	if err != nil {
		onError(err)
		return
	}
}

// Return error promise if occurred
func WasmErrorObjIf(err error, promise js.Value) bool {
	if err != nil {
		var errorConstructor = js.Global().Get("Error")     // Error constructor
		var errorObject = errorConstructor.New(err.Error()) // Error object
		promise.Invoke(errorObject)                         // Reject promise
		return true
	}
	return false // No error
}

// AsyncHTTPHandler creates a http request handler for a promise
func asyncHttpHandler(fetchURL, method string, cb func(r *http.Request) (js.Value, error)) js.Func {
	var handler = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var resolve = args[0]
		var reject = args[1]
		go func() {
			// Create request
			var res, err = http.DefaultClient.Do(&http.Request{
				Method: fetchURL,
				URL:    &url.URL{Path: method},
			})

			// Return error if one occurred
			if ret := WasmErrorObjIf(err, reject); ret {
				return
			}

			// Close body when done
			defer res.Body.Close()

			// Set cookies
			err = cookies.GlobalJar.SetHTTPCookies(res.Cookies())
			// Return error if one occurred
			if ret := WasmErrorObjIf(err, reject); ret {
				return
			}

			// Get the response from the callback, response will be a js.Value
			data, err := cb(res.Request)
			// Return error if one occurred
			if ret := WasmErrorObjIf(err, reject); ret {
				return
			}

			// Create response object
			var responseConstructor = js.Global().Get("Response")
			var response = responseConstructor.New(data)

			// Resolve the promise
			resolve.Invoke(response)
		}()
		return nil
	})
	return handler
}
