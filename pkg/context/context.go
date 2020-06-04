package context

import (
	ctx "context"
)

type Context struct {
	ctx.Context
}

func New() *Context {
	return Wrap(ctx.Background())
}

func Wrap(ctx ctx.Context) *Context {
	return &Context{
		Context: ctx,
	}
}
