# Repositories

Import your repositories

## Usage

```sh
$ github-terraform org repos --help                 

Import organization repositories

Usage:
  github-terraform organization repositories [flags]

Aliases:
  repositories, repos

Flags:
  -d, --dest string    Path that will contains the output files (default "./output")
  -h, --help           help for repositories
  -n, --name string    Repository name.
  -o, --org string     Organization name.
      --page int       Current page (default 1)
      --per-page int   Items per page (default 100)
      --token string   Github token. This property is not necessary if you already exported GITHUB_TOKEN
  -t, --type string    Repository type. Could be public or private
```

**Example**

```sh
$ github-terraform org repos -o my-org -d ./output/my-repos --per-page 5 -t public
```