# create-go-workspace

Simple CLI to generate the directory structure for go workspaces.

## Get started

To install, use `go get`:

```bash
$ go get -u github.com/mdaverde/create-go-workspace
```

## Usage

```bash
$ create-go-workspace github.com/mdaverde/great-idea
Creating github.com/mdaverde/great-idea workspace...
Created: /Users/mdaverde/Development/great-idea
Created: /Users/mdaverde/Development/great-idea/src
Created: /Users/mdaverde/Development/great-idea/src/github.com/mdaverde/great-idea
Created: /Users/mdaverde/Development/great-idea/bin
Created: /Users/mdaverde/Development/great-idea/pkg
Created: /Users/mdaverde/Development/great-idea/.envrc
Created: /Users/mdaverde/Development/great-idea/src/github.com/mdaverde/great-idea/main.go
Created: /Users/mdaverde/Development/great-idea/src/github.com/mdaverde/great-idea/README.md
Done.
```

## Generates

- `src/`, `bin/`, `pkg/` directories
- `main.go`
- [.envrc](http://tammersaleh.com/posts/manage-your-gopath-with-direnv/) to manage setting $GOPATH
- README.md


## Directory Structure

```bash
$ tree -a great-idea
great-idea
├── .envrc
├── bin
├── pkg
└── src
    └── github.com
        └── mdaverde
            └── great-idea
                ├── README.md
                └── main.go

6 directories, 3 files
```

## Options

```
--silent, -s   suppress output (default: false)
--dir-env      generate .envrc (default: true)
--main-go      generate main.go (default: true)
--read-me      generate README.md (default: true)
--help, -h     show help
--version, -v  print the version
```


## Contribution

1. Fork ([https://github.com/mdaverde/create-go-workspace/fork](https://github.com/mdaverde/create-go-workspace/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -w *.go`
1. Create a new Pull Request


## Author

[mdaverde](https://github.com/mdaverde)
