package buttons

import (
	"strconv"

	"github.com/Nigel2392/gen/elems"
	"github.com/Nigel2392/gen/genstring"
	"github.com/Nigel2392/gen/predef/csshelpers"
)

const BUTTON_PREFIX = "gohtml-button-"

// General colors for use in CSS. Can be overridden.
var (
	CARD_CONTAINER   = "card-container-" // + string
	COLOR_MAIN       = csshelpers.COLOR_MAIN
	COLOR_BLACK      = csshelpers.COLOR_BLACK
	COLOR_GREY       = csshelpers.COLOR_GREY
	COLOR_ONE        = csshelpers.COLOR_ONE
	COLOR_TWO        = csshelpers.COLOR_TWO
	COLOR_THREE      = csshelpers.COLOR_THREE
	COLOR_FOUR       = csshelpers.COLOR_FOUR
	COLOR_FIVE       = csshelpers.COLOR_FIVE
	COLOR_SIX        = csshelpers.COLOR_SIX
	COLOR_BROWN      = csshelpers.COLOR_BROWN
	COLOR_DARK_GRAY  = csshelpers.COLOR_DARK_GRAY
	BACKGROUND_COLOR = csshelpers.BACKGROUND_COLOR
)

// Uses COLOR_MAIN for text, and a supplied color.
//
//   - This button creates a specific style tag for each button created.
//
//   - If the maincolor is white or black, the three colors will be automatically replaced to:
//
//     COLOR_BLACK
//
//     COLOR_BROWN
//
//     COLOR_DARK_GRAY
//
// https://www.udemy.com/course/css-animation-transitions-and-transforms-creativity-course/
func RotatorSwap(href, text, mainColorHex string, widthPX int) (*elems.Element, *elems.Element) {
	var idName = genstring.RandStringBytesMaskImprSrcUnsafe(10)
	var btnContainer = elems.Div()
	var a = btnContainer.A(text).Href(href).Class(BUTTON_PREFIX + "rotator-swap").ID(idName)
	var COLOR1, COLOR2, COLOR3 csshelpers.Color
	COLOR1 = csshelpers.UnsafeHex(mainColorHex)
	var colors = COLOR1.Triadic()
	COLOR2, COLOR3 = colors[0], colors[1]

	if COLOR1.Hex() == COLOR2.Hex() && COLOR2.Hex() == COLOR3.Hex() {
		COLOR1 = csshelpers.UnsafeHex(COLOR_BLACK)
		COLOR2 = csshelpers.UnsafeHex(COLOR_BROWN)
		COLOR3 = csshelpers.UnsafeHex(COLOR_DARK_GRAY)
	}

	var height = int(float64(widthPX) / 2.5)
	// var height = widthPX / 2

	btnContainer.StyleBlock(`
	#` + idName + ` {
		padding: ` + strconv.Itoa(height/10) + `px ` + strconv.Itoa(widthPX/10) + `px;
		text-decoration: none;
		text-transform: uppercase;
		background-color: ` + COLOR1.Hex() + `;
		color: ` + COLOR_MAIN + `;
		font-size: ` + strconv.Itoa(widthPX/10) + `px;
		letter-spacing: ` + strconv.Itoa(widthPX/100) + `px;
		width: ` + strconv.Itoa(widthPX) + `px;
		height: ` + strconv.Itoa(height) + `px;
		position: relative;
	}
	#` + idName + `:before,
	#` + idName + `:after {
		content: "` + text + `";
		position: absolute;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		color: ` + COLOR_MAIN + `;
		display: flex;
		justify-content: center;
		align-items: center;
		transform-origin: top;
		transform: rotateX(270deg);
		transition: all 0.5s;
	}
	#` + idName + `:before {
		background-color: ` + COLOR2.Hex() + `;
	}
	#` + idName + `:after {
		background-color: ` + COLOR3.Hex() + `;
		transition-delay: 0.25s;
	}
	#` + idName + `:hover:before,
	#` + idName + `:hover:after {
		transform: rotateX(0deg);
	}`)
	return btnContainer, a
}

// Uses a user supplied color.
//   - Needs a style tag to be added to the dom, use the EnclosingSideDotsCss function.
func EnclosingSideDots(href, text string) *elems.Element {
	var link = elems.A(text).Href(href).TextAfter().Class(BUTTON_PREFIX + "enclosing-side-dots")
	link.Span()
	return link
}

// RotatingBackground returns a button with a rotating background.
//   - The background rotates 120 degrees when the mouse hovers over the button.
func RotatingBackground(href, text string) *elems.Element {
	var link = elems.A(text).Href(href).Class(BUTTON_PREFIX + "expanding-cut")
	link.Span()
	return link
}
