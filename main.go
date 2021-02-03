package main

import (
	log "log"
	"os"
	"strings"

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

	for i := 0; i < len(repos); i++ {
		repo := repos[i]

		if strings.HasPrefix(*repo.Name, "terraform-azurerm") {
			log.Printf("Downloading README from %s", *repo.Name)
			c.DownloadDoc(repo)
		}
	}

}
