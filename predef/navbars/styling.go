package navbars

import (
	"strconv"
	"strings"

	"github.com/Nigel2392/gen/elems"
	"github.com/Nigel2392/gen/predef/csshelpers"
)

const DefaultNavStyle = `
.gohtml-navbar{
	background-color: #333;
  	overflow: hidden;
}
.gohtml-navbar .gohtml-navigation-link {
	float: left;
	display: block;
	color: #f2f2f2;
	text-align: center;
	padding: 14px 16px;
	text-decoration: none;
	font-size: 17px;
}
.gohtml-navbar .gohtml-navigation-link:hover {
	background: #ddd;
	color: black;
}
.gohtml-navbar .gohtml-navigation-link.gohtml-active {
	background-color: #9200ff;
	font-size: 20px;
	padding: 11px 13px;
	color: white;
}
.gohtml-left{
	float: left;
}
.gohtml-right{
	float: right;
}`

const DefaultSideNavStyle = `
.gohtml-sidenav {
	height: 100%;
	width: 0;
	position: fixed;
	z-index: 1;
	top: 0;
	left: 0;
	background-color: white;
	overflow-x: hidden;
	transition: 0.5s;
	padding-top: 10px;
}
.gohtml-main-content{
	transition: 0.3s;
}
.gohtml-sidenav .gohtml-navigation-link {
	padding: 8px 8px 8px 32px;
	text-decoration: none;
	font-size: 15px;
	color: #818181;
	display: block;
	transition: 0.3s;
	width: 100%;
}
.gohtml-sidenav .gohtml-content-nav .gohtml-navigation-link:hover {
	color: #000;
	background-color: #e0e0e0;
	text-decoration: none;
	position: relative;
}
.gohtml-navigation-link svg {
	position: absolute;
	right: 10px;
}
#gohtml-openbtn {
	cursor: pointer;
	margin-top: 5px;
	position: fixed;
	top: 5px;
	left: 5px;
}
#gohtml-closebtn {
	font-size: 46px;
	color: #333333;
	cursor: pointer;
	text-decoration: none;
	margin-right: 10px;
}
.gohtml-active {
	background-color: #e0e0e0;
	color: #000;
	text-decoration: none;
}
.gohtml-footer{
	position: absolute;
	bottom: 0;
	left: 0;
	width: 100%;
}`

func OfficialNavbarCss(height int) *elems.Element {
	var fontSize = int(float64(height) / 2.5)
	var paddingVertical = int(float64(height) / 4)
	var PaddingHorizontal = int(float64(height) / 2)
	return elems.StyleBlock(`
		.` + CLASS_NAVBAR + `-official {
			background-color: ` + COLOR_ONE + `;
			width: 100%;
			height: ` + strconv.Itoa(height) + `px;
			padding: 0;
			margin: 0;
			overflow: hidden;
			text-align: center;
		}
		.` + CLASS_NAVBAR + `-official .` + CLASS_NAVIGATION_LOGO + ` {
			height: 100%;
			width: auto;
			margin-left: ` + strconv.Itoa(PaddingHorizontal) + `px;
			margin-right: ` + strconv.Itoa(PaddingHorizontal) + `px;
		}
		.` + CLASS_NAVBAR_LEFT + ` {
			float: left;
			text-align: left;
			position: relative;
			height: ` + strconv.Itoa(height) + `px;
		}
		.` + CLASS_NAVBAR_RIGHT + ` {
			float: right;
			text-align: right;
			position: relative;
			height: ` + strconv.Itoa(height) + `px;
		}
		.` + CLASS_ACTIVE + ` {
			color: ` + COLOR_MAIN + `;
			background-color: ` + COLOR_ONE_LIGHTER + `;
		}
		.` + CLASS_NAVBAR + `-official .` + CLASS_NAVIGATION_LINK + ` {
			text-decoration: none;
			color: ` + COLOR_MAIN + `;
			font-size: ` + strconv.Itoa(fontSize) + `px;
			padding: ` + strconv.Itoa(paddingVertical*2) + `px ` + strconv.Itoa(PaddingHorizontal) + `px;
			height: ` + strconv.Itoa(height) + `px;
			line-height: ` + strconv.Itoa(height) + `px;
			vertical-align: middle;
			transition: all 0.3s;
		}
		.` + CLASS_NAVBAR + `-official .` + CLASS_NAVIGATION_LINK + `:hover {
			color: ` + COLOR_MAIN + `;
			background-color: ` + COLOR_ONE_LIGHTER + `;
			font-size: ` + strconv.Itoa(fontSize+6) + `px;
		}
	`)
}

