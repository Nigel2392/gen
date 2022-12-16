package elems

import (
	"fmt"
	"html/template"
	"runtime/debug"
	"strings"
)

func ParseToTemplate(buffer Buffer, templateData any) Buffer {
	var bufferdata = buffer.String()
	var tmpl, err = template.New("html").Parse(bufferdata)
	if err != nil {
		panic(err)
	}
	var buf Buffer = new(strings.Builder)
	// Push to buffer
	err = tmpl.Execute(buf, templateData)
	if err != nil {
		panic(err)
	}
	return buf
}

func PrintRecover() any {
	if r := recover(); r != nil {
		println(string(debug.Stack()))
		println("///////////////////////////////////////////")
		println("///")
		println(fmt.Sprintf("///	%v", r))
		println("///")
		println("///////////////////////////////////////////")
		return r
	}
	return nil
}
