# couchbase-cloud-go-client

[![Go Reference](https://pkg.go.dev/badge/github.com/couchbaselabs/couchbase-cloud-go-client.svg)](https://pkg.go.dev/github.com/couchbaselabs/couchbase-cloud-go-client)

Work in progress Go wrapper for the Couchbase Cloud REST API

## Usage

`go get github.com/couchbaselabs/couchbase-cloud-go-client`

### Import
```go
import (
    "github.com/couchbaselabs/couchbase-cloud-go-client/couchbasecloud"
)
```

### Client

```go
client := couchbasecloud.NewClient(os.Getenv("CBC_ACCESS_KEY"), os.Getenv("CBC_SECRET_KEY"))
```

### List all clouds
Options can be specified with the `ListCloudOptions` type (some are optional)
```go
return client.ListCloudPages(nil, func(clouds couchbasecloud.Clouds, b bool) bool {
    for _, cloud := range clouds {
        fmt.Println(cloud.Name)
    }
})
```

### List all clusters
Options can be specified with the `ListClustersOptions` type (some are optional)
```go
return client.ListClusterPages(nil, func(clusters couchbasecloud.Clusters, b bool) bool {
    for _, cloud := range clouds {
        fmt.Println(cluster.Name)
    }
})
```

## TODO

- Unit/integration testing
- Add documentation strings to all public types and functions so that documentation can be auto-generated from the code.
- Add linter https://github.com/golangci/golangci-lint to help maintain code quality.
- Auto-generate boilerplate code for the rest of the client using OpenAPI spec for the Couchbase Cloud API. This could 
  also be used to generate a server stub for unit testing or as an emulator for integration level tests.