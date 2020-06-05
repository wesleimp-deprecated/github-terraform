package repositorycollaborator

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"github.com/wesleimp/github-terraform/internal/output"
	"github.com/wesleimp/github-terraform/internal/tmpl"
	"github.com/wesleimp/github-terraform/pkg/context"
	"github.com/wesleimp/github-terraform/pkg/templates"
)

// Import repository collaborators
func Import(ctx *context.Context) error {
	if ctx.Config.RepositoryCollaborator.Repo == "" || ctx.Config.RepositoryCollaborator.Owner == "" {
		return errors.New("Repository and owner property should be informed to import collaborators")
	}

	err := importRepo(ctx, ctx.Config.RepositoryCollaborator.Owner, ctx.Config.RepositoryCollaborator.Repo)
	if err != nil {
		return err
	}

	return nil
}

func importRepo(ctx *context.Context, owner, repo string) error {
	color.New(color.Bold).Printf("Importing collabortors for %s/%s\n", owner, repo)
	cc, _, err := ctx.Client.Repositories.ListCollaborators(ctx, owner, repo, &github.ListCollaboratorsOptions{
		ListOptions: github.ListOptions{
			Page:    ctx.Config.RepositoryCollaborator.Page,
			PerPage: ctx.Config.RepositoryCollaborator.PerPage,
		},
	})
	if err != nil {
		return errors.Wrapf(err, "Error listing repository collaborators for %s/%s", owner, repo)
	}

	for _, c := range cc {
		permission := getPermission(c.GetPermissions())
		content, err := tmpl.New().WithFields(tmpl.Fields{
			"Repository": repo,
			"Username":   c.GetLogin(),
			"Permission": permission,
		}).Apply(templates.RepositoryCollaborator)
		if err != nil {
			return err
		}

		err = output.Save(ctx.Config.RepositoryCollaborator.Dest, fmt.Sprintf("%s_%s", repo, c.GetLogin()), content)
		if err != nil {
			return err
		}
	}

	return nil
}

func getPermission(perms map[string]bool) string {
	if admin, ok := perms["admin"]; ok && admin {
		return "admin"
	}

	if triage, ok := perms["triage"]; ok && triage {
		return "triage"
	}

	if maintain, ok := perms["maintain"]; ok && maintain {
		return "maintain"
	}

	if push, ok := perms["push"]; ok && push {
		return "push"
	}

	if pull, ok := perms["pull"]; ok && pull {
		return "pull"
	}

	return ""
}
