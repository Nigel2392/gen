package elems

import (
	"reflect"
	"regexp"
	"strings"
)

func StyleSheet(href string, rel string) *Element {
	sheet := Elem("link")
	sheet.Custom("rel", rel)
	sheet.Custom("href", href)
	sheet.NoClose()
	return sheet
}
func Script(src string, typ ...string) *Element {
	script := Elem("script")
	script.Custom("src", src)
	if len(typ) > 0 {
		script.A_Type(typ[0])
	}
	return script
}
func ScriptInline(sourceCode string, typ ...string) *Element {
	script := Elem("script")
	script.InnerText(sourceCode)
	if len(typ) > 0 {
		script.A_Type(typ[0])
	}
	return script
}
func StyleBlock(t ...string) *Element {
	var v = Elem("style")
	v.A_Type("text/css")
	if len(t) > 0 {
		var sourceCode = strings.Join(t, "\n")
		sourceCode = strings.ReplaceAll(sourceCode, "\n\n", "\n")
		var re = regexp.MustCompile(`\s+([a-z-]+ *:[^:;]+;)\n`)
		sourceCode = re.ReplaceAllString(sourceCode, "$1")
		v.InnerHTML(sourceCode)
	}
	return v
}
func A(t ...string) *Element {
	var v = Elem("a", t...)
	return v
}
func Abbr(t ...string) *Element {
	var v = Elem("abbr", t...)
	return v
}
func Address(t ...string) *Element {
	var v = Elem("address", t...)
	return v
}
func Area(t ...string) *Element {
	var v = Elem("area", t...)
	return v
}
func Article(t ...string) *Element {
	var v = Elem("article", t...)
	return v
}
func Aside(t ...string) *Element {
	var v = Elem("aside", t...)
	return v
}
func Audio(t ...string) *Element {
	var v = Elem("audio", t...)
	return v
}
func B(t ...string) *Element {
	var v = Elem("b", t...)
	return v
}
func Bdi(t ...string) *Element {
	var v = Elem("bdi", t...)
	return v
}
func Bdo(t ...string) *Element {
	var v = Elem("bdo", t...)
	return v
}
func Blockquote(t ...string) *Element {
	var v = Elem("blockquote", t...)
	return v
}
func Body(t ...string) *Element {
	var v = Elem("body", t...)
	return v
}
func Br() *Element {
	var v = Elem("br").NoClose()
	return v
}
func Button(t ...string) *Element {
	var v = Elem("button", t...)
	return v
}
func Canvas(t ...string) *Element {
	var v = Elem("canvas", t...)
	return v
}
func Caption(t ...string) *Element {
	var v = Elem("caption", t...)
	return v
}
func Cite(t ...string) *Element {
	var v = Elem("cite", t...)
	return v
}
func Code(t ...string) *Element {
	var v = Elem("code", t...)
	return v
}
func Col(t ...string) *Element {
	var v = Elem("col", t...)
	return v
}
func Colgroup(t ...string) *Element {
	var v = Elem("colgroup", t...)
	return v
}
func Command(t ...string) *Element {
	var v = Elem("command", t...)
	return v
}
func Datalist(t ...string) *Element {
	var v = Elem("datalist", t...)
	return v
}
func Dd(t ...string) *Element {
	var v = Elem("dd", t...)
	return v
}
func Del(t ...string) *Element {
	var v = Elem("del", t...)
	return v
}
func Details(t ...string) *Element {
	var v = Elem("details", t...)
	return v
}
func Dfn(t ...string) *Element {
	var v = Elem("dfn", t...)
	return v
}
func Div(t ...string) *Element {
	var v = Elem("div", t...)
	return v
}
func Dl(t ...string) *Element {
	var v = Elem("dl", t...)
	return v
}
func Dt(t ...string) *Element {
	var v = Elem("dt", t...)
	return v
}
func Em(t ...string) *Element {
	var v = Elem("em", t...)
	return v
}
func Embed(t ...string) *Element {
	var v = Elem("embed", t...)
	return v
}
func Fieldset(t ...string) *Element {
	var v = Elem("fieldset", t...)
	return v
}
func Figcaption(t ...string) *Element {
	var v = Elem("figcaption", t...)
	return v
}
func Figure(t ...string) *Element {
	var v = Elem("figure", t...)
	return v
}
func Footer(t ...string) *Element {
	var v = Elem("footer", t...)
	return v
}
func Form(t ...string) *Element {
	var v = Elem("form", t...)
	return v
}
func H1(t ...string) *Element {
	var v = Elem("h1", t...)
	return v
}
func H2(t ...string) *Element {
	var v = Elem("h2", t...)
	return v
}
func H3(t ...string) *Element {
	var v = Elem("h3", t...)
	return v
}
func H4(t ...string) *Element {
	var v = Elem("h4", t...)
	return v
}
func H5(t ...string) *Element {
	var v = Elem("h5", t...)
	return v
}
func H6(t ...string) *Element {
	var v = Elem("h6", t...)
	return v
}
func Header(t ...string) *Element {
	var v = Elem("header", t...)
	return v
}
func Hr() *Element {
	var v = Elem("hr").NoClose()
	return v
}
func I(t ...string) *Element {
	var v = Elem("i", t...)
	return v
}
func Iframe(t ...string) *Element {
	var v = Elem("iframe", t...)
	return v
}
func Img(src string, alt ...string) *Element {
	var v = Elem("img").Src(src).NoClose()
	if len(alt) != 0 {
		v.Alt(alt[0])
	}
	return v
}

