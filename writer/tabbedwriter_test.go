package writer

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlatEmptyLines(t *testing.T) {
	var b bytes.Buffer
	w := MakeTabbedWriter("  ", &b)
	w.EndOfLine()
	w.EndOfLine()
	w.EndOfLine()
	assert.Equal(t, "", b.String())
}

func TestFlatFullLines(t *testing.T) {
	var b bytes.Buffer
	w := MakeTabbedWriter("  ", &b)
	w.WriteLine("ab")
	w.WriteLine("cde")
	w.WriteLine("f")
	assert.Equal(t, "ab\ncde\nf\n", b.String())
}

func TestFlatGap1(t *testing.T) {
	var b bytes.Buffer
	w := MakeTabbedWriter("  ", &b)
	w.WriteLine("ab")
	w.EndOfLine()
	w.WriteLine("f")
	assert.Equal(t, "ab\n\nf\n", b.String())
}

func TestFlatGap2(t *testing.T) {
	var b bytes.Buffer
	w := MakeTabbedWriter("  ", &b)
	w.WriteLine("ab")
	w.EndOfLine()
	w.EndOfLine()
	w.WriteLine("f")
	assert.Equal(t, "ab\n\n\nf\n", b.String())
}

func TestSpaceIdent1(t *testing.T) {
	var b bytes.Buffer
	w := MakeTabbedWriter("  ", &b)
	w.Indent()
	w.WriteLine("x")
	w.Dedent()
	assert.Equal(t, "  x\n", b.String())
}

func TestSpaceIdent2(t *testing.T) {
	var b bytes.Buffer
	w := MakeTabbedWriter("  ", &b)
	w.Indent()
	w.Indent()
	w.WriteLine("x")
	w.Dedent()
	w.Dedent()
	assert.Equal(t, "    x\n", b.String())
}

func TestTabIdent1(t *testing.T) {
	var b bytes.Buffer
	w := MakeTabbedWriter("\t", &b)
	w.Indent()
	w.WriteLine("x")
	w.Dedent()
	assert.Equal(t, "\tx\n", b.String())
}

func TestTabIdent2(t *testing.T) {
	var b bytes.Buffer
	w := MakeTabbedWriter("\t", &b)
	w.Indent()
	w.Indent()
	w.WriteLine("x")
	w.Dedent()
	w.Dedent()
	assert.Equal(t, "\t\tx\n", b.String())
}
