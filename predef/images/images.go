package images

import (
	"strconv"

	"github.com/Nigel2392/gen/elems"
	"github.com/Nigel2392/gen/predef/csshelpers"
)

// General colors for use in CSS. Can be overridden.
var (
	IMAGE_CONTAINER  = "image-container-" // + string
	COLOR_BLACK      = csshelpers.COLOR_BLACK
	COLOR_MAIN       = csshelpers.COLOR_MAIN
	COLOR_DARK_GRAY  = csshelpers.COLOR_DARK_GRAY
	COLOR_SIX        = csshelpers.COLOR_SIX
	BACKGROUND_COLOR = csshelpers.BACKGROUND_COLOR
)

func TextOnHoverCss(sideSizePX int) *elems.Element {
	return elems.StyleBlock(`
	.` + IMAGE_CONTAINER + `text-on-hover {
		width: ` + strconv.Itoa(sideSizePX) + `px;
		height: ` + strconv.Itoa(sideSizePX) + `px;
		position: relative;
		overflow: hidden;
	}
	.` + IMAGE_CONTAINER + `text-on-hover img {
		width: 100%;
		transition: all 0.5s;
	}
	.` + IMAGE_CONTAINER + `text-on-hover div {
		position: absolute;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		color: white;
		opacity: 0;
		transition: all 0.5s;
	}
	.` + IMAGE_CONTAINER + `text-on-hover div h1 {
		text-transform: uppercase;
		margin: 0;
	}
	.` + IMAGE_CONTAINER + `text-on-hover div p {
		font-size: 18px;
		text-transform: Capitalize;
	}
	.` + IMAGE_CONTAINER + `text-on-hover:hover div  {
		opacity: 1;
		background-color: ` + BACKGROUND_COLOR + `;
	}
	.` + IMAGE_CONTAINER + `text-on-hover:hover img {
		transform: scale(1.3);
		rotate: 15deg;
		filter: blur(3px);
	}
	`)
}

func TextOnHover(src, caption, text string) *elems.Element {
	var containerName = IMAGE_CONTAINER + "text-on-hover"
	var container = elems.Div().Class(containerName)
	container.Img(src)
	var cptn = container.Div()
	cptn.H1(caption)
	cptn.P(text)
	return container
}

func TextWithLinkBtn(src, href, buttontext, caption, text string) (*elems.Element, *elems.Element) {
	var containerName = IMAGE_CONTAINER + "text-with-link-btn"
	var container = elems.Div().Class(containerName)
	container.Img(src)
	var cptn = container.Div()
	cptn.H2(caption)
	cptn.P(text)
	var anchor = cptn.A(buttontext).Href(href)
	return container, anchor
}

func TextWithLinkBtnCss(widthPX int, src string) *elems.Element {
	var height = int(widthPX / 3 * 2)
	return elems.StyleBlock(`
	.` + IMAGE_CONTAINER + `text-with-link-btn {
		width: ` + strconv.Itoa(widthPX) + `px;
		height: ` + strconv.Itoa(height) + `px;
		position: relative;
		overflow: hidden;
	}
	.` + IMAGE_CONTAINER + `text-with-link-btn img {
		width: 100%;
	}
	.` + IMAGE_CONTAINER + `text-with-link-btn:before {
		content: "";
		position: absolute;
		top: 0;
		left: 0;
		width: 120%;
		height: 205%;
		background-color: ` + BACKGROUND_COLOR + `;
		transition: all 0.25s;
		transform: translate(` + strconv.Itoa(int((widthPX + height/3))) + `px, ` + strconv.Itoa(height) + `px) rotate(45deg);
	}
	.` + IMAGE_CONTAINER + `text-with-link-btn:hover:before {
		transform: translate(-` + strconv.Itoa(int(height/6)) + `px, -` + strconv.Itoa(int(widthPX/3)) + `px) rotate(45deg);
	}
	.` + IMAGE_CONTAINER + `text-with-link-btn div {
		position: absolute;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		display: flex;
		flex-direction: column;
		justify-content: space-around;
		color: white;
		text-transform: uppercase;
	}
	.` + IMAGE_CONTAINER + `text-with-link-btn div h2 {
		text-align: center;
		font-size: ` + strconv.Itoa(int(height/3/10*2)) + `px;
		padding: 20px;
		margin-top: 20px;
		background-color: ` + BACKGROUND_COLOR + `;
		transform: translate(` + strconv.Itoa(int(height/3*2)) + `px, -` + strconv.Itoa(height/3*2) + `px);
		transition: all 0.25s;
	}
	.` + IMAGE_CONTAINER + `text-with-link-btn:hover div h2 {
		transform: translate(0,0);
		transition-delay: 0.3s;
	}
	.` + IMAGE_CONTAINER + `text-with-link-btn div p {
		text-transform: capitalize;
		font-size: ` + strconv.Itoa(int(float64(height)/float64(3)/float64(10)*1.5)) + `px;
		width: 80%;
		align-self: center;
		text-align: center;
		transition: all 0.25s;
		transform: translate(-` + strconv.Itoa(int(height/3*2)) + `px, ` + strconv.Itoa(height/3*2) + `px);
	}
	.` + IMAGE_CONTAINER + `text-with-link-btn:hover div p {
		transform: translate(0,0);
		transition-delay: 0.4s;
	}
	.` + IMAGE_CONTAINER + `text-with-link-btn div a {
		text-decoration: none;
		background-color: ` + COLOR_BLACK + `;
		color: ` + COLOR_MAIN + `;
		width: ` + strconv.Itoa(int(height/3)) + `px;
		margin-bottom: 20px;
		align-self: center;
		text-align: center;
		transition: all 0.25s;
		font-size: ` + strconv.Itoa(int(float64(height)/float64(3)/float64(10)*1.75)) + `px;
		transform: translateY(` + strconv.Itoa(int(height/3)) + `px);
	}
	.` + IMAGE_CONTAINER + `text-with-link-btn:hover div a {
		transform: translateY(0);
		transition-delay: 0.5s;
	}
	`)
}

