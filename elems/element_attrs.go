package elems

import (
	"strconv"
)

func (e *Element) Style(styles ...string) *Element {
	e.semicolattrs["style"] = append(e.semicolattrs["style"], styles...)
	return e
}

func (e *Element) Src(src string) *Element {
	e.attrs["src"] = append(e.attrs["src"], src)
	return e
}

func (e *Element) Href(href string) *Element {
	e.attrs["href"] = append(e.attrs["href"], href)
	return e
}

func (e *Element) Class(classes ...string) *Element {
	e.attrs["class"] = append(e.attrs["class"], classes...)
	return e
}

func (e *Element) ID(id string) *Element {
	e.attrs["id"] = append(e.attrs["id"], id)
	return e
}

func (e *Element) Custom(key string, values ...string) *Element {
	e.attrs[key] = append(e.attrs[key], values...)
	return e
}

func (e *Element) NoClose() *Element {
	e.autoClose = false

	return e
}

// Accept is the attr_accept attribute. (Autogenerated)
func (e *Element) A_Accept(attr_accept ...string) *Element {
	e.attrs["accept"] = append(e.attrs["accept"], attr_accept...)
	return e
}

// Accesskey is the attr_accesskey attribute. (Autogenerated)
func (e *Element) A_Accesskey(attr_accesskey ...string) *Element {
	e.attrs["accesskey"] = append(e.attrs["accesskey"], attr_accesskey...)
	return e
}

// Action is the attr_action attribute. (Autogenerated)
func (e *Element) Action(attr_action ...string) *Element {
	e.attrs["action"] = append(e.attrs["action"], attr_action...)
	return e
}

// Align is the attr_align attribute. (Autogenerated)
func (e *Element) A_Align(attr_align ...string) *Element {
	e.attrs["align"] = append(e.attrs["align"], attr_align...)
	return e
}

// Alt is the attr_alt attribute. (Autogenerated)
func (e *Element) Alt(attr_alt string) *Element {
	e.attrs["alt"] = append(e.attrs["alt"], attr_alt)
	return e
}

// Async is the attr_async attribute. (Autogenerated)
func (e *Element) A_Async() *Element {
	e.boolattrs["async"] = true
	return e
}

// Autocomplete is the attr_autocomplete attribute. (Autogenerated)
func (e *Element) A_Autocomplete(def bool) *Element {
	e.boolattrs["autocomplete"] = def
	return e
}

// Autofocus is the attr_autofocus attribute. (Autogenerated)
func (e *Element) A_Autofocus() *Element {
	e.boolattrs["autofocus"] = true
	return e
}

// Autoplay is the attr_autoplay attribute. (Autogenerated)
func (e *Element) A_Autoplay() *Element {
	e.boolattrs["autoplay"] = true
	return e
}

// Bgcolor is the attr_bgcolor attribute. (Autogenerated)
func (e *Element) Bgcolor(attr_bgcolor string) *Element {
	e.attrs["bgcolor"] = append(e.attrs["bgcolor"], attr_bgcolor)
	return e
}

func (e *Element) BackgroundColor(color string) *Element {
	return e.Style("background-color:" + color)
}

// Border is the attr_border attribute. (Autogenerated)
func (e *Element) Border(attr_border ...string) *Element {
	e.attrs["border"] = append(e.attrs["border"], attr_border...)
	return e
}

// Charset is the attr_charset attribute. (Autogenerated)
func (e *Element) Charset(attr_charset ...string) *Element {
	e.attrs["charset"] = append(e.attrs["charset"], attr_charset...)
	return e
}

// Checked is the attr_checked attribute. (Autogenerated)
func (e *Element) Checked() *Element {
	e.boolattrs["checked"] = true
	return e
}

// Cite is the attr_cite attribute. (Autogenerated)
func (e *Element) A_Cite(attr_cite ...string) *Element {
	e.attrs["cite"] = append(e.attrs["cite"], attr_cite...)
	return e
}

// Color is the attr_color attribute. (Autogenerated)
func (e *Element) Color(attr_color string) *Element {
	e.attrs["color"] = append(e.attrs["color"], attr_color)
	return e
}

// Cols is the attr_cols attribute. (Autogenerated)
func (e *Element) A_Cols(attr_cols string) *Element {
	e.attrs["cols"] = append(e.attrs["cols"], attr_cols)
	return e
}

// Colspan is the attr_colspan attribute. (Autogenerated)
func (e *Element) A_Colspan(attr_colspan string) *Element {
	e.attrs["colspan"] = append(e.attrs["colspan"], attr_colspan)
	return e
}

// Content is the attr_content attribute. (Autogenerated)
func (e *Element) A_Content(attr_content ...string) *Element {
	e.attrs["content"] = append(e.attrs["content"], attr_content...)
	return e
}

// Contenteditable is the attr_contenteditable attribute. (Autogenerated)
func (e *Element) A_Contenteditable() *Element {
	e.boolattrs["contenteditable"] = true
	return e
}

