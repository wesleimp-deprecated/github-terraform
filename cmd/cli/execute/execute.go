package execute

import (
	"fmt"

	"github.com/apex/log"
	"github.com/spf13/cobra"
	"github.com/wesleimp/github-terraform/cmd/cli/error"
	"github.com/wesleimp/github-terraform/cmd/cli/issuelabel"
	"github.com/wesleimp/github-terraform/cmd/cli/membership"
	"github.com/wesleimp/github-terraform/cmd/cli/repository"
	"github.com/wesleimp/github-terraform/cmd/cli/repositorycollaborator"
	"github.com/wesleimp/github-terraform/cmd/cli/teams"
)

// Run starts the execution
func Run(version string, exit func(int), args []string) {
	fmt.Println()
	defer fmt.Println()
	newRootCmd(version, exit).Run(args)
}

func (cmd *rootCmd) Run(args []string) {
	cmd.cmd.SetArgs(args)

	if err := cmd.cmd.Execute(); err != nil {
		var code = 1
		var msg = "command failed"
		if eerr, ok := err.(*error.Exit); ok {
			code = eerr.Code
			if eerr.Details != "" {
				msg = eerr.Details
			}
		}

		log.WithError(err).Error(msg)
		cmd.exit(code)
	}
}

type rootCmd struct {
	cmd  *cobra.Command
	exit func(int)
}

func newRootCmd(version string, exit func(int)) *rootCmd {
	var root = &rootCmd{
		exit: exit,
	}
	var cmd = &cobra.Command{
		Use:           "github-terraform",
		Short:         "Import your github into terraform",
		Version:       version,
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cmd.AddCommand(
		issuelabel.NewCmd().Cmd,
		repository.NewCmd().Cmd,
		repositorycollaborator.NewCmd().Cmd,
		teams.NewCmd().Cmd,
		membership.NewCmd().Cmd,
	)

	root.cmd = cmd
	return root
}
