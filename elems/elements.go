package elems

import (
	"strconv"
	"strings"
	"sync"
)

type Buffer interface {
	WriteString(string) (int, error)
	String() string
	Reset()
	Len() int
	Write(p []byte) (int, error)
	Grow(n int)
}

// https://github.com/golang/go/blob/master/src/html/escape.go
var htmlEscaper = strings.NewReplacer(
	`&`, "&amp;",
	`'`, "&#39;", // "&#39;" is shorter than "&apos;" and apos was not in HTML until HTML5.
	`<`, "&lt;",
	`>`, "&gt;",
	`"`, "&#34;", // "&#34;" is shorter than "&quot;".
)

const SPACING string = "    "

// List of multiple elements
type Elements []*Element

// Loop over each element in Elements.
// Return false to stop the loop
func (e Elements) ForEach(fn func(*Element) bool) {
	for _, child := range e {
		if !fn(child) {
			return
		}
	}
}

// Loop recursively over each element in Elements.
// Return false to stop the loop
func (e Elements) RecursiveForEach(maxLevel int, fn func(*Element) bool) {
	recursiveForEachElement(0, maxLevel, e, fn)
}

func recursiveForEachElement(curLevel, maxLevel int, elems []*Element, fn func(*Element) bool) {
	for _, e := range elems {
		if !fn(e) {
			return
		}
		if curLevel == 0 || curLevel < maxLevel {
			recursiveForEachElement(curLevel+1, maxLevel, e.Children(), fn)
		}
	}
}

// String representation of Elements
func (e Elements) String() string {
	elems := make([]string, 0)
	for _, e := range e {
		elems = append(elems, e.Type)
	}
	return "<<ELEMENTS: " + strings.Join(elems, ", ") + ">>"
}

// Element is a single HTML element.
type Element struct {
	text         string
	Type         string
	textAfter    bool
	autoClose    bool
	boolattrs    map[string]bool
	attrs        map[string][]string
	semicolattrs map[string][]string
	inner        Elements
}

// Get inner elements by classname
func (e *Element) GetByClassname(className string) Elements {
	var elems Elements
	var mu = sync.Mutex{}
	e.AsyncForEach(func(e *Element) bool {
		for _, c := range e.GetClasses() {
			if c == className {
				mu.Lock()
				elems = append(elems, e)
				mu.Unlock()
				return false
			}
		}
		return false
	})
	return elems
}

// Get inner elements by ID
func (e *Element) GetByID(id string) *Element {
	var elem *Element
	e.AsyncForEach(func(e *Element) bool {
		if e.GetID() == id {
			elem = e
			return true
		}
		return false
	})
	return elem
}

// Delete inner elements by ID
func (e *Element) DeleteInnerElementByID(id string) {
	var wg = sync.WaitGroup{}
	var mu = sync.Mutex{}
	go deleteInnerElementByID(e, id, &wg, &mu)
}

func deleteInnerElementByID(e *Element, id string, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	wg.Add(1)

	for i, child := range e.inner {
		if child.GetID() == id {
			mu.Lock()
			e.inner = append(e.inner[:i], e.inner[i+1:]...)
			mu.Unlock()
			return
		}
	}
	for _, child := range e.inner {
		go deleteInnerElementByID(child, id, wg, mu)
	}
}

// Get all classes of the element
func (e *Element) GetClasses() []string {
	return e.attrs["class"]
}

// Get the ID of the element
func (e *Element) GetID() string {
	attr, ok := e.attrs["id"]
	if ok {
		return attr[0]
	}
	return ""
}

// Add an element to the inner elements
func (e *Element) Add(elems ...*Element) *Element {
	if e.inner == nil {
		e.inner = make([]*Element, 0)
	}
	e.inner = append(e.inner, elems...)
	return e
}

// Set the text of the element
// Autoescape possible HTML input.
func (e *Element) InnerText(text string) *Element {
	e.text = htmlEscaper.Replace(text)
	return e
}

// Set raw HTML of the element.
func (e *Element) InnerHTML(html string) *Element {
	e.text = html
	return e
}

// Get the inner text of the element.
func (e *Element) GetText() string {
	return e.text
}

// Text will be rendered after the children are rendered.
func (e *Element) TextAfter() *Element {
	e.textAfter = true
	return e
}

// Set an attribute of the element.
func (e *Element) SetAttr(name string, values ...string) *Element {
	if e.attrs == nil {
		e.attrs = make(map[string][]string)
	}
	e.attrs[name] = append(e.attrs[name], values...)
	return e
}

