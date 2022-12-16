package elems_test

import (
	"testing"

	"github.com/Nigel2392/gen/elems"
)

func getHTML() *elems.Element {
	html := elems.Div()
	html.StyleSheet("test.css", "stylesheet").NoClose()
	div := html.Div()
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
		elems.P("Test"),
		elems.P("Test"),
		elems.P("Test"),
		elems.P("Test"),
		elems.P("Test"),
		elems.P("Test"),
		elems.Div().Class("row").Add(
			elems.Div().Class("col").Add(
				elems.P("Test"),
				elems.P("Test"),
				elems.P("Test"),
			),
			elems.Div().Class("col").Add(
				elems.P("Test"),
				elems.P("Test"),
				elems.P("Test"),
			),
			elems.Div().Class("col").Add(
				elems.P("Test"),
				elems.P("Test"),
				elems.P("Test"),
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
		elems.P("Test"),
		elems.P("Test"),
		elems.P("Test"),
		elems.Div().Class("row").Add(
			elems.Div().Class("col").Add(
				elems.P("Test"),
				elems.P("Test"),
				elems.P("Test"),
			),
			elems.Div().Class("col").Add(
				elems.P("Test"),
				elems.P("Test").Add(
					elems.Span("Test"),
					elems.Section("Test").Add(
						elems.Div().Add(
							elems.P("Test"),
							elems.P("Test"),
							elems.P("Test"),
						),
					),
				),
				elems.P("Test"),
			),
			elems.Div().Class("col").Add(
				elems.P("Test"),
				elems.P("Test"),
				elems.P("Test"),
			),
		).Add(
			elems.P("Test"),
			elems.P("Test"),
			elems.P("Test"),
			elems.Div().Class("row").Add(
				elems.Div().Class("col").Add(
					elems.P("Test"),
					elems.P("Test").Add(
						elems.P("Test"),
						elems.P("Test"),
						elems.P("Test"),
						elems.Div().Class("row").Add(
							elems.Div().Class("col").Add(
								elems.P("Test"),
								elems.P("Test"),
								elems.P("Test"),
							),
							elems.Div().Class("col").Add(
								elems.P("Test"),
								elems.P("Test").Add(
									elems.Span("Test"),
									elems.Section("Test").Add(
										elems.Div().Add(
											elems.P("Test"),
											elems.P("Test"),
											elems.P("Test"),
										),
									),
								),
								elems.P("Test"),
							),
							elems.Div().Class("col").Add(
								elems.P("Test"),
								elems.P("Test"),
								elems.P("Test"),
							),
						),
					),
					elems.P("Test"),
				),
				elems.Div().Class("col").Add(
					elems.P("Test"),
					elems.P("Test").Add(
						elems.Span("Test"),
						elems.Section("Test").Add(
							elems.Div().Add(
								elems.P("Test"),
								elems.P("Test"),
								elems.P("Test").Add(
									elems.P("Test"),
									elems.P("Test"),
									elems.P("Test"),
									elems.Div().Class("row").Add(
										elems.Div().Class("col").Add(
											elems.P("Test"),
											elems.P("Test"),
											elems.P("Test"),
										),
										elems.Div().Class("col").Add(
											elems.P("Test"),
											elems.P("Test").Add(
												elems.Span("Test"),
												elems.Section("Test").Add(
													elems.Div().Add(
														elems.P("Test"),
														elems.P("Test"),
														elems.P("Test").Add(
															elems.P("Test"),
															elems.P("Test"),
															elems.P("Test"),
															elems.Div().Class("row").Add(
																elems.Div().Class("col").Add(
																	elems.P("Test"),
																	elems.P("Test"),
																	elems.P("Test"),
																),
																elems.Div().Class("col").Add(
																	elems.P("Test"),
																	elems.P("Test").Add(
																		elems.Span("Test"),
																		elems.Section("Test").Add(
																			elems.Div().Add(
																				elems.P("Test"),
																				elems.P("Test"),
																				elems.P("Test"),
																			),
																		),
																	),
																	elems.P("Test"),
																),
																elems.Div().Class("col").Add(
																	elems.P("Test").Class("test"),
																	elems.P("Test").Class("test"),
																	elems.P("Test").Class("test"),
																),
															),
														),
													),
												),
											),
											elems.P("Test"),
										),
										elems.Div().Class("col").Add(
											elems.P("Test").Class("test"),
											elems.P("Test").Class("test"),
											elems.P("Test").Class("test"),
										),
									),
								),
							),
						),
					),
					elems.P("Test"),
				),
				elems.Div().Class("col").Add(
					elems.P("Test"),
					elems.P("Test"),
					elems.P("Test"),
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
	return html
}

func BenchmarkFindID(b *testing.B) {
	html := getHTML()
	for i := 0; i < b.N; i++ {
		_ = html.GetByID("GET_THIS_ID")
		// b.Log(html.GetByID("GET_THIS_ID").String())
	}
}

func BenchmarkFindClasses(b *testing.B) {
	html := getHTML()
	for i := 0; i < b.N; i++ {
		_ = html.GetByClassname("test")
		// b.Log(html.GetByClassname("test").String())
	}
}

func TestFormToStruct(t *testing.T) {
	var data = make(map[string]string)
	data["Name"] = "Jane"
	// data["time"] = time.Now().Format("2006-01-02T15:04")
	data["Age"] = "22"
	data["Friend_Age"] = "25"    //TODO TODO TODO
	data["Friend_Name"] = "John" //TODO TODO TODO
	data["Friend_Other_Age"] = "25"
	data["Friend_Other_Name"] = "Jippy"

	type Me struct {
		Name   string
		Age    int
		Friend struct {
			Name  string
			Age   int
			Other *struct {
				Name string
				Age  int
			}
		}
	}
	var testForm Me
	elems.FormDataToStruct(data, &testForm)
	if testForm.Name != "Jane" {
		t.Error("Name is not Jane")
	}
	if testForm.Age != 22 {
		t.Error("Age is not 22")
	}
	if testForm.Friend.Name != "John" {
		t.Error("Friend name is not John")
	}
	if testForm.Friend.Age != 25 {
		t.Error("Friend age is not 25")
	}
	if testForm.Friend.Other == nil {
		t.Error("Other is nil")
	}
	if testForm.Friend.Other.Name != "Jippy" {
		t.Error("Friend other name is not Jippy")
	}
	if testForm.Friend.Other.Age != 25 {
		t.Error("Friend other age is not 25")
	}
	t.Log(testForm)
}
