# HTMLgen!

A package to help you write WASM, or HTML inside of Golang!


## Installation
Easily install  with the following command:
```
go get github.com/Nigel2392/gen
```

## Usage

### Example of form submissions:
```go
var body = g.Div()
var form = body.Form()
form.Method("POST").Action("")
form.Label("Name")
form.Input("text")                       // Type
form.Input("text", "name")               // Type, name and id
form.Input("submit", "submit", "Submit") // Type, name and display text
var _, rdy = form.WasmFormSubmit(func(data map[string]string) {
	println(fmt.Sprintf("Form submitted: %v", data))
})
// Generate the HTML inside of wasm, append it to and existing html element.
body.WasmGenerate(elems.ID_MAIN_CONTENT, rdy)
```