package gen

import (
	"github.com/Nigel2392/gen/elems"
)

func (h *HTML) render(buffer elems.Buffer) elems.Buffer {
	// Creates duplicates upon multiple calls to render
	buffer.WriteString("<!DOCTYPE html>\r\n<html>\r\n")
	h.Head.Generate("    ", buffer)
	h.Body.Generate("    ", buffer)
	buffer.WriteString("</html>\r\n")
	return buffer
}

func (h *HTML) RenderBuffer(buffer elems.Buffer, templateData any) elems.Buffer {
	buffer = h.render(buffer)
	return elems.ParseToTemplate(buffer, templateData)
}
