package gen

import (
	"strings"

	"github.com/Nigel2392/gen/elems"
)

type HTML struct {
	Head   *elems.Element
	Body   *elems.Element
	Buffer elems.Buffer
}

func NewHTML(title string, buf ...elems.Buffer) *HTML {
	var htmlbuffer elems.Buffer
	if len(buf) > 0 {
		htmlbuffer = buf[0]
	} else {
		// htmlbuffer = new(strings.Builder)
		htmlbuffer = new(strings.Builder)
	}
	var html = &HTML{
		Head:   elems.Elem("head"),
		Body:   elems.Elem("body"),
		Buffer: htmlbuffer,
	}
	html.Head.Add(elems.Elem("title", title))
	return html
}

func (h *HTML) String() string {
	return h.render(h.Buffer).String()
}

// Render to template.
func (h *HTML) HTML(templateData any, buffer ...elems.Buffer) string {
	if len(buffer) > 0 {
		return h.RenderBuffer(buffer[0], templateData).String()
	}
	return h.RenderBuffer(h.Buffer, templateData).String()
}

func (h *HTML) Release() {
	h.Buffer.Reset()
}

func (h *HTML) SetTitle(str string) {
	h.Head.Add(elems.Elem("title", str))
}

func (h *HTML) Add(e ...*elems.Element) *elems.Element {
	if h.Body == nil {
		h.Body = elems.Elem("body")
	}
	return h.Body.Add(e...)
}

func (h *HTML) AddHead(e ...*elems.Element) *elems.Element {
	return h.Head.Add(e...)
}

func (h *HTML) StyleSheet(href string, rel string) *elems.Element {
	sheet := elems.Elem("link")
	sheet.Custom("rel", rel)
	sheet.Custom("href", href)
	h.Head.Add(sheet)
	return sheet
}

func (h *HTML) Style(sourceCode string) *elems.Element {
	style := elems.Elem("style")
	style.InnerText(sourceCode)
	h.Head.Add(style)
	return style
}

func (h *HTML) Script(src string, typ ...string) *elems.Element {
	script := elems.Elem("script")
	script.Custom("src", src)
	if len(typ) > 0 {
		script.Custom("type", typ[0])
	}
	h.Head.Add(script)
	return script
}

func (h *HTML) ScriptInline(sourceCode string, typ ...string) *elems.Element {
	script := elems.Elem("script")
	script.InnerText(sourceCode)
	if len(typ) > 0 {
		script.A_Type(typ[0])
	}
	h.Head.Add(script)
	return script
}

func (h *HTML) ScriptOnLoad(sourceCode string, typ ...string) *elems.Element {
	script := elems.Elem("script")
	if len(typ) > 0 {
		script.A_Type(typ[0])
	}
	waiter := "document.addEventListener('DOMContentLoaded', function() {"
	ending := "});"
	script.InnerText(waiter + sourceCode + ending)
	h.Head.Add(script)
	return script
}
