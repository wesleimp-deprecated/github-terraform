package templates

// Repository template
var Repository = `resource "github_repository" "repository_{{ .Name }}" {
  name               = "{{ .Name }}"
  description        = "{{ .Description }}"
  private            = {{ .Private }}
  has_wiki           = {{ .HasWiki }}
  has_projects       = {{ .HasProjects }}
  has_downloads      = {{ .HasDownloads }}
  has_issues         = {{ .HasIssues }}
  archived           = {{ .Archived }}
  allow_merge_commit = {{ .AllowMergeCommit }}
  allow_rebase_merge = {{ .AllowRebaseMerge }}
  allow_squash_merge = {{ .AllowSquashMerge }}
  auto_init          = {{ .AutoInit }}
  gitignore_template = "{{ .GitignoreTemplate }}"
  license_template   = "{{ .LicenseTemplate }}"
  homepage_url       = "{{ .HomepageURL }}"
  default_brach      = "{{ .DefaultBranch }}"
  topics             = "{{ .Topics }}"
}
`

// RepositoryCollaborator template
var RepositoryCollaborator = `resource "github_repository_collaborator" "repo_collaborator_{{ .Repository }}_{{ .Username }}" {
  repository = "{{ .Repository }}"
  username   = "{{ .Username }}"
  permission = "{{ .Permission }}"
}
`

// RepositoryProject template
var RepositoryProject = `resource "github_repository_project" "project_{{ .Name }}" {
  name       = "{{ .Name }}"
  repository = "{{ .Repository }}"
  body       = "{{ .Body }}"
}
`
