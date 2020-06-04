package config

type Repository struct {
	Name string
	Org  string
	User string
	Dest string
	Type string
}

type Config struct {
	Repository Repository
}
