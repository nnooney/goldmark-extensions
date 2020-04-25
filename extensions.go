package extensions

import (
	"github.com/nnooney/goldmark-extensions/extension"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type tufte struct {
}

// Tufte is an extension that converts markdown output to render with Tufte CSS.
var Tufte = &tufte{}

func (e *tufte) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithASTTransformers(
			util.Prioritized(extension.NewSectionASTTransformer(), 999),
		),
	)
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(extension.NewSectionHTMLRenderer(), 500),
	))
}
