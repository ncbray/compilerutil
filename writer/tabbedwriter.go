package writer

import (
	"io"
)

type TabbedWriter struct {
	indent             string
	out                io.Writer
	bufferedEmptyLines int
	dirty              bool
	indentLevel        int
}

func (writer *TabbedWriter) Indent() {
	if writer.dirty {
		panic("Invalid operation while dirty.")
	}
	writer.indentLevel += 1
}

func (writer *TabbedWriter) Dedent() {
	if writer.dirty {
		panic("Invalid operation while dirty.")
	}
	writer.indentLevel -= 1
}

func (writer *TabbedWriter) rawWrite(text string) (n int, err error) {
	return writer.out.Write([]byte(text))
}

func (writer *TabbedWriter) beginningOfLine() {
	if writer.dirty {
		panic("Invalid operation while dirty.")
	}
	for i := 0; i < writer.bufferedEmptyLines; i++ {
		writer.rawWrite("\n")
	}
	writer.bufferedEmptyLines = 0
	for i := 0; i < writer.indentLevel; i++ {
		writer.rawWrite(writer.indent)
	}
	writer.dirty = true
}

func (writer *TabbedWriter) Write(p []byte) (n int, err error) {
	if !writer.dirty {
		writer.beginningOfLine()
	}
	return writer.out.Write(p)
}

func (writer *TabbedWriter) WriteString(text string) (n int, err error) {
	if !writer.dirty {
		writer.beginningOfLine()
	}
	return writer.rawWrite(text)
}

func (writer *TabbedWriter) WriteLine(text string) {
	if writer.dirty {
		panic("Invalid operation while dirty.")
	}
	writer.WriteString(text)
	writer.EndOfLine()
}

func (writer *TabbedWriter) EndOfLine() {
	if writer.dirty {
		writer.rawWrite("\n")
		writer.dirty = false
	} else {
		writer.bufferedEmptyLines += 1
	}
}

func MakeTabbedWriter(indent string, out io.Writer) *TabbedWriter {
	return &TabbedWriter{
		indent: indent,
		out:    out,
	}
}
