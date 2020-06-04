package repository

import "github.com/spf13/cobra"

// Cmd config
type Cmd struct {
	Cmd     *cobra.Command
	options options
}

type options struct {
	name  string
	owner string
	dest  string
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

	commands.Flags().StringVarP(&root.options.name, "name", "n", "", `Repository name. If multiple, should be splited by ","`)
	commands.Flags().StringVarP(&root.options.owner, "owner", "o", "", "Repository owner")
	commands.Flags().StringVarP(&root.options.dest, "dest", "d", "", "Path that will contains the output files")

	root.Cmd = commands
	return root
}

func run(cmd *cobra.Command, args []string) error {
	return nil
}
