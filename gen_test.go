package gen_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	h "github.com/Nigel2392/gen"
	g "github.com/Nigel2392/gen/elems"
	"github.com/Nigel2392/gen/predef/navbars"
)

func TestGetByID(t *testing.T) {
	html := h.NewHTML("Test")
	html.StyleSheet("test.css", "stylesheet").NoClose()
	html.ScriptInline("alert('test');")
	html.Script("test.js")
	root := html.Body.Div()
	root.P("Test")
	root.Class("test")
	var div = root.Div().Alt("test!").BackgroundColor("black").Width("100px").Height("100px")
	div = div.Div().Alt("test!").BackgroundColor("black").Width("100px").Height("100px")
	div = div.Div().Alt("test!").BackgroundColor("black").Width("100px").Height("100px")
	div = div.Div().Alt("test!").BackgroundColor("black").Width("100px").Height("100px")
	div = div.Div().Alt("test!").BackgroundColor("black").Width("100px").Height("100px")
	div = div.Div().Alt("test!").BackgroundColor("black").Width("100px").Height("100px")
	div = div.Div().Alt("test!").BackgroundColor("black").Width("100px").Height("100px")
	div.Add(
		g.P("Test"),
		g.P("Test"),
		g.P("Test"),
		g.P("Test"),
		g.P("Test"),
		g.P("Test"),
		g.Div().Class("row").Add(
			g.Div().Class("col").Add(
				g.P("Test"),
				g.P("Test"),
				g.P("Test"),
			),
			g.Div().Class("col").Add(
				g.P("Test"),
				g.P("Test"),
				g.P("Test"),
			),
			g.Div().Class("col").Add(
				g.P("Test"),
				g.P("Test"),
				g.P("Test"),
			),
		),
	)
	div = div.Div().Alt("test!").Bgcolor("black").Width("100px").Height("100px")
	div = div.Div().Alt("test!").Bgcolor("black").Width("100px").Height("100px")
	div = div.Div().Alt("test!").Bgcolor("black").Width("100px").Height("100px")
	div = div.Div().Alt("test!").Bgcolor("black").Width("100px").Height("100px")
	div = div.Div().Alt("test!").Bgcolor("black").Width("100px").Height("100px")
	div = div.Div().Alt("test!").Bgcolor("black").Width("100px").Height("100px")
	div = div.Div().Alt("test!").Bgcolor("black").Width("100px").Height("100px")
	form := div.Form()
	form.Method("POST")
	form.Action("test")
	form.Class("test")
	form.Input().A_Type("text").Name("test").ID("test")
	div = div.Div()
	div = div.Div()
	div = div.Div()
	div = div.Div()
	div = div.Div()
	div = div.Div()
	div.Add(
		g.P("Test"),
		g.P("Test"),
		g.P("Test"),
		g.Div().Class("row").Add(
			g.Div().Class("col").Add(
				g.P("Test"),
				g.P("Test"),
				g.P("Test"),
			),
			g.Div().Class("col").Add(
				g.P("Test"),
				g.P("Test").Add(
					g.Span("Test"),
					g.Section("Test").Add(
						g.Div().Add(
							g.P("Test"),
							g.P("Test"),
							g.P("Test"),
						),
					),
				),
				g.P("Test"),
			),
			g.Div().Class("col").Add(
				g.P("Test"),
				g.P("Test"),
				g.P("Test"),
			),
		).Add(
			g.P("Test"),
			g.P("Test"),
			g.P("Test"),
			g.Div().Class("row").Add(
				g.Div().Class("col").Add(
					g.P("Test"),
					g.P("Test").Add(
						g.P("Test"),
						g.P("Test"),
						g.P("Test"),
						g.Div().Class("row").Add(
							g.Div().Class("col").Add(
								g.P("Test"),
								g.P("Test"),
								g.P("Test"),
							),
							g.Div().Class("col").Add(
								g.P("Test"),
								g.P("Test").Add(
									g.Span("Test"),
									g.Section("Test").Add(
										g.Div().Add(
											g.P("Test"),
											g.P("Test"),
											g.P("Test"),
										),
									),
								),
								g.P("Test"),
							),
							g.Div().Class("col").Add(
								g.P("Test"),
								g.P("Test"),
								g.P("Test"),
							),
						),
					),
					g.P("Test"),
				),
				g.Div().Class("col").Add(
					g.P("Test"),
					g.P("Test").Add(
						g.Span("Test"),
						g.Section("Test").Add(
							g.Div().Add(
								g.P("Test"),
								g.P("Test"),
								g.P("Test").Add(
									g.P("Test"),
									g.P("Test"),
									g.P("Test"),
									g.Div().Class("row").Add(
										g.Div().Class("col").Add(
											g.P("Test"),
											g.P("Test"),
											g.P("Test"),
										),
										g.Div().Class("col").Add(
											g.P("Test"),
											g.P("Test").Add(
												g.Span("Test"),
												g.Section("Test").Add(
													g.Div().Add(
														g.P("Test"),
														g.P("Test"),
														g.P("Test").Add(
															g.P("Test"),
															g.P("Test"),
															g.P("Test"),
															g.Div().Class("row").Add(
																g.Div().Class("col").Add(
																	g.P("Test"),
																	g.P("Test"),
																	g.P("Test"),
																),
																g.Div().Class("col").Add(
																	g.P("Test"),
																	g.P("Test").Add(
																		g.Span("Test"),
																		g.Section("Test").Add(
																			g.Div().Add(
																				g.P("Test"),
																				g.P("Test"),
																				g.P("Test"),
																			),
																		),
																	),
																	g.P("Test"),
																),
																g.Div().Class("col").Add(
																	g.P("Test"),
																	g.P("Test"),
																	g.P("Test"),
																),
															),
														),
													),
												),
											),
											g.P("Test"),
										),
										g.Div().Class("col").Add(
											g.P("Test").Class("test"),
											g.P("Test").Class("test"),
											g.P("Test").Class("test"),
										),
									),
								),
							),
						),
					),
					g.P("Test"),
				),
				g.Div().Class("col").Add(
					g.P("Test"),
					g.P("Test"),
					g.P("Test"),
				),
			),
		),
	)
	div = div.Div()
	div = div.Div().ID("GET_THIS_ID")
	div = div.Div()
	h1 := div.H1("Test")
	h1.Class("test")
	h1.Class("test2")
	h1.ID("test").Class("test")

	tim := time.Now()
	fmt.Println(root.GetByID("GET_THIS_ID").String())
	fmt.Printf("Time since: %v nanoseconds\n", time.Since(tim).Nanoseconds())
	time.Sleep(1 * time.Second)
	tim = time.Now()
	fmt.Println(root.GetByClassname("test").String())
	fmt.Printf("Time since: %v nanoseconds\n", time.Since(tim).Nanoseconds())
	//tim = time.Now()
	//fmt.Println(root.AsyncGetByClassname("test").String())
	//fmt.Printf("Time since: %v nanoseconds\n", time.Since(tim).Nanoseconds())
	//tim = time.Now()
	//fmt.Println(root.GetByIDSYNC("GET_THIS_ID").String())
	//fmt.Printf("Time since: %v\n", time.Since(tim).Microseconds())

}

