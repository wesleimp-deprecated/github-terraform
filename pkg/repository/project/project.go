package project

import (
	"github.com/fatih/color"
	"github.com/google/go-github/github"
	"github.com/wesleimp/github-terraform/internal/output"
	"github.com/wesleimp/github-terraform/internal/tmpl"
	"github.com/wesleimp/github-terraform/pkg/context"
	"github.com/wesleimp/github-terraform/pkg/templates"
)

// Import repository project
func Import(ctx *context.Context) error {
	color.New(color.Bold).Println("Importing repository project")
	pp, _, err := ctx.Client.Repositories.ListProjects(ctx, ctx.Config.Repository.Project.Owner, ctx.Config.Repository.Project.Repo, &github.ProjectListOptions{
		State: ctx.Config.Repository.Project.State,
		ListOptions: github.ListOptions{
			Page:    ctx.Config.Repository.Project.Page,
			PerPage: ctx.Config.Repository.Project.PerPage,
		},
	})
	if err != nil {
		return err
	}

	for _, p := range pp {
		content, err := tmpl.New().WithFields(tmpl.Fields{
			"Name":       p.GetName(),
			"Repository": ctx.Config.Repository.Project.Repo,
			"Body":       p.GetBody(),
		}).Apply(templates.RepositoryProject)
		if err != nil {
			return err
		}

		err = output.Save(ctx.Config.Repository.Project.Dest, p.GetName(), content)
		if err != nil {
			return err
		}
	}

	return nil
}