// Set a boolean attribute of the element.
func (e *Element) SetBool(name string, val bool) *Element {
	if e.boolattrs == nil {
		e.boolattrs = make(map[string]bool)
	}
	e.boolattrs[name] = val
	return e
}

// Set attributes split by a semicolon.
func (e *Element) SetSemiColon(name string, values ...string) *Element {
	if e.semicolattrs == nil {
		e.semicolattrs = make(map[string][]string)
	}
	e.semicolattrs[name] = append(e.semicolattrs[name], values...)
	return e
}

// Delete an attribute of the element.
func (e *Element) DelAttr(name string) *Element {
	delete(e.attrs, name)
	delete(e.boolattrs, name)
	delete(e.semicolattrs, name)
	return e
}

// Generate the Element and children as HTML.
func (e *Element) Generate(spacing string, buf Buffer) {
	var length int
	if e.Type != "" {
		length += len(e.Type) + len(spacing) + 4
		if e.text != "" {
			length += len(e.text) + len(spacing) + len(SPACING) + 2
		}
		if !e.autoClose {
			length += len(spacing) + 5 + len(e.Type)
		}
		buf.Grow(length)

		buf.WriteString(spacing)
		buf.WriteString("<")
		buf.WriteString(e.Type)

		writeAttrs(buf, e.attrs, " ")
		writeAttrs(buf, e.semicolattrs, ";")

		for k, attr := range e.boolattrs {
			// buf.WriteString(" " + k + "=\"" + strconv.FormatBool(attr) + "\"")
			if attr {
				buf.Grow(10 + len(k))
			} else {
				buf.Grow(11 + len(k))
			}
			buf.WriteString(" ")
			buf.WriteString(k)
			buf.WriteString("=\"")
			buf.WriteString(strconv.FormatBool(attr))
			buf.WriteString("\"")
		}
		buf.WriteString(">\n")
	}

	if !e.textAfter {
		if e.text != "" {
			// buf.WriteString(spacing + SPACING + e.text + "\n")
			buf.WriteString(spacing)
			buf.WriteString(SPACING)
			buf.WriteString(e.text)
			buf.WriteString("\n")
		}
	}
	// Loop over inner html elements
	for _, e := range e.inner {
		e.Generate(spacing+SPACING, buf)
	}

	if e.textAfter {
		if e.text != "" {
			// buf.WriteString(spacing + SPACING + e.text + "\n")
			buf.WriteString(spacing)
			buf.WriteString(SPACING)
			buf.WriteString(e.text)
			buf.WriteString("\n")
		}
	}
	if e.Type != "" {
		if e.autoClose {
			// buf.WriteString(spacing + "</" + e.Type + ">\n")
			buf.WriteString(spacing)
			buf.WriteString("</")
			buf.WriteString(e.Type)
			buf.WriteString(">\n")
		}
	}
}

// Write the attributes to the buffer.
func writeAttrs(buf Buffer, attrs map[string][]string, sep string) {
	for k, attr := range attrs {
		// buf.WriteString(" " + k + "=\"" + strings.Join(attr, sep) + "\"")
		n := len(sep)*(len(attr)-1) + len(k) + 6
		for i := 0; i < len(attr); i++ {
			n += len(attr[i])
		}
		buf.Grow(n)
		buf.WriteString(" ")
		buf.WriteString(k)
		buf.WriteString("=\"")
		// buf.WriteString(strings.Join(attr, sep))
		if len(attr) == 1 {
			buf.WriteString(attr[0])
		} else {
			buf.WriteString(attr[0])
			for _, s := range attr[1:] {
				buf.WriteString(sep)
				buf.WriteString(s)
			}
		}
		buf.WriteString("\"")
	}
}

// Create a new Element.
func Elem(t string, inner ...string) *Element {
	e := newElemNoPtr(t, inner...)
	return &e
}

// Create a new Element without a pointer.
func newElemNoPtr(t string, inner ...string) Element {
	e := Element{
		Type:         t,
		autoClose:    true,
		attrs:        make(map[string][]string),
		boolattrs:    make(map[string]bool),
		semicolattrs: make(map[string][]string),
	}
	if len(inner) > 0 {
		e.text = htmlEscaper.Replace(inner[0])
	}
	return e
}

// Render to template.
func (e *Element) HTML(templateData any, buffer ...Buffer) string {
	if len(buffer) > 0 {
		return e.RenderBuffer(buffer[0], templateData).String()
	}
	// Create new buffer
	var buf = &strings.Builder{}
	return e.RenderBuffer(buf, templateData).String()
}

