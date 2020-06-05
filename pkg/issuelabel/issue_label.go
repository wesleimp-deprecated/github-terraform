package issuelabel

import (
	"strings"

	"github.com/fatih/color"
	"github.com/google/go-github/github"
	"github.com/pkg/errors"

	"github.com/wesleimp/github-terraform/internal/output"
	"github.com/wesleimp/github-terraform/internal/tmpl"
	"github.com/wesleimp/github-terraform/pkg/context"
	"github.com/wesleimp/github-terraform/pkg/templates"
)

// Import issue labels
func Import(ctx *context.Context) error {
	if ctx.Config.IssueLabel.Repo == "" || ctx.Config.IssueLabel.Owner == "" {
		return errors.New("Repository and owner property should be informed to import issue labels")
	}

	err := importIssueLabels(ctx, ctx.Config.IssueLabel.Owner, ctx.Config.IssueLabel.Repo)
	if err != nil {
		return err
	}

	return nil
}

func importIssueLabels(ctx *context.Context, owner, repo string) error {
	color.New(color.Bold).Printf("Importing issue labels for %s/%s", owner, repo)

	ll, _, err := ctx.Client.Issues.ListLabels(ctx, owner, repo, &github.ListOptions{
		Page:    ctx.Config.IssueLabel.Page,
		PerPage: ctx.Config.IssueLabel.PerPage,
	})
	if err != nil {
		return errors.Wrapf(err, "Error listing issue labels for %s/%s", owner, repo)
	}

	for _, l := range ll {
		resName := strings.ReplaceAll(l.GetName(), " ", "_")
		content, err := tmpl.New().WithFields(tmpl.Fields{
			"Repository":   repo,
			"Name":         l.GetName(),
			"ResourceName": resName,
			"Color":        l.GetColor(),
			"Description":  l.GetDescription(),
			"URL":          l.GetURL(),
		}).Apply(templates.IssueLabel)
		if err != nil {
			return err
		}

		err = output.Save(ctx.Config.IssueLabel.Dest, resName, content)
		if err != nil {
			return err
		}
	}

	return nil
}
