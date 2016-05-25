package writer

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStandardDeclaration(t *testing.T) {
	var b bytes.Buffer
	w := MakeXMLWriter("  ", &b)
	w.StandardDeclaration()
	assert.Equal(t, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n", b.String())
}

func TestElement(t *testing.T) {
	var b bytes.Buffer
	w := MakeXMLWriter("  ", &b)
	w.Element("a", []Attr{})
	assert.Equal(t, "<a/>\n", b.String())
}

func TestElementAttr(t *testing.T) {
	var b bytes.Buffer
	w := MakeXMLWriter("  ", &b)
	w.Element("ab", []Attr{{"cd", "ef"}, {"gh", "ij"}})
	assert.Equal(t, "<ab cd=\"ef\" gh=\"ij\"/>\n", b.String())
}

func TestBeginEndInline(t *testing.T) {
	var b bytes.Buffer
	w := MakeXMLWriter("  ", &b)
	w.Begin("a", []Attr{{"b", "c"}}, false)
	w.End()
	assert.Equal(t, "<a b=\"c\"></a>", b.String())
}

func TestBeginEndBlock(t *testing.T) {
	var b bytes.Buffer
	w := MakeXMLWriter("  ", &b)
	w.Begin("a", []Attr{{"b", "c"}}, true)
	w.End()
	assert.Equal(t, "<a b=\"c\">\n</a>\n", b.String())
}

func TestChildrenInline(t *testing.T) {
	var b bytes.Buffer
	w := MakeXMLWriter("  ", &b)
	w.Begin("a", []Attr{}, false)
	w.Begin("b", []Attr{}, false)
	w.Begin("c", []Attr{}, false)
	w.End()
	w.End()
	w.End()
	w.EndOfLine()
	assert.Equal(t, "<a><b><c></c></b></a>\n", b.String())
}

func TestChildrenBlock(t *testing.T) {
	var b bytes.Buffer
	w := MakeXMLWriter("  ", &b)
	w.Begin("a", []Attr{}, true)
	w.Begin("b", []Attr{}, true)
	w.Begin("c", []Attr{}, true)
	w.End()
	w.End()
	w.End()
	assert.Equal(t, "<a>\n  <b>\n    <c>\n    </c>\n  </b>\n</a>\n", b.String())
}