// Contextmenu is the attr_contextmenu attribute. (Autogenerated)
func (e *Element) A_Contextmenu(attr_contextmenu ...string) *Element {
	e.attrs["contextmenu"] = append(e.attrs["contextmenu"], attr_contextmenu...)
	return e
}

// Controls is the attr_controls attribute. (Autogenerated)
func (e *Element) A_Controls(attr_controls ...string) *Element {
	e.attrs["controls"] = append(e.attrs["controls"], attr_controls...)
	return e
}

// Coords is the attr_coords attribute. (Autogenerated)
func (e *Element) A_Coords(attr_coords ...string) *Element {
	e.attrs["coords"] = append(e.attrs["coords"], attr_coords...)
	return e
}

// Data is the attr_data attribute. (Autogenerated)
func (e *Element) A_Data(attr_data ...string) *Element {
	e.attrs["data"] = append(e.attrs["data"], attr_data...)
	return e
}

// Datetime is the attr_datetime attribute. (Autogenerated)
func (e *Element) A_Datetime(attr_datetime ...string) *Element {
	e.attrs["datetime"] = append(e.attrs["datetime"], attr_datetime...)
	return e
}

// Default is the attr_default attribute. (Autogenerated)
func (e *Element) A_Default(attr_default ...string) *Element {
	e.attrs["default"] = append(e.attrs["default"], attr_default...)
	return e
}

// Defer is the attr_defer attribute. (Autogenerated)
func (e *Element) A_Defer() *Element {
	e.boolattrs["defer"] = true
	return e
}

// Dir is the attr_dir attribute. (Autogenerated)
func (e *Element) A_Dir(attr_dir ...string) *Element {
	e.attrs["dir"] = append(e.attrs["dir"], attr_dir...)
	return e
}

// Dirname is the attr_dirname attribute. (Autogenerated)
func (e *Element) A_Dirname(attr_dirname ...string) *Element {
	e.attrs["dirname"] = append(e.attrs["dirname"], attr_dirname...)
	return e
}

// Disabled is the attr_disabled attribute. (Autogenerated)
func (e *Element) A_Disabled() *Element {
	e.boolattrs["disabled"] = true
	return e
}

// Download is the attr_download attribute. (Autogenerated)
func (e *Element) A_Download(attr_download ...string) *Element {
	e.attrs["download"] = append(e.attrs["download"], attr_download...)
	return e
}

// Draggable is the attr_draggable attribute. (Autogenerated)
func (e *Element) A_Draggable() *Element {
	e.boolattrs["draggable"] = true
	return e
}

// Dropzone is the attr_dropzone attribute. (Autogenerated)
func (e *Element) A_Dropzone(attr_dropzone ...string) *Element {
	e.attrs["dropzone"] = append(e.attrs["dropzone"], attr_dropzone...)
	return e
}

// Enctype is the attr_enctype attribute. (Autogenerated)
func (e *Element) A_Enctype(attr_enctype ...string) *Element {
	e.attrs["enctype"] = append(e.attrs["enctype"], attr_enctype...)
	return e
}

// For is the attr_for attribute. (Autogenerated)
func (e *Element) A_For(attr_for ...string) *Element {
	e.attrs["for"] = append(e.attrs["for"], attr_for...)
	return e
}

// Form is the attr_form attribute. (Autogenerated)
func (e *Element) A_Form(attr_form ...string) *Element {
	e.attrs["form"] = append(e.attrs["form"], attr_form...)
	return e
}

// Formaction is the attr_formaction attribute. (Autogenerated)
func (e *Element) A_Formaction(attr_formaction ...string) *Element {
	e.attrs["formaction"] = append(e.attrs["formaction"], attr_formaction...)
	return e
}

// Headers is the attr_headers attribute. (Autogenerated)
func (e *Element) A_Headers(attr_headers ...string) *Element {
	e.attrs["headers"] = append(e.attrs["headers"], attr_headers...)
	return e
}

// Height is the attr_height attribute. (Autogenerated)
func (e *Element) Height(attr_height string) *Element {
	e.attrs["height"] = append(e.attrs["height"], attr_height)
	return e
}

// Hidden is the attr_hidden attribute. (Autogenerated)
func (e *Element) Hidden() *Element {
	e.boolattrs["hidden"] = true
	return e
}

// High is the attr_high attribute. (Autogenerated)
func (e *Element) A_High(attr_high ...string) *Element {
	e.attrs["high"] = append(e.attrs["high"], attr_high...)
	return e
}

// Hreflang is the attr_hreflang attribute. (Autogenerated)
func (e *Element) A_Hreflang(attr_hreflang ...string) *Element {
	e.attrs["hreflang"] = append(e.attrs["hreflang"], attr_hreflang...)
	return e
}

// Http is the attr_http attribute. (Autogenerated)
func (e *Element) A_Http(attr_http ...string) *Element {
	e.attrs["http"] = append(e.attrs["http"], attr_http...)
	return e
}

// Ismap is the attr_ismap attribute. (Autogenerated)
func (e *Element) A_Ismap(attr_ismap ...string) *Element {
	e.attrs["ismap"] = append(e.attrs["ismap"], attr_ismap...)
	return e
}

