package client

import (
	"github.com/google/go-github/github"
	"github.com/wesleimp/github-terraform/pkg/context"
	"golang.org/x/oauth2"
)

// New creates a new github client
func New(ctx *context.Context) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ctx.Token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	return client
}
