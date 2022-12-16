//go:build js && wasm
// +build js,wasm

package elems

import (
	"math/rand"
	"strconv"
	"strings"
	"syscall/js"
	"time"
)

var GoHTML js.Value

func init() {
	rand.Seed(time.Now().UnixNano())
	var window = js.Global().Get("window")
	window.Set("gohtml", map[string]interface{}{})
	GoHTML = window.Get("gohtml")
	// js.Global().Get("document").Call("addEventListener", "DOMContentLoaded", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	// importScript imports a Script for use in JS
	GoHTML.Set("importScript", document_ImportScript)
	// Add an inline script with sourcecode to the head.
	GoHTML.Set("inlineScript", document_RawScript)
	// removeScript removes a script from the document.
	GoHTML.Set("removeScript", document_RemoveScript)
	// Get a list of all generated gohtml-classes, and their associated JS elements.
}

var document_RemoveScript = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	defer PrintRecover()
	var name = args[0].String()
	if name == "" {
		println("Must provide an argument!")
		return nil
	}
	var script = js.Global().Get("document").Call("getElementsByTagName", "script")
	for i := 0; i < script.Length(); i++ {
		if strings.EqualFold(script.Index(i).Get("src").String(), name) {
			script.Index(i).Call("remove")
			println("Removed script " + name + " " + strconv.Itoa(i))
		}
	}
	if script.Length() <= 0 {
		println("No scripts found!")
		var scriptStr string
		for i := 0; i < script.Length(); i++ {
			scriptStr += script.Index(i).Get("src").String() + "\n"
		}
		println(scriptStr)
	}
	return nil
})

var document_RawScript = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	defer PrintRecover()
	var document = js.Global().Get("document")
	var sourceCode = args[0].String()
	if len(args) < 2 {
		panic("Must provide an ID!")
	}
	var id = args[1].String()
	var script = document.Call("createElement", "script")
	script.Set("type", "text/javascript")
	script.Set("innerHTML", sourceCode)
	script.Set("id", id)
	document.Get("head").Call("appendChild", script)
	return nil
})

var document_ImportScript = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	defer PrintRecover()
	var document = js.Global().Get("document")
	var name = args[0].String()
	if name == "" {
		println("Must provide an argument!")
		return nil
	}
	var script = document.Call("createElement", "script")
	if len(args) > 0 {
		script.Set("src", args[0].String())
	}
	if len(args) > 1 {
		script.Set("type", args[1].String())
	} else {
		script.Set("type", "text/javascript")
	}
	document.Get("head").Call("appendChild", script)
	return nil
})
