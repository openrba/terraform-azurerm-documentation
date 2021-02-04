package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"

	"github.com/google/go-github/v33/github"
	"golang.org/x/oauth2"
)

// GithubClient for accessing the github Api
type GithubClient interface {
	GetRepos(o string) []*github.Repository
	DownloadDoc(repo *github.Repository)
	DownloadDocs(repos []*github.Repository)
}

// githubClient TODO
type githubClient struct {
	*github.Client                 // Client base
	ctx            context.Context // Context
}

func validate(t string) {
	if len(t) < 20 {
		log.Fatal("Invalid token. Must be 20 characters")
	}
}

// NewClient Creates new instance of GithubClient
func NewClient(t string) GithubClient {
	log.Print("Connecting to Github")
	validate(t)
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: t},
	)
	tc := oauth2.NewClient(ctx, ts)

	return &githubClient{Client: github.NewClient(tc), ctx: context.Background()}
}

// GetRepos TODO
func (client *githubClient) GetRepos(o string) []*github.Repository {
	log.Printf("Getting repositories for org %s", o)
	repos, _, err := client.Repositories.ListByOrg(client.ctx, o, nil)

	if err != nil {
		log.Fatal(err)
	}

	return repos
}

// DownloadDoc TODO
func (client *githubClient) DownloadDoc(repo *github.Repository) {
	rc, resp, err := client.Repositories.DownloadContents(client.ctx, org, *repo.Name, "README.md", nil)

	if err != nil {
		log.Print(resp)
		return
	}

	// Create readme in docs/{repo}.md
	file := path.Join("docs", fmt.Sprintf("%s.md", *repo.Name))
	// Open a string writer for the file
	w, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Copy reader
	defer w.Close()
	_, err = io.Copy(w, rc)
	if err != nil {
		log.Fatal(err)
		return
	}
}

// DownloadDocs TODO
func (client *githubClient) DownloadDocs(repos []*github.Repository) {
	for i := 0; i < len(repos); i++ {
		repo := repos[i]

		if strings.HasPrefix(*repo.Name, "terraform-azurerm") {
			log.Printf("Downloading README from %s", *repo.Name)
			client.DownloadDoc(repo)
		}
	}
}
