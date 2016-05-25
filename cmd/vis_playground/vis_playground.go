package main

import (
	"github.com/ncbray/compilerutil/viz"
	"os"
)

func main() {
	viz.WriteSVG(os.Stdout)
}