// Kind is the attr_kind attribute. (Autogenerated)
func (e *Element) A_Kind(attr_kind ...string) *Element {
	e.attrs["kind"] = append(e.attrs["kind"], attr_kind...)
	return e
}

// Label is the attr_label attribute. (Autogenerated)
func (e *Element) A_Label(attr_label ...string) *Element {
	e.attrs["label"] = append(e.attrs["label"], attr_label...)
	return e
}

// Lang is the attr_lang attribute. (Autogenerated)
func (e *Element) A_Lang(attr_lang ...string) *Element {
	e.attrs["lang"] = append(e.attrs["lang"], attr_lang...)
	return e
}

// List is the attr_list attribute. (Autogenerated)
func (e *Element) A_List(attr_list ...string) *Element {
	e.attrs["list"] = append(e.attrs["list"], attr_list...)
	return e
}

// Loop is the attr_loop attribute. (Autogenerated)
func (e *Element) A_Loop(attr_loop ...string) *Element {
	e.attrs["loop"] = append(e.attrs["loop"], attr_loop...)
	return e
}

// Low is the attr_low attribute. (Autogenerated)
func (e *Element) A_Low(attr_low ...string) *Element {
	e.attrs["low"] = append(e.attrs["low"], attr_low...)
	return e
}

// Max is the attr_max attribute. (Autogenerated)
func (e *Element) A_Max(attr_max ...string) *Element {
	e.attrs["max"] = append(e.attrs["max"], attr_max...)
	return e
}

// Maxlength is the attr_maxlength attribute. (Autogenerated)
func (e *Element) A_Maxlength(attr_maxlength ...string) *Element {
	e.attrs["maxlength"] = append(e.attrs["maxlength"], attr_maxlength...)
	return e
}

// Media is the attr_media attribute. (Autogenerated)
func (e *Element) A_Media(attr_media ...string) *Element {
	e.attrs["media"] = append(e.attrs["media"], attr_media...)
	return e
}

// Method is the attr_method attribute. (Autogenerated)
func (e *Element) Method(attr_method ...string) *Element {
	e.attrs["method"] = append(e.attrs["method"], attr_method...)
	return e
}

// Min is the attr_min attribute. (Autogenerated)
func (e *Element) A_Min(attr_min ...string) *Element {
	e.attrs["min"] = append(e.attrs["min"], attr_min...)
	return e
}

// Multiple is the attr_multiple attribute. (Autogenerated)
func (e *Element) A_Multiple(attr_multiple ...string) *Element {
	e.attrs["multiple"] = append(e.attrs["multiple"], attr_multiple...)
	return e
}

// Muted is the attr_muted attribute. (Autogenerated)
func (e *Element) A_Muted() *Element {
	e.boolattrs["muted"] = true
	return e
}

// Name is the attr_name attribute. (Autogenerated)
func (e *Element) Name(attr_name string) *Element {
	e.attrs["name"] = append(e.attrs["name"], attr_name)
	return e
}

// Novalidate is the attr_novalidate attribute. (Autogenerated)
func (e *Element) A_Novalidate() *Element {
	e.boolattrs["novalidate"] = true
	return e
}

// Onabort is the attr_onabort attribute. (Autogenerated)
func (e *Element) Onabort(attr_onabort ...string) *Element {
	e.semicolattrs["onabort"] = append(e.semicolattrs["onabort"], attr_onabort...)
	return e
}

// Onafterprint is the attr_onafterprint attribute. (Autogenerated)
func (e *Element) Onafterprint(attr_onafterprint ...string) *Element {
	e.semicolattrs["onafterprint"] = append(e.semicolattrs["onafterprint"], attr_onafterprint...)
	return e
}

// Onbeforeprint is the attr_onbeforeprint attribute. (Autogenerated)
func (e *Element) Onbeforeprint(attr_onbeforeprint ...string) *Element {
	e.semicolattrs["onbeforeprint"] = append(e.semicolattrs["onbeforeprint"], attr_onbeforeprint...)
	return e
}

// Onbeforeunload is the attr_onbeforeunload attribute. (Autogenerated)
func (e *Element) Onbeforeunload(attr_onbeforeunload ...string) *Element {
	e.semicolattrs["onbeforeunload"] = append(e.semicolattrs["onbeforeunload"], attr_onbeforeunload...)
	return e
}

// Onblur is the attr_onblur attribute. (Autogenerated)
func (e *Element) Onblur(attr_onblur ...string) *Element {
	e.semicolattrs["onblur"] = append(e.semicolattrs["onblur"], attr_onblur...)
	return e
}

// Oncanplay is the attr_oncanplay attribute. (Autogenerated)
func (e *Element) Oncanplay(attr_oncanplay ...string) *Element {
	e.semicolattrs["oncanplay"] = append(e.semicolattrs["oncanplay"], attr_oncanplay...)
	return e
}

