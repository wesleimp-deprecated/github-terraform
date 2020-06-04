package repository

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/wesleimp/github-terraform/pkg/context"
)

// Cmd config
type Cmd struct {
	Cmd     *cobra.Command
	options options
}

type options struct {
	name     string
	org      string
	user     string
	dest     string
	repoType string
	token    string
}

// NewCmd creates a repository cmd
func NewCmd() *Cmd {
	root := &Cmd{}

	var commands = &cobra.Command{
		Use:           "repositories",
		Aliases:       []string{"repos"},
		Short:         "Import repositories",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE:          run,
	}

	commands.Flags().StringVarP(&root.options.name, "name", "n", "", `Repository name. The name must contains owner/repo`)
	commands.Flags().StringVarP(&root.options.org, "org", "o", "", "Repository organization")
	commands.Flags().StringVarP(&root.options.user, "user", "u", "", "Repository user")
	commands.Flags().StringVarP(&root.options.dest, "dest", "d", "./output", "Path that will contains the output files")
	commands.Flags().StringVarP(&root.options.repoType, "type", "t", "", "Repository type. Could be public or private")
	commands.Flags().StringVar(&root.options.token, "token", "", "Github token. This property is not necessary if you already exported GITHUB_TOKEN")

	root.Cmd = commands
	return root
}

func run(cmd *cobra.Command, args []string) error {
	return nil
}

func setupContext(ctx *context.Context, o options) *context.Context {
	ctx.Config.Repository.Dest = o.dest
	ctx.Config.Repository.Name = o.name
	ctx.Config.Repository.User = o.user
	ctx.Config.Repository.Org = o.org
	ctx.Config.Repository.Type = o.repoType

	if o.token == "" {
		ctx.Token = os.Getenv("GITHUB_TOKEN")
	} else {
		ctx.Token = o.token
	}

	return ctx
}