func TestStructToForm(t *testing.T) {
	var testForm = struct {
		Name   string
		Email  string
		Time   time.Time
		Age    int
		Friend struct {
			Name string
			Age  int
			Time time.Time
		}
	}{
		Name:  "Nigel",
		Email: "nigel2392@gmail.com",
		Time:  time.Now(),
		Age:   22,
		Friend: struct {
			Name string
			Age  int
			Time time.Time
		}{
			Name: "Nigel",
			Age:  22,
			Time: time.Now(),
		},
	}
	var form = g.StructToForm(&testForm, "", "")
	var buf strings.Builder
	form.Generate(" ", &buf)
	t.Log(buf.String())
}

func TestHTML(t *testing.T) {
	html := h.NewHTML("Test")
	html.Body.BackgroundColor("white")
	url_home := g.URL{
		Name:        "Home",
		URL:         "/home",
		Icon:        "",
		LeftOrRight: true,
	}
	url_about := g.URL{
		Name:        "About",
		URL:         "/about",
		Icon:        "",
		LeftOrRight: false,
	}
	url_contact := g.URL{
		Name:        "Contact",
		URL:         "/contact",
		Icon:        "",
		LeftOrRight: false,
	}
	urls := g.URLs{url_home, url_about, url_contact}
	navbar, _ := navbars.NavBar(nil, urls)
	html.Add(navbar)

	html.StyleSheet("test.css", "stylesheet").NoClose()
	html.ScriptInline("alert('test');")
	html.Script("test.js")
	div := html.Body.Div()
	div.P("Test")
	div.Class("test")
	div = div.Div()
	div = div.Div()
	div = div.Div()
	h1 := div.H1("Test")
	h1.Class("test")
	h1.Class("test2")
	h1.ID("test")
	h1.Style("color:red", "font-size:12px")

	for i := 0; i < 3; i++ {
		t.Log(html.String())
		t.Log("LENGTH OF BUFFER:", html.Buffer.Len())
		html.Head.ForEach(func(e *g.Element) {
			t.Logf("HEADER: %s \n", e.Type)
		})
		html.Body.ForEach(func(e *g.Element) {
			t.Logf("BODY: %s \n", e.Type)
		})
		html.Buffer = &strings.Builder{}
	}
}

