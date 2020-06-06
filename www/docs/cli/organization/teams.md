# Teams

Import teams

## Usage

```sh
$ github-terraform org teams --help     

Import organization teams

Usage:
  github-terraform organization teams [flags]

Flags:
  -d, --dest string    Path that will contains the output files (default "./output")
  -h, --help           help for teams
  -n, --name string    Team name.
  -o, --org string     Team organization name
      --page int       Current page (default 1)
      --per-page int   Items per page (default 100)
      --token string   Github token. This property is not necessary if you already exported $GITHUB_TOKEN
```

**Example**

```sh
$ github-terraform org teams -o my-org -d ./output/teams --per-page 5
```