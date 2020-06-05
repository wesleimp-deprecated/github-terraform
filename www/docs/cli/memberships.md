# Memberships

Import memberships for the authenticated user

## Usage

```sh
$ github-terraform memberships --help

Import organization memberships for the authenticated user

Usage:
  github-terraform memberships [flags]

Flags:
  -d, --dest string    Path that will contains the output files (default "./output")
  -h, --help           help for memberships
      --page int       Current page (default 1)
      --per-page int   Items per page (default 100)
  -s, --state string   Filter memberships to include only those with the specified state. Possible values are: "active", "pending"
      --token string   Github token. This property is not necessary if you already exported $GITHUB_TOKEN
```