func BenchmarkHTML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		html := h.NewHTML("Test")
		html.StyleSheet("test.css", "stylesheet").NoClose()
		html.ScriptInline("alert('test');")
		html.Script("test.js")
		div := html.Body.Div()
		div.P("Test")
		div.Class("test")
		div = div.Div().Alt("test!").BackgroundColor("black").Width("100px").Height("100px")
		div = div.Div().Alt("test!").BackgroundColor("black").Width("100px").Height("100px")
		div = div.Div().Alt("test!").BackgroundColor("black").Width("100px").Height("100px")
		div = div.Div().Alt("test!").BackgroundColor("black").Width("100px").Height("100px")
		div = div.Div().Alt("test!").BackgroundColor("black").Width("100px").Height("100px")
		div = div.Div().Alt("test!").BackgroundColor("black").Width("100px").Height("100px")
		div = div.Div().Alt("test!").BackgroundColor("black").Width("100px").Height("100px")
		div.Add(
			g.P("Test"),
			g.P("Test"),
			g.P("Test"),
			g.P("Test"),
			g.P("Test"),
			g.P("Test"),
			g.Div().Class("row").Add(
				g.Div().Class("col").Add(
					g.P("Test"),
					g.P("Test"),
					g.P("Test"),
				),
				g.Div().Class("col").Add(
					g.P("Test"),
					g.P("Test"),
					g.P("Test"),
				),
				g.Div().Class("col").Add(
					g.P("Test"),
					g.P("Test"),
					g.P("Test"),
				),
			),
		)
		div = div.Div().Alt("test!").Bgcolor("black").Width("100px").Height("100px")
		div = div.Div().Alt("test!").Bgcolor("black").Width("100px").Height("100px")
		div = div.Div().Alt("test!").Bgcolor("black").Width("100px").Height("100px")
		div = div.Div().Alt("test!").Bgcolor("black").Width("100px").Height("100px")
		div = div.Div().Alt("test!").Bgcolor("black").Width("100px").Height("100px")
		div = div.Div().Alt("test!").Bgcolor("black").Width("100px").Height("100px")
		div = div.Div().Alt("test!").Bgcolor("black").Width("100px").Height("100px")
		form := div.Form()
		form.Method("POST")
		form.Action("test")
		form.Class("test")
		form.Input().A_Type("text").Name("test").ID("test")
		div = div.Div()
		div = div.Div()
		div = div.Div()
		div = div.Div()
		div = div.Div()
		div = div.Div()
		div.Add(
			g.P("Test"),
			g.P("Test"),
			g.P("Test"),
			g.Div().Class("row").Add(
				g.Div().Class("col").Add(
					g.P("Test"),
					g.P("Test"),
					g.P("Test"),
				),
				g.Div().Class("col").Add(
					g.P("Test"),
					g.P("Test").Add(
						g.Span("Test"),
						g.Section("Test").Add(
							g.Div().Add(
								g.P("Test"),
								g.P("Test"),
								g.P("Test"),
							),
						),
					),
					g.P("Test"),
				),
				g.Div().Class("col").Add(
					g.P("Test"),
					g.P("Test"),
					g.P("Test"),
				),
			).Add(
				g.P("Test"),
				g.P("Test"),
				g.P("Test"),
				g.Div().Class("row").Add(
					g.Div().Class("col").Add(
						g.P("Test"),
						g.P("Test").Add(
							g.P("Test"),
							g.P("Test"),
							g.P("Test"),
							g.Div().Class("row").Add(
								g.Div().Class("col").Add(
									g.P("Test"),
									g.P("Test"),
									g.P("Test"),
								),
								g.Div().Class("col").Add(
									g.P("Test"),
									g.P("Test").Add(
										g.Span("Test"),
										g.Section("Test").Add(
											g.Div().Add(
												g.P("Test"),
												g.P("Test"),
												g.P("Test"),
											),
										),
									),
									g.P("Test"),
								),
								g.Div().Class("col").Add(
									g.P("Test"),
									g.P("Test"),
									g.P("Test"),
								),
							),
						),
						g.P("Test"),
					),
					g.Div().Class("col").Add(
						g.P("Test"),
						g.P("Test").Add(
							g.Span("Test"),
							g.Section("Test").Add(
								g.Div().Add(
									g.P("Test"),
									g.P("Test"),
									g.P("Test").Add(
										g.P("Test"),
										g.P("Test"),
										g.P("Test"),
										g.Div().Class("row").Add(
											g.Div().Class("col").Add(
												g.P("Test"),
												g.P("Test"),
												g.P("Test"),
											),
											g.Div().Class("col").Add(
												g.P("Test"),
												g.P("Test").Add(
													g.Span("Test"),
													g.Section("Test").Add(
														g.Div().Add(
															g.P("Test"),
															g.P("Test"),
															g.P("Test").Add(
																g.P("Test"),
																g.P("Test"),
																g.P("Test"),
																g.Div().Class("row").Add(
																	g.Div().Class("col").Add(
																		g.P("Test"),
																		g.P("Test"),
																		g.P("Test"),
																	),
																	g.Div().Class("col").Add(
																		g.P("Test"),
																		g.P("Test").Add(
																			g.Span("Test"),
																			g.Section("Test").Add(
																				g.Div().Add(
																					g.P("Test"),
																					g.P("Test"),
																					g.P("Test"),
																				),
																			),
																		),
																		g.P("Test"),
																	),
																	g.Div().Class("col").Add(
																		g.P("Test"),
																		g.P("Test"),
																		g.P("Test"),
																	),
																),
															),
														),
													),
												),
												g.P("Test"),
											),
											g.Div().Class("col").Add(
												g.P("Test"),
												g.P("Test"),
												g.P("Test"),
											),
										),
									),
								),
							),
						),
						g.P("Test"),
					),
					g.Div().Class("col").Add(
						g.P("Test"),
						g.P("Test"),
						g.P("Test"),
					),
				),
			),
		)
		div = div.Div()

		div = div.Div()
		div = div.Div()
		h1 := div.H1("Test")
		h1.Class("test")
		h1.Class("test2")
		h1.ID("GET_THIS_ID")
		h1.Style("color:red", "font-size:12px")
		_ = html.String()
		html.Release()
	}
}
func BenchmarkHTMLTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		html := h.NewHTML("Test")
		html.StyleSheet("test.css", "stylesheet").NoClose()
		html.ScriptInline("alert('test');")
		html.Script("test.js")
		for i := 0; i < 10000; i++ {
			div := html.Body.Div()
			div = div.Div()
			div = div.Div()
			div.H1("Test")
		}
		b.StartTimer()
		_ = html.String()
	}
}
