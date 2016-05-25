package viz

import (
	"fmt"
	"github.com/ncbray/compilerutil/writer"
	"io"
	"strings"
)

type Node struct {
	X     int
	Y     int
	W     int
	H     int
	Label string
}

type Attr struct {
	Name  string
	Value string
}

func WriteSVG(out io.Writer) {

	nodes := []*Node{
		{X: 0, Y: 0, W: 60, H: 60, Label: "Hello1\n!"},
		{X: 200, Y: 0, W: 60, H: 60, Label: "Hello2\n!"},

		{X: 0, Y: 200, W: 60, H: 60, Label: "Hello3\n!"},
		{X: 200, Y: 200, W: 60, H: 60, Label: "Hello4\n!"},
	}

	xout := writer.MakeXMLWriter("  ", out)

	xout.StandardDeclaration()
	xout.Begin("svg", []writer.Attr{
		{"width", "500"},
		{"height", "500"},
		{"xmlns", "http://www.w3.org/2000/svg"},
		{"xmlns:xlink", "http://www.w3.org/1999/xlink"},
	}, true)

	xout.Begin("title", nil, false)
	xout.WriteString("CFG")
	xout.End()
	xout.EndOfLine()

	xout.Begin("style", []writer.Attr{{"type", "text/css"}}, false)
	style := xout.BeginTabbedCData()
	style.WriteLine(".basic_block {")
	style.Indent()
	style.WriteLine("fill:orange;")
	style.WriteLine("stroke:gray;")
	style.WriteLine("stroke-width:2px;")
	style.Dedent()
	style.WriteLine("}")
	style.WriteLine(".basic_text {")
	style.Indent()
	style.WriteLine("font-family:sans-serif;")
	style.WriteLine("font-size:14px;")
	style.WriteLine("text-anchor:middle;")
	style.Dedent()
	style.WriteLine("}")
	style.WriteLine(".basic_arrow {")
	style.Indent()
	style.WriteLine("fill:none;")
	style.WriteLine("stroke:red;")
	style.WriteLine("stroke-width:1px;")
	style.WriteLine("marker-end:url(#arrow);")
	style.Dedent()
	style.WriteLine("}")

	xout.EndTabbedCData()
	xout.End()
	xout.EndOfLine()

	xout.Begin("defs", []writer.Attr{}, true)

	xout.Begin("marker", []writer.Attr{
		{"id", "arrow"},
		{"markerWidth", "10"},
		{"markerHeight", "10"},
		{"refX", "8"},
		{"refY", "3"},
		{"orient", "auto"},
		{"markerUnits", "strokeWidth"},
	}, true)
	xout.Element("path", []writer.Attr{
		{"d", "M0,0 L0,6 L9,3 z"},
		{"fill", "#ff0000"},
	})
	xout.End()

	xout.End()

	font_size := 14
	line_spacing := font_size + 2

	for _, node := range nodes {
		xout.Element("rect", []writer.Attr{
			{"x", fmt.Sprintf("%d", node.X)},
			{"y", fmt.Sprintf("%d", node.Y)},
			{"width", fmt.Sprintf("%d", node.W)},
			{"height", fmt.Sprintf("%d", node.H)},
			{"class", "basic_block"},
		})
		x := node.X + node.W/2
		y := node.Y + line_spacing
		lines := strings.Split(node.Label, "\n")
		for _, line := range lines {
			xout.Begin("text", []writer.Attr{
				{"x", fmt.Sprintf("%d", x)},
				{"y", fmt.Sprintf("%d", y)},
				{"class", "basic_text"},
			}, false)
			xout.WriteString(line)
			xout.End()
			xout.EndOfLine()
			y += line_spacing
		}
	}

	xout.Element("polyline", []writer.Attr{
		{"points", "60,60 200,60"},
		{"class", "basic_arrow"},
	})
	xout.Element("polyline", []writer.Attr{
		{"points", "60,60 200,200"},
		{"class", "basic_arrow"},
	})
	xout.Element("polyline", []writer.Attr{
		{"points", "60,60 60,200"},
		{"class", "basic_arrow"},
	})

	xout.End()
}
