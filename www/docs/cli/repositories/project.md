# Projects

Import repository projects

## Usage

```sh 
$ github-terraform repository project --help

Import repository project

Usage:
  github-terraform repositories project [flags]

Flags:
  -d, --dest string    Path that will contains the output files (default "./output")
  -h, --help           help for project
  -o, --owner string   Repository owner
      --page int       Current page (default 1)
      --per-page int   Items per page (default 100)
  -r, --repo string    Repository name
      --token string   Github token. This property is not necessary if you already exported $GITHUB_TOKEN
```

**Example**

```sh
$ github-terraform repos project -o my-org -r my-repo
```