// Oncanplaythrough is the attr_oncanplaythrough attribute. (Autogenerated)
func (e *Element) Oncanplaythrough(attr_oncanplaythrough ...string) *Element {
	e.semicolattrs["oncanplaythrough"] = append(e.semicolattrs["oncanplaythrough"], attr_oncanplaythrough...)
	return e
}

// Onchange is the attr_onchange attribute. (Autogenerated)
func (e *Element) Onchange(attr_onchange ...string) *Element {
	e.semicolattrs["onchange"] = append(e.semicolattrs["onchange"], attr_onchange...)
	return e
}

// Onclick is the attr_onclick attribute. (Autogenerated)
func (e *Element) Onclick(attr_onclick ...string) *Element {
	e.semicolattrs["onclick"] = append(e.semicolattrs["onclick"], attr_onclick...)
	return e
}

// Oncontextmenu is the attr_oncontextmenu attribute. (Autogenerated)
func (e *Element) Oncontextmenu(attr_oncontextmenu ...string) *Element {
	e.semicolattrs["oncontextmenu"] = append(e.semicolattrs["oncontextmenu"], attr_oncontextmenu...)
	return e
}

// Oncopy is the attr_oncopy attribute. (Autogenerated)
func (e *Element) Oncopy(attr_oncopy ...string) *Element {
	e.semicolattrs["oncopy"] = append(e.semicolattrs["oncopy"], attr_oncopy...)
	return e
}

// Oncuechange is the attr_oncuechange attribute. (Autogenerated)
func (e *Element) Oncuechange(attr_oncuechange ...string) *Element {
	e.semicolattrs["oncuechange"] = append(e.semicolattrs["oncuechange"], attr_oncuechange...)
	return e
}

// Oncut is the attr_oncut attribute. (Autogenerated)
func (e *Element) Oncut(attr_oncut ...string) *Element {
	e.semicolattrs["oncut"] = append(e.semicolattrs["oncut"], attr_oncut...)
	return e
}

// Ondblclick is the attr_ondblclick attribute. (Autogenerated)
func (e *Element) Ondblclick(attr_ondblclick ...string) *Element {
	e.semicolattrs["ondblclick"] = append(e.semicolattrs["ondblclick"], attr_ondblclick...)
	return e
}

// Ondrag is the attr_ondrag attribute. (Autogenerated)
func (e *Element) Ondrag(attr_ondrag ...string) *Element {
	e.semicolattrs["ondrag"] = append(e.semicolattrs["ondrag"], attr_ondrag...)
	return e
}

// Ondragend is the attr_ondragend attribute. (Autogenerated)
func (e *Element) Ondragend(attr_ondragend ...string) *Element {
	e.semicolattrs["ondragend"] = append(e.semicolattrs["ondragend"], attr_ondragend...)
	return e
}

// Ondragenter is the attr_ondragenter attribute. (Autogenerated)
func (e *Element) Ondragenter(attr_ondragenter ...string) *Element {
	e.semicolattrs["ondragenter"] = append(e.semicolattrs["ondragenter"], attr_ondragenter...)
	return e
}

// Ondragleave is the attr_ondragleave attribute. (Autogenerated)
func (e *Element) Ondragleave(attr_ondragleave ...string) *Element {
	e.semicolattrs["ondragleave"] = append(e.semicolattrs["ondragleave"], attr_ondragleave...)
	return e
}

// Ondragover is the attr_ondragover attribute. (Autogenerated)
func (e *Element) Ondragover(attr_ondragover ...string) *Element {
	e.semicolattrs["ondragover"] = append(e.semicolattrs["ondragover"], attr_ondragover...)
	return e
}

// Ondragstart is the attr_ondragstart attribute. (Autogenerated)
func (e *Element) Ondragstart(attr_ondragstart ...string) *Element {
	e.semicolattrs["ondragstart"] = append(e.semicolattrs["ondragstart"], attr_ondragstart...)
	return e
}

// Ondrop is the attr_ondrop attribute. (Autogenerated)
func (e *Element) Ondrop(attr_ondrop ...string) *Element {
	e.semicolattrs["ondrop"] = append(e.semicolattrs["ondrop"], attr_ondrop...)
	return e
}

// Ondurationchange is the attr_ondurationchange attribute. (Autogenerated)
func (e *Element) Ondurationchange(attr_ondurationchange ...string) *Element {
	e.semicolattrs["ondurationchange"] = append(e.semicolattrs["ondurationchange"], attr_ondurationchange...)
	return e
}

// Onemptied is the attr_onemptied attribute. (Autogenerated)
func (e *Element) Onemptied(attr_onemptied ...string) *Element {
	e.semicolattrs["onemptied"] = append(e.semicolattrs["onemptied"], attr_onemptied...)
	return e
}

// Onended is the attr_onended attribute. (Autogenerated)
func (e *Element) Onended(attr_onended ...string) *Element {
	e.semicolattrs["onended"] = append(e.semicolattrs["onended"], attr_onended...)
	return e
}

// Onerror is the attr_onerror attribute. (Autogenerated)
func (e *Element) Onerror(attr_onerror ...string) *Element {
	e.semicolattrs["onerror"] = append(e.semicolattrs["onerror"], attr_onerror...)
	return e
}

