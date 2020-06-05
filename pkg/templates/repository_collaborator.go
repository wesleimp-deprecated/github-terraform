package templates

// RepositoryCollaborator template
var RepositoryCollaborator = `resource "github_repository_collaborator" "repo_collaborator_{{ .Repository }}_{{ .Username }}" {
  repository = "{{ .Repository }}"
  username   = "{{ .Username }}"
  permission = "{{ .Permission }}"
}
`
