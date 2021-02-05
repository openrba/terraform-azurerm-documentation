package main

import (
	"log"
	"path"
	"strings"

	"github.com/atsushinee/go-markdown-generator/doc"
	"github.com/google/go-github/v33/github"
)

// DocGenerator for generating markdown
type DocGenerator interface {
	WriteMenuItems(repos []*github.Repository)
	ExportReadme()
}

// docGenerator TODO
type docGenerator struct {
	*doc.MarkDownDoc
}

// NewDocGenerator Creates instance of docGenerator
func NewDocGenerator() DocGenerator {
	dg := docGenerator{doc.NewMarkDown()}
	dg.WriteLevel1Title("Terraform Azurerm Documentation")
	return &dg
}

func (g *docGenerator) WriteMenuItem(title string, path string) {
	g.Write("* ").WriteLink(title, path).WriteLines(1)
}

func (g *docGenerator) WriteMenuItems(repos []*github.Repository) {
	g.Write("Links: ").WriteLines(2)
	for _, repo := range repos {
		if strings.HasPrefix(*repo.Name, providerPrefix) {
			name := SanitizeName(*repo.Name)
			g.WriteMenuItem(strings.ReplaceAll(name, "-", " "), name)
		}
	}
}

func (g *docGenerator) ExportReadme() {
	err := g.Export(path.Join("docs", "HOME.md"))
	if err != nil {
		log.Fatal(err)
	}
}
