package templates

// Membership template
var Membership = `resource "github_membership" "membership_{{ .Name }}" {
  username = "{{ .Name }}"
  role     = "{{ .Role }}"
}
`