// String representation of element
func (e *Element) String() string {
	return e.All().String()
}

// Render element to buffer
func (e *Element) RenderBuffer(buf Buffer, templateData any) Buffer {
	e.Generate("", buf)
	return ParseToTemplate(buf, templateData)
}

// Append a child element.
func (e *Element) AppendChild(child *Element) {
	e.inner = append(e.inner, child)
}

// Loop over all inner elements recursively.
func (e *Element) ForEach(fn func(*Element)) {
	fn(e)
	for _, child := range e.inner {
		child.ForEach(fn)
	}
}

// Loop over all inner elements recursively asynchronously.
func (e *Element) AsyncForEach(fn func(*Element) bool) {
	var wg sync.WaitGroup
	var terminate = make(chan struct{}, 1)
	wg.Add(1)
	go e.asyncForEach(fn, &wg, terminate)
	wg.Wait()
	for len(terminate) > 0 {
		<-terminate
	}
	close(terminate)
}

func (e *Element) asyncForEach(fn func(*Element) bool, wg *sync.WaitGroup, terminate chan struct{}) {
	defer wg.Done()
	if fn(e) {
		terminate <- struct{}{}
		return
	}
	select {
	case <-terminate:
		return
	default:
		wg.Add(len(e.inner))
		for _, child := range e.inner {
			go child.asyncForEach(fn, wg, terminate)
		}
	}
}

// Get the element's children.
func (e *Element) Children() Elements {
	return e.inner
}

// Return all inner elements in no particular order.
func (e *Element) All() Elements {
	var all Elements
	e.ForEach(func(child *Element) {
		all = append(all, child)
	})
	return all
}

// Append a script to an element.
// If the GOOS == JS and GOARCH == wasm, the script will be appended to the document head instead.
func (e *Element) Script(src string, typ ...string) *Element {
	script := Elem("script")
	script.Custom("src", src)
	if len(typ) > 0 {
		script.Custom("type", typ[0])
	}
	e.inner = append(e.inner, script)
	return script
}

// Append a script with sourcecode to an element.
//   - If the GOOS == JS and GOARCH == wasm, the script will be appended to the document head instead.
//   - Returns nil if GOOS == JS and GOARCH == wasm.
func (e *Element) ScriptInline(sourceCode string, typ ...string) *Element {
	script := Elem("script")
	script.InnerText(sourceCode)
	if len(typ) > 0 {
		script.Custom("type", typ[0])
	}
	e.inner = append(e.inner, script)
	return script
}

func (e *Element) StyleSheet(src string, rel string) *Element {
	v := StyleSheet(src, rel)
	e.inner = append(e.inner, v)
	return v
}

////////////////////////////////////////////////////////////////////////
//
/////
//////// Element types
/////
//
////////////////////////////////////////////////////////////////////////

func (e *Element) StyleBlock(sourceCode ...string) *Element {
	v := StyleBlock(sourceCode...)
	e.AppendChild(v)
	return v
}