// Returns style element and offset height for next elements.
// If no btnCol is given, we will take the complementary color of the background color.
func WailsFrameCss(widthBtn int, backcol string, btnCol ...string) (*elems.Element, int) {
	var padding_vert = widthBtn / 4
	var padding_hor = widthBtn / 4

	colr, err := csshelpers.Hex(backcol)
	if err != nil {
		panic(err)
	}
	var complementaryColor = colr.Complementary().Hex()
	if strings.EqualFold(complementaryColor, backcol) {
		if r, g, b := colr.RGB255(); r == 0 && g == 0 && b == 0 {
			complementaryColor = "#FFFFFF"
		} else {
			if r > 127 && g > 127 && b > 127 {
				complementaryColor = "#000000"
			} else {
				complementaryColor = "#FFFFFF"
			}
		}
	}
	if len(btnCol) > 0 {
		complementaryColor = btnCol[0]
	}
	return elems.StyleBlock(`
		body {
			margin: 0;
			padding: 0;
		}
		.gohtml-noselect {
			-webkit-touch-callout: none; /* iOS Safari */
			  -webkit-user-select: none; /* Safari */
			   -khtml-user-select: none; /* Konqueror HTML */
				 -moz-user-select: none; /* Old versions of Firefox */
				  -ms-user-select: none; /* Internet Explorer/Edge */
					  user-select: none; /* Non-prefixed version, currently
											supported by Chrome, Edge, Opera and Firefox */
		}
		.` + CLASS_NAVBAR + `-framed-wails {
			background-color: ` + backcol + `;
			width: 100%;
			height: ` + strconv.Itoa(widthBtn+(padding_hor*2)) + `px;
			position: fixed;
			top: 0;
			left: 0;
			z-index: 1000;
			overflow: hidden;
		}
		#gohtml-window-control-box {
			word-spacing: -5;
			position: fixed;
			top: 0;
			right: 10px;
			margin: 0;
		}
		#gohtml-window-control-box div {
			height: ` + strconv.Itoa(widthBtn) + `px;
			width: ` + strconv.Itoa(widthBtn) + `px;
			display: inline-block;
			margin: ` + strconv.Itoa(padding_vert) + `px ` + strconv.Itoa(padding_hor) + `px;
		}
		#gohtml-window-control-minimize {
			background-color: ` + backcol + `;
			border-radius: 50%;
			position: relative;
			transition: all 0.3s
		}
		#gohtml-window-control-minimize span {
			position: absolute;
			top: 65%;
			left: 50%;
			transform: translate(-50%, -50%);
			width: 50%;
			height: ` + strconv.Itoa(widthBtn/5) + `px;
			background-color: ` + complementaryColor + `;
			z-index: 1;
			transition: all 0.3s
		}
		#gohtml-window-control-minimize:hover {
			cursor: pointer;
			background-color: ` + complementaryColor + `;
		}
		#gohtml-window-control-minimize:hover span {
			background-color: ` + backcol + `;
		}
		#gohtml-window-control-maximize {
			background-color: ` + backcol + `;
			border-radius: 50%;
			position: relative;
			transition: all 0.3s;
		}
		#gohtml-window-control-maximize span {
			position: absolute;
			top: 50%;
			left: 50%;
			transform: translate(-50%, -50%);
			width: ` + strconv.Itoa(int(float64(widthBtn)/2.2)) + `px;
			height: ` + strconv.Itoa(int(float64(widthBtn)/2.2)) + `px;
			border: ` + strconv.Itoa(widthBtn/10) + `px solid ` + complementaryColor + `;
			transition: all 0.3s;
		}
		#gohtml-window-control-maximize:hover {
			cursor: pointer;
			background-color: ` + complementaryColor + `;
		}
		#gohtml-window-control-maximize:hover span {
			border: ` + strconv.Itoa(widthBtn/10) + `px solid ` + backcol + `;
		}

		#gohtml-window-control-close {
			background-color:  ` + backcol + `;
			border-radius: 50%;
			position: relative;
			transition: all 0.3s;
		}
		#gohtml-window-control-close span {
			position: absolute;
			top: 50%;
			left: 50%;
			transform: translate(-50%, -50%);
			transition: all 0.3s;
		}
		#gohtml-window-control-close span:nth-child(1),
		#gohtml-window-control-close span:nth-child(2) {
			content: "";
			position: absolute;
			top: 50%;
			left: 50%;
			width: 50%;
			height: ` + strconv.Itoa(widthBtn/6) + `px;
			background-color: ` + complementaryColor + `;
			z-index: 1;
			transition: all 0.3s;
		}
		#gohtml-window-control-close span:nth-child(1) {
			transform: translate(-50%, -50%) rotate(45deg);
		}
		#gohtml-window-control-close span:nth-child(2) {
			transform: translate(-50%, -50%) rotate(-45deg);
		}
		#gohtml-window-control-close:hover {
			cursor: pointer;
			background-color: ` + complementaryColor + `;
		}
		#gohtml-window-control-close:hover span:nth-child(1),
		#gohtml-window-control-close:hover span:nth-child(2) {
			background-color: ` + backcol + `;
		}
		.gohtml-window-control-left {
			float:left;
			text-align:left;
			display: flex;
			flex-direction: row;
			align-items: center;
		}
		.gohtml-window-control-favicon{
			width: ` + strconv.Itoa(widthBtn+(padding_hor*2)) + `px;
			height: ` + strconv.Itoa(widthBtn+(padding_hor*2)) + `px
			background-color: ` + backcol + `;
			margin-left: ` + strconv.Itoa(padding_hor) + `px;
			margin-right: ` + strconv.Itoa(padding_hor) + `px;
		}
		.` + CLASS_NAVIGATION_LINK + `-control {
			color: ` + complementaryColor + `;
			text-decoration: none;
			text-transform: capitalize;
			font-size: ` + strconv.Itoa(int(float64(widthBtn)/1.5)) + `px;
			font-family: sans-serif;
			font-weight: 600;
			transition: all 0.3s;
			padding: ` + strconv.Itoa(padding_vert) + `px ` + strconv.Itoa(padding_hor) + `px;
			height: ` + strconv.Itoa(widthBtn+(padding_hor*2)) + `px;
		}
		.` + CLASS_NAVIGATION_LINK + `-control:hover {
			color: ` + backcol + `;
			background-color: ` + complementaryColor + `;
		}
		.` + CLASS_ACTIVE + ` {
			color: ` + backcol + `;
			background-color: ` + complementaryColor + `;
		}
	`), widthBtn + (padding_hor * 2)
}

