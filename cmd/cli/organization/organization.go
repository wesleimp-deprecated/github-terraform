package organization

import (
	"github.com/spf13/cobra"
	"github.com/wesleimp/github-terraform/cmd/cli/organization/repositories"
	"github.com/wesleimp/github-terraform/cmd/cli/organization/teams"
)

type Cmd struct {
	Cmd *cobra.Command
}

func NewCmd() *Cmd {
	orgCmd := &Cmd{}

	var command = &cobra.Command{
		Use:           "organization",
		Short:         "Import your organization",
		Aliases:       []string{"org"},
		SilenceUsage:  true,
		SilenceErrors: false,
	}

	command.AddCommand(
		repositories.NewCmd().Cmd,
		teams.NewCmd().Cmd,
	)

	orgCmd.Cmd = command
	return orgCmd
}
