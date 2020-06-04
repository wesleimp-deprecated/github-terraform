package config

type Repository struct {
	Name    string
	Org     string
	User    string
	Dest    string
	Type    string
	PerPage int
	Page    int
}

type Config struct {
	Repository Repository
}
