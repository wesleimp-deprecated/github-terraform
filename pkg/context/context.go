package context

import (
	ctx "context"

	"github.com/google/go-github/github"
	"github.com/wesleimp/github-terraform/pkg/config"
)

type Context struct {
	ctx.Context
	Token  string
	Config *config.Config
	Client *github.Client
}

// New creates new context
func New() *Context {
	return Wrap(ctx.Background())
}

// Wrap context
func Wrap(ctx ctx.Context) *Context {
	return &Context{
		Context: ctx,
	}
}
