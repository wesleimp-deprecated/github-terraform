package repository

import (
	"strings"

	"github.com/apex/log"
	"github.com/fatih/color"
	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"github.com/wesleimp/github-terraform/internal/output"
	"github.com/wesleimp/github-terraform/internal/tmpl"
	"github.com/wesleimp/github-terraform/pkg/context"
)

// Import repositories
func Import(ctx *context.Context) error {
	log.WithFields(log.Fields{
		"name":     ctx.Config.Repository.Name,
		"org":      ctx.Config.Repository.Org,
		"user":     ctx.Config.Repository.User,
		"type":     ctx.Config.Repository.Type,
		"per-page": ctx.Config.Repository.PerPage,
		"page":     ctx.Config.Repository.Page,
	}).Debug("Importing repositorires")

	if len(ctx.Config.Repository.Type) > 0 {
		if ctx.Config.Repository.Type != "private" || ctx.Config.Repository.Type != "public" {
			return errors.New("Invalid repository type. Should be private or public")
		}
	}

	if ctx.Config.Repository.Name != "" {
		return importRepoByName(ctx, ctx.Config.Repository.Name)
	}

	if ctx.Config.Repository.Org != "" {
		err := importReposByOrg(ctx, ctx.Config.Repository.Org)
		if err != nil {
			return err
		}
	}

	if ctx.Config.Repository.Name != "" {
		err := importReposByOrg(ctx, ctx.Config.Repository.Name)
		if err != nil {
			return err
		}
	}

	return nil
}

func importRepoByName(ctx *context.Context, name string) error {
	ownerRepo := strings.Split(name, "/")
	if len(ownerRepo) != 2 {
		return errors.New("Invalid repository name for %s. The name must be owner/repo")
	}

	err := importRepo(ctx, ownerRepo[0], ownerRepo[1])
	if err != nil {
		return err
	}

	return nil
}

func importReposByOrg(ctx *context.Context, org string) error {
	rr, _, err := ctx.Client.Repositories.ListByOrg(ctx, org, &github.RepositoryListByOrgOptions{
		Type: ctx.Config.Repository.Type,
		ListOptions: github.ListOptions{
			PerPage: ctx.Config.Repository.PerPage,
			Page:    ctx.Config.Repository.Page,
		},
	})
	if err != nil {
		return errors.Wrap(err, "Error listing repos by org")
	}

	for _, r := range rr {
		err := importRepo(ctx, ctx.Config.Repository.Org, r.GetName())
		if err != nil {
			return err
		}
	}

	return nil
}

func importReposByUser(ctx *context.Context, user string) error {
	rr, _, err := ctx.Client.Repositories.List(ctx, user, &github.RepositoryListOptions{
		Type: ctx.Config.Repository.Type,
		ListOptions: github.ListOptions{
			PerPage: ctx.Config.Repository.PerPage,
			Page:    ctx.Config.Repository.Page,
		},
	})

	if err != nil {
		return errors.Wrap(err, "Error listing repos by user")
	}

	for _, r := range rr {
		err := importRepo(ctx, ctx.Config.Repository.User, r.GetName())
		if err != nil {
			return err
		}
	}

	return nil
}

func importRepo(ctx *context.Context, owner, repo string) error {
	color.New(color.Bold).Printf("Importing %s/%s\n", owner, repo)
	r, _, err := ctx.Client.Repositories.Get(ctx, owner, repo)
	if err != nil {
		return errors.Wrapf(err, "Error getting repository %s/%s", owner, repo)
	}

	content, err := tmpl.New().WithFields(tmpl.Fields{
		"Name":              r.GetName(),
		"Description":       r.GetDescription(),
		"Private":           r.GetPrivate(),
		"AllowMergeCommit":  r.GetAllowMergeCommit(),
		"AllowRebaseMerge":  r.GetAllowRebaseMerge(),
		"AllowSquashMerge":  r.GetAllowSquashMerge(),
		"Archived":          r.GetArchived(),
		"AutoInit":          r.GetAutoInit(),
		"GitignoreTemplate": r.GetGitignoreTemplate(),
		"LicenseTemplate":   r.GetLicenseTemplate(),
		"HasDownloads":      r.GetHasDownloads(),
		"HasIssues":         r.GetHasIssues(),
		"HasProjects":       r.GetHasProjects(),
		"HasWiki":           r.GetHasWiki(),
		"HomepageURL":       r.GetHomepage(),
	}).Apply(Template)
	if err != nil {
		return err
	}

	err = output.Save(ctx.Config.Repository.Dest, r.GetName(), content)
	if err != nil {
		return errors.Wrapf(err, "Error on save output file. Repo: %s", r.GetFullName())
	}

	return nil
}
