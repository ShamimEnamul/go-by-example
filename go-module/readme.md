### Go Module
- To solve the version problem against the dependencies

# Itroduce go module 

```
    go mod init [module name]
    go mod tidy // import all dependencies
```

#### check go dependency list
```bash
go list -m all

```

inside the go.mod file there should have indirect dependencies too. Besides if you introduce any dependency inside the project, then the go module file import the dependency according to version


**Summary**
Go modules are the future of dependency management in Go. Module functionality is now available in all supported Go versions (that is, in Go 1.11 and Go 1.12).

This post introduced these workflows using Go modules:

- `go mod init` creates a new module, initializing the go.mod file that describes it.
- `go build`, go test, and other package-building commands add new dependencies to go.mod as needed.
- `go list -m all` prints the current module’s dependencies.
- `go get` changes the required version of a dependency (or adds a new dependency).
- `go mod tidy` removes unused dependencies.