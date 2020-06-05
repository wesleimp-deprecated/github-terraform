package templates

// Teams template
var Teams = `resource "github_team" "team_{{ .Name }}" {
  name           = "{{ .Name }}"
  description    = "{{ .Description }}"
  privacy        = "{{ .Privacy }}"
  {{ if gt .ParentID 0}}parent_team_id = {{ .ParentID }}{{- end}}
}
`