// Onfocus is the attr_onfocus attribute. (Autogenerated)
func (e *Element) Onfocus(attr_onfocus ...string) *Element {
	e.semicolattrs["onfocus"] = append(e.semicolattrs["onfocus"], attr_onfocus...)
	return e
}

// Onhashchange is the attr_onhashchange attribute. (Autogenerated)
func (e *Element) Onhashchange(attr_onhashchange ...string) *Element {
	e.semicolattrs["onhashchange"] = append(e.semicolattrs["onhashchange"], attr_onhashchange...)
	return e
}

// Oninput is the attr_oninput attribute. (Autogenerated)
func (e *Element) Oninput(attr_oninput ...string) *Element {
	e.semicolattrs["oninput"] = append(e.semicolattrs["oninput"], attr_oninput...)
	return e
}

// Oninvalid is the attr_oninvalid attribute. (Autogenerated)
func (e *Element) Oninvalid(attr_oninvalid ...string) *Element {
	e.semicolattrs["oninvalid"] = append(e.semicolattrs["oninvalid"], attr_oninvalid...)
	return e
}

// Onkeydown is the attr_onkeydown attribute. (Autogenerated)
func (e *Element) Onkeydown(attr_onkeydown ...string) *Element {
	e.semicolattrs["onkeydown"] = append(e.semicolattrs["onkeydown"], attr_onkeydown...)
	return e
}

// Onkeypress is the attr_onkeypress attribute. (Autogenerated)
func (e *Element) Onkeypress(attr_onkeypress ...string) *Element {
	e.semicolattrs["onkeypress"] = append(e.semicolattrs["onkeypress"], attr_onkeypress...)
	return e
}

// Onkeyup is the attr_onkeyup attribute. (Autogenerated)
func (e *Element) Onkeyup(attr_onkeyup ...string) *Element {
	e.semicolattrs["onkeyup"] = append(e.semicolattrs["onkeyup"], attr_onkeyup...)
	return e
}

// Onload is the attr_onload attribute. (Autogenerated)
func (e *Element) Onload(attr_onload ...string) *Element {
	e.semicolattrs["onload"] = append(e.semicolattrs["onload"], attr_onload...)
	return e
}

// Onloadeddata is the attr_onloadeddata attribute. (Autogenerated)
func (e *Element) Onloadeddata(attr_onloadeddata ...string) *Element {
	e.semicolattrs["onloadeddata"] = append(e.semicolattrs["onloadeddata"], attr_onloadeddata...)
	return e
}

// Onloadedmetadata is the attr_onloadedmetadata attribute. (Autogenerated)
func (e *Element) Onloadedmetadata(attr_onloadedmetadata ...string) *Element {
	e.semicolattrs["onloadedmetadata"] = append(e.semicolattrs["onloadedmetadata"], attr_onloadedmetadata...)
	return e
}

// Onloadstart is the attr_onloadstart attribute. (Autogenerated)
func (e *Element) Onloadstart(attr_onloadstart ...string) *Element {
	e.semicolattrs["onloadstart"] = append(e.semicolattrs["onloadstart"], attr_onloadstart...)
	return e
}

// Onmousedown is the attr_onmousedown attribute. (Autogenerated)
func (e *Element) Onmousedown(attr_onmousedown ...string) *Element {
	e.semicolattrs["onmousedown"] = append(e.semicolattrs["onmousedown"], attr_onmousedown...)
	return e
}

// Onmousemove is the attr_onmousemove attribute. (Autogenerated)
func (e *Element) Onmousemove(attr_onmousemove ...string) *Element {
	e.semicolattrs["onmousemove"] = append(e.semicolattrs["onmousemove"], attr_onmousemove...)
	return e
}

// Onmouseout is the attr_onmouseout attribute. (Autogenerated)
func (e *Element) Onmouseout(attr_onmouseout ...string) *Element {
	e.semicolattrs["mouseout"] = append(e.semicolattrs["mouseout"], attr_onmouseout...)
	return e
}

// Onmouseover is the attr_onmouseover attribute. (Autogenerated)
func (e *Element) Onmouseover(attr_onmouseover ...string) *Element {
	e.semicolattrs["onmouseover"] = append(e.semicolattrs["onmouseover"], attr_onmouseover...)
	return e
}

// Onmouseup is the attr_onmouseup attribute. (Autogenerated)
func (e *Element) Onmouseup(attr_onmouseup ...string) *Element {
	e.semicolattrs["onmouseup"] = append(e.semicolattrs["onmouseup"], attr_onmouseup...)
	return e
}

// Onmousewheel is the attr_onmousewheel attribute. (Autogenerated)
func (e *Element) Onmousewheel(attr_onmousewheel ...string) *Element {
	e.semicolattrs["onmousewheel"] = append(e.semicolattrs["onmousewheel"], attr_onmousewheel...)
	return e
}

// Onoffline is the attr_onoffline attribute. (Autogenerated)
func (e *Element) Onoffline(attr_onoffline ...string) *Element {
	e.semicolattrs["onoffline"] = append(e.semicolattrs["onoffline"], attr_onoffline...)
	return e
}

