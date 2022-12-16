package buttons

import (
	"strconv"
	"strings"

	"github.com/Nigel2392/gen/elems"
	"github.com/Nigel2392/gen/predef/csshelpers"
)

// Uses a user supplied color for the EnclosingSideDots button.
//   - Complementary color is used for the text.
func EnclosingSideDotsCss(mainColorHex string, widthPX int) *elems.Element {
	var cC, err = csshelpers.Hex(mainColorHex)
	if err != nil {
		panic(err)
	}
	var height = int(float64(widthPX) / 2.75)
	var fontSize = int(float64(widthPX) / 8.46153846154)
	var complementaryColor = cC.Complementary().Hex()
	if strings.EqualFold(complementaryColor, mainColorHex) {
		if r, g, b := cC.RGB255(); r == 0 && g == 0 && b == 0 {
			complementaryColor = "#FFFFFF"
		} else {
			complementaryColor = "#000000"
		}
	}
	return elems.StyleBlock(`
	.` + BUTTON_PREFIX + `enclosing-side-dots {
		width: ` + strconv.Itoa(widthPX) + `px;
		height: ` + strconv.Itoa(height) + `px;
		color: ` + mainColorHex + `;
		background-color: transparent;
		font-size: ` + strconv.Itoa(fontSize) + `px;
		text-decoration: none;
		text-transform: uppercase;
		text-align: center;
		line-height: ` + strconv.Itoa(height) + `px;
		padding: ` + strconv.Itoa(widthPX/11) + `px ` + strconv.Itoa(widthPX/11*2) + `px;
		font-family: "Lucida Console", Monaco, monospace;
		font-weight: 700;
		letter-spacing: ` + strconv.Itoa(widthPX/22/10*4) + `px;
		transition: all 0.5s;
		position: relative;
		border-bottom: ` + strconv.Itoa(widthPX/220*2) + `px solid ` + mainColorHex + `;
		border-radius: ` + strconv.Itoa(widthPX*10/220) + `px;
	}
	.` + BUTTON_PREFIX + `enclosing-side-dots:before,
	.` + BUTTON_PREFIX + `enclosing-side-dots:after {
		content: "";
		position: absolute;
		top: 50%;
		width: ` + strconv.Itoa(widthPX/22) + `px;
		height: ` + strconv.Itoa(widthPX/22) + `px;
		background-color: ` + mainColorHex + `;
		border-radius: 50%;
		transform: translateY(-50%);
		transition: all 0.3s;
		z-index: -1;
		opacity: 0;
	}
	.` + BUTTON_PREFIX + `enclosing-side-dots:before {
		left: 0;
		box-shadow: -` + strconv.Itoa(widthPX/22*10) + `px 0 0 ` + mainColorHex + `;
	}
	.` + BUTTON_PREFIX + `enclosing-side-dots:after {
		right: 0;
		box-shadow: ` + strconv.Itoa(widthPX/22*10) + `px 0 0 ` + mainColorHex + `;
	}
	.` + BUTTON_PREFIX + `enclosing-side-dots:hover:before {
		left: 50%;
		box-shadow: ` + strconv.Itoa(widthPX/22*3) + `px 0 0 ` + mainColorHex + `;
		transform: translateX(-50%) translateY(-50%);
		opacity: 1;
	}
	.` + BUTTON_PREFIX + `enclosing-side-dots:hover:after {
		right: 50%;
		box-shadow: -` + strconv.Itoa(widthPX/22*3) + `px 0 0 ` + mainColorHex + `;
		transform: translateX(50%) translateY(-50%);
		opacity: 1;
	}
	.` + BUTTON_PREFIX + `enclosing-side-dots span {
		position: absolute;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		background-color: ` + mainColorHex + `;
		border-radius: ` + strconv.Itoa(widthPX*10/220) + `px;
		transform: scale(0);
		transition: all 0.3s;
		z-index: -1;
		box-sizing: border-box;
	}
	.` + BUTTON_PREFIX + `enclosing-side-dots:hover span {
		transform: scale(1);
		transition-delay: 0.4s;
	}
	.` + BUTTON_PREFIX + `enclosing-side-dots:hover {
		color: ` + complementaryColor + `;
		transition-delay: 0.4s;
	}
	`)
}

// RotatingBackgroundCss returns the CSS for the RotatingBackground button.
// The mainColorHex is the color of the button, and the widthPX is the width of the button in pixels.
// Text color is automatically set to the complementary color of the mainColorHex.
func RotatingBackgroundCss(mainColorHex string, widthPX int) *elems.Element {
	var cC, err = csshelpers.Hex(mainColorHex)
	if err != nil {
		panic(err)
	}
	var height = int(float64(widthPX) / 2.66666666667)
	var fontSize = int(float64(widthPX) / 6.4)
	var complementaryColor = cC.Complementary().Hex()
	if strings.EqualFold(complementaryColor, mainColorHex) {
		if r, g, b := cC.RGB255(); r == 0 && g == 0 && b == 0 {
			complementaryColor = "#FFFFFF"
		} else {
			complementaryColor = "#000000"
		}
	}
	return elems.StyleBlock(`
	.` + BUTTON_PREFIX + `expanding-cut {
		text-decoration: none;
		width: ` + strconv.Itoa(widthPX) + `px;
		height: ` + strconv.Itoa(height) + `px;
		color: ` + mainColorHex + `;
		line-height: ` + strconv.Itoa(height) + `px;
		text-align: center;
		text-transform: uppercase;
		font-size: ` + strconv.Itoa(fontSize) + `px;
		padding: ` + strconv.Itoa(widthPX/20) + `px ` + strconv.Itoa(widthPX/10) + `px;
		position: relative;
		transition: all 0.3s;
	}
	.` + BUTTON_PREFIX + `expanding-cut span {
		position: absolute;
		width: 100%;
		height: 100%;
		top: 0;
		left: 0;
		background: transparent;
		border-radius: ` + strconv.Itoa(widthPX/20) + `px;
		overflow: hidden;
		border: ` + strconv.Itoa(widthPX*2/160) + `px solid ` + mainColorHex + `;
		z-index: -1;
	}
	.` + BUTTON_PREFIX + `expanding-cut span:before {
		content: "";
		position: absolute;
		width: 8%;
		height: 500%;
		top: 50%;
		left: 50%;
		transform: translate(-50%, -50%) rotate(-60deg);
		transition: all 0.3s;
		opacity: 0;
		filter: blur(5px);
	}
	.` + BUTTON_PREFIX + `expanding-cut:hover span:before {
		width: 100%;
		background-color: ` + mainColorHex + `;
		transform: translate(-50%, -50%) rotate(60deg);
		opacity: 1;
		filter: blur(0);
	}
	.` + BUTTON_PREFIX + `expanding-cut:hover {
		color: ` + complementaryColor + `;
	}
	`)
}
