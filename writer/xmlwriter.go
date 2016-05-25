package writer

import (
	"encoding/xml"
	"io"
)

type Attr struct {
	Name  string
	Value string
}

type scopeInfo struct {
	tag      string
	indented bool
}

type XMLWriter struct {
	tagStack []scopeInfo
	out      *TabbedWriter
}

func (w *XMLWriter) StandardDeclaration() {
	w.out.WriteLine("<?xml version=\"1.0\" encoding=\"UTF-8\"?>")
}

func (w *XMLWriter) Element(name string, attrs []Attr) {
	w.out.WriteString("<")
	w.out.WriteString(name)

	for _, attr := range attrs {
		w.out.WriteString(" ")
		w.out.WriteString(attr.Name)
		w.out.WriteString("=\"")
		xml.EscapeText(w.out.out, []byte(attr.Value))
		w.out.WriteString("\"")
	}
	w.out.WriteString("/>")
	w.out.EndOfLine()
}

func (w *XMLWriter) Begin(name string, attrs []Attr, block bool) {
	w.out.WriteString("<")
	w.out.WriteString(name)

	for _, attr := range attrs {
		w.out.WriteString(" ")
		w.out.WriteString(attr.Name)
		w.out.WriteString("=\"")
		xml.EscapeText(w.out, []byte(attr.Value))
		w.out.WriteString("\"")
	}
	w.out.WriteString(">")
	if block {
		w.out.EndOfLine()
		w.out.Indent()
	}
	w.tagStack = append(w.tagStack, scopeInfo{name, block})
}

func (w *XMLWriter) End() {
	info := w.tagStack[len(w.tagStack)-1]
	w.tagStack = w.tagStack[:len(w.tagStack)-1]
	if info.indented {
		w.out.Dedent()
	}
	w.out.WriteString("</")
	w.out.WriteString(info.tag)
	w.out.WriteString(">")
	if info.indented {
		w.out.EndOfLine()
	}
}

func (w *XMLWriter) BeginTabbedCData() *TabbedWriter {
	w.out.WriteString("<![CDATA[")
	w.out.EndOfLine()
	w.out.Indent()
	return w.out
}

func (w *XMLWriter) EndTabbedCData() {
	w.out.Dedent()
	w.out.WriteString("]]>")
}

func (w *XMLWriter) Write(p []byte) {
	xml.EscapeText(w.out, p)
}

func (w *XMLWriter) WriteString(text string) {
	xml.EscapeText(w.out, []byte(text))
}

func (w *XMLWriter) WriteLine(text string) {
	xml.EscapeText(w.out, []byte(text))
	w.out.EndOfLine()
}

func (w *XMLWriter) EndOfLine() {
	w.out.EndOfLine()
}

func MakeXMLWriter(indent string, out io.Writer) *XMLWriter {
	return &XMLWriter{
		out: MakeTabbedWriter(indent, out),
	}
}
