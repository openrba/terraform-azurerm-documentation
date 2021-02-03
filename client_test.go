package main

import (
	"reflect"
	"testing"

	"github.com/google/go-github/v33/github"
)

func Test_validate(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validate(tt.args.t)
		})
	}
}

func TestNewClient(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want GithubClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_githubClient_GetRepos(t *testing.T) {
	type args struct {
		o string
	}
	tests := []struct {
		name   string
		client *githubClient
		args   args
		want   []*github.Repository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.client.GetRepos(tt.args.o); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("githubClient.GetRepos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_githubClient_DownloadDoc(t *testing.T) {
	type args struct {
		repo *github.Repository
	}
	tests := []struct {
		name   string
		client *githubClient
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.client.DownloadDoc(tt.args.repo)
		})
	}
}
