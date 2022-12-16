//go:build js && wasm && wails
// +build js,wasm,wails

package navbars

import (
	"syscall/js"

	"github.com/Nigel2392/gen/elems"
	"github.com/Nigel2392/gen/wailsext"
)

func WailsControlListeners(elements elems.Elements) []elems.Ready {
	var (
		minimize = elements[0]
		maximize = elements[1]
		close    = elements[2]
	)

	var _, rdy1 = minimize.AddEventListener("click", func(this js.Value, args []js.Value) interface{} {
		defer elems.PrintRecover()
		wailsext.WindowMinimise()
		return nil
	})
	var _, rdy2 = maximize.AddEventListener("click", func(this js.Value, args []js.Value) interface{} {
		defer elems.PrintRecover()
		wailsext.WindowIsMaximised(func(v bool) {
			if v {
				wailsext.WindowUnmaximise()
			} else {
				wailsext.WindowMaximise()
			}
		})
		return nil
	})
	var _, rdy3 = close.AddEventListener("click", func(this js.Value, args []js.Value) interface{} {
		defer elems.PrintRecover()
		wailsext.ExitApp()
		return nil
	})
	return []elems.Ready{rdy1, rdy2, rdy3}
}
