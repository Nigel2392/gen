package navbars

import (
	"github.com/Nigel2392/gen/elems"
	"github.com/Nigel2392/gen/predef/csshelpers"
)

const (
	ID_MAIN_CONTENT           = "gohtml-main-content"
	CLASS_NAVBAR              = "gohtml-navbar"
	CLASS_NAVBAR_LEFT         = "gohtml-left"
	CLASS_NAVBAR_RIGHT        = "gohtml-right"
	CLASS_NAVIGATION_LINK     = "gohtml-navigation-link"
	ID_SIDE_OPEN_BTN          = "gohtml-openbtn"
	ID_SIDE_CLOSE_BTN         = "gohtml-closebtn"
	ID_SIDE_NAVBAR            = "gohtml-sidenav"
	CLASS_SIDE_NAVBAR         = "gohtml-sidenav"
	CLASS_SIDE_NAVBAR_FOOTER  = "gohtml-footer"
	CLASS_SIDE_NAVBAR_CONTENT = "gohtml-content-nav"
	CLASS_NAVIGATION_LOGO     = "gohtml-navigation-logo"
	CLASS_ACTIVE              = "gohtml-active"
	ONCLICK_ACTIVE_RAW        = "this.classList.toggle('" + CLASS_ACTIVE + "');let other = document.getElementsByClassName('" + CLASS_ACTIVE + "');for (let i = 0; i < other.length; i++) {if (other[i] != this) {other[i].classList.remove('" + CLASS_ACTIVE + "');}}"
)

// General colors for use in CSS. Can be overridden.
var (
	CARD_CONTAINER    = "card-container-"            // + string
	COLOR_MAIN        = csshelpers.COLOR_MAIN        // Main color - white
	COLOR_BLACK       = csshelpers.COLOR_BLACK       // Black
	COLOR_GREY        = csshelpers.COLOR_GREY        // Grey
	COLOR_ONE         = csshelpers.COLOR_ONE         // COLOR_ONE - dark blue
	COLOR_ONE_LIGHTER = csshelpers.COLOR_ONE_LIGHTER // COLOR_ONE_LIGHTER - lighter dark blue
	COLOR_TWO         = csshelpers.COLOR_TWO         // COLOR_TWO - yellow
	COLOR_THREE       = csshelpers.COLOR_THREE       // COLOR_THREE - green
	COLOR_FOUR        = csshelpers.COLOR_FOUR        // COLOR_FOUR - light blue
	COLOR_FIVE        = csshelpers.COLOR_FIVE        // COLOR_FIVE - dark green
	COLOR_SIX         = csshelpers.COLOR_SIX         // COLOR_SIX - red
	COLOR_BROWN       = csshelpers.COLOR_BROWN       // COLOR_BROWN - brown
	COLOR_DARK_GRAY   = csshelpers.COLOR_DARK_GRAY   // COLOR_DARK_GRAY - dark gray
	BACKGROUND_COLOR  = csshelpers.BACKGROUND_COLOR  // BACKGROUND_COLOR - background color
)

type Logo struct {
	URL         elems.URL
	Image       string
	LeftOrRight bool
}