// Ononline is the attr_ononline attribute. (Autogenerated)
func (e *Element) Ononline(attr_ononline ...string) *Element {
	e.semicolattrs["ononline"] = append(e.semicolattrs["ononline"], attr_ononline...)
	return e
}

// Onpagehide is the attr_onpagehide attribute. (Autogenerated)
func (e *Element) Onpagehide(attr_onpagehide ...string) *Element {
	e.semicolattrs["onpagehide"] = append(e.semicolattrs["onpagehide"], attr_onpagehide...)
	return e
}

// Onpageshow is the attr_onpageshow attribute. (Autogenerated)
func (e *Element) Onpageshow(attr_onpageshow ...string) *Element {
	e.semicolattrs["onpageshow"] = append(e.semicolattrs["onpageshow"], attr_onpageshow...)
	return e
}

// Onpaste is the attr_onpaste attribute. (Autogenerated)
func (e *Element) Onpaste(attr_onpaste ...string) *Element {
	e.semicolattrs["onpaste"] = append(e.semicolattrs["onpaste"], attr_onpaste...)
	return e
}

// Onpause is the attr_onpause attribute. (Autogenerated)
func (e *Element) Onpause(attr_onpause ...string) *Element {
	e.semicolattrs["onpause"] = append(e.semicolattrs["onpause"], attr_onpause...)
	return e
}

// Onplay is the attr_onplay attribute. (Autogenerated)
func (e *Element) Onplay(attr_onplay ...string) *Element {
	e.semicolattrs["onplay"] = append(e.semicolattrs["onplay"], attr_onplay...)
	return e
}

// Onplaying is the attr_onplaying attribute. (Autogenerated)
func (e *Element) Onplaying(attr_onplaying ...string) *Element {
	e.semicolattrs["onplaying"] = append(e.semicolattrs["onplaying"], attr_onplaying...)
	return e
}

// Onpopstate is the attr_onpopstate attribute. (Autogenerated)
func (e *Element) Onpopstate(attr_onpopstate ...string) *Element {
	e.semicolattrs["onpopstate"] = append(e.semicolattrs["onpopstate"], attr_onpopstate...)
	return e
}

// Onprogress is the attr_onprogress attribute. (Autogenerated)
func (e *Element) Onprogress(attr_onprogress ...string) *Element {
	e.semicolattrs["progress"] = append(e.semicolattrs["progress"], attr_onprogress...)
	return e
}

// Onratechange is the attr_onratechange attribute. (Autogenerated)
func (e *Element) Onratechange(attr_onratechange ...string) *Element {
	e.semicolattrs["ratechange"] = append(e.semicolattrs["ratechange"], attr_onratechange...)
	return e
}

// Onreset is the attr_onreset attribute. (Autogenerated)
func (e *Element) Onreset(attr_onreset ...string) *Element {
	e.semicolattrs["reset"] = append(e.semicolattrs["reset"], attr_onreset...)
	return e
}

// Onresize is the attr_onresize attribute. (Autogenerated)
func (e *Element) Onresize(attr_onresize ...string) *Element {
	e.semicolattrs["resize"] = append(e.semicolattrs["resize"], attr_onresize...)
	return e
}

// Onscroll is the attr_onscroll attribute. (Autogenerated)
func (e *Element) Onscroll(attr_onscroll ...string) *Element {
	e.semicolattrs["scroll"] = append(e.semicolattrs["scroll"], attr_onscroll...)
	return e
}

// Onsearch is the attr_onsearch attribute. (Autogenerated)
func (e *Element) Onsearch(attr_onsearch ...string) *Element {
	e.semicolattrs["search"] = append(e.semicolattrs["search"], attr_onsearch...)
	return e
}

// Onseeked is the attr_onseeked attribute. (Autogenerated)
func (e *Element) Onseeked(attr_onseeked ...string) *Element {
	e.semicolattrs["seeked"] = append(e.semicolattrs["seeked"], attr_onseeked...)
	return e
}

// Onseeking is the attr_onseeking attribute. (Autogenerated)
func (e *Element) Onseeking(attr_onseeking ...string) *Element {
	e.semicolattrs["seeking"] = append(e.semicolattrs["seeking"], attr_onseeking...)
	return e
}

// Onselect is the attr_onselect attribute. (Autogenerated)
func (e *Element) Onselect(attr_onselect ...string) *Element {
	e.semicolattrs["select"] = append(e.semicolattrs["select"], attr_onselect...)
	return e
}

// Onshow is the attr_onshow attribute. (Autogenerated)
func (e *Element) Onshow(attr_onshow ...string) *Element {
	e.semicolattrs["show"] = append(e.semicolattrs["show"], attr_onshow...)
	return e
}

// Onstalled is the attr_onstalled attribute. (Autogenerated)
func (e *Element) Onstalled(attr_onstalled ...string) *Element {
	e.semicolattrs["stalled"] = append(e.semicolattrs["stalled"], attr_onstalled...)
	return e
}

