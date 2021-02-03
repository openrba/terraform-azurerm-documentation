package main

import (
	"log"

	"github.com/atsushinee/go-markdown-generator/doc"
)

// Tool for generating markdown
type DocGenerator interface {
}

// docGenerator TODO
type docGenerator struct {
	*doc.MarkDownDoc
}

// NewDocGenerator: Creates instance of docGenerator
func NewDocGenerator() DocGenerator {
	return &docGenerator{doc.NewMarkDown()}
}

func (g *docGenerator) AppendMenuItem(title string, path string) {
	g.WriteLinkLine(title, path)
}

func (g *docGenerator) WriteIndex() {

	g.WriteTitle("Summary", doc.LevelTitle).
		WriteLines(2)

	g.WriteTitle("Author", doc.LevelNormal).
		WriteCodeLine("lichun")

	g.WriteTitle("Website", doc.LevelNormal)
	g.WriteLinkLine("lichunorz", "https://lichunorz.com")

	t := doc.NewTable(4, 4)
	t.SetTitle(0, "Version")
	t.SetTitle(1, "Date")
	t.SetTitle(2, "Creator")
	t.SetTitle(3, "Remarks")
	t.SetContent(0, 0, "v1")
	t.SetContent(0, 1, "2019-11-21")
	t.SetContent(0, 2, "Lee")
	t.SetContent(0, 3, "无")
	g.WriteTable(t)
	err := g.Export("README.md")
	if err != nil {
		log.Fatal(err)
	}
}