// Returns the navbar body to append to your HTML, the main-content area and all URLS
func SideNavBar(logo *Logo, urls elems.URLs, footer_urls ...elems.URL) (*elems.Element, *elems.Element, elems.Elements) {
	var urlElems = make(elems.Elements, len(urls))
	for i, url := range urls {
		var urlElem = elems.A(" " + url.Name).Href(url.URL).Class(CLASS_NAVIGATION_LINK)
		if url.Icon != "" {
			urlElem.Add(elems.I().Class(url.Icon)).TextAfter()
		}
		urlElems[i] = urlElem
	}
	var footerElems = make(elems.Elements, len(footer_urls))
	for i, url := range footer_urls {
		var urlElem = elems.A(" " + url.Name).Href(url.URL).Class(CLASS_NAVIGATION_LINK)
		if url.Icon != "" {
			urlElem.Add(elems.I().Class(url.Icon)).TextAfter()
		}
		footerElems[i] = urlElem
	}
	// Selectors
	//.ID("gohtml-openbtn")
	//.ID("gohtml-sidenav")
	//.Class("gohtml-sidenav")
	//.Class("gohtml-footer")
	//.ID("gohtml-main-content")
	//.ID("gohtml-closebtn")

	var body = elems.Elem("div")

	body.Add(
		elems.StyleBlock(DefaultSideNavStyle),
	)

	var navbar_btn = elems.Span().InnerHTML("&#9776;").Style("font-size:40px").ID(ID_SIDE_OPEN_BTN).Onclick("document.getElementById(`" + ID_SIDE_NAVBAR + "`).style.width = `250px`;document.getElementById(`" + ID_MAIN_CONTENT + "`).style.marginLeft = `250px`;")
	var sidenav = elems.Elem("div").ID(ID_SIDE_NAVBAR).Class(CLASS_SIDE_NAVBAR)
	var inner_footer = elems.Elem("div").Class(CLASS_SIDE_NAVBAR_FOOTER)
	var content_div = elems.Elem("div").ID(ID_MAIN_CONTENT)

	inner_footer.Add(footerElems...)

	var sideNavCloseBtnDiv = elems.Div().Style("width: 100%")
	var logoElem *elems.Element
	if logo != nil {
		logoElem = elems.A().Href(logo.URL.URL).Style("width: 70%", "height: auto", "margin-left:15px")
		if logo.Image != "" {
			logoElem.Add(elems.Img(logo.Image).Style("width: 70%", "height: auto"))
		}
		sideNavCloseBtnDiv.Add(logoElem)
	}

	var allURLS = make(elems.Elements, len(urlElems)+len(footerElems))
	var n = copy(allURLS, urlElems)
	copy(allURLS[n:], footerElems)
	if logo != nil {
		allURLS = append(allURLS, logoElem)
	}

	sidenav.Add(
		sideNavCloseBtnDiv.Add(
			elems.A().InnerHTML("&times;").Href("javascript:void(0)").ID(ID_SIDE_CLOSE_BTN).Style("float:right").Onclick("document.getElementById(`" + ID_SIDE_NAVBAR + "`).style.width = `0`;document.getElementById(`" + ID_MAIN_CONTENT + "`).style.marginLeft= `0`;")),
	)

	sidenav.Add(elems.Div().Class(CLASS_SIDE_NAVBAR_CONTENT).Add(append(urlElems, inner_footer)...))

	body.Add(navbar_btn, sidenav, content_div)

	return body, content_div, allURLS
}

func NavBar(logo *Logo, urls elems.URLs) (*elems.Element, elems.Elements) {
	nav := elems.Elem("div")
	nav.Class(CLASS_NAVBAR)
	left := elems.Elem("div")
	right := elems.Elem("div")
	left.Class(CLASS_NAVBAR_LEFT)
	right.Class(CLASS_NAVBAR_RIGHT)

	var elements elems.Elements
	// Selectors
	//.Class("gohtml-navbar")
	//.Class("gohtml-left")
	//.Class("gohtml-right")
	//.Class("gohtml-navigation-link")
	var logoElem *elems.Element
	if logo != nil {
		logoElem = elems.A().Class(CLASS_NAVIGATION_LINK).Href(logo.URL.URL)
		if logo.Image != "" {
			logoElem.Add(elems.Img(logo.Image))
		}
		elements = append(elements, logoElem)
		if logo.LeftOrRight {
			left.Add(logoElem.Style("float:left !important"))
		} else {
			right.Add(logoElem.Style("float:right !important"))
		}
	}

	for _, url := range urls {
		var urlElem *elems.Element = elems.Elem("a").Class(CLASS_NAVIGATION_LINK).InnerText(" " + url.Name).Href(url.URL)
		if url.Icon != "" {
			urlElem.Add(elems.I().Class(url.Icon).Style("margin-left:5px;"))
		}
		elements = append(elements, urlElem)
		if url.LeftOrRight {
			left.Add(urlElem.Style("float:left !important"))
		} else {
			right.Add(urlElem.Style("float:right !important"))
		}
	}

	nav.Add(left, right)
	return nav, elements
}

