package writer

import (
	"io"
	"strings"
	"unicode/utf8"
)

type TabbedWriter struct {
	indent             string
	out                io.Writer
	bufferedEmptyLines int
	dirty              bool
	indentLevel        int
}

func (writer *TabbedWriter) failIfDirty() {
	if writer.dirty {
		panic("Invalid operation while dirty.")
	}
}

func (writer *TabbedWriter) Indent() {
	writer.failIfDirty()
	writer.indentLevel += 1
}

func (writer *TabbedWriter) Dedent() {
	writer.failIfDirty()
	writer.indentLevel -= 1
}

func (writer *TabbedWriter) rawWrite(text string) (n int, err error) {
	return writer.out.Write([]byte(text))
}

func (writer *TabbedWriter) beginningOfLine() {
	writer.failIfDirty()
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
	writer.failIfDirty()
	writer.WriteString(text)
	writer.EndOfLine()
}

func splitIndent(line string) (string, int) {
	spc := 0
	for len(line) > 0 {
		r, sz := utf8.DecodeRuneInString(line)
		if r == '\t' {
			spc += 4
		} else if r == ' ' {
			spc += 1
		} else {
			break
		}
		line = line[sz:]
	}
	return line, spc
}

func (writer *TabbedWriter) WriteChunk(text string) {
	writer.failIfDirty()
	lines := strings.Split(text, "\n")
	// Strip right spaces
	for i := 0; i < len(lines); i++ {
		lines[i] = strings.TrimRight(lines[i], "\t ")
	}
	// Discard trailing lines
	for len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	// Discard preceding lines
	for len(lines) > 0 && lines[0] == "" {
		lines = lines[1:]
	}
	if len(lines) == 0 {
		return
	}
	line, spc := splitIndent(lines[0])
	writer.WriteLine(line)
	indents := 0
	for _, line := range lines[1:] {
		line, newSpc := splitIndent(line)
		if line == "" {
			writer.EndOfLine()
			continue
		}
		// Change indentation level as needed.
		// TODO: require consistency.
		if newSpc > spc {
			writer.Indent()
			indents += 1
			spc = newSpc
		} else if newSpc < spc {
			writer.Dedent()
			indents -= 1
			if indents < 0 {
				panic("Bad indentation.")
			}
			spc = newSpc
		}
		writer.WriteLine(line)
	}
	if indents != 0 {
		panic("Bad indentation.")
	}
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
