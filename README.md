# Go-Paperspace

[![GoDoc](https://godoc.org/github.com/TimDurward/go-paperspace?status.svg)](https://godoc.org/github.com/TimDurward/go-paperspace)
 
Go-Paperspace is a Go client library for accessing the Paperspace API.

You can view the client API docs here: https://godoc.org/github.com/TimDurward/go-paperspace

You can view Paperspace API docs here: https://www.paperspace.com/api

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

opts := &paperspace.MachineOptions{
  Region:      "West Coast (CA1)",
  MachineType: "GPU+",
}

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
  
machine, _, err := paperspaceClient.Machines.Create(m)
```


# 🔨 WIP 🔧
### Done:
1. Paperspace Client
2. Base http Requests and Response handling.
3. Scripts methods semi complete.
4. Some of Machines methods complete.

### Not Done:
1. Most of Machine methods
2. Network methods
3. Project methods
4. Jobs methods
5. Tests
6. Versioning

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
