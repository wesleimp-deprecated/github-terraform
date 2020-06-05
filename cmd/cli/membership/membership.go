package membership

import (
	"os"

	"github.com/apex/log"
	"github.com/fatih/color"
	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	eerror "github.com/wesleimp/github-terraform/cmd/cli/error"
	"github.com/wesleimp/github-terraform/pkg/config"
	"github.com/wesleimp/github-terraform/pkg/context"
	"github.com/wesleimp/github-terraform/pkg/membership"
	"golang.org/x/oauth2"
)

// Cmd config
type Cmd struct {
	Cmd     *cobra.Command
	options options
}

type options struct {
	state   string
	dest    string
	token   string
	perPage int
	page    int
}

// NewCmd creates a repository cmd
func NewCmd() *Cmd {
	root := &Cmd{}

	var commands = &cobra.Command{
		Use:           "memberships",
		Short:         "Import organization memberships for the authenticated user",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Info(color.New(color.Bold).Sprint("Importing..."))

			_, err := startImport(root.options)
			if err != nil {
				return eerror.Wrap(err, color.New(color.Bold).Sprintf("Import error"))
			}

			log.Infof(color.New(color.Bold).Sprintf("Import Succeeded"))
			return nil
		},
	}

	//
	commands.Flags().StringVarP(&root.options.state, "state", "s", "", `Filter memberships to include only those with the specified state. Possible values are: "active", "pending"`)
	commands.Flags().StringVarP(&root.options.dest, "dest", "d", "./output", "Path that will contains the output files")
	commands.Flags().StringVar(&root.options.token, "token", "", "Github token. This property is not necessary if you already exported $GITHUB_TOKEN")
	commands.Flags().IntVar(&root.options.perPage, "per-page", 100, "Items per page")
	commands.Flags().IntVar(&root.options.page, "page", 1, "Current page")

	root.Cmd = commands
	return root
}

func startImport(o options) (*context.Context, error) {
	ctx := context.New(&config.Config{
		Membership: config.Membership{},
	})
	ctx = setupContext(ctx, o)

	err := membership.Import(ctx)
	if err != nil {
		return ctx, err
	}

	return ctx, nil
}

func setupContext(ctx *context.Context, o options) *context.Context {
	ctx.Config.Membership.State = o.state
	ctx.Config.Membership.Dest = o.dest
	ctx.Config.Membership.PerPage = o.perPage
	ctx.Config.Membership.Page = o.page

	if o.token == "" {
		ctx.Token = os.Getenv("GITHUB_TOKEN")
	} else {
		ctx.Token = o.token
	}

	ctx.Client = setupClient(ctx)

	return ctx
}

func setupClient(ctx *context.Context) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ctx.Token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	return client
}
