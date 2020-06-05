package config

// Config struct
type Config struct {
	Repository Repository
	Team       Team
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
