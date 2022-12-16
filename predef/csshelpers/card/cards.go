package cards

import (
	"strconv"

	"github.com/Nigel2392/gen/elems"
	"github.com/Nigel2392/gen/predef/csshelpers"
)

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
	BACKGROUND_COLOR = csshelpers.BACKGROUND_COLOR
)

// https://www.udemy.com/course/css-animation-transitions-and-transforms-creativity-course/
func NumberedCardCss(widthPX int) *elems.Element {
	var height = widthPX / 3 * 4
	return elems.StyleBlock(`
	.` + CARD_CONTAINER + `gohtml-card-numbered {
		position: relative;
		width: ` + strconv.Itoa(widthPX) + `px;
		height: ` + strconv.Itoa(height) + `px;
		background-color: ` + COLOR_MAIN + `;
		transition: 0.5s;
	}
	.` + CARD_CONTAINER + `gohtml-card-numbered:hover {
		box-shadow: 0 30px 50px rgba(0,0,0,0.2);
	}
	.` + CARD_CONTAINER + `gohtml-card-numbered div {
		position: absolute;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		display: flex;
		justify-content: center;
		align-items: center;
		flex-direction: column;
	}
	.` + CARD_CONTAINER + `gohtml-card-numbered div:nth-child(1) {
		padding: 20px;
		text-align: center;
		background-color: ` + COLOR_MAIN + `;
		color: ` + COLOR_ONE + `;
	}
	.` + CARD_CONTAINER + `gohtml-card-numbered div:nth-child(1) h2 {
		margin-top: ` + strconv.Itoa(widthPX/30*8) + `px;
		font-size: ` + strconv.Itoa(widthPX/30*2) + `px;
	}
	.` + CARD_CONTAINER + `gohtml-card-numbered div:nth-child(1) p {
		font-size: ` + strconv.Itoa(widthPX/30*1) + `px;
	}
	.` + CARD_CONTAINER + `gohtml-card-numbered div:nth-child(1) a {
		display: inline-block;
		margin-top: ` + strconv.Itoa(widthPX/30*2) + `px;
		padding: ` + strconv.Itoa(widthPX/30/2) + `px ` + strconv.Itoa(widthPX/30) + `px;
		text-decoration: none;
		color: ` + COLOR_ONE + `;
		border: ` + strconv.Itoa(widthPX/30*2/10) + `px solid ` + COLOR_ONE + `;
		font-weight: bold;
		transition: 0.3s;
		font-size: ` + strconv.Itoa(int(float64(widthPX)/float64(30)*1.5)) + `px;
	}
	.` + CARD_CONTAINER + `gohtml-card-numbered div:nth-child(2) {
		background-color: ` + COLOR_ONE + `;
		transition: 0.5s;
	}
	.` + CARD_CONTAINER + `gohtml-card-numbered div:nth-child(2) h2 {
		font-size: ` + strconv.Itoa(widthPX/30*12) + `px;
		background-color: ` + COLOR_ONE + `;
		color: ` + COLOR_MAIN + `;
		transition: 0.5s;
	}
	.` + CARD_CONTAINER + `gohtml-card-numbered:hover div:nth-child(2) {
		background-color: ` + COLOR_MAIN + `;
		height: ` + strconv.Itoa(height/40*8) + `px;
		width: ` + strconv.Itoa(widthPX/30*8) + `px;
		border-radius: 50%;
		top: ` + strconv.Itoa(height/40*4) + `px;
		left: 50%;
		transform: translateX(-50%);
		border: ` + strconv.Itoa(widthPX/30*2/10) + `px solid ` + COLOR_ONE + `;
	}
	.` + CARD_CONTAINER + `gohtml-card-numbered:hover div:nth-child(2) h2 {
		background-color: ` + COLOR_MAIN + `;
		color: ` + COLOR_ONE + `;
		font-size: ` + strconv.Itoa(widthPX/10) + `px;
	}
	.` + CARD_CONTAINER + `gohtml-card-numbered div:nth-child(1) a:hover {
		background-color: ` + COLOR_ONE + `;
		color: ` + COLOR_MAIN + `;
	}
	`)
}

// https://www.udemy.com/course/css-animation-transitions-and-transforms-creativity-course/
func NumberedCard(number int, href, btnText, caption, text string) (*elems.Element, *elems.Element) {
	var card = elems.Div().Class(CARD_CONTAINER + "gohtml-card-numbered")
	var face = card.Div()
	var content = face.Div()
	content.H2(caption)
	content.P(text)
	var btn = content.A(btnText).Href(href)
	var face2 = card.Div()
	face2.H2(strconv.Itoa(number))
	return card, btn
}