// width: calc(100% - ` + strconv.Itoa(int(float64(widthBtn+(padding_hor*2))*2.5)) + `px);

//  .gohtml-window-control-middle {
//  	height: ` + strconv.Itoa(widthBtn+(padding_hor*2)) + `px;
//  	background-color: ` + backcol + `;
//  	position: absolute;
//  	left: 50%;
//  	top: 50%;
//  	transform: translate(-50%, -30%);
//  }

func WailsFrameButtonsOnlyCss(widthBtn int, backcol string, btnCol ...string) (*elems.Element, int) {
	var padding_vert = widthBtn / 4
	var padding_hor = widthBtn / 4

	colr, err := csshelpers.Hex(backcol)
	if err != nil {
		panic(err)
	}
	var complementaryColor = colr.Complementary().Hex()
	if strings.EqualFold(complementaryColor, backcol) {
		if r, g, b := colr.RGB255(); r == 0 && g == 0 && b == 0 {
			complementaryColor = "#FFFFFF"
		} else {
			if r > 127 && g > 127 && b > 127 {
				complementaryColor = "#000000"
			} else {
				complementaryColor = "#FFFFFF"
			}
		}
	}
	if len(btnCol) > 0 {
		complementaryColor = btnCol[0]
	}
	return elems.StyleBlock(`
	#gohtml-window-control-box {
		word-spacing: -5;
		position: fixed;
		top: 0;
		right: 0;
		margin: 0;
		background-color: ` + backcol + `;
		border-radius: 0 0 0 ` + strconv.Itoa(widthBtn/2) + `px;
		padding-left:` + strconv.Itoa(widthBtn/2) + `px;
		padding-right:` + strconv.Itoa(widthBtn/3) + `px;
		padding-top:` + strconv.Itoa(widthBtn*2/30) + `px;
		z-index: 1000;
	}
	#gohtml-window-control-box div {
		height: ` + strconv.Itoa(widthBtn) + `px;
		width: ` + strconv.Itoa(widthBtn) + `px;
		display: inline-block;
		margin: ` + strconv.Itoa(padding_vert) + `px ` + strconv.Itoa(padding_hor) + `px;
	}
	#gohtml-window-control-minimize {
		background-color: ` + backcol + `;
		border-radius: 50%;
		position: relative;
		transition: all 0.3s
	}
	#gohtml-window-control-minimize span {
		position: absolute;
		top: 65%;
		left: 50%;
		transform: translate(-50%, -50%);
		width: 50%;
		height: ` + strconv.Itoa(widthBtn/5) + `px;
		background-color: ` + complementaryColor + `;
		z-index: 1;
		transition: all 0.3s
	}
	#gohtml-window-control-minimize:hover {
		cursor: pointer;
		background-color: ` + complementaryColor + `;
	}
	#gohtml-window-control-minimize:hover span {
		background-color: ` + backcol + `;
	}
	#gohtml-window-control-maximize {
		background-color: ` + backcol + `;
		border-radius: 50%;
		position: relative;
		transition: all 0.3s;
	}
	#gohtml-window-control-maximize span {
		position: absolute;
		top: 50%;
		left: 50%;
		transform: translate(-50%, -50%);
		width: ` + strconv.Itoa(int(float64(widthBtn)/2.2)) + `px;
		height: ` + strconv.Itoa(int(float64(widthBtn)/2.2)) + `px;
		border: ` + strconv.Itoa(widthBtn/10) + `px solid ` + complementaryColor + `;
		transition: all 0.3s;
	}
	#gohtml-window-control-maximize:hover {
		cursor: pointer;
		background-color: ` + complementaryColor + `;
	}
	#gohtml-window-control-maximize:hover span {
		border: ` + strconv.Itoa(widthBtn/10) + `px solid ` + backcol + `;
	}

	#gohtml-window-control-close {
		background-color:  ` + backcol + `;
		border-radius: 50%;
		position: relative;
		transition: all 0.3s;
	}
	#gohtml-window-control-close span {
		position: absolute;
		top: 50%;
		left: 50%;
		transform: translate(-50%, -50%);
		transition: all 0.3s;
	}
	#gohtml-window-control-close span:nth-child(1),
	#gohtml-window-control-close span:nth-child(2) {
		content: "";
		position: absolute;
		top: 50%;
		left: 50%;
		width: 50%;
		height: ` + strconv.Itoa(widthBtn/6) + `px;
		background-color: ` + complementaryColor + `;
		z-index: 1;
		transition: all 0.3s;
	}
	#gohtml-window-control-close span:nth-child(1) {
		transform: translate(-50%, -50%) rotate(45deg);
	}
	#gohtml-window-control-close span:nth-child(2) {
		transform: translate(-50%, -50%) rotate(-45deg);
	}
	#gohtml-window-control-close:hover {
		cursor: pointer;
		background-color: ` + complementaryColor + `;
	}
	#gohtml-window-control-close:hover span:nth-child(1),
	#gohtml-window-control-close:hover span:nth-child(2) {
		background-color: ` + backcol + `;
	}
`), widthBtn + (padding_hor * 2)
}
