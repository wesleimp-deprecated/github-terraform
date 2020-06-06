package repositories

import (
	"strings"

	"github.com/apex/log"
	"github.com/fatih/color"
	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"github.com/wesleimp/github-terraform/internal/output"
	"github.com/wesleimp/github-terraform/internal/tmpl"
	"github.com/wesleimp/github-terraform/pkg/context"
	"github.com/wesleimp/github-terraform/pkg/templates"
)

// Import repositories
func Import(ctx *context.Context) error {
	log.WithFields(log.Fields{
		"name":     ctx.Config.Organization.Repository.Name,
		"org":      ctx.Config.Organization.Repository.Org,
		"type":     ctx.Config.Organization.Repository.Type,
		"per-page": ctx.Config.Organization.Repository.PerPage,
		"page":     ctx.Config.Organization.Repository.Page,
	}).Debug("Importing repositorires")

	if ctx.Config.Organization.Repository.Type != "" &&
		ctx.Config.Organization.Repository.Type != "private" &&
		ctx.Config.Organization.Repository.Type != "public" {
		return errors.New("Invalid repository type. Should be private or public")
	}

	if ctx.Config.Organization.Repository.Name != "" {
		return importRepo(ctx, ctx.Config.Organization.Repository.Org, ctx.Config.Organization.Repository.Name)
	}

	if ctx.Config.Organization.Repository.Org != "" {
		err := importReposByOrg(ctx, ctx.Config.Organization.Repository.Org)
		if err != nil {
			return err
		}
	}

	return nil
}

func importReposByOrg(ctx *context.Context, org string) error {
	rr, _, err := ctx.Client.Repositories.ListByOrg(ctx, org, &github.RepositoryListByOrgOptions{
		Type: ctx.Config.Organization.Repository.Type,
		ListOptions: github.ListOptions{
			PerPage: ctx.Config.Organization.Repository.PerPage,
			Page:    ctx.Config.Organization.Repository.Page,
		},
	})
	if err != nil {
		return errors.Wrap(err, "Error listing repos by org")
	}

	for _, r := range rr {
		err := importRepo(ctx, org, r.GetName())
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
		"DefaultBranch":     r.GetDefaultBranch(),
		"Topics":            strings.Join(r.Topics, ","),
	}).Apply(templates.Repository)
	if err != nil {
		return err
	}

	err = output.Save(ctx.Config.Organization.Repository.Dest, r.GetName(), content)
	if err != nil {
		return errors.Wrapf(err, "Error on save output file. Repo: %s", r.GetFullName())
	}

	return nil
}
