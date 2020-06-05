# Issue Lables

Import repository issue labels

## Usage

```sh 
$ github-terraform issue-labels --help

Import repository issue labels

Usage:
  github-terraform issue-labels [flags]

Flags:
  -d, --dest string    Path that will contains the output files (default "./output")
  -h, --help           help for issue-labels
  -o, --owner string   Repository owner
      --page int       Current page (default 1)
      --per-page int   Items per page (default 100)
  -r, --repo string    Repository name
      --token string   Github token. This property is not necessary if you already exported $GITHUB_TOKEN
```

**Example**

```
$ github-terraform issue-labels -o my-org -r my-repo
```