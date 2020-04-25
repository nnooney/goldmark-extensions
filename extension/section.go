// Package extension is an extension for goldmark (https://github.com/yuin/goldmark).
package extension

import (
	"github.com/nnooney/goldmark-extensions/extension/ast"
	"github.com/yuin/goldmark"
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

// Define an AST Transformer to convert a document into sections.
type sectionASTTransformer struct {
}

var defaultSectionASTTransformer = &sectionASTTransformer{}

// NewSectionASTTransformer returns a new parser.ASTTransformer that groups
// heading and contents into sections.
func NewSectionASTTransformer() parser.ASTTransformer {
	return defaultSectionASTTransformer
}

func (a *sectionASTTransformer) Transform(
	node *gast.Document, reader text.Reader, pc parser.Context) {
	// Iterate through the direct children of the document and insert section
	// nodes each time a heading is encountered.
	sectionNode := ast.NewSection()
	node.AppendChild(node, sectionNode)
	for inode := node.FirstChild(); inode != nil; {
		next := inode.NextSibling()
		if inode.Kind().String() == "Section" {
			// At this point we're iterating over nodes we already inserted, so
			// break out of the loop.
			break
		} else if inode.Kind().String() == "Heading" && inode.(*gast.Heading).Level < 3 {
			// If the node is a heading, close/start a new list.
			if sectionNode.ChildCount() > 0 {
				sectionNode = ast.NewSection()
				node.AppendChild(node, sectionNode)
			}
		}
		// Add the node to the temporary list.
		node.RemoveChild(node, inode)
		sectionNode.AppendChild(sectionNode, inode)
		inode = next
	}
}

// sectionHTMLRenderer is a renderer.NodeRenderer implementation that renders
// Section nodes.
type sectionHTMLRenderer struct {
	html.Config
}

// NewSectionHTMLRenderer returns a new sectionHTMLRenderer
func NewSectionHTMLRenderer(opts ...html.Option) renderer.NodeRenderer {
	r := &sectionHTMLRenderer{
		Config: html.NewConfig(),
	}
	for _, opt := range opts {
		opt.SetHTMLOption(&r.Config)
	}
	return r
}

// RegisterFuncs implements renderer.NodeRenderer.RegisterFuncs.
func (r *sectionHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindSection, r.renderSection)
}

func (r *sectionHTMLRenderer) renderSection(w util.BufWriter, source []byte, node gast.Node, entering bool) (gast.WalkStatus, error) {
	if entering {
		_, _ = w.WriteString("<section>\n")
	} else {
		_, _ = w.WriteString("</section>\n")
	}
	return gast.WalkContinue, nil
}

type section struct {
}

// Section is an extension that groups headings and paragraphs into sections.
var Section = &section{}

func (e *section) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithASTTransformers(
			util.Prioritized(NewSectionASTTransformer(), 999),
		),
	)
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(NewSectionHTMLRenderer(), 500),
	))
}
