package writer

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestEmptyChunk(t *testing.T) {
	var b bytes.Buffer
	w := MakeTabbedWriter("\t", &b)
	w.Indent()
	w.WriteChunk("")
	w.Dedent()
	assert.Equal(t, "", b.String())
}

func TestSimpleChunk(t *testing.T) {
	var b bytes.Buffer
	w := MakeTabbedWriter("\t", &b)
	w.Indent()
	w.WriteChunk("foo")
	w.Dedent()
	assert.Equal(t, "\tfoo\n", b.String())
}

func TestTrimChunk(t *testing.T) {
	var b bytes.Buffer
	w := MakeTabbedWriter("\t", &b)
	w.Indent()
	w.WriteChunk("\n\nfoo\n\n")
	w.Dedent()
	assert.Equal(t, "\tfoo\n", b.String())
}

func TestRetabChunk(t *testing.T) {
	var b bytes.Buffer
	w := MakeTabbedWriter("\t", &b)
	w.Indent()
	w.WriteChunk("  foo")
	w.Dedent()
	assert.Equal(t, "\tfoo\n", b.String())
}

func TestTwoChunk(t *testing.T) {
	var b bytes.Buffer
	w := MakeTabbedWriter("\t", &b)
	w.Indent()
	w.WriteChunk("foo\nbar")
	w.Dedent()
	assert.Equal(t, "\tfoo\n\tbar\n", b.String())
}

func TestSplitChunk(t *testing.T) {
	var b bytes.Buffer
	w := MakeTabbedWriter("\t", &b)
	w.Indent()
	w.WriteChunk("  foo\n\n  bar")
	w.Dedent()
	assert.Equal(t, "\tfoo\n\n\tbar\n", b.String())
}

func TestIndentDedentChunk(t *testing.T) {
	var b bytes.Buffer
	w := MakeTabbedWriter("\t", &b)
	w.Indent()
	w.WriteChunk("  foo\n    bar\n  baz")
	w.Dedent()
	assert.Equal(t, "\tfoo\n\t\tbar\n\tbaz\n", b.String())
}

func TestIndentDedentTabsChunk(t *testing.T) {
	var b bytes.Buffer
	w := MakeTabbedWriter("\t", &b)
	w.Indent()
	w.WriteChunk("\tfoo\n\t\tbar\n\tbaz")
	w.Dedent()
	assert.Equal(t, "\tfoo\n\t\tbar\n\tbaz\n", b.String())
}

func TestIndentIndentChunk(t *testing.T) {
	var b bytes.Buffer
	w := MakeTabbedWriter("\t", &b)
	w.Indent()
	w.WriteChunk("  foo\n    bar\n  baz")
	w.Dedent()
	assert.Equal(t, "\tfoo\n\t\tbar\n\tbaz\n", b.String())
}

/*
func TestIndentResetChunk(t *testing.T) {
	var b bytes.Buffer
	w := MakeTabbedWriter("\t", &b)
	w.Indent()
	w.WriteChunk("  foo\n    bar")
	w.WriteLine("baz")
	w.Dedent()
	assert.Equal(t, "\tfoo\n\t\tbar\n\tbaz\n", b.String())
}
*/