// Onstorage is the attr_onstorage attribute. (Autogenerated)
func (e *Element) Onstorage(attr_onstorage ...string) *Element {
	e.semicolattrs["storage"] = append(e.semicolattrs["storage"], attr_onstorage...)
	return e
}

// Onsubmit is the attr_onsubmit attribute. (Autogenerated)
func (e *Element) Onsubmit(attr_onsubmit ...string) *Element {
	e.semicolattrs["submit"] = append(e.semicolattrs["submit"], attr_onsubmit...)
	return e
}

// Onsuspend is the attr_onsuspend attribute. (Autogenerated)
func (e *Element) Onsuspend(attr_onsuspend ...string) *Element {
	e.semicolattrs["suspend"] = append(e.semicolattrs["suspend"], attr_onsuspend...)
	return e
}

// Ontimeupdate is the attr_ontimeupdate attribute. (Autogenerated)
func (e *Element) Ontimeupdate(attr_ontimeupdate ...string) *Element {
	e.semicolattrs["timeupdate"] = append(e.semicolattrs["timeupdate"], attr_ontimeupdate...)
	return e
}

// Ontoggle is the attr_ontoggle attribute. (Autogenerated)
func (e *Element) Ontoggle(attr_ontoggle ...string) *Element {
	e.semicolattrs["toggle"] = append(e.semicolattrs["toggle"], attr_ontoggle...)
	return e
}

// Onunload is the attr_onunload attribute. (Autogenerated)
func (e *Element) Onunload(attr_onunload ...string) *Element {
	e.semicolattrs["unload"] = append(e.semicolattrs["unload"], attr_onunload...)
	return e
}

// Onvolumechange is the attr_onvolumechange attribute. (Autogenerated)
func (e *Element) Onvolumechange(attr_onvolumechange ...string) *Element {
	e.semicolattrs["volumechange"] = append(e.semicolattrs["volumechange"], attr_onvolumechange...)
	return e
}

// Onwaiting is the attr_onwaiting attribute. (Autogenerated)
func (e *Element) Onwaiting(attr_onwaiting ...string) *Element {
	e.semicolattrs["waiting"] = append(e.semicolattrs["waiting"], attr_onwaiting...)
	return e
}

// Onwheel is the attr_onwheel attribute. (Autogenerated)
func (e *Element) Onwheel(attr_onwheel ...string) *Element {
	e.semicolattrs["wheel"] = append(e.semicolattrs["wheel"], attr_onwheel...)
	return e
}

// Open is the attr_open attribute. (Autogenerated)
func (e *Element) A_Open(attr_open ...string) *Element {
	e.attrs["open"] = append(e.attrs["open"], attr_open...)
	return e
}

// Optimum is the attr_optimum attribute. (Autogenerated)
func (e *Element) A_Optimum(attr_optimum ...string) *Element {
	e.attrs["optimum"] = append(e.attrs["optimum"], attr_optimum...)
	return e
}

// Pattern is the attr_pattern attribute. (Autogenerated)
func (e *Element) A_Pattern(attr_pattern ...string) *Element {
	e.attrs["pattern"] = append(e.attrs["pattern"], attr_pattern...)
	return e
}

// Placeholder is the attr_placeholder attribute. (Autogenerated)
func (e *Element) A_Placeholder(attr_placeholder ...string) *Element {
	e.attrs["placeholder"] = append(e.attrs["placeholder"], attr_placeholder...)
	return e
}

// Poster is the attr_poster attribute. (Autogenerated)
func (e *Element) A_Poster(attr_poster ...string) *Element {
	e.attrs["poster"] = append(e.attrs["poster"], attr_poster...)
	return e
}

// Preload is the attr_preload attribute. (Autogenerated)
func (e *Element) A_Preload(attr_preload bool) *Element {
	if attr_preload {
		e.attrs["preload"] = append(e.attrs["preload"], "preload")
	} else {
		e.attrs["preload"] = append(e.attrs["preload"], "nopreload")
	}
	return e
}

// Readonly is the attr_readonly attribute. (Autogenerated)
func (e *Element) A_Readonly() *Element {
	e.boolattrs["readonly"] = true
	return e
}

// Rel is the attr_rel attribute. (Autogenerated)
func (e *Element) A_Rel(attr_rel string) *Element {
	e.attrs["rel"] = append(e.attrs["rel"], attr_rel)
	return e
}

// Required is the attr_required attribute. (Autogenerated)
func (e *Element) A_Required() *Element {
	e.boolattrs["required"] = true
	return e
}

// Reversed is the attr_reversed attribute. (Autogenerated)
func (e *Element) A_Reversed() *Element {
	e.boolattrs["reversed"] = true
	return e
}

// Rows is the attr_rows attribute. (Autogenerated)
func (e *Element) A_Rows(attr_rows string) *Element {
	e.attrs["rows"] = append(e.attrs["rows"], attr_rows)
	return e
}

// Rowspan is the attr_rowspan attribute. (Autogenerated)
func (e *Element) A_Rowspan(attr_rowspan string) *Element {
	e.attrs["rowspan"] = append(e.attrs["rowspan"], attr_rowspan)
	return e
}

