package repository

// Template is a generic template for all repositories
var Template = `resource "github_repository" "repository_{{ .Name }}" {
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
}
`
