package menus

import (
	"strconv"
	"strings"

	"github.com/Nigel2392/gen/elems"
	"github.com/Nigel2392/gen/predef/csshelpers"
)

// Provide a fullscreen menu with a curtain drop effect.
// Style is included in the returned element.
// CLASSES AND IDS
//
//	".gohtml-menu-container"
//	"#gohtml-menu-open-btn"
//	"#gohtml-menu"
//	"#gohtml-menu-close-btn"
func MenuCurtainDrop(urls elems.URLs, btnWidth int, curtainColor string, compColor ...string) (*elems.Element, elems.Elements) {
	// Menu container
	var menu_container = elems.Div().Class("gohtml-menu-container")
	// Open button
	menu_container.Div().InnerHTML("&#9776;").ID("gohtml-menu-open-btn").Onclick("document.getElementById('gohtml-menu').style.transform = 'translateX(0)';document.getElementById('gohtml-menu-open-btn').style.display = 'none'")
	// Menu
	var menu = menu_container.Div().ID("gohtml-menu")
	// Close button
	var closeBtn = menu.Div().ID("gohtml-menu-close-btn").Onclick("document.getElementById('gohtml-menu').style.transform = 'translateX(-100%)';document.getElementById('gohtml-menu-open-btn').style.display = 'block'")
	closeBtn.Span()
	closeBtn.Span()
	// Urls for the navigation menu
	var urlItems = make(elems.Elements, len(urls))
	if len(urls) > 0 {
		var ul = menu.Ul()
		for i, url := range urls {
			var urlItem = elems.A(url.Name).Href(url.URL)
			urlItem.Span()
			urlItem.Span()
			urlItem.Span()
			urlItem.Span()
			ul.Li().Add(urlItem)
			urlItems[i] = urlItem
		}
	}
	// Get the main color
	colr, err := csshelpers.Hex(curtainColor)
	if err != nil {
		panic(err)
	}
	// Get the complementary color
	var complementaryColor = colr.Complementary().Hex()
	var r, g, b = colr.RGB255()
	// Check if the color is valid, and not the same.
	if strings.EqualFold(complementaryColor, curtainColor) {
		if r == 0 && g == 0 && b == 0 {
			complementaryColor = "#FFFFFF"
		} else {
			if r > 127 && g > 127 && b > 127 {
				complementaryColor = "#000000"
			} else {
				complementaryColor = "#FFFFFF"
			}
		}
	}
	// If a complementary color is provided, use that instead.
	if len(compColor) > 0 {
		complementaryColor = compColor[0]
	}
	// Calculate the height of the menu buttons
	var btnHeight = btnWidth / 3

	// Convert the RGB values to strings for use in CSS rgba(r, g, b, a)
	var r_str, g_str, b_str = strconv.Itoa(int(r)), strconv.Itoa(int(g)), strconv.Itoa(int(b))

	// Add the style to the menu container
	menu_container.StyleBlock(`
		#gohtml-menu-open-btn,
		#gohtml-menu-close-btn{
			position: fixed;
			top: 10px;
			left: 10px;
			font-size: ` + strconv.Itoa(btnWidth/4) + `px;
			cursor: pointer;
		}
		#gohtml-menu-open-btn {
			color: ` + curtainColor + `;
		}
		#gohtml-menu-close-btn span {
			background-color: ` + complementaryColor + `;
		}
		#gohtml-menu-close-btn {
			top: calc(10px + ` + strconv.Itoa(btnHeight/8) + `px);
			background-color:  ` + complementaryColor + `;
			border-radius: 50%;
			transition: all 0.3s;
			width: ` + strconv.Itoa(btnWidth/4) + `px;
			height: ` + strconv.Itoa(btnWidth/4) + `px;
		}
		#gohtml-menu-close-btn span {
			position: absolute;
			top: 50%;
			left: 50%;
			transform: translate(-50%, -50%);
			transition: all 0.3s;
		}
		#gohtml-menu-close-btn span:nth-child(1),
		#gohtml-menu-close-btn span:nth-child(2) {
			content: "";
			position: absolute;
			top: 50%;
			left: 50%;
			width: 60%;
			height: ` + strconv.Itoa(btnHeight/8) + `px;
			background-color: ` + curtainColor + `;
			z-index: 1;
			transition: all 0.3s;
		}
		#gohtml-menu-close-btn span:nth-child(1) {
			transform: translate(-50%, -50%) rotate(45deg);
		}
		#gohtml-menu-close-btn span:nth-child(2) {
			transform: translate(-50%, -50%) rotate(-45deg);
		}
		#gohtml-menu-close-btn:hover {
			cursor: pointer;
			background-color: ` + curtainColor + `;
		}
		#gohtml-menu-close-btn:hover span:nth-child(1),
		#gohtml-menu-close-btn:hover span:nth-child(2) {
			background-color: ` + complementaryColor + `;
		}
		#gohtml-menu {
			position: fixed;
			top: 0;
			left: 0;
			width: 100%;
			height: 100%;
			background-color: rgba(` + r_str + `, ` + g_str + `, ` + b_str + `, 0.5);
			display: flex;
			justify-content: center;
			align-items: center;
			transform: translateX(-100%);
			transition: all 0.4s;
		}
		#gohtml-menu ul {
			margin: 0;
			padding: 0;
			display: flex;
			list-style: none;
		}
		#gohtml-menu ul li a {
			display: block;
			width: ` + strconv.Itoa(btnWidth) + `px;
			height: ` + strconv.Itoa(btnHeight) + `px;
			line-height: ` + strconv.Itoa(btnHeight) + `px;
			text-align: center;
			color: ` + complementaryColor + `;
			text-decoration: none;
			text-transform: uppercase;
			position: relative;
			transition: all 0.4s;
			border-top: 1px solid ` + complementaryColor + `;
			border-bottom: 1px solid ` + complementaryColor + `;
			letter-spacing: ` + strconv.Itoa(btnWidth*2/120) + `px;
    		font-weight: 800;
		}
		#gohtml-menu ul li:first-child a {
			border-left: 1px solid ` + complementaryColor + `;
		}
		#gohtml-menu ul li:last-child a {
			border-right: 1px solid ` + complementaryColor + `;
		}
		#gohtml-menu ul li a:hover {
			color: ` + curtainColor + `;
		}
		#gohtml-menu ul li a:hover span {
			transform: scaleY(1);
		}
		#gohtml-menu ul li span {
			background-color: ` + complementaryColor + `;
			position: absolute;
			left: 0;
			top: 0;
			width: 25%;
			height: 100%;
			z-index: -1;
			transform: scaleY(0);
			transition: all 0.5s;
			transform-origin: top;
		}
		#gohtml-menu ul li span:nth-child(1) {
			transition-delay: 0.2s;
		}
		#gohtml-menu ul li span:nth-child(2) {
			left: 25%;
			transition-delay: 0.1s;
		}
		#gohtml-menu ul li span:nth-child(3) {
			left: 50%;
		}
		#gohtml-menu ul li span:nth-child(4) {
			left: 75%;
			transition-delay: 0.3s;
		}
		@media screen and (min-width: 1367px) {
			#gohtml-menu ul li a {
				font-size: ` + strconv.Itoa(btnWidth/8) + `px;
				width: ` + strconv.Itoa(btnWidth) + `px;
				height: ` + strconv.Itoa(btnHeight) + `px;
				line-height: ` + strconv.Itoa(btnHeight) + `px;
				letter-spacing: ` + strconv.Itoa(btnWidth*2/120) + `px;
			}
		}
		@media screen and (max-width: 1366px) {
			#gohtml-menu ul li a {
				font-size: ` + strconv.Itoa(btnWidth/14) + `px;
				width: ` + strconv.Itoa(int(float64(btnWidth)/1.5)) + `px;
				height: ` + strconv.Itoa(int(float64(btnHeight)/1.5)) + `px;
				line-height: ` + strconv.Itoa(int(float64(btnHeight)/1.5)) + `px;
				letter-spacing: ` + strconv.Itoa(btnWidth/60) + `px;
			}
		}
		@media screen and (max-width: 768px) {
			#gohtml-menu ul li a {
				font-size: ` + strconv.Itoa(btnWidth/22) + `px;
				width: ` + strconv.Itoa(int(float64(btnWidth)/3)) + `px;
				height: ` + strconv.Itoa(int(float64(btnHeight)/3)) + `px;
				line-height: ` + strconv.Itoa(int(float64(btnHeight)/3)) + `px;
				letter-spacing: ` + strconv.Itoa(btnWidth/180) + `px;
			}
		}
		@media screen and (max-width: 380px) {
			#gohtml-menu ul li a {
				font-size: ` + strconv.Itoa(btnWidth/60) + `px;
				width: ` + strconv.Itoa(btnWidth/12) + `px;
				height: ` + strconv.Itoa(btnHeight/12) + `px;
				line-height: ` + strconv.Itoa(btnHeight/12) + `px;
				letter-spacing: ` + strconv.Itoa(btnWidth/240) + `px;
			}
		}
	`)
	// Return the menu container and the urlItems
	return menu_container, urlItems
}