func (e *Element) A(t ...string) *Element {
	v := A(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Abbr(t ...string) *Element {
	v := Abbr(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Address(t ...string) *Element {
	v := Address(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Area(t ...string) *Element {
	v := Area(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Article(t ...string) *Element {
	v := Article(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Aside(t ...string) *Element {
	v := Aside(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Audio(t ...string) *Element {
	v := Audio(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) B(t ...string) *Element {
	v := B(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Bdi(t ...string) *Element {
	v := Bdi(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Bdo(t ...string) *Element {
	v := Bdo(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Blockquote(t ...string) *Element {
	v := Blockquote(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Body(t ...string) *Element {
	v := Body(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Br(t ...string) *Element {
	v := Br()
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Button(t ...string) *Element {
	v := Button(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Canvas(t ...string) *Element {
	v := Canvas(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Caption(t ...string) *Element {
	v := Caption(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Cite(t ...string) *Element {
	v := Cite(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Code(t ...string) *Element {
	v := Code(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Col(t ...string) *Element {
	v := Col(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Colgroup(t ...string) *Element {
	v := Colgroup(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Command(t ...string) *Element {
	v := Command(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Datalist(t ...string) *Element {
	v := Datalist(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Dd(t ...string) *Element {
	v := Dd(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Del(t ...string) *Element {
	v := Del(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Details(t ...string) *Element {
	v := Details(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Dfn(t ...string) *Element {
	v := Dfn(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Div(t ...string) *Element {
	v := Div(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Dl(t ...string) *Element {
	v := Dl(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Dt(t ...string) *Element {
	v := Dt(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Em(t ...string) *Element {
	v := Em(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Embed(t ...string) *Element {
	v := Embed(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Fieldset(t ...string) *Element {
	v := Fieldset(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Figcaption(t ...string) *Element {
	v := Figcaption(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Figure(t ...string) *Element {
	v := Figure(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Footer(t ...string) *Element {
	v := Footer(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Form(t ...string) *Element {
	v := Form(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) H1(t ...string) *Element {
	v := H1(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) H2(t ...string) *Element {
	v := H2(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) H3(t ...string) *Element {
	v := H3(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) H4(t ...string) *Element {
	v := H4(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) H5(t ...string) *Element {
	v := H5(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) H6(t ...string) *Element {
	v := H6(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Header(t ...string) *Element {
	v := Header(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Hr() *Element {
	v := Hr()
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) I(t ...string) *Element {
	v := I(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Iframe(t ...string) *Element {
	v := Iframe(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Img(src string, alt ...string) *Element {
	v := Img(src, alt...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Input(t ...string) *Element {
	v := Input(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) InputWithLabel(t ...string) (*Element, *Element, *Element) {
	div, label, input := InputWithLabel(t...)
	e.inner = append(e.inner, div)
	return div, label, input
}

func (e *Element) Ins(t ...string) *Element {
	v := Ins(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Kbd(t ...string) *Element {
	v := Kbd(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Keygen(t ...string) *Element {
	v := Keygen(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Label(t ...string) *Element {
	v := Label(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Legend(t ...string) *Element {
	v := Legend(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Li(t ...string) *Element {
	v := Li(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Main(t ...string) *Element {
	v := Main(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Map(t ...string) *Element {
	v := Map(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Mark(t ...string) *Element {
	v := Mark(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Menu(t ...string) *Element {
	v := Menu(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Meter(t ...string) *Element {
	v := Meter(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Nav(t ...string) *Element {
	v := Nav(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Object(t ...string) *Element {
	v := Object(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Ol(t ...string) *Element {
	v := Ol(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Optgroup(t ...string) *Element {
	v := Optgroup(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Option(t ...string) *Element {
	v := Option(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Output(t ...string) *Element {
	v := Output(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) P(t ...string) *Element {
	v := P(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Param(t ...string) *Element {
	v := Param(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Pre(t ...string) *Element {
	v := Pre(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Progress(t ...string) *Element {
	v := Progress(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Q(t ...string) *Element {
	v := Q(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Rp(t ...string) *Element {
	v := Rp(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Rt(t ...string) *Element {
	v := Rt(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Ruby(t ...string) *Element {
	v := Ruby(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) S(t ...string) *Element {
	v := S(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Samp(t ...string) *Element {
	v := Samp(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Section(t ...string) *Element {
	v := Section(t...)
	e.inner = append(e.inner, v)
	return v
}

// Optional default values:
// - Classname[0], DisplayItem[1]
func (e *Element) Select(name string, listItems []any, deflt ...string) *Element {
	v := Select(name, listItems, deflt...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Small(t ...string) *Element {
	v := Small(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Source(t ...string) *Element {
	v := Source(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Span(t ...string) *Element {
	v := Span(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Strong(t ...string) *Element {
	v := Strong(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Sub(t ...string) *Element {
	v := Sub(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Summary(t ...string) *Element {
	v := Summary(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Sup(t ...string) *Element {
	v := Sup(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Table(t ...string) *Element {
	v := Table(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Tbody(t ...string) *Element {
	v := Tbody(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Td(t ...string) *Element {
	v := Td(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Textarea(t ...string) *Element {
	v := Textarea(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Tfoot(t ...string) *Element {
	v := Tfoot(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Th(t ...string) *Element {
	v := Th(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Thead(t ...string) *Element {
	v := Thead(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Time(t ...string) *Element {
	v := Time(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Tr(t ...string) *Element {
	v := Tr(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Track(t ...string) *Element {
	v := Track(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) U(t ...string) *Element {
	v := U(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Ul(t ...string) *Element {
	v := Ul(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Var(t ...string) *Element {
	v := Var(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Video(t ...string) *Element {
	v := Video(t...)
	e.inner = append(e.inner, v)
	return v
}

func (e *Element) Wbr(t ...string) *Element {
	v := Wbr(t...)
	e.inner = append(e.inner, v)
	return v
}
