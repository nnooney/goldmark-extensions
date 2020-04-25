package ast

import (
	"bufio"
	"io"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

// A Renderer struct is an implementation of renderer.Renderer that renders
// nodes as raw text.
type Renderer struct{}

// NewRenderer returns a new Renderer with the given options.
func NewRenderer() renderer.Renderer {
	r := &Renderer{}

	return r
}

var indentLevel int = 0

// Render implements Renderer.Render.
func (r *Renderer) Render(w io.Writer, source []byte, n ast.Node) error {
	writer, ok := w.(util.BufWriter)
	if !ok {
		writer = bufio.NewWriter(w)
	}

	err := ast.Walk(n, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering {
			node.Dump(source, indentLevel)
			indentLevel++
		} else {
			indentLevel--
		}
		result := ast.WalkContinue
		if node.Type() == ast.TypeBlock || node.Kind().String() == "Document" {
			result = ast.WalkSkipChildren
		}
		return result, nil
	})
	if err != nil {
		return err
	}
	return writer.Flush()
}

// AddOptions implements Renderer.AddOptions.
func (r *Renderer) AddOptions(opts ...renderer.Option) {}
