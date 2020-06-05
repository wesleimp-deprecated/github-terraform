package membership

import (
	"errors"

	"github.com/fatih/color"
	"github.com/google/go-github/github"
	"github.com/wesleimp/github-terraform/internal/output"
	"github.com/wesleimp/github-terraform/internal/tmpl"
	"github.com/wesleimp/github-terraform/pkg/context"
	"github.com/wesleimp/github-terraform/pkg/templates"
)

// Import org memberships
func Import(ctx *context.Context) error {
	color.New(color.Bold).Printf("Importing organization memberships\n")

	if ctx.Config.Membership.State != "" &&
		ctx.Config.Membership.State != "active" &&
		ctx.Config.Membership.State != "pending" {
		return errors.New(`Ivalid state argument. State shoud be "active" or "pending" when informed`)
	}

	mm, _, err := ctx.Client.Organizations.ListOrgMemberships(ctx, &github.ListOrgMembershipsOptions{
		State: ctx.Config.Membership.State,
		ListOptions: github.ListOptions{
			Page:    ctx.Config.Membership.Page,
			PerPage: ctx.Config.Membership.PerPage,
		},
	})
	if err != nil {
		return err
	}

	for _, m := range mm {
		var username = m.GetUser().GetName()
		if username == "" {
			continue
		}

		content, err := tmpl.New().WithFields(tmpl.Fields{
			"Name": username,
			"Role": m.GetRole(),
		}).Apply(templates.Membership)
		if err != nil {
			return err
		}

		err = output.Save(ctx.Config.Membership.Dest, username, content)
		if err != nil {
			return err
		}
	}

	return nil
}