// TYPE -> (NAME, ID) -> PLACEHOLDER -> VALUE
func Input(t ...string) *Element {
	var v *Element = Elem("input")
	switch len(t) {
	case 0:
		v.A_Type("text")
	case 1:
		v.A_Type(t[0])
	case 2:
		v.A_Type(t[0])
		v.Name(t[1])
		v.ID(t[1])
	case 3:
		v.A_Type(t[0])
		v.Name(t[1])
		v.ID(t[1])
		v.A_Placeholder(t[2])
	case 4:
		v.A_Type(t[0])
		v.Name(t[1])
		v.ID(t[1])
		v.A_Placeholder(t[2])
		v.A_Value(t[3])
	}
	v.NoClose()
	return v
}

// TYPE -> (NAME, ID) -> PLACEHOLDER -> VALUE -> TEXT
func InputWithLabel(t ...string) (*Element, *Element, *Element) {
	var div = Div()
	var label = Label()
	var input = Input()
	switch len(t) {
	case 0:
		input.A_Type("text")
	case 1:
		input.A_Type(t[0])
	case 2:
		input.A_Type(t[0])
		input.Name(t[1])
		input.ID(t[1])
	case 3:
		input.A_Type(t[0])
		input.Name(t[1])
		input.ID(t[1])
		input.A_Placeholder(t[2])
		label.InnerText(t[2])
	case 4:
		input.A_Type(t[0])
		input.Name(t[1])
		input.ID(t[1])
		input.A_Placeholder(t[2])
		input.A_Value(t[3])
		label.InnerText(t[2])
	}
	input.NoClose()
	div.Add(label, input)
	return div, label, input
}