func OfficialNavbar(logo *Logo, urls elems.URLs, footer_urls ...elems.URL) (*elems.Element, elems.Elements) {
	var navbar_container = elems.Div().Class(CLASS_NAVBAR + "-official")
	var navbar = navbar_container.Div().Class(CLASS_NAVBAR + "-official-inner")
	var navbar_left = navbar.Div().Class(CLASS_NAVBAR_LEFT)
	var navbar_right = navbar.Div().Class(CLASS_NAVBAR_RIGHT)
	var url_elems = make(elems.Elements, len(urls)+len(footer_urls))

	if logo != nil {
		if logo.LeftOrRight {
			var logo_link = elems.A().Href(logo.URL.URL)
			logo_link.Img(logo.Image).Class(CLASS_NAVIGATION_LOGO)
			navbar_left.Add(logo_link)
		}
	}

	for i, url := range urls {
		var element = elems.A(url.Name).Class(CLASS_NAVIGATION_LINK).Href(url.URL)
		if url.LeftOrRight {
			navbar_left.Add(element)
		} else {
			navbar_right.Add(element)
		}
		url_elems[i] = element
	}

	if logo != nil {
		if !logo.LeftOrRight {
			var logo_link = elems.A().Href(logo.URL.URL)
			logo_link.Img(logo.Image).Class(CLASS_NAVIGATION_LOGO)
			navbar_right.Add(logo_link)
		}
	}

	for i, url := range footer_urls {
		var element = elems.A(url.Name).Class(CLASS_NAVIGATION_LINK + "-footer").Href(url.URL)
		if url.LeftOrRight {
			navbar_left.Add(element)
		} else {
			navbar_right.Add(element)
		}
		url_elems[i+len(urls)] = element
	}

	return navbar_container, url_elems
}

// Returns a window control frame for frameless wails applications.
// Allows for urls to be passed in for the menu.
// Returns the frame container, the control buttons, and the urls.
//
// CLASSES:
//
// "gohtml-window-control-left"
// "gohtml-window-control-favicon"
// "gohtml-window-control-middle"
// "gohtml-window-control-right"
// "gohtml-navigation-link-control"
//
// IDs:
//
//	"gohtml-window-control-box"
//	"gohtml-window-control-minimize"
//	"gohtml-window-control-maximize"
//	"gohtml-window-control-close"
func WailsFrame(imgUrl string, menu elems.URLs) (*elems.Element, elems.Elements, elems.Elements) {
	var frame_container = elems.Div().Class(CLASS_NAVBAR + "-framed-wails").Style("--wails-draggable:drag")
	var frame_left = frame_container.Div().Class("gohtml-window-control-left", "gohtml-noselect")
	frame_left.Img(imgUrl).Class("gohtml-window-control-favicon", "gohtml-noselect")
	// var frame_middle = frame_container.Div().Class("gohtml-window-control-middle", "gohtml-noselect")
	var urlElems = make(elems.Elements, len(menu))
	for i, url := range menu {
		var element = elems.A(url.Name).Class(CLASS_NAVIGATION_LINK + "-control").Href(url.URL)
		// if url.LeftOrRight {
		frame_left.Add(element)
		// } else {
		// frame_middle.Add(element)
		// }
		urlElems[i] = element
	}
	// Create the window controls
	var frame_right = frame_container.Div().Class("gohtml-window-control-right", "gohtml-noselect")
	var window_control_box = frame_right.Div().ID("gohtml-window-control-box")
	var minimize_btn = window_control_box.Div().ID("gohtml-window-control-minimize").Class("gohtml-noselect")
	var maximize_btn = window_control_box.Div().ID("gohtml-window-control-maximize").Class("gohtml-noselect")
	var close_btn = window_control_box.Div().ID("gohtml-window-control-close").Class("gohtml-noselect")
	// Create spans for graphics
	minimize_btn.Span()
	maximize_btn.Span()
	close_btn.Span()
	close_btn.Span()
	// Set the control button elements
	var controlsList = elems.Elements{minimize_btn, maximize_btn, close_btn}
	// Return
	return frame_container, controlsList, urlElems
}
func WailsFrameButtonsOnly() (*elems.Element, elems.Elements) {
	var frame_container = elems.Div().Class(CLASS_NAVBAR + "-framed-wails")
	// Create the window controls
	var window_control_box = frame_container.Div().ID("gohtml-window-control-box").Style("--wails-draggable:drag")
	var minimize_btn = window_control_box.Div().ID("gohtml-window-control-minimize").Class("gohtml-noselect")
	var maximize_btn = window_control_box.Div().ID("gohtml-window-control-maximize").Class("gohtml-noselect")
	var close_btn = window_control_box.Div().ID("gohtml-window-control-close").Class("gohtml-noselect")
	// Create spans for graphics
	minimize_btn.Span()
	maximize_btn.Span()
	close_btn.Span()
	close_btn.Span()
	// Set the control button elements
	var controlsList = elems.Elements{minimize_btn, maximize_btn, close_btn}
	// Return
	return frame_container, controlsList
}
