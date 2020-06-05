package config

// Config struct
type Config struct {
	Repository             Repository
	RepositoryCollaborator RepositoryCollaborator
	Team                   Team
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
