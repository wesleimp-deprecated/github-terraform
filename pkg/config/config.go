package config

// Config struct
type Config struct {
	Repository             Repository
	RepositoryCollaborator RepositoryCollaborator
	Team                   Team
	IssueLabel             IssueLabel
	Membership             Membership
}

// Repository config
type Repository struct {
	Name    string
	Org     string
	User    string
	Dest    string
	Type    string
	PerPage int
	Page    int
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

// RepositoryCollaborator config
type RepositoryCollaborator struct {
	Repo    string
	Owner   string
	Dest    string
	PerPage int
	Page    int
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
