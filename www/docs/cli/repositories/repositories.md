# Repositories

Import your repositories

## Usage

```sh
$ github-terraform repos --help

Import user repositories

Usage:
  github-terraform repositories [flags]
  github-terraform repositories [command]

Aliases:
  repositories, repos

Available Commands:
  collaborator Import repository collaborator

Flags:
  -d, --dest string    Path that will contains the output files (default "./output")
  -h, --help           help for repositories
  -n, --name string    Repository name. The name must contains owner/repo
      --page int       Current page (default 1)
      --per-page int   Items per page (default 100)
      --token string   Github token. This property is not necessary if you already exported GITHUB_TOKEN
  -t, --type string    Repository type. Could be public or private
  -u, --user string    Repository user

Use "github-terraform repositories [command] --help" for more information about a command.
```

**Example**

```sh
$ github-terraform repos -u wesleimp -d ./output/my-repos --per-page 5 -t public
```