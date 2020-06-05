package templates

// IssueLabel template
var IssueLabel = `resource "github_issue_label" "issue_label_{{ .Repository }}_{{ .ResourceName }}" {
  repository  = "{{ .Repository }}"
  name        = "{{ .Name }}"
	color       = "{{ .Color }}"
	description = "{{ .Description }}"
	url         = "{{ .URL }}"
}
`
