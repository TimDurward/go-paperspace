# go-paperspace

[![GoDoc](https://godoc.org/github.com/TimDurward/go-paperspace?status.svg)](https://godoc.org/github.com/TimDurward/go-paperspace)
 
go-paperspace is a Go client library for accessing the [Paperspace API](https://www.paperspace.com/api).

# Usage
```go
import "github.com/timdurward/go-paperspace"
```

Construct a new Paperspace client, then use the various services on the client to access different parts of the Paperspace API. For example:

```go
config := &paperspace.Config{
  APIKey: "<API KEY>",
}

// Default http.Client
client, _ := paperspace.NewClient(config, nil)

// Create Machine
// https://paperspace.github.io/paperspace-node/machines.html#.create
m := &paperspace.MachinesRequest{
  Region:      "West Coast (CA1)",
  MachineType: "Air",
  Size:        50,
  BillingType: "monthly",
  MachineName: "Provisioned using Go API client.",
  TemplateID:  "<template_id>",
}
  
machine, _, err := client.Machines.Create(m)
```


# 🔨 WIP 🔧
### Client v1 Coverage
- [x] Base Paperspace Client
- [x] Machines Service
- [x] Networks Service
- [ ] Scripts Service
- [ ] Templates Service
- [ ] Jobs Service
- [ ] Tests
- [ ] Versioning

# Contributing
#### Feel free to submit a Github issue in regards to questions.

If you submit a pull request, please keep the following guidelines in mind:

1. Code should be `go fmt` compliant.
2. Types, structs and funcs should be documented.

## Getting set up

Assuming your `$GOPATH` is set up according to your desires, run:

```sh
go get github.com/timdurward/go-paperspace
```
