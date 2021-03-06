package ast

import (
	gast "github.com/yuin/goldmark/ast"
)

// A Section represents a section of content.
type Section struct {
	gast.BaseBlock
}

// Dump implements Node.Dump.
func (n *Section) Dump(source []byte, level int) {
	m := map[string]string{}
	gast.DumpHelper(n, source, level, m, nil)
}

// KindSection is the NodeKind of the Section Node.
var KindSection = gast.NewNodeKind("Section")

// Kind implements Node.Kind.
func (n *Section) Kind() gast.NodeKind {
	return KindSection
}

// NewSection returns a new Section node.
func NewSection() *Section {
	return &Section{}
}