func LayeredImageCss(widthPX int) *elems.Element {
	var height = int(float64(widthPX) * 1.09375)
	return elems.StyleBlock(`
	.` + IMAGE_CONTAINER + `layered-image-container {
		position: relative;
		width: ` + strconv.Itoa(widthPX) + `px;
		height: ` + strconv.Itoa(height) + `px;
		background-color: ` + COLOR_MAIN + `;
		border-radius: ` + strconv.Itoa(widthPX/80) + `px;
		box-shadow: 0 2px 10px ` + BACKGROUND_COLOR + `;
	}
	.` + IMAGE_CONTAINER + `layered-image-container:before,
	.` + IMAGE_CONTAINER + `layered-image-container:after {
		content: "";
		position: absolute;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		border-radius: ` + strconv.Itoa(widthPX/80) + `px;
		background-color: ` + COLOR_MAIN + `;
		transition: 0.5s;
		z-index: -1;
	}
	.` + IMAGE_CONTAINER + `layered-image-container:hover:before {
		transform: rotate(20deg);
		box-shadow: 0 2px 20px ` + BACKGROUND_COLOR + `;
	}
	.` + IMAGE_CONTAINER + `layered-image-container:hover:after {
		transform: rotate(10deg);
		box-shadow: 0 2px 20px ` + BACKGROUND_COLOR + `;
	}
	.` + IMAGE_CONTAINER + `layered-image-container div:nth-child(1) {
		position: absolute;
		top: ` + strconv.Itoa(widthPX/32) + `px;
		left: ` + strconv.Itoa(widthPX/32) + `px;
		right: ` + strconv.Itoa(widthPX/32) + `px;
		bottom: ` + strconv.Itoa(widthPX/32) + `px;
		background-color: ` + COLOR_DARK_GRAY + `;
		transition: 0.5s;
		z-index: 2;
	}
	.` + IMAGE_CONTAINER + `layered-image-container:hover div:nth-child(1) {
		bottom: ` + strconv.Itoa(widthPX/4) + `px;
	}
	.` + IMAGE_CONTAINER + `layered-image-container div:nth-child(1) img {
		position: absolute;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		object-fit: cover;
	}
	.` + IMAGE_CONTAINER + `layered-image-container div:nth-child(2) {
		position: absolute;
		left: ` + strconv.Itoa(widthPX/32) + `px;
		right: ` + strconv.Itoa(widthPX/32) + `px;
		bottom: ` + strconv.Itoa(widthPX/32) + `px;
		height: ` + strconv.Itoa(int(float64(widthPX)/5.33333)) + `px;
		text-align: center;
	}
	.` + IMAGE_CONTAINER + `layered-image-container div:nth-child(2) h2 {
		margin: 0;
		padding: 0;
		font-weight: 600;
		font-size: ` + strconv.Itoa(int(float64(widthPX)/16)) + `px;
		color: ` + COLOR_DARK_GRAY + `;
		text-transform: uppercase;
	}
	.` + IMAGE_CONTAINER + `layered-image-container div:nth-child(2) h2 span {
		font-weight: 500;
		font-size: ` + strconv.Itoa(int(float64(widthPX)/20)) + `px;
		color: ` + COLOR_SIX + `;
	}
	`)
}

func LayeredImage(URL, imageURL, name, title string) *elems.Element {
	var anchor = elems.A().Href(URL)
	var container = anchor.Div().Class(IMAGE_CONTAINER + "layered-image-container")
	var nth_child_one = container.Div()
	nth_child_one.Img(imageURL, name)
	var nth_child_two = container.Div()
	var nameElem = nth_child_two.H2(name)
	nameElem.Br().NoClose()
	nameElem.Span(title)
	return anchor
}
