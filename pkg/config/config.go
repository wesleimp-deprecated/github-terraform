package config

// Config struct
type Config struct {
	Organization Organization
	Repository   Repository
	Team         Team
	IssueLabel   IssueLabel
	Membership   Membership
}

// Organization Config
type Organization struct {
	Repository OrganizationRepository
}

// PageOptions pagination options
type PageOptions struct {
	PerPage int
	Page    int
}

// OrganizationRepository config
type OrganizationRepository struct {
	Name string
	Org  string
	Dest string
	Type string
	PageOptions
}

// Repository config
type Repository struct {
	Name string
	User string
	Dest string
	Type string
	PageOptions

	Collaborator
	Project
}

// Team config
type Team struct {
	Name    string
	Org     string
	Dest    string
	Token   string
	PerPage int
	Page    int
}

// Collaborator config
type Collaborator struct {
	Repo  string
	Owner string
	Dest  string
	PageOptions
}

// Project cofig
type Project struct {
	Repo  string
	Owner string
	State string
	Dest  string
	Token string
	PageOptions
}

// IssueLabel config
type IssueLabel struct {
	Repo    string
	Owner   string
	Dest    string
	PerPage int
	Page    int
}

// Membership config
type Membership struct {
	State   string
	PerPage int
	Page    int
	Dest    string
}