// Sandbox is the attr_sandbox attribute. (Autogenerated)
func (e *Element) A_Sandbox(attr_sandbox ...string) *Element {
	e.attrs["sandbox"] = append(e.attrs["sandbox"], attr_sandbox...)
	return e
}

// Scope is the attr_scope attribute. (Autogenerated)
func (e *Element) A_Scope(attr_scope ...string) *Element {
	e.attrs["scope"] = append(e.attrs["scope"], attr_scope...)
	return e
}

// Selected is the attr_selected attribute. (Autogenerated)
func (e *Element) A_Selected() *Element {
	e.boolattrs["selected"] = true
	return e
}

// Shape is the attr_shape attribute. (Autogenerated)
func (e *Element) A_Shape(attr_shape ...string) *Element {
	e.attrs["shape"] = append(e.attrs["shape"], attr_shape...)
	return e
}

// Size is the attr_size attribute. (Autogenerated)
func (e *Element) A_Size(attr_size ...string) *Element {
	e.attrs["size"] = append(e.attrs["size"], attr_size...)
	return e
}

// Sizes is the attr_sizes attribute. (Autogenerated)
func (e *Element) A_Sizes(attr_sizes ...string) *Element {
	e.attrs["sizes"] = append(e.attrs["sizes"], attr_sizes...)
	return e
}

// Span is the attr_span attribute. (Autogenerated)
func (e *Element) A_Span(attr_span ...string) *Element {
	e.attrs["span"] = append(e.attrs["span"], attr_span...)
	return e
}

// Spellcheck is the attr_spellcheck attribute. (Autogenerated)
func (e *Element) A_Spellcheck(attr_spellcheck bool) *Element {
	e.attrs["spellcheck"] = append(e.attrs["spellcheck"], strconv.FormatBool(attr_spellcheck))
	return e
}

// Srcdoc is the attr_srcdoc attribute. (Autogenerated)
func (e *Element) A_Srcdoc(attr_srcdoc ...string) *Element {
	e.attrs["srcdoc"] = append(e.attrs["srcdoc"], attr_srcdoc...)
	return e
}

// Srclang is the attr_srclang attribute. (Autogenerated)
func (e *Element) A_Srclang(attr_srclang ...string) *Element {
	e.attrs["srclang"] = append(e.attrs["srclang"], attr_srclang...)
	return e
}

// Srcset is the attr_srcset attribute. (Autogenerated)
func (e *Element) A_Srcset(attr_srcset ...string) *Element {
	e.attrs["srcset"] = append(e.attrs["srcset"], attr_srcset...)
	return e
}

// Start is the attr_start attribute. (Autogenerated)
func (e *Element) A_Start(attr_start ...string) *Element {
	e.attrs["start"] = append(e.attrs["start"], attr_start...)
	return e
}

// Step is the attr_step attribute. (Autogenerated)
func (e *Element) A_Step(attr_step ...string) *Element {
	e.attrs["step"] = append(e.attrs["step"], attr_step...)
	return e
}

// Tabindex is the attr_tabindex attribute. (Autogenerated)
func (e *Element) A_Tabindex(attr_tabindex ...string) *Element {
	e.attrs["tabindex"] = append(e.attrs["tabindex"], attr_tabindex...)
	return e
}

// Target is the attr_target attribute. (Autogenerated)
func (e *Element) A_Target(attr_target ...string) *Element {
	e.attrs["target"] = append(e.attrs["target"], attr_target...)
	return e
}

// Title is the attr_title attribute. (Autogenerated)
func (e *Element) A_Title(attr_title string) *Element {
	e.attrs["title"] = []string{attr_title}
	return e
}

// Translate is the attr_translate attribute. (Autogenerated)
func (e *Element) A_Translate(attr_translate bool) *Element {
	if attr_translate {
		e.attrs["translate"] = append(e.attrs["translate"], "yes")
	} else {
		e.attrs["translate"] = append(e.attrs["translate"], "no")
	}
	return e
}

// Type is the attr_type attribute. (Autogenerated)
func (e *Element) A_Type(attr_type string) *Element {
	e.attrs["type"] = []string{attr_type}
	return e
}

// Usemap is the attr_usemap attribute. (Autogenerated)
func (e *Element) A_Usemap(attr_usemap ...string) *Element {
	e.attrs["usemap"] = append(e.attrs["usemap"], attr_usemap...)
	return e
}

// Value is the attr_value attribute. (Autogenerated)
func (e *Element) A_Value(attr_value string) *Element {
	e.attrs["value"] = []string{attr_value}
	return e
}

// Width is the attr_width attribute. (Autogenerated)
func (e *Element) Width(attr_width string) *Element {
	e.attrs["width"] = []string{attr_width}
	return e
}

// Wrap is the attr_wrap attribute. (Autogenerated)
func (e *Element) A_Wrap(attr_wrap ...string) *Element {
	e.attrs["wrap"] = append(e.attrs["wrap"], attr_wrap...)
	return e
}
