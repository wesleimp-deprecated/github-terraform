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
func New(conf *config.Config) *Context {
	return Wrap(ctx.Background(), conf)
}

// Wrap context
func Wrap(ctx ctx.Context, conf *config.Config) *Context {
	return &Context{
		Context: ctx,
		Config:  conf,
	}
}
