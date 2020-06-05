# Install

## Compiling from source

### Clone

```sh
$ git clone git@github.com:wesleimp/github-terraform

$ cd github-terraform
```

### Download dependencies

```sh
$ go get ./...
```

### Build

```sh
$ go build -o github-terraform .
```

### Verify it works

```sh
$ ./github-terraform --help
```