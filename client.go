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

const providerPrefix = "terraform-azurerm"

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
	opts := github.RepositoryListByOrgOptions{
		Type: "all",
		ListOptions: github.ListOptions{
			Page:    0,
			PerPage: 1000,
		},
	}
	repos, _, err := client.Repositories.ListByOrg(client.ctx, o, &opts)

	if err != nil {
		log.Fatal(err)
	}

	return repos
}

// SanitizeName formats the repo text to a pretty title name
func SanitizeName(name string) string {
	sb := strings.Builder{}

	ss := strings.Split(strings.TrimPrefix(name, providerPrefix+"-"), "-")
	for i, s := range ss {
		sb.WriteString(strings.Title(strings.ToLower(s)))
		if i < len(ss)-1 {
			sb.WriteString("-")
		}
	}
	return sb.String()
}

// DownloadDoc downloads the README.md for a given repository
func (client *githubClient) DownloadDoc(repo *github.Repository) {
	rc, resp, err := client.Repositories.DownloadContents(client.ctx, org, *repo.Name, "README.md", nil)

	if err != nil {
		log.Print(resp)
		return
	}

	// Create readme in docs/{repo}.md

	file := path.Join("docs", fmt.Sprintf("%s.md", SanitizeName(*repo.Name)))
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

// DownloadDocs downloads all README.md for given set of repositories
func (client *githubClient) DownloadDocs(repos []*github.Repository) {
	for i := 0; i < len(repos); i++ {
		repo := repos[i]

		if strings.HasPrefix(*repo.Name, providerPrefix) {
			log.Printf("Downloading README from %s", *repo.Name)
			client.DownloadDoc(repo)
		}
	}
}