// TODO: 	// VerticalCardCss returns the css for a vertical card with the given width in pixels
// TODO: 	func VerticalCardCss(widthPX int) *elems.Element {
// TODO: 		var heightPX = widthPX / 3 * 2
// TODO: 		return elems.StyleBlock(`
// TODO: 		.` + CARD_CONTAINER + `gohtml-card-expand {
// TODO: 			position: relative;
// TODO: 			cursor: pointer;
// TODO: 			width: ` + strconv.Itoa(widthPX) + `px;
// TODO: 		}
// TODO: 		.` + CARD_CONTAINER + `gohtml-card-expand div {
// TODO: 			position: relative;
// TODO: 			width: ` + strconv.Itoa(widthPX) + `px;
// TODO: 			height: ` + strconv.Itoa(heightPX) + `px;
// TODO: 			flex-direction: column;
// TODO: 			justify-content: center;
// TODO: 			align-items: center;
// TODO: 		}
// TODO: 		.` + CARD_CONTAINER + `gohtml-card-expand div:nth-child(1) {
// TODO: 			background-color: ` + BACKGROUND_COLOR + `
// TODO: 			align-items: center;
// TODO: 			transition: 0.5s;
// TODO: 			transform: translateY(` + strconv.Itoa(heightPX/2) + `px);
// TODO: 			z-index: 1;
// TODO: 		}
// TODO: 		.` + CARD_CONTAINER + `gohtml-card-expand div:nth-child(2) {
// TODO: 			background-color: ` + COLOR_MAIN + `;
// TODO: 			box-sizing: border-box;
// TODO: 			padding: ` + strconv.Itoa(heightPX/10) + `px;
// TODO: 			box-shadow: 0 ` + strconv.Itoa(heightPX/10) + `px ` + strconv.Itoa(heightPX/20*5) + `px 0 rgba(0,0,0,0.8);
// TODO: 			transition: 0.5s;
// TODO: 			transform: translateY(-` + strconv.Itoa(heightPX/2) + `px);
// TODO: 		}
// TODO: 		.` + CARD_CONTAINER + `gohtml-card-expand:hover div:nth-child(1) {
// TODO: 			background-color: ` + COLOR_SIX + `;
// TODO: 			transform: translateY(0);
// TODO: 		}
// TODO: 		.` + CARD_CONTAINER + `gohtml-card-expand:hover div:nth-child(2) {
// TODO: 			transform: translateY(0);
// TODO: 		}
// TODO: 		.` + CARD_CONTAINER + `gohtml-card-expand div div {
// TODO: 			opacity: 0.2;
// TODO: 			transition: 0.5s;
// TODO: 		}
// TODO: 		.` + CARD_CONTAINER + `gohtml-card-expand:hover div:nth-child(1) div {
// TODO: 			opacity: 1;
// TODO: 		}
// TODO: 		.` + CARD_CONTAINER + `gohtml-card-expand:hover div:nth-child(1) div img {
// TODO: 			width: ` + strconv.Itoa(heightPX/2) + `px;
// TODO: 		}
// TODO: 		.` + CARD_CONTAINER + `gohtml-card-expand:hover div:nth-child(1) div h3 {
// TODO: 			margin: 10px 0 0;
// TODO: 			padding: 0;
// TODO: 			color: ` + COLOR_MAIN + `;
// TODO: 			text-align: center;
// TODO: 			font-size: ` + strconv.Itoa(int(float64(heightPX)/float64(20)*2.5)) + `px;
// TODO: 		}
// TODO: 		.` + CARD_CONTAINER + `gohtml-card-expand div:nth-child(2) div p {
// TODO: 			margin: 0;
// TODO: 			padding: 0;
// TODO: 		}
// TODO: 		.` + CARD_CONTAINER + `gohtml-card-expand div:nth-child(2) div a {
// TODO: 			margin: 15px 0 0;
// TODO: 			display: inline-block;
// TODO: 			text-decoration: none;
// TODO: 			font-weight: 900;
// TODO: 			color: ` + COLOR_GREY + `;
// TODO: 			padding;  ` + strconv.Itoa(int(float64(heightPX)/float64(20)*0.5)) + `px;
// TODO: 			border: ` + strconv.Itoa(int(float64(heightPX)/float64(20)*0.2)) + `px solid ` + COLOR_GREY + `;
// TODO: 			transition: 0.3s;
// TODO: 		}
// TODO: 		.` + CARD_CONTAINER + `gohtml-card-expand div:nth-child(2) div a:hover {
// TODO: 			background-color: ` + COLOR_GREY + `;
// TODO: 			color: ` + COLOR_MAIN + `;
// TODO: 		}
// TODO: 		`)
// TODO: 	}
// TODO:
// TODO: 	// Return an expanding html card element.
// TODO: 	func VerticalExpand(src, href, btnText, caption, text string) (*elems.Element, *elems.Element) {
// TODO: 		var card = elems.Div().Class(CARD_CONTAINER + "gohtml-card-expand")
// TODO: 		var face = card.Div()
// TODO: 		var imageBox = face.Div()
// TODO: 		imageBox.Img().Src(src)
// TODO: 		imageBox.H3(caption)
// TODO: 		var face2 = card.Div()
// TODO: 		var content = face2.Div()
// TODO: 		content.P(text)
// TODO: 		var anchor = content.A(btnText).Href(href)
// TODO: 		return card, anchor
// TODO: 	}
