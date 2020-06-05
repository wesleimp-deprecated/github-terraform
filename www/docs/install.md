# Install

## Using the pre-compiled binaries

### Snapcraft

```sh
$ sudo snap install --classic github-terraform
```

### deb/rpm

Download the `.deb` or `.rpm` from the [releases page](https://github.com/wesleimp/github-terraform/releases) and install with `dpkg -i` and `rpm -i` respectively.

### Manually

Download the pre-compiled binaries from the [releases page](https://github.com/wesleimp/github-terraform/releases) and move to the desired location.

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