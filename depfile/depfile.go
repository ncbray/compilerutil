package depfile

import (
	"bytes"
	"io"
	"regexp"
)

type DepfileBuilder struct {
	output   string
	inputs   []string
	inputLut map[string]bool
}

func CreateDepfileBuilder(output string) *DepfileBuilder {
	return &DepfileBuilder{
		output:   output,
		inputs:   []string{},
		inputLut: map[string]bool{},
	}
}

func (b *DepfileBuilder) Add(input string) {
	_, exists := b.inputLut[input]
	if exists {
		return
	}
	b.inputs = append(b.inputs, input)
	b.inputLut[input] = true
}

var specialPathChars = regexp.MustCompile("([ \\\\])")

func escapeFileName(path string) string {
	return specialPathChars.ReplaceAllString(path, "\\$1")
}

func (b *DepfileBuilder) Write(out io.Writer) {
	out.Write([]byte(escapeFileName(b.output)))
	out.Write([]byte(":"))
	for _, input := range b.inputs {
		out.Write([]byte(" "))
		out.Write([]byte(escapeFileName(input)))
	}
	out.Write([]byte("\n"))
}

func (b *DepfileBuilder) String() string {
	var buffer bytes.Buffer
	b.Write(&buffer)
	return buffer.String()
}
