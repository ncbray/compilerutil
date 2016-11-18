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
		{Name: "width", Value: "500"},
		{Name: "height", Value: "500"},
		{Name: "xmlns", Value: "http://www.w3.org/2000/svg"},
		{Name: "xmlns:xlink", Value: "http://www.w3.org/1999/xlink"},
	}, true)

	xout.Begin("title", nil, false)
	xout.WriteString("CFG")
	xout.End()
	xout.EndOfLine()

	xout.Begin("style", []writer.Attr{{Name: "type", Value: "text/css"}}, false)
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
		{Name: "id", Value: "arrow"},
		{Name: "markerWidth", Value: "10"},
		{Name: "markerHeight", Value: "10"},
		{Name: "refX", Value: "8"},
		{Name: "refY", Value: "3"},
		{Name: "orient", Value: "auto"},
		{Name: "markerUnits", Value: "strokeWidth"},
	}, true)
	xout.Element("path", []writer.Attr{
		{Name: "d", Value: "M0,0 L0,6 L9,3 z"},
		{Name: "fill", Value: "#ff0000"},
	})
	xout.End()

	xout.End()

	font_size := 14
	line_spacing := font_size + 2

	for _, node := range nodes {
		xout.Element("rect", []writer.Attr{
			{Name: "x", Value: fmt.Sprintf("%d", node.X)},
			{Name: "y", Value: fmt.Sprintf("%d", node.Y)},
			{Name: "width", Value: fmt.Sprintf("%d", node.W)},
			{Name: "height", Value: fmt.Sprintf("%d", node.H)},
			{Name: "class", Value: "basic_block"},
		})
		x := node.X + node.W/2
		y := node.Y + line_spacing
		lines := strings.Split(node.Label, "\n")
		for _, line := range lines {
			xout.Begin("text", []writer.Attr{
				{Name: "x", Value: fmt.Sprintf("%d", x)},
				{Name: "y", Value: fmt.Sprintf("%d", y)},
				{Name: "class", Value: "basic_text"},
			}, false)
			xout.WriteString(line)
			xout.End()
			xout.EndOfLine()
			y += line_spacing
		}
	}

	xout.Element("polyline", []writer.Attr{
		{Name: "points", Value: "60,60 200,60"},
		{Name: "class", Value: "basic_arrow"},
	})
	xout.Element("polyline", []writer.Attr{
		{Name: "points", Value: "60,60 200,200"},
		{Name: "class", Value: "basic_arrow"},
	})
	xout.Element("polyline", []writer.Attr{
		{Name: "points", Value: "60,60 60,200"},
		{Name: "class", Value: "basic_arrow"},
	})

	xout.End()
}
