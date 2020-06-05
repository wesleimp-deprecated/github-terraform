package teams

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

// Import teams
func Import(ctx *context.Context) error {
	log.WithFields(log.Fields{
		"Org": ctx.Config.Team.Org,
	}).Debug("Importing teams")

	tt, _, err := ctx.Client.Teams.ListTeams(ctx, ctx.Config.Team.Org, &github.ListOptions{
		PerPage: ctx.Config.Team.PerPage,
		Page:    ctx.Config.Team.Page,
	})
	if err != nil {
		return errors.Wrapf(err, "Error listing organizations teams. Org: %s", ctx.Config.Team.Org)
	}

	for _, t := range tt {
		color.New(color.Bold).Printf("Importing %s team", t.GetName())

		resName := strings.ReplaceAll(t.GetName(), " ", "_")
		resName = strings.ReplaceAll(t.GetName(), "/", "_")

		content, err := tmpl.New().WithFields(tmpl.Fields{
			"ResourceName": resName,
			"Name":         t.GetName(),
			"Description":  t.GetDescription(),
			"Privacy":      t.GetPrivacy(),
			"ParentID":     t.GetParent().GetID(),
		}).Apply(templates.Teams)
		if err != nil {
			return err
		}

		err = output.Save(ctx.Config.Team.Dest, resName, content)
		if err != nil {
			return err
		}
	}

	return nil
}
