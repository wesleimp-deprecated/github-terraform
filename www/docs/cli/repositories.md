# Repositories

Import your repositories

## Usage

```sh
$ github-terraform repos --help

Import repositories

Usage:
  github-terraform repositories [flags]

Aliases:
  repositories, repos

Flags:
  -d, --dest string    Path that will contains the output files (default "./output")
  -h, --help           help for repositories
  -n, --name string    Repository name. The name must contains owner/repo
  -o, --org string     Repository organization
      --page int       Current page (default 1)
      --per-page int   Items per page (default 100)
      --token string   Github token. This property is not necessary if you already exported $GITHUB_TOKEN
  -t, --type string    Repository type. Could be public or private
  -u, --user string    Repository user

```

**Example**

```sh
$ github-terraform repos -u wesleimp -d ./output/my-repos --per-page 5 -t public
```