func Ins(t ...string) *Element {
	var v = Elem("ins", t...)
	return v
}
func Kbd(t ...string) *Element {
	var v = Elem("kbd", t...)
	return v
}
func Keygen(t ...string) *Element {
	var v = Elem("keygen", t...)
	return v
}
func Label(t ...string) *Element {
	var v = Elem("label", t...)
	return v
}
func Legend(t ...string) *Element {
	var v = Elem("legend", t...)
	return v
}
func Li(t ...string) *Element {
	var v = Elem("li", t...)
	return v
}
func Main(t ...string) *Element {
	var v = Elem("main", t...)
	return v
}
func Map(t ...string) *Element {
	var v = Elem("map", t...)
	return v
}
func Mark(t ...string) *Element {
	var v = Elem("mark", t...)
	return v
}
func Menu(t ...string) *Element {
	var v = Elem("menu", t...)
	return v
}
func Meter(t ...string) *Element {
	var v = Elem("meter", t...)
	return v
}
func Nav(t ...string) *Element {
	var v = Elem("nav", t...)
	return v
}
func Object(t ...string) *Element {
	var v = Elem("object", t...)
	return v
}
func Ol(t ...string) *Element {
	var v = Elem("ol", t...)
	return v
}
func Optgroup(t ...string) *Element {
	var v = Elem("optgroup", t...)
	return v
}
func Option(t ...string) *Element {
	var v = Elem("option", t...)
	return v
}
func Output(t ...string) *Element {
	var v = Elem("output", t...)
	return v
}
func P(t ...string) *Element {
	var v = Elem("p", t...)
	return v
}
func Param(t ...string) *Element {
	var v = Elem("param", t...)
	return v
}
func Pre(t ...string) *Element {
	var v = Elem("pre", t...)
	return v
}
func Progress(t ...string) *Element {
	var v = Elem("progress", t...)
	return v
}
func Q(t ...string) *Element {
	var v = Elem("q", t...)
	return v
}
func Rp(t ...string) *Element {
	var v = Elem("rp", t...)
	return v
}
func Rt(t ...string) *Element {
	var v = Elem("rt", t...)
	return v
}
func Ruby(t ...string) *Element {
	var v = Elem("ruby", t...)
	return v
}
func S(t ...string) *Element {
	var v = Elem("s", t...)
	return v
}
func Samp(t ...string) *Element {
	var v = Elem("samp", t...)
	return v
}
func Section(t ...string) *Element {
	var v = Elem("section", t...)
	return v
}

// Takes the text to be displayed and a list of items to be selected from
func Select(name string, listItems []any, deflt ...string) *Element {
	var v = Elem("select").Name(name).ID(name)
	if len(deflt) > 0 {
		v.Class(deflt[0])
	}
	if len(deflt) > 1 {
		v.Add(Option(deflt[1]).A_Disabled())
	}
	for _, item := range listItems {
		var refl_item = reflect.ValueOf(item)
		var str_val = valueToString(refl_item)
		v.Add(Option(str_val))
	}
	return v
}

func Small(t ...string) *Element {
	var v = Elem("small", t...)
	return v
}
func Source(t ...string) *Element {
	var v = Elem("source", t...)
	return v
}
func Span(t ...string) *Element {
	var v = Elem("span", t...)
	return v
}
func Strong(t ...string) *Element {
	var v = Elem("strong", t...)
	return v
}
func Sub(t ...string) *Element {
	var v = Elem("sub", t...)
	return v
}
func Summary(t ...string) *Element {
	var v = Elem("summary", t...)
	return v
}
func Sup(t ...string) *Element {
	var v = Elem("sup", t...)
	return v
}
func Table(t ...string) *Element {
	var v = Elem("table", t...)
	return v
}
func Tbody(t ...string) *Element {
	var v = Elem("tbody", t...)
	return v
}
func Td(t ...string) *Element {
	var v = Elem("td", t...)
	return v
}
func Textarea(t ...string) *Element {
	var v = Elem("textarea", t...)
	return v
}
func Tfoot(t ...string) *Element {
	var v = Elem("tfoot", t...)
	return v
}
func Th(t ...string) *Element {
	var v = Elem("th", t...)
	return v
}
func Thead(t ...string) *Element {
	var v = Elem("thead", t...)
	return v
}
func Time(t ...string) *Element {
	var v = Elem("time", t...)
	return v
}
func Tr(t ...string) *Element {
	var v = Elem("tr", t...)
	return v
}
func Track(t ...string) *Element {
	var v = Elem("track", t...)
	return v
}
func U(t ...string) *Element {
	var v = Elem("u", t...)
	return v
}
func Ul(t ...string) *Element {
	var v = Elem("ul", t...)
	return v
}
func Var(t ...string) *Element {
	var v = Elem("var", t...)
	return v
}
func Video(t ...string) *Element {
	var v = Elem("video", t...)
	return v
}
func Wbr(t ...string) *Element {
	var v = Elem("wbr", t...)
	return v
}
