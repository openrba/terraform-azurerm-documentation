package main

import (
	"os"
)

var (
	org   = os.Getenv("ORG")
	token = os.Getenv("GH_TOKEN")
)

func main() {
	c := NewClient(token)

	repos := c.GetRepos(org)

	c.DownloadDocs(repos)

	d := NewDocGenerator()
	d.WriteMenuItems(repos)
	d.ExportReadme()
}
