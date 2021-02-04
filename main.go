package main

import (
	log "log"
	"os"

	"github.com/atsushinee/go-markdown-generator/doc"
)

var (
	org   = os.Getenv("ORG")
	token = os.Getenv("GITHUB_TOKEN")
)

func md() {

	book := doc.NewMarkDown()
	book.WriteTitle("Summary", doc.LevelTitle).WriteLines(2)

	// book.WriteMultiCode(string(code), "go")

	err := book.Export("SUMMARY.md")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	c := NewClient(token)

	repos := c.GetRepos(org)

	c.DownloadDocs(repos)

}