func StackedRotation(urls elems.URLs, btnWidth int, bgCol, textCol, hoverCol string) (*elems.Element, elems.Elements) {
	// Menu container
	var menu_container = elems.Div().Class("gohtml-menu-container")
	// Open button
	menu_container.Div().InnerHTML("&#9776;").ID("gohtml-menu-open-btn").Onclick("document.getElementById('gohtml-menu').style.transform = 'translateX(0)';document.getElementById('gohtml-menu-open-btn').style.display = 'none'")
	// Menu
	var menu = menu_container.Div().ID("gohtml-menu")
	// Close button
	var closeBtn = menu.Div().ID("gohtml-menu-close-btn").Onclick("document.getElementById('gohtml-menu').style.transform = 'translateX(-100%)';document.getElementById('gohtml-menu-open-btn').style.display = 'block'")
	closeBtn.Span()
	closeBtn.Span()

	var urlItems = make(elems.Elements, len(urls))
	if len(urls) > 0 {
		var ul = menu.Ul()
		for i, url := range urls {
			var urlItem = elems.A().Href(url.URL)
			urlItem.Span(url.Name)
			urlItem.Span(url.Name)
			ul.Li().Add(urlItem)
			urlItems[i] = urlItem
		}
	}

	bgColor, err := csshelpers.Hex(bgCol)
	if err != nil {
		panic(err)
	}

	textColor, err := csshelpers.Hex(textCol)
	if err != nil {
		panic(err)
	}

	hoverColor, err := csshelpers.Hex(hoverCol)
	if err != nil {
		panic(err)
	}

	var r, g, b = bgColor.RGB255()

	var r_str, g_str, b_str = strconv.Itoa(int(r)), strconv.Itoa(int(g)), strconv.Itoa(int(b))

	var hoverTextCol = hoverColor.Complementary().Hex()
	if strings.EqualFold(hoverTextCol, hoverCol) {
		if r, g, b := textColor.RGB255(); r == 0 && g == 0 && b == 0 {
			hoverTextCol = "#FFFFFF"
		} else {
			if r > 127 && g > 127 && b > 127 {
				hoverTextCol = "#000000"
			} else {
				hoverTextCol = "#FFFFFF"
			}
		}
	}

	var btnHeight = btnWidth / 3

	// Add the style to the menu container
	menu_container.StyleBlock(`
		#gohtml-menu-open-btn,
		#gohtml-menu-close-btn{
			position: fixed;
			top: 10px;
			left: 10px;
			font-size: ` + strconv.Itoa(btnWidth/8) + `px;
			cursor: pointer;
		}
		#gohtml-menu-open-btn {
			color: ` + textColor.Hex() + `;
		}
		#gohtml-menu-close-btn span {
			background-color: ` + bgColor.Hex() + `;
		}
		#gohtml-menu-close-btn {
			top: calc(10px + ` + strconv.Itoa(btnWidth/16) + `px);
			background-color:  ` + bgColor.Hex() + `;
			border-radius: 50%;
			transition: all 0.3s;
			width: ` + strconv.Itoa(btnWidth/4) + `px;
			height: ` + strconv.Itoa(btnWidth/4) + `px;
		}
		#gohtml-menu-close-btn span {
			position: absolute;
			top: 50%;
			left: 50%;
			transform: translate(-50%, -50%);
			transition: all 0.3s;
		}
		#gohtml-menu-close-btn span:nth-child(1),
		#gohtml-menu-close-btn span:nth-child(2) {
			content: "";
			position: absolute;
			top: 50%;
			left: 50%;
			width: 60%;
			height: ` + strconv.Itoa(btnHeight/8) + `px;
			background-color: ` + textColor.Hex() + `;
			z-index: 1;
			transition: all 0.3s;
		}
		#gohtml-menu-close-btn span:nth-child(1) {
			transform: translate(-50%, -50%) rotate(45deg);
		}
		#gohtml-menu-close-btn span:nth-child(2) {
			transform: translate(-50%, -50%) rotate(-45deg);
		}
		#gohtml-menu-close-btn:hover {
			cursor: pointer;
			background-color: ` + textColor.Hex() + `;
		}
		#gohtml-menu-close-btn:hover span:nth-child(1),
		#gohtml-menu-close-btn:hover span:nth-child(2) {
			background-color: ` + bgColor.Hex() + `;
		}
		#gohtml-menu {
			position: fixed;
			top: 0;
			left: 0;
			width: 100%;
			height: 100%;
			background-color: rgba(` + r_str + `, ` + g_str + `, ` + b_str + `, 0.5);
			display: flex;
			justify-content: center;
			align-items: center;
			transform: translateX(-100%);
			transition: all 0.4s;
		}
		#gohtml-menu ul {
			list-style: none;
			display: flex;
			flex-direction: column;
			align-items: center;
			justify-content: center;
		}
		#gohtml-menu ul li {
			position: relative;
			list-style: none;
			margin: 5px;
		}
		#gohtml-menu ul li a {
			position: relative;
			display: block;
			height: ` + strconv.Itoa(btnHeight) + `px;
			width: ` + strconv.Itoa(btnWidth) + `px;
			text-decoration: none;
			padding: ` + strconv.Itoa(btnHeight/14) + `px ` + strconv.Itoa(btnHeight/6) + `px;
			letter-spacing: 2px;
			overflow: hidden;
		}
		#gohtml-menu ul li a span {
			position: relative;
			width: 100%;
			height: 100%;
			display: flex;
			align-items: center;
			justify-content: center;
			transition: 0.5s;
			color: ` + textColor.Hex() + `;
			font-size: ` + strconv.Itoa(btnHeight/3) + `px;
			text-transform: capitalize;
		}
		#gohtml-menu ul li a:hover span:nth-child(1) {
			transform: translateY(-100%);
			transition-delay: 0s;
			color: ` + hoverTextCol + `
		}
		#gohtml-menu ul li a:hover span:nth-child(2) {
			transform: translateY(-100%);
			transition-delay: 0s;
			color: ` + hoverTextCol + `
		}
		#gohtml-menu ul li a:hover:before  {
			transform: scaleX(100%);
			transition-delay: 0.5s;
		}
		#gohtml-menu ul li a:before {
			content: "";
			position: absolute;
			top: 0;
			left: 0;
			width: 100%;
			height: 100%;
			background-color: ` + hoverColor.Hex() + `;
			transform: scaleX(0);
			transform-origin: left;
			transition: 0.5s;
			transition-delay: 0.5s;
		}
	`)

	return menu_container, urlItems
}
