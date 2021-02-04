package main

import (
	"log"

	"github.com/atsushinee/go-markdown-generator/doc"
)

// DocGenerator for generating markdown
type DocGenerator interface {
	WriteMenuItem(title string, path string)
	ExportReadme()
}

// docGenerator TODO
type docGenerator struct {
	*doc.MarkDownDoc
}

// NewDocGenerator Creates instance of docGenerator
func NewDocGenerator() DocGenerator {
	return &docGenerator{doc.NewMarkDown()}
}

func (g *docGenerator) WriteMenuItem(title string, path string) {
	g.Write("* ").WriteLink(title, path).WriteLines(1)
}

func (g *docGenerator) ExportReadme() {
	err := g.Export("README.md")
	if err != nil {
		log.Fatal(err)
	}